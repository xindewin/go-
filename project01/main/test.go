package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	l:=len(str)
	var slice []byte
	for i:=l-1;i>=0;i--{
		slice=append(slice,str[i])
	}
	fmt.Printf("%q",slice)


}


/*func main() {
	var str []int
	var i int
	for i=0;;i++{
		var a int
		fmt.Scan(&a)
		str = append(str, a)
		//fmt.Println(str[i])
		if str[i]==0{
			break
		}

	}
	fmt.Println(str[0])
	fmt.Println(str[1])
	fmt.Println(str[2])
	//str[2]='a'
	fmt.Println(str)

}*/
