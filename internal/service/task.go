package service

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Boh1mean/workmateTask/internal/model"
	"github.com/google/uuid"
)

type TaskUsecase interface {
	NewTask() *model.Task
	GetTask(id string) (*model.Task, bool)
	DeleteTask(id string) error
}

type TaskService struct {
	store TaskStorage
}

func NewTaskService(store TaskStorage) *TaskService {
	return &TaskService{
		store: store,
	}
}

func (s *TaskService) NewTask() *model.Task {
	id := uuid.New().String()
	duration := time.Duration(3*60+rand.Intn(2*60)) * time.Second

	task := &model.Task{
		ID:        id,
		Status:    "pending",
		Duration:  duration,
		CreatedAt: time.Now(),
	}

	if err := s.store.CreateTask(task); err != nil {
		log.Printf("failed to create task: %v", err)
		return nil
	}
	go func() {
		s.store.SetStatus(task.ID, "running")
		time.Sleep(duration)
		s.store.SetStatus(task.ID, "completed")
	}()

	return task
}

func (s *TaskService) GetTask(id string) (*model.Task, bool) {
	task, ok := s.store.GetTask(id)
	if !ok {
		return nil, false
	}

	return task, true
}

func (s *TaskService) DeleteTask(id string) error {
	if err := s.store.DeleteTask(id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
