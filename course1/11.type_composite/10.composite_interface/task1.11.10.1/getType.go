package main

func getType(i interface{}) string {
	switch i.(type) {
	default:
		return "Пустой интерфейс"
	case bool:
		return "bool"
	case int:
		return "int"
	case string:
		return "string"
	case []int:
		return "[]int"
	case []string:
		return "[]string"
	case float64:
		return "float"
	}
}
