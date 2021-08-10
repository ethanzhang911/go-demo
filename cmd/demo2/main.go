package main

import (
	"fmt"
	"net/http"
)

func handlerPing(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("收到ping")
	select {
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Fprintf(w, "%s", "context收到")
			fmt.Println("context收到")
		}
	}
	//fmt.Fprintf(w, "done")
	////body, _ := req.GetBody()
	//getbody, _ := req.GetBody()
	//body, _ := io.ReadAll(getbody)
	//fmt.Fprintf(w, "%s\n", body)
}
func main() {
	httpServer := &http.Server{
		Addr: "172.18.3.1:8080",
	}
	http.HandleFunc("/ping", handlerPing)
	httpServer.ListenAndServe()

}
