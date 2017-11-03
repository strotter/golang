package main

type Test1 struct {
	data    []string
	numbers []int32
	data2   string
}

type Test2 struct {
	data    []*string
	numbers []int32
	data2   string
}

func Alloc1() []Test1 {
	a := make([]Test1, 100)
	return a
}

func Alloc2() []Test2 {
	a := make([]Test2, 100)
	return a
}

func main() {
}
