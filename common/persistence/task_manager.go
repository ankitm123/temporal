package persistence

import (
	"context"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"
	persistencespb "go.temporal.io/server/api/persistence/v1"
	"go.temporal.io/server/common/persistence/serialization"
	"go.temporal.io/server/common/primitives/timestamp"
)

// Subqueue zero corresponds to "the queue" before migrating metadata to subqueues.
// For SQL: metadata operations apply to subqueue zero only, while tasks are stored in
// multiple subqueues.
// For Cassandra: subqueues are represented in the row type.
const SubqueueZero = 0

type taskManagerImpl struct {
	taskStore  TaskStore
	serializer serialization.Serializer
}

// NewTaskManager creates a new instance of TaskManager
func NewTaskManager(
	store TaskStore,
	serializer serialization.Serializer,
) TaskManager {
	return &taskManagerImpl{
		taskStore:  store,
		serializer: serializer,
	}
}

func (m *taskManagerImpl) Close() {
	m.taskStore.Close()
}

func (m *taskManagerImpl) GetName() string {
	return m.taskStore.GetName()
}

func (m *taskManagerImpl) CreateTaskQueue(
	ctx context.Context,
	request *CreateTaskQueueRequest,
) (*CreateTaskQueueResponse, error) {
	taskQueueInfo := request.TaskQueueInfo
	if taskQueueInfo.LastUpdateTime == nil {
		panic("CreateTaskQueue encountered LastUpdateTime not set")
	}
	if taskQueueInfo.ExpiryTime == nil && taskQueueInfo.GetKind() == enumspb.TASK_QUEUE_KIND_STICKY {
		panic("CreateTaskQueue encountered ExpiryTime not set for sticky task queue")
	}
	taskQueueInfoBlob, err := m.serializer.TaskQueueInfoToBlob(taskQueueInfo)
	if err != nil {
		return nil, err
	}

	internalRequest := &InternalCreateTaskQueueRequest{
		NamespaceID:   request.TaskQueueInfo.GetNamespaceId(),
		TaskQueue:     request.TaskQueueInfo.GetName(),
		TaskType:      request.TaskQueueInfo.GetTaskType(),
		TaskQueueKind: request.TaskQueueInfo.GetKind(),
		RangeID:       request.RangeID,
		ExpiryTime:    taskQueueInfo.ExpiryTime,
		TaskQueueInfo: taskQueueInfoBlob,
	}
	if err := m.taskStore.CreateTaskQueue(ctx, internalRequest); err != nil {
		return nil, err
	}
	return &CreateTaskQueueResponse{}, nil
}

func (m *taskManagerImpl) UpdateTaskQueue(
	ctx context.Context,
	request *UpdateTaskQueueRequest,
) (*UpdateTaskQueueResponse, error) {
	taskQueueInfo := request.TaskQueueInfo
	if taskQueueInfo.LastUpdateTime == nil {
		panic("UpdateTaskQueue encountered LastUpdateTime not set")
	}
	if taskQueueInfo.ExpiryTime == nil && taskQueueInfo.GetKind() == enumspb.TASK_QUEUE_KIND_STICKY {
		panic("UpdateTaskQueue encountered ExpiryTime not set for sticky task queue")
	}
	taskQueueInfoBlob, err := m.serializer.TaskQueueInfoToBlob(taskQueueInfo)
	if err != nil {
		return nil, err
	}

	internalRequest := &InternalUpdateTaskQueueRequest{
		NamespaceID:   request.TaskQueueInfo.GetNamespaceId(),
		TaskQueue:     request.TaskQueueInfo.GetName(),
		TaskType:      request.TaskQueueInfo.GetTaskType(),
		RangeID:       request.RangeID,
		TaskQueueInfo: taskQueueInfoBlob,

		TaskQueueKind: request.TaskQueueInfo.GetKind(),
		ExpiryTime:    taskQueueInfo.ExpiryTime,

		PrevRangeID: request.PrevRangeID,
	}
	return m.taskStore.UpdateTaskQueue(ctx, internalRequest)
}

