package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	//logger, _ := zap.NewProduction()
	var (
		expr *cronexpr.Expression
		err  error
	)

	if expr, err = cronexpr.Parse("*/1 * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	now := time.Now()
	next_time := expr.Next(now)
	fmt.Printf("%s, %s", now, next_time)
	time.AfterFunc(next_time.Sub(now), func() {
		fmt.Println("被调度了:", next_time)
	})
	time.Sleep(1 * time.Minute)
}
