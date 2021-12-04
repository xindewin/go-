package main

import (
	"fmt"
)

func re(a *int) {
    var sum int = 0
	var y int
	for i:=0.0;*a>0;i++{
		y=*a%10
		*a=*a/10
		sum=sum*10+y
	}
	*a=sum
}

func restring()  {
	
}

func main()  {
	var b int
   fmt.Scan(&b)
	var c *int
	c=&b
	re(c)
	fmt.Println(b)
}