func (m *taskManagerImpl) GetTaskQueue(
	ctx context.Context,
	request *GetTaskQueueRequest,
) (*GetTaskQueueResponse, error) {
	response, err := m.taskStore.GetTaskQueue(ctx, &InternalGetTaskQueueRequest{
		NamespaceID: request.NamespaceID,
		TaskQueue:   request.TaskQueue,
		TaskType:    request.TaskType,
	})
	if err != nil {
		return nil, err
	}

	taskQueueInfo, err := m.serializer.TaskQueueInfoFromBlob(response.TaskQueueInfo)
	if err != nil {
		return nil, err
	}
	return &GetTaskQueueResponse{
		TaskQueueInfo: taskQueueInfo,
		RangeID:       response.RangeID,
	}, nil
}

func (m *taskManagerImpl) ListTaskQueue(
	ctx context.Context,
	request *ListTaskQueueRequest,
) (*ListTaskQueueResponse, error) {
	internalResp, err := m.taskStore.ListTaskQueue(ctx, request)
	if err != nil {
		return nil, err
	}
	taskQueues := make([]*PersistedTaskQueueInfo, len(internalResp.Items))
	for i, item := range internalResp.Items {
		tqi, err := m.serializer.TaskQueueInfoFromBlob(item.TaskQueue)
		if err != nil {
			return nil, err
		}
		taskQueues[i] = &PersistedTaskQueueInfo{
			Data:    tqi,
			RangeID: item.RangeID,
		}

	}
	return &ListTaskQueueResponse{
		Items:         taskQueues,
		NextPageToken: internalResp.NextPageToken,
	}, nil
}

func (m *taskManagerImpl) DeleteTaskQueue(
	ctx context.Context,
	request *DeleteTaskQueueRequest,
) error {
	return m.taskStore.DeleteTaskQueue(ctx, request)
}

func (m *taskManagerImpl) CreateTasks(
	ctx context.Context,
	request *CreateTasksRequest,
) (*CreateTasksResponse, error) {
	taskQueueInfo := request.TaskQueueInfo.Data
	taskQueueInfo.LastUpdateTime = timestamp.TimeNowPtrUtc()
	taskQueueInfoBlob, err := m.serializer.TaskQueueInfoToBlob(taskQueueInfo)
	if err != nil {
		return nil, err
	}

	tasks := make([]*InternalCreateTask, len(request.Tasks))
	for i, task := range request.Tasks {
		taskBlob, err := m.serializer.TaskInfoToBlob(task)
		if err != nil {
			return nil, serviceerror.NewUnavailablef("CreateTasks operation failed during serialization. Error : %v", err)
		}
		tasks[i] = &InternalCreateTask{
			TaskPass:   task.TaskPass,
			TaskId:     task.TaskId,
			ExpiryTime: task.Data.ExpiryTime,
			Task:       taskBlob,
		}
		if i < len(request.Subqueues) {
			tasks[i].Subqueue = request.Subqueues[i]
		}
	}
	internalRequest := &InternalCreateTasksRequest{
		NamespaceID:   request.TaskQueueInfo.Data.GetNamespaceId(),
		TaskQueue:     request.TaskQueueInfo.Data.GetName(),
		TaskType:      request.TaskQueueInfo.Data.GetTaskType(),
		RangeID:       request.TaskQueueInfo.RangeID,
		TaskQueueInfo: taskQueueInfoBlob,
		Tasks:         tasks,
	}
	return m.taskStore.CreateTasks(ctx, internalRequest)
}

