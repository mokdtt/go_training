package display

type T struct {
	s string
	n int
}

func Example_mapStruct() {
	m := make(map[T]int)
	key := T{"a", 1}
	m[key] = 0
	Display("mapStruct", m)
	// Output:
	//Display mapStruct (map[display.T]int):
	//mapStruct[{s:"a",n:1}] = 0
}

func Example_mapArray() {
	m := make(map[[3]string]int)
	key := [3]string{"a", "b", "c"}
	m[key] = 0
	Display("mapArray", m)
	// Output:
	//Display mapArray (map[[3]string]int):
	//mapArray[["a","b","c"]] = 0
}
