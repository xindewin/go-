package main

import "fmt"

type animal struct {
	height float32
	weight float32


}

func (a animal)eat()  {
	fmt.Println("eat")
}

type dog struct {
	animal

}

func (d dog)jump()  {

}
func main()  {
	a := animal{}
	a.eat()
	d := dog{}
	d.weight = 8
	d.eat()
}