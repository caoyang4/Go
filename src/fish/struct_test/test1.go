package main

import (
	"encoding/json"
	"fish/struct_test/model"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int
	Addr string
}

// 结构体方法
func (stu Student) spread() {
	fmt.Println(stu.Name, "is a good man")
}

func (stu Student) makeFriend(name string) {
	fmt.Println(stu.Name, "will make friend with", name)
}

func (stu Student) getDesc() string {
	return stu.Name + " go!"
}

func (stu Student) String() string {
	desc := fmt.Sprintf("{name:%v age:%v addr:%v}", stu.Name, stu.Age, stu.Addr)
	return desc
}

// 结构体是值类型
func main() {

	stu := make(map[string]Student, 5)
	stu["tom"] = Student{
		Name: "tom",
		Age:  18,
		Addr: "La",
	}
	fmt.Println(stu)

	s1 := Student{
		Name: "tom",
		Age:  18,
		Addr: "La",
	}

	s2 := s1 // 结构体默认值类型，默认值拷贝
	s2.Age = 16

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	var s3 *Student = new(Student)
	(*s3).Age = 10
	// go设计者为方便，底层进行优化
	s3.Addr = "gz"
	fmt.Println("s3:", *s3)

	var s4 *Student = &Student{}
	(*s4).Age = 6
	s4.Addr = "beijing"
	fmt.Println("s4:", *s4)

	fmt.Println()
	// Marshal 序列化底层使用反射
	jsonStu, _ := json.Marshal(s1)
	fmt.Println(string(jsonStu))

	fmt.Println()

	fmt.Println("结构体方法测试")
	s1.spread()
	s1.makeFriend("kobe")
	desc := s1.getDesc()
	fmt.Println(desc)

	james := model.NewPeople("james", "male", 38)
	fmt.Println(james)
}
