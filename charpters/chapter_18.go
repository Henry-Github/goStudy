package main

import "fmt"

// interface

type Common interface {
	// interface 里面提供一些函数
	calculateArea() int
	calculatePerimeter() int
}

type CommonPtr interface {
	calculateArea() int
	calculatePerimeter() int
	disPlay()
}

type Rect struct {
	width  int
	length int
}

func (r Rect) calculateArea() int {
	return r.width * r.length
}

func (r Rect) calculatePerimeter() int {
	return 2 * (r.length + r.width)
}

type RectPtr struct {
	width  int
	length int
}

func (r *RectPtr) calculateArea() int {
	return r.width * r.length
}

func (r *RectPtr) calculatePerimeter() int {
	return 2 * (r.length + r.width)
}

func (r *RectPtr) disPlay() {
	fmt.Println(r.width, r.length)
}

func calculate(c Common) {
	fmt.Println(c.calculateArea())
	fmt.Println(c.calculatePerimeter())
}

func calculatePtr(c CommonPtr) {
	fmt.Println(c.calculateArea())
	fmt.Println(c.calculatePerimeter())
	c.disPlay()

}

func main() {
	r := Rect{2, 4}
	calculate(r)
	rptr := RectPtr{2, 3}
	calculatePtr(&rptr)
}
