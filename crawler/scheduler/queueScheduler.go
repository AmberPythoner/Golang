package scheduler

import "crawler/engine"

type QueueScheduler struct {
	workerChan  chan chan engine.Request
	requestChan chan engine.Request
}

func (s *QueueScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

//func (s *QueueScheduler) WorkerChan(a chan engine.Request) {
//	s.workerChan <- a
//}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestsQueue []engine.Request
		var workQueue []chan engine.Request
		for {
			var activateRequests engine.Request
			var activateWork chan engine.Request
			if len(requestsQueue) > 0 && len(workQueue) > 0 {
				activateRequests = requestsQueue[0]
				activateWork = workQueue[0]
			}
			select {
			case r := <-s.requestChan:
				requestsQueue = append(requestsQueue, r)
			case w := <-s.workerChan:
				workQueue = append(workQueue, w)
			case activateWork <- activateRequests:
				requestsQueue = requestsQueue[1:]
				workQueue = workQueue[1:]
			}
		}
	}()
}
