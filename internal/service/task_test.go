package service

import (
	"testing"
	"time"

	"github.com/Boh1mean/workmateTask/internal/model"
)

func TestNewTask(t *testing.T) {

	store := NewMemoryStorage()
	service := NewTaskService(store)

	task := service.NewTask()

	if task == nil {
		t.Fatal("expected non-nil task")
	}

	if task.ID == "" {
		t.Error("expected non-empty task ID")
	}

	if task.Status != "pending" {
		t.Errorf("expected status 'pending', got '%s'", task.Status)
	}

	if task.Duration < 180*time.Second || task.Duration > 300*time.Second {
		t.Errorf("expected duration between 180s and 300s, got %v", task.Duration)
	}

	if time.Since(task.CreatedAt) > time.Second {
		t.Errorf("CreatedAt is too far in the past: %v", task.CreatedAt)
	}

	storedTask, ok := store.GetTask(task.ID)
	if !ok {
		t.Errorf("task with ID %s was not stored", task.ID)
	}
	if storedTask.ID != task.ID {
		t.Errorf("stored task ID mismatch: got %s, want %s", storedTask.ID, task.ID)
	}

}

func TestGetTask(t *testing.T) {
	store := NewMemoryStorage()
	taskService := NewTaskService(store)

	task := &model.Task{
		ID:        "test-id",
		Status:    "pending",
		Duration:  2 * time.Second,
		CreatedAt: time.Now(),
	}

	if err := store.CreateTask(task); err != nil {
		t.Fatalf("failed to create task: %v", err)
	}

	gotTask, ok := taskService.GetTask("test-id")
	if !ok {
		t.Errorf("expected task to be found, got ok=false")
	}
	if gotTask.ID != task.ID {
		t.Errorf("expected ID %s, got %s", task.ID, gotTask.ID)
	}

	gotTask, ok = taskService.GetTask("non-existent-id")
	if ok {
		t.Errorf("expected task to not be found, got ok=true")
	}
	if gotTask != nil {
		t.Errorf("expected nil task, got %v", gotTask)
	}
}
