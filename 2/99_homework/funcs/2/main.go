package main

func showMeTheType(i interface{}) string {
	//	return ""

	//	if _, ok := i.(int); ok {
	//		return "int"
	//	} else if _, ok := i.(uint); ok {
	//		return "uint"
	//	} else if _, ok := i.(int8); ok {
	//		return "int8"
	//	} else if _, ok := i.(float64); ok {
	//		return "float64"
	//	} else if _, ok := i.(string); ok {
	//		return "string"
	//	} else if _, ok := i.(int32); ok {
	//		return "int32"
	//	} else if _, ok := i.([]int); ok {
	//		return "[]int"
	//	} else if _, ok := i.(map[string]bool); ok {
	//		return "map[string]bool"
	//	}

	switch i.(type) {
	case int:
		return "int"
	case uint:
		return "uint"
	case int8:
		return "int8"
	case float64:
		return "float64"
	case string:
		return "string"
	case int32:
		return "int32"
	case []int:
		return "[]int"
	case map[string]bool:
		return "map[string]bool"
	default:
		return ""
	}

}

func main() {
}
