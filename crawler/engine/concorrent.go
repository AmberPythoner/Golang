package engine

import "fmt"

type ConcorrentEngine struct {
	Scheduler Scheduler
	WorkInt   int
}
type Scheduler interface {
	Submit(Request)
	WorkChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (e ConcorrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkInt; i++ {
		createWork(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for _, v := range seeds {
		// requests channel 里面继续分配新的任务
		e.Scheduler.Submit(v)
	}
	for {
		result := <-out
		for _, item := range result.Item {
			fmt.Println("get the item ", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWork(in chan Request, out chan ParseResult, scheduler Scheduler) {
	go func() {
		for {
			//in := make(chan Request)
			// 通过 scheduler 里面的select 使得in channel 里面有requests任务就去 执行
			scheduler.WorkerReady(in)
			// 拿到任务执行
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
