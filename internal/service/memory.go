package service

import (
	"fmt"
	"sync"

	"github.com/Boh1mean/workmateTask/internal/model"
)

type TaskStorage interface {
	CreateTask(task *model.Task) error
	GetTask(id string) (*model.Task, bool)
	DeleteTask(id string) error
	SetStatus(id string, status string) bool
}

type MemoryStore struct {
	tasks map[string]*model.Task
	mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStore {
	return &MemoryStore{
		tasks: make(map[string]*model.Task),
	}
}

func (ms *MemoryStore) CreateTask(task *model.Task) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.tasks[task.ID] = task
	return nil
}

func (ms *MemoryStore) GetTask(id string) (*model.Task, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	task, ok := ms.tasks[id]
	return task, ok
}

func (ms *MemoryStore) DeleteTask(id string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if _, ok := ms.tasks[id]; !ok {
		return fmt.Errorf("task with id %s not found", id)
	}

	delete(ms.tasks, id)

	return nil
}

func (ms *MemoryStore) SetStatus(id string, status string) bool {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	task, ok := ms.tasks[id]
	if !ok {
		return false
	}

	task.Status = status
	return true
}
