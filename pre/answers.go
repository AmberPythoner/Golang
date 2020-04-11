package main

//@todo exit status 1 hello,2s，在win下，子协程里sleep10秒还是，还是能输出hello
import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

//将子进程的输出输出给结构体，供父进程查看
type result struct {
	err    error
	output []byte
}

/*要完成的任务：
一个任务运行太久，怎样kill掉
1. 在一个协程里执行一个command，sleep4;echo hello
2. 在main方法里,1秒的时候，kill 上面的command

*/
func main() {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		//定义channel 供两个协程传递
		resultChan chan *result
		res        *result
	)
	resultChan = make(chan *result, 1000)

	/*上下文,左边开头的上下文，继承了右边括号里的上下文,
	context里有一个channel byte，channel会放在context对象中，cancelFunc会把这个channel关掉，而子进程的返回值是一个cmd指针，而通过cmdContext构造的对象,实际内部也维护了context对象，且通过select语法监听ctx的Done(),也就是监听这个channel是否被关闭，也就说主进程一旦调用了cancelFunc，这个channel就会被关闭，会被CommandContext返回的cmd监听到（通过select），然后就执行kill 子进程pid杀掉
	*/
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			output []byte
			err    error
		)
		//select {case <-ctx.Done()}
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 10; echo hello,2s")
		//执行捕获输出，将结果通过channel传递给main协程
		output, err = cmd.CombinedOutput()
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	//一秒后取消上下文，子协程没有时间输出 hello，2s，会被直接kill
	time.Sleep(1 * time.Second)
	cancelFunc()
	//在main协程里，等待子协程的退出，并打印任务执行结果

	res = <-resultChan

	fmt.Println(res.err, string(res.output))

	fmt.Println(runtime.NumCPU())
}
