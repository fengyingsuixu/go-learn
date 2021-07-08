package chapter03

import (
	"fmt"
	"unicode/utf8"
)

func StringPrint(){
	var s string = "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println("goodbye" + s[5:])


	s1 := "hello, 世界"
	fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))

	for i := 0; i < len(s1); i++{
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n",i, r)
		i += size
	}

	for i ,r := range s1{
		fmt.Printf("%d\t%q\t%d\n",i, r,r)
	}
}
