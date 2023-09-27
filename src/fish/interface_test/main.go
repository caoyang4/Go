package main

import "fmt"

type Usb interface {
	// 声明两个方法
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("phone start")
}

func (p Phone) Stop() {
	fmt.Println("phone stop")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("camera start")
}

func (c Camera) Stop() {
	fmt.Println("camera stop")
}

type Computer struct {
	Name string
}

func (c *Computer) work(usb Usb) {
	usb.Start()
	fmt.Println(c.Name, "work")
	usb.Stop()
}

func main() {

	computer := Computer{"huawei"}
	phone := Phone{}
	camera := Camera{}
	computer.work(phone)
	fmt.Println()
	computer.work(camera)

}
