package main
import "fmt"

func isPalindrome(x int) bool {
	m:=x
	var sum int
	var n int
	for ;x>0;{
		n=x%10
		x=x/10
		sum=sum*10+n
	}
	if sum==m{
		return true
	} else{
		return false
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("%v",isPalindrome(a))
}
