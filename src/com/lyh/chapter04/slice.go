package chapter04

import (
	"fmt"
)

func PrintSlice(){
	months := [...]string{1: " January",2: " February",3: " March",
	4: " April",5: " May",6: " June",
	7: " July",8: " August",9: " September",
	10: " October",11: " November", 12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2,summer)

	for _,s := range Q2{
		for _,t := range summer{
			if s == t{
				fmt.Printf("%s appears in both\n",s)
			}
		}

	}
}

