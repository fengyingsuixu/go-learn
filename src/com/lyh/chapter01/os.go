package chapter01

import (
	"fmt"
	"os"
	"strings"
)


func OsParam(){
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func OsParam1(){
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func OsParam2(){

	fmt.Println(os.Args[1:])
}

func OsParam3(){

	fmt.Println(strings.Join(os.Args[1:]," "))
}