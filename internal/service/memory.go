package service

import (
	"sync"

	"github.com/Boh1mean/workmateTask/internal/model"
)

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

func (ms *MemoryStore) DeleteTask(id string) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	delete(ms.tasks, id)
}
