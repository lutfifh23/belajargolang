package main

import (
	"fmt"
	"reflect"
)

type Person struct {
}

func (p *Person) Walk() {
	fmt.Println("Saya Berjalan")
}

func (p *Person) Sum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	type VariableBuatanSaya string

	var a VariableBuatanSaya = "Aurellia"

	r := reflect.ValueOf(a)
	fmt.Println(r)
	fmt.Println(r.Type())
	fmt.Println(r.Kind())
	fmt.Printf("%T\n", a)

	var p *Person = &Person{}
	p.Walk()

	// rp => struct
	rp := reflect.ValueOf(p)
	fmt.Println(rp)

	// rpw => function / method
	// rpw := rp.MethodByName("Walk")
	// fmt.Println(rpw)

	// rpw.Call([]reflect.Value{})

	rpw := rp.MethodByName("Sum")
	fmt.Println(rpw)

	rpw.Call([]reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
}
