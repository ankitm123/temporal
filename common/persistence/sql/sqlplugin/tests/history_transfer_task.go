package tests

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/server/common/persistence/sql/sqlplugin"
	"go.temporal.io/server/common/shuffle"
)

type (
	historyHistoryTransferTaskSuite struct {
		suite.Suite
		*require.Assertions

		store sqlplugin.HistoryTransferTask
	}
)

const (
	testHistoryTransferTaskEncoding = "random encoding"
)

var (
	testHistoryTransferTaskData = []byte("random history transfer task data")
)

func NewHistoryTransferTaskSuite(
	t *testing.T,
	store sqlplugin.HistoryTransferTask,
) *historyHistoryTransferTaskSuite {
	return &historyHistoryTransferTaskSuite{
		Assertions: require.New(t),
		store:      store,
	}
}

func (s *historyHistoryTransferTaskSuite) SetupSuite() {

}

func (s *historyHistoryTransferTaskSuite) TearDownSuite() {

}

func (s *historyHistoryTransferTaskSuite) SetupTest() {
	s.Assertions = require.New(s.T())
}

func (s *historyHistoryTransferTaskSuite) TearDownTest() {

}

func (s *historyHistoryTransferTaskSuite) TestInsert_Single_Success() {
	shardID := rand.Int31()
	taskID := int64(1)

	task := s.newRandomTransferTaskRow(shardID, taskID)
	result, err := s.store.InsertIntoTransferTasks(newExecutionContext(), []sqlplugin.TransferTasksRow{task})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))
}

func (s *historyHistoryTransferTaskSuite) TestInsert_Multiple_Success() {
	shardID := rand.Int31()
	taskID := int64(1)

	task1 := s.newRandomTransferTaskRow(shardID, taskID)
	taskID++
	task2 := s.newRandomTransferTaskRow(shardID, taskID)
	result, err := s.store.InsertIntoTransferTasks(newExecutionContext(), []sqlplugin.TransferTasksRow{task1, task2})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(2, int(rowsAffected))
}

func (s *historyHistoryTransferTaskSuite) TestInsert_Single_Fail_Duplicate() {
	shardID := rand.Int31()
	taskID := int64(1)

	task := s.newRandomTransferTaskRow(shardID, taskID)
	result, err := s.store.InsertIntoTransferTasks(newExecutionContext(), []sqlplugin.TransferTasksRow{task})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))

	task = s.newRandomTransferTaskRow(shardID, taskID)
	_, err = s.store.InsertIntoTransferTasks(newExecutionContext(), []sqlplugin.TransferTasksRow{task})
	s.Error(err) // TODO persistence layer should do proper error translation
}

func (s *historyHistoryTransferTaskSuite) TestInsert_Multiple_Fail_Duplicate() {
	shardID := rand.Int31()
	taskID := int64(1)

	task1 := s.newRandomTransferTaskRow(shardID, taskID)
	taskID++
	task2 := s.newRandomTransferTaskRow(shardID, taskID)
	result, err := s.store.InsertIntoTransferTasks(newExecutionContext(), []sqlplugin.TransferTasksRow{task1, task2})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(2, int(rowsAffected))

	task2 = s.newRandomTransferTaskRow(shardID, taskID)
	taskID++
	task3 := s.newRandomTransferTaskRow(shardID, taskID)
	_, err = s.store.InsertIntoTransferTasks(newExecutionContext(), []sqlplugin.TransferTasksRow{task2, task3})
	s.Error(err) // TODO persistence layer should do proper error translation
}

func (s *historyHistoryTransferTaskSuite) TestInsertSelect_Single() {
	shardID := rand.Int31()
	taskID := int64(1)

	task := s.newRandomTransferTaskRow(shardID, taskID)
	result, err := s.store.InsertIntoTransferTasks(newExecutionContext(), []sqlplugin.TransferTasksRow{task})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))

	rangeFilter := sqlplugin.TransferTasksRangeFilter{
		ShardID:            shardID,
		InclusiveMinTaskID: taskID,
		ExclusiveMaxTaskID: taskID + 1,
		PageSize:           1,
	}
	rows, err := s.store.RangeSelectFromTransferTasks(newExecutionContext(), rangeFilter)
	s.NoError(err)
	for index := range rows {
		rows[index].ShardID = shardID
	}
	s.Equal([]sqlplugin.TransferTasksRow{task}, rows)
}

func (s *historyHistoryTransferTaskSuite) TestInsertSelect_Multiple() {
	numTasks := 20

	shardID := rand.Int31()
	minTaskID := int64(1)
	taskID := minTaskID
	maxTaskID := taskID + int64(numTasks)

	var tasks []sqlplugin.TransferTasksRow
	for i := 0; i < numTasks; i++ {
		task := s.newRandomTransferTaskRow(shardID, taskID)
		taskID++
		tasks = append(tasks, task)
	}
	result, err := s.store.InsertIntoTransferTasks(newExecutionContext(), tasks)
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(numTasks, int(rowsAffected))

	for _, pageSize := range []int{numTasks / 2, numTasks * 2} {
		filter := sqlplugin.TransferTasksRangeFilter{
			ShardID:            shardID,
			InclusiveMinTaskID: minTaskID,
			ExclusiveMaxTaskID: maxTaskID,
			PageSize:           pageSize,
		}
		rows, err := s.store.RangeSelectFromTransferTasks(newExecutionContext(), filter)
		s.NoError(err)
		s.NotEmpty(rows)
		s.True(len(rows) <= filter.PageSize)
		for index := range rows {
			rows[index].ShardID = shardID
		}
		s.Equal(tasks[:min(numTasks, pageSize)], rows)
	}
}

