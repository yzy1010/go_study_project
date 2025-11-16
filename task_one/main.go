package main

import (
	"fmt"
	"math"
)

// Shape 接口，包含 Area 和 Perimeter 方法
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 实现 Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 实现 Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Employee 组合 Person，并添加 EmployeeID
type Employee struct {
	Person
	EmployeeID string
}

// PrintInfo 输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %s\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	// 演示 Shape 接口
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 2.5}

	shapes := []Shape{r, c}
	for _, s := range shapes {
		fmt.Printf("%T -> Area: %.2f, Perimeter: %.2f\n", s, s.Area(), s.Perimeter())
	}

	// 演示组合的 Person/Employee
	emp := Employee{Person: Person{Name: "Alice", Age: 30}, EmployeeID: "E12345"}
	emp.PrintInfo()
}
