package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

var logger = log.New()

// 详细说明了关于context如何在多进程之间传输，看样子还是有问题拿不到server返回的值
//
func main() {
	ctx := context.Background()
	logger.Out = os.Stdout
	ctx, canceled := context.WithTimeout(ctx, 2*time.Second)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, "http://58.33.40.242:10030/ping", nil)
	log.WithFields(log.Fields{"time": time.Now()}).Infof("%s", "生成了NewRequest")
	if err != nil {
		log.Fatalln(err)
	}
	req = req.WithContext(ctx)
	log.WithFields(log.Fields{"time": time.Now()}).Infof("%s", "传入了context")
	// 这里是真正的执行，不过是异步的，它会阻塞。
	// 另外一点非常关键，这里问题出现在server端不返回,导致context deadline exceeded
	resq, err := client.Do(req)
	log.WithFields(log.Fields{"time": time.Now()}).Infof("%s", "执行完http.DefaultClient.Do(req)")

	if err != nil {
		log.Fatalln(err)
	}

	resqBody, err := io.ReadAll(resq.Body)
	fmt.Printf("%s", resqBody)
	defer resq.Body.Close()
	canceled()
	log.WithFields(log.Fields{"time": time.Now()}).Infof("%s", "调用了canceled")

}
