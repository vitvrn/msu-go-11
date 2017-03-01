package main

import "fmt"

type Flyer interface {
	Fly()
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {

}

type Mig45 struct{}

func (m Mig45) Fly() {
	fmt.Println("Mig Flied away")
}

func main() {
	duckPlane := Bird{"Duck plane"}

	GoFly(duckPlane)
}

func GoFly(f Flyer) {
	f.Fly()
	if b, ok := f.(Bird); ok {
		fmt.Println(b.Name)
	}
}