func (m *taskManagerImpl) GetTasks(
	ctx context.Context,
	request *GetTasksRequest,
) (*GetTasksResponse, error) {
	if request.InclusiveMinTaskID >= request.ExclusiveMaxTaskID {
		return &GetTasksResponse{}, nil
	}

	internalResp, err := m.taskStore.GetTasks(ctx, request)
	if err != nil {
		return nil, err
	}
	tasks := make([]*persistencespb.AllocatedTaskInfo, len(internalResp.Tasks))
	for i, taskBlob := range internalResp.Tasks {
		task, err := m.serializer.TaskInfoFromBlob(taskBlob)
		if err != nil {
			return nil, serviceerror.NewUnavailablef("GetTasks failed to deserialize task: %s", err.Error())
		}
		tasks[i] = task
	}
	return &GetTasksResponse{Tasks: tasks, NextPageToken: internalResp.NextPageToken}, nil
}

func (m *taskManagerImpl) CompleteTasksLessThan(
	ctx context.Context,
	request *CompleteTasksLessThanRequest,
) (int, error) {
	return m.taskStore.CompleteTasksLessThan(ctx, request)
}

// GetTaskQueueUserData implements TaskManager
func (m *taskManagerImpl) GetTaskQueueUserData(ctx context.Context, request *GetTaskQueueUserDataRequest) (*GetTaskQueueUserDataResponse, error) {
	response, err := m.taskStore.GetTaskQueueUserData(ctx, request)
	if err != nil {
		return nil, err
	}
	data, err := m.serializer.TaskQueueUserDataFromBlob(response.UserData)
	if err != nil {
		return nil, err
	}
	return &GetTaskQueueUserDataResponse{UserData: &persistencespb.VersionedTaskQueueUserData{Version: response.Version, Data: data}}, nil
}

// UpdateTaskQueueUserData implements TaskManager
func (m *taskManagerImpl) UpdateTaskQueueUserData(ctx context.Context, request *UpdateTaskQueueUserDataRequest) error {
	internalRequest := &InternalUpdateTaskQueueUserDataRequest{
		NamespaceID: request.NamespaceID,
		Updates:     make(map[string]*InternalSingleTaskQueueUserDataUpdate, len(request.Updates)),
	}
	for taskQueue, update := range request.Updates {
		userData, err := m.serializer.TaskQueueUserDataToBlob(update.UserData.Data)
		if err != nil {
			return err
		}
		internalRequest.Updates[taskQueue] = &InternalSingleTaskQueueUserDataUpdate{
			Version:         update.UserData.Version,
			UserData:        userData,
			BuildIdsAdded:   update.BuildIdsAdded,
			BuildIdsRemoved: update.BuildIdsRemoved,
			Applied:         update.Applied,
			Conflicting:     update.Conflicting,
		}
	}
	return m.taskStore.UpdateTaskQueueUserData(ctx, internalRequest)
}

func (m *taskManagerImpl) ListTaskQueueUserDataEntries(ctx context.Context, request *ListTaskQueueUserDataEntriesRequest) (*ListTaskQueueUserDataEntriesResponse, error) {
	response, err := m.taskStore.ListTaskQueueUserDataEntries(ctx, request)
	if err != nil {
		return nil, err
	}
	entries := make([]*TaskQueueUserDataEntry, len(response.Entries))
	for i, entry := range response.Entries {
		data, err := m.serializer.TaskQueueUserDataFromBlob(entry.Data)
		if err != nil {
			return nil, err
		}
		entries[i] = &TaskQueueUserDataEntry{
			TaskQueue: entry.TaskQueue,
			UserData: &persistencespb.VersionedTaskQueueUserData{
				Data:    data,
				Version: entry.Version,
			},
		}
	}
	return &ListTaskQueueUserDataEntriesResponse{
		NextPageToken: response.NextPageToken,
		Entries:       entries,
	}, nil
}

func (m *taskManagerImpl) GetTaskQueuesByBuildId(ctx context.Context, request *GetTaskQueuesByBuildIdRequest) ([]string, error) {
	return m.taskStore.GetTaskQueuesByBuildId(ctx, request)
}

func (m *taskManagerImpl) CountTaskQueuesByBuildId(ctx context.Context, request *CountTaskQueuesByBuildIdRequest) (int, error) {
	return m.taskStore.CountTaskQueuesByBuildId(ctx, request)
}
