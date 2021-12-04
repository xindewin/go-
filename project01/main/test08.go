package main

import "fmt"

func main() {

	ages := map[string]int{
		"xxx":123 ,
		"yyy":234234 ,
		"zczx":21323 ,
	}

	for i,s:=range ages{
		fmt.Printf("[%q]=>[%d]\n",i,s)
	}


}
