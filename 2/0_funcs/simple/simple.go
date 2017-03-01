package main

import "fmt"

func main() {
	showMeTheMoney()
}

func showMeTheMoney() {
	fmt.Printf("$$$$")
}

func sum(i int, j int) int {
	return i + j
}

func sumLight(i, j int) int {
	return i + j
}

func sumMore(stuff ...int) (res int) {
	for i := range stuff {
		res += stuff[i]
	}
	return
}

func sumOnlyNatural(stuff ...int) (int, error) {
	res := 0
	for i := range stuff {
		if stuff[i] < 0 {
			return 0, fmt.Errorf("Only natural numbers expected - given %d", stuff[i])
		}
		res += stuff[i]
	}
	return res, nil
}

func sumErrorStuff(i int, st string) {
	if i <= 0 {
		fmt.Println("Nope")
		return
	}

	switch st {
	case "pron":

	}

}
