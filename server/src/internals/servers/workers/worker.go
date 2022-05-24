package workers

import (
	"fmt"

	"github.com/boybulbousbemperor/go-saturn/src/internals/servers/tasks"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Queue   queue.Queue
	TaskMap map[uuid.UUID]tasks.Task
	Tasks   int
}

func (w *Worker) CollectStats() {
	fmt.Println("CollectStats()")
}

func (w *Worker) StartTask() {
	fmt.Println("StartTask()")
}

func (w *Worker) RunTask() {
	fmt.Println("RunTask()")
}

func (w *Worker) StopTask() {
	fmt.Println("StopTask()")
}
