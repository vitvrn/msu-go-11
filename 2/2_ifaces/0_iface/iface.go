package main

import "fmt"

type Flyer interface {
	Fly()
	Greet()
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name + " is flying")
}

func (b Bird) Greet() {
	fmt.Println("Hey there")
}

func DoFly(f Flyer) {
	f.Greet()
	f.Fly()
}

type Mig45 struct{}

func (m Mig45) Fly() {
	fmt.Println("Mig Flied away")
}

//vit:
func (m Mig45) Greet() {}

func main() {
	duckPlane := Bird{"Duck plane"}

	GoFly(duckPlane)

	//vit:
	mig := Mig45{}
	GoFly(mig)
}

func GoFly(f Flyer) {
	f.Fly()
	//b := f.(Bird)

	if b, ok := f.(Bird); ok {
		fmt.Println(b.Name)
	}
}
