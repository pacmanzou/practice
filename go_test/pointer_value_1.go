package main

import "fmt"

type Person interface {
	getName() string
	getGrade() float64
	getGradeDiv()
	getGradeMul()
}

type Student struct {
	name  string
	grade float64
}

func (s Student) getName() string {
	return s.name
}
func (s Student) getGrade() float64 {
	return s.grade
}
func (s Student) getGradeDiv() {
	s.grade = s.grade / 10
}

func (s *Student) getGradeMul() {
	s.grade = s.grade * 10
}

func main() {

	var c Person
	// 当实现接口的结构体里面有指针方法时，不能把值对象赋值给该接口类型的变量
	// cannot use (Student literal) (value of type Student) as Person value in assignment: missing method getGradeMul (getGradeMul has pointer receiver)
	// c = Student{"merry", 70} // panic
	c = &Student{"merry", 70}
	fmt.Println(c.getGrade())
	c.getGradeDiv()
	fmt.Println(c.getGrade())
	c.getGradeMul()
	fmt.Println(c.getGrade())
	fmt.Println()

	fmt.Println("指针对象")
	a := &Student{"jarry", 82.5}
	fmt.Println("初始值", a.getGrade())
	a.getGradeDiv()
	fmt.Println("调用值方法div修改后", a.getGrade())
	fmt.Println("初始值", a.getGrade())
	a.getGradeMul()
	fmt.Println("调用指针方法mul修改后", a.getGrade())

	fmt.Println()
	fmt.Println("值对象")
	b := Student{"jack", 92.5}
	fmt.Println("初始值", b.getGrade())
	b.getGradeDiv()
	fmt.Println("调用值方法div修改后", b.getGrade())
	fmt.Println("初始值", b.getGrade())
	b.getGradeMul()
	fmt.Println("调用指针方法mul修改后", b.getGrade())
}
