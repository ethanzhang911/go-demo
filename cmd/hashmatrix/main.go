package main

import (
	"fmt"
	"reflect"
)

func main() {
	tmp := [9][9]int{}
	showmatrix(tmp)
}

func showmatrix(arg interface{}) {
	value := reflect.ValueOf(arg)
	switch value.Kind() {
	case reflect.Array:
		rowNum := value.Len()
		colNum := value.Index(0).Len()
		for i := 0; i < rowNum; i++ {
			for j := 0; j < colNum; j++ {
				fmt.Printf("%v  ", value.Index(i).Index(j))
			}
			fmt.Println()
		}

	}
}
