package model

type Video struct {
	Id   int64  `uri:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
