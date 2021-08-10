package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cuDir, _ := os.Getwd()
	fmt.Printf("当前路径是%v\n", cuDir)
	s1 := bufio.NewScanner(strings.NewReader("my name is ethan, trust me!"))
	for s1.Scan() {
		fmt.Printf("%v\n", s1.Text())
	}
}