func (s *historyHistoryTransferTaskSuite) TestDeleteSelect_Single() {
	shardID := rand.Int31()
	taskID := int64(1)

	filter := sqlplugin.TransferTasksFilter{
		ShardID: shardID,
		TaskID:  taskID,
	}
	result, err := s.store.DeleteFromTransferTasks(newExecutionContext(), filter)
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(0, int(rowsAffected))

	rangeFilter := sqlplugin.TransferTasksRangeFilter{
		ShardID:            shardID,
		InclusiveMinTaskID: taskID,
		ExclusiveMaxTaskID: taskID + 1,
		PageSize:           1,
	}
	rows, err := s.store.RangeSelectFromTransferTasks(newExecutionContext(), rangeFilter)
	s.NoError(err)
	for index := range rows {
		rows[index].ShardID = shardID
	}
	s.Equal([]sqlplugin.TransferTasksRow(nil), rows)
}

func (s *historyHistoryTransferTaskSuite) TestDeleteSelect_Multiple() {
	shardID := rand.Int31()
	minTaskID := int64(1)
	maxTaskID := int64(101)

	filter := sqlplugin.TransferTasksRangeFilter{
		ShardID:            shardID,
		InclusiveMinTaskID: minTaskID,
		ExclusiveMaxTaskID: maxTaskID,
		PageSize:           int(maxTaskID - minTaskID),
	}
	result, err := s.store.RangeDeleteFromTransferTasks(newExecutionContext(), filter)
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(0, int(rowsAffected))

	rows, err := s.store.RangeSelectFromTransferTasks(newExecutionContext(), filter)
	s.NoError(err)
	for index := range rows {
		rows[index].ShardID = shardID
	}
	s.Equal([]sqlplugin.TransferTasksRow(nil), rows)
}

func (s *historyHistoryTransferTaskSuite) TestInsertDeleteSelect_Single() {
	shardID := rand.Int31()
	taskID := int64(1)

	task := s.newRandomTransferTaskRow(shardID, taskID)
	result, err := s.store.InsertIntoTransferTasks(newExecutionContext(), []sqlplugin.TransferTasksRow{task})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))

	filter := sqlplugin.TransferTasksFilter{
		ShardID: shardID,
		TaskID:  taskID,
	}
	result, err = s.store.DeleteFromTransferTasks(newExecutionContext(), filter)
	s.NoError(err)
	rowsAffected, err = result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))

	rangeFilter := sqlplugin.TransferTasksRangeFilter{
		ShardID:            shardID,
		InclusiveMinTaskID: taskID,
		ExclusiveMaxTaskID: taskID + 1,
		PageSize:           1,
	}
	rows, err := s.store.RangeSelectFromTransferTasks(newExecutionContext(), rangeFilter)
	s.NoError(err)
	for index := range rows {
		rows[index].ShardID = shardID
	}
	s.Equal([]sqlplugin.TransferTasksRow(nil), rows)
}

func (s *historyHistoryTransferTaskSuite) TestInsertDeleteSelect_Multiple() {
	numTasks := 20

	shardID := rand.Int31()
	minTaskID := int64(1)
	taskID := minTaskID
	maxTaskID := taskID + int64(numTasks)

	var tasks []sqlplugin.TransferTasksRow
	for i := 0; i < numTasks; i++ {
		task := s.newRandomTransferTaskRow(shardID, taskID)
		taskID++
		tasks = append(tasks, task)
	}
	result, err := s.store.InsertIntoTransferTasks(newExecutionContext(), tasks)
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(numTasks, int(rowsAffected))

	filter := sqlplugin.TransferTasksRangeFilter{
		ShardID:            shardID,
		InclusiveMinTaskID: minTaskID,
		ExclusiveMaxTaskID: maxTaskID,
		PageSize:           int(maxTaskID - minTaskID),
	}
	result, err = s.store.RangeDeleteFromTransferTasks(newExecutionContext(), filter)
	s.NoError(err)
	rowsAffected, err = result.RowsAffected()
	s.NoError(err)
	s.Equal(numTasks, int(rowsAffected))

	rows, err := s.store.RangeSelectFromTransferTasks(newExecutionContext(), filter)
	s.NoError(err)
	for index := range rows {
		rows[index].ShardID = shardID
	}
	s.Equal([]sqlplugin.TransferTasksRow(nil), rows)
}

func (s *historyHistoryTransferTaskSuite) newRandomTransferTaskRow(
	shardID int32,
	taskID int64,
) sqlplugin.TransferTasksRow {
	return sqlplugin.TransferTasksRow{
		ShardID:      shardID,
		TaskID:       taskID,
		Data:         shuffle.Bytes(testHistoryTransferTaskData),
		DataEncoding: testHistoryTransferTaskEncoding,
	}
}
