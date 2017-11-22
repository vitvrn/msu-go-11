package main

import "fmt"
import "strconv"
import "sort"
//import "main_test"

func ReturnInt() int {
    return 1
}

func ReturnFloat() float32 {
    return 1.1
}

func ReturnIntArray() [3]int {
    return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
    return []int{1, 2, 3}
}

func IntSliceToString(incoming []int) string {
//    return string(incoming)
    var result string
    for _, elem := range incoming {
        fmt.Println(elem)
//        result += string(elem)
        result += strconv.FormatInt(int64(elem), 10)
    }
    return result
}

func MergeSlices(floatslice []float32, intslice []int32) []int {
//    var result []int
    result := make([]int, 0, len(floatslice) + len(intslice))
    for _, elem := range floatslice {
        result = append(result, int(elem))
    }
    for _, elem := range intslice {
        result = append(result, int(elem))
    }
    return result
}

func GetMapValuesSortedByKey(themap map[int]string) []string {
    result := make([]string, len(themap))
//    for key, value := range themap {
//        result[key-1] = value
//    }
    i := 0
    keys := make([]int, len(themap))
    for key := range themap {
        keys[i] = key
        i++
    }
    sort.Ints(keys)
    for j, key := range keys {
        result[j] = themap[key]
    }
    return result
}

func main() {
    println("Hello World!")
//    fmt.Println(IntSliceToString([]int{17, 23, 100500}))
//    fmt.Println(string(17))
    var input = map[int]string {
    9:  "Сентябрь",
    1:  "Январь",
    2:  "Февраль",
    10: "Октябрь",
    5:  "Май",
    7:  "Июль",
    8:  "Август",
    12: "Декарь",
    3:  "Март",
    6:  "Июнь",
    4:  "Апрель",
    11: "Ноябрь",
    }
    println("input len: ", len(input))
//    var keys = []int {2, 1, 3}
//    for _, key := range keys {
//        println(key)
//    }
    fmt.Println(GetMapValuesSortedByKey(input))
}