package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerPing(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("收到ping")
	select {
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Fprintf(w, "%s", "context收到")
			fmt.Println(err)
		} else {
			fmt.Fprintf(w, "%s", "context收到")
			fmt.Println("context收到")
		}
		return
	case <-time.After(6 * time.Second):
		fmt.Fprintf(w, "%s", "开始干活了")
		return
	}
}
func main() {
	httpServer := &http.Server{
		Addr: "172.18.3.1:8080",
	}
	http.HandleFunc("/ping", handlerPing)
	httpServer.ListenAndServe()

}
