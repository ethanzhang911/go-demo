package main

import (
	"errors"
	"fmt"
	"log"
)

func division(a int, b int) (quotient int, remainder int, err error) {
	if b <= 0 {
		err = errors.New("除数不能为0或负数")
		return
	}
	quotient = a / b
	remainder = a % b
	err = nil
	return
}

func grade(a int) string {
	if a < 60 {
		return "C"
	} else if a < 90 {
		return "B"
	} else {
		return "A"
	}
}

func grade2(a int) string {
	switch {
	case a < 60:
		return "C"
	case a > 90:
		return "A"
	case a >= 60 && a <= 90:
		return "B"
	}
	return "未定义"
}

const (
	Sunday     = iota //0
	Monday            //1
	Tuesday           //2
	Wedenesday        //3
	Thursday          //4
	Friday            //5
	Saturday          //6
)

type say interface {
	say()
	change(name string)
	getname() string
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Printf("大家好我的名字叫%s\n", p.name)
}

func (p *person) change(name string) {
	p.name = name
}

func (p person) getname() string {
	return p.name
}

func main() {
	var a say
	a = &person{
		name: "张三",
	}
	//a.say()
	a.change("王五")
	fmt.Printf("我的名字叫%s\n", a.getname())

	map1 := make(map[int]string, 5)
	map1[1] = "张三"

	v, ok := map1[1]
	if ok {
		fmt.Printf("%s\n", v)
	} else {
		fmt.Println("在map中未找到")
	}

	fmt.Println("星期天=", Sunday)
	quotient, remainder, err := division(10, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("商为%d,余数为%d\n", quotient, remainder)
	fmt.Printf("考试等级为%s\n", grade(99))
	fmt.Printf("考试等级为%s\n", grade2(90))

	slice1 := []string{"A", "B", "C", "D", "E"}

	if func(a string, tobesearched []string) bool {
		for _, v := range tobesearched {
			if v == a {
				return true
			}
		}
		return false
	}("F", slice1) {
		fmt.Printf("%s", "已查找到")
	} else {
		fmt.Printf("%s", "未查找到")
	}

}
