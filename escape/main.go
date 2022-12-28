package main

type Person struct {
	Name   string
	EnName string
}

func main() {
	badBoy := &Person{
		Name:   "法外狂徒张三",
		EnName: "ZhangSan",
	}
	_ = badBoy
}
