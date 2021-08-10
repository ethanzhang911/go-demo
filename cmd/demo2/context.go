package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func subFunc(c context.Context) {
	defer sg.Done()
	select {
	case <-c.Done():
		if c.Err() != nil {
			fmt.Printf("%s", "收到")
		} else {
			fmt.Printf("%s", c.Err().Error())
		}
		return
	}
}

var sg sync.WaitGroup

func main() {
	ctx := context.Background()
	ctx, canceled := context.WithTimeout(ctx, 3*time.Second)
	//defer canceled()
	sg.Add(1)
	go subFunc(ctx)
	sg.Wait()
	canceled()
	canceled()
}
