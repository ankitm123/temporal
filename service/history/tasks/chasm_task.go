// The MIT License
//
// Copyright (c) 2025 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2025 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tasks

import (
	"time"

	enumsspb "go.temporal.io/server/api/enums/v1"
	persistencespb "go.temporal.io/server/api/persistence/v1"
	"go.temporal.io/server/common/definition"
)

// ChasmTaskPure is a task that applies a batch of state changes to a CHASM
// entity. All components within the entity's tree will have their pending state
// changes applied.
type ChasmTaskPure struct {
	definition.WorkflowKey // Task interface assumes WorkflowKey.
	VisibilityTimestamp    time.Time
	TaskID                 int64
	Category               Category // Set based on the task's queue.
}

var _ Task = &ChasmTaskPure{}

func (t *ChasmTaskPure) GetTaskID() int64 {
	return t.TaskID
}

func (t *ChasmTaskPure) SetTaskID(id int64) {
	t.TaskID = id
}

func (t *ChasmTaskPure) GetVisibilityTime() time.Time {
	return t.VisibilityTimestamp
}

func (t *ChasmTaskPure) SetVisibilityTime(timestamp time.Time) {
	t.VisibilityTimestamp = timestamp
}

func (t *ChasmTaskPure) GetCategory() Category {
	return t.Category
}

func (t *ChasmTaskPure) GetType() enumsspb.TaskType {
	return enumsspb.TASK_TYPE_CHASM_PURE
}

func (t *ChasmTaskPure) GetKey() Key {
	if t.GetCategory().Type() == CategoryTypeScheduled {
		return NewKey(t.VisibilityTimestamp, t.TaskID)
	}

	return NewImmediateKey(t.TaskID)
}

// ChasmTask is a task with side effects generated by a CHASM component.
type ChasmTask struct {
	definition.WorkflowKey // Task interface assumes WorkflowKey.
	VisibilityTimestamp    time.Time
	TaskID                 int64
	Category               Category // Set based on the task's queue.
	Destination            string   // Set for outbound tasks.
	Info                   *persistencespb.ChasmTaskInfo
}

var _ Task = &ChasmTask{}

func (t *ChasmTask) GetCategory() Category {
	return t.Category
}

func (t *ChasmTask) GetType() enumsspb.TaskType {
	return enumsspb.TASK_TYPE_CHASM
}

func (t *ChasmTask) GetKey() Key {
	if t.GetCategory().Type() == CategoryTypeScheduled {
		return NewKey(t.VisibilityTimestamp, t.TaskID)
	}

	return NewImmediateKey(t.TaskID)
}

func (t *ChasmTask) GetTaskID() int64 {
	return t.TaskID
}

func (t *ChasmTask) SetTaskID(id int64) {
	t.TaskID = id
}

func (t *ChasmTask) GetVisibilityTime() time.Time {
	return t.VisibilityTimestamp
}

func (t *ChasmTask) SetVisibilityTime(timestamp time.Time) {
	t.VisibilityTimestamp = timestamp
}
