package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan(a chan engine.Request) {
	s.workerChan = a
}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
	go func() { s.workerChan <- r }()
}
