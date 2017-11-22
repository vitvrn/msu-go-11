package main

import (
	"fmt"
	//	"sort" //vit
)

type MyStruct struct {
	Num  int
	Name string
}

type MyInt int

//vit:
//type MyMyInt MyInt

type withFiles bool

func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	//vit:
	//	*m = *m + MyInt(i)
	*m = *m + i
}

type mySliceStruct []MyStruct

//vit:
//func (m mySliceStruct) Less(i int) bool
//func (m mySliceStruct) Len() int
//func (m mySliceStruct) Swap(i, j int)
func (m mySliceStruct) Less(i, j int) bool { return m[i].Num < m[j].Num }
func (m mySliceStruct) Len() int           { return len(m) }
func (m mySliceStruct) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func main() {
	i := MyInt(0)

	i.add(3)
	i.showYourSelf()
}

func sorter(sl []MyStruct) []MyStruct {
	//vit:
	//	sort.Slice(sl, func(i, j int) {})
	return sl
}
