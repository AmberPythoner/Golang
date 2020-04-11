package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	output []byte
	err    error
}

func main() {
	resChan := make(chan *result, 10000)
	ctx, canFunc := context.WithCancel(context.TODO())

	go func() {
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;touch 3344.txt;")
		// 执行任务, 捕获输出
		output, err := cmd.CombinedOutput()

		// 把任务输出结果, 传给main协程
		resChan <- &result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)
	canFunc()
	res := <-resChan
	fmt.Println(res.err, string(res.output))
}
