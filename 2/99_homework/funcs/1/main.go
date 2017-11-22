package main

import "fmt"

type memoizeFunction func(int, ...int) interface{}

// TODO реализовать
//var fibonacci memoizeFunction
//var romanForDecimal memoizeFunction
var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

//TODO Write memoization function

func memoize(function memoizeFunction) memoizeFunction {
	//	var mem = make(map[int]int)
	var mem = make(map[int]interface{})
	a := func(param int, none ...int) interface{} {
		//закомментировать вывод для больших param???:
		fmt.Println(mem)
		//		var result interface{}
		if result, ok := mem[param]; ok {
			return result
		} else {
			//			result = function(param).(int) //?м.б. объявить result как interface{}
			result = function(param)
			mem[param] = result
			return result
		}
	}
	return a
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {

	fibonacci = func(p int, pp ...int) interface{} {
		//для начала обычная ф-ция Фибоначчи
		if p < 1 {
			panic(fmt.Errorf("Fibonacci is defined only for naturals"))
		}
		if p == 1 || p == 2 {
			return 1
		} else {
			return fibonacci(p-2).(int) + fibonacci(p-1).(int) //??? fib(p-2) нужно вычислять раньше fib(p-1)
		}
	}
	//теперь мемоизированная
	fibonacci = memoize(fibonacci)

	romanForDecimal = func(p int, pp ...int) interface{} {
		if p >= 4000 || p < 1 {
			panic(fmt.Errorf("Not allowed input: %v", p))
		}
		var result string
		//		for M := p / 1000; M > 0; M-- {
		//			result += "M"
		//		}

		//		if p > 1000
		MM := p / 1000
		for i := MM; i > 0; i-- { //assert(M<4) == true //p<4000
			result += "M"
		}
		p -= MM * 1000
		//---------------------v- 900 -v---------------------------
		if p >= 900 {
			result += "CM"
			p -= 900
		}
		if p >= 500 { //(DD==0 || DD==1) == true
			result += "D"
			p -= 500
		}
		//		if p > 100
		CC := p / 100
		//---------------------v- 400 -v---------------------------
		if CC == 4 {
			result += "CD"
			//p -= 400
		} else {
			for i := CC; i > 0; i-- {
				result += "C"
			}
			//p -= CC * 100
		}
		p -= CC * 100 //comment for optimize ???
		//---------------------- 90 ----------------------------
		if p >= 90 {
			result += "XC"
			p -= 90
		}
		if p >= 50 { //(LL==0 || LL==1) == true
			result += "L"
			p -= 50
		}
		XX := p / 10
		//---------------------- 40 ----------------------------
		if XX == 4 {
			result += "XL"
		} else {
			for i := XX; i > 0; i-- {
				result += "X"
			}
		}
		p -= XX * 10
		//---------------------- 9 ----------------------------
		if p == 9 {
			return result + "IX"
		}
		if p >= 5 {
			result += "V"
			p -= 5
		}
		//---------------------v- 4 -v---------------------------
		if p == 4 {
			return result + "IV"
		}
		for p > 0 {
			result += "I"
			p--
		}
		return result
	}
	romanForDecimal = memoize(romanForDecimal)

}

func main() {
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
	//	fmt.Println(fibonacci(40))
	//	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))

	//	for i := 1; i < 1000; i++ {
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
	//	}
}
