package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, canceled := context.WithTimeout(ctx, 3*time.Second)
	defer canceled()
	req, err := http.NewRequest(http.MethodGet, "http://58.33.40.242:10030/ping", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req = req.WithContext(ctx)
	time.Sleep(time.Second * 5)
	resq, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	resqBody, err := io.ReadAll(resq.Body)
	fmt.Printf("%s", resqBody)
	defer resq.Body.Close()

}
