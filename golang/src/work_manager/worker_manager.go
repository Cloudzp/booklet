package work_manager

import (
	"log"
)

type Element struct {
}

type WorkerManager struct {
}

func (m *WorkerManager) Start() {

}

func (m *WorkerManager) Stop() {

}

type worker struct {
	name      string
	stop      chan struct{}
	runWorker func()
	intervals int
}

func (w *worker) start() {
	for w.process() {
	}
}

func (w *worker) process() bool {
	defer HandleCrash()

	select {

	case <-w.stop:
		log.Printf("worker[%s] be stopped! ", w.name)
		return false
	default:
		w.runWorker()
	}

	return true
}
