package managers

import (
	"fmt"

	"github.com/boybulbousbemperor/go-saturn/src/internals/servers/tasks"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending           queue.Queue
	TasksMap          map[string][]tasks.Task
	EventMap          map[string][]tasks.TaskedJobEvent
	Workers           []string
	WorkerIntoTaskMap map[string][]uuid.UUID
	TaskIntoWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("SelectWorker()")
}

func (m *Manager) SendWork() {
	fmt.Println("SendWork()")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("UpdateTasks()")
}
