package main

import "fmt"

// map in golang
func main() {
	var myMap = make(map[string]int)

	fmt.Println("my map = ", myMap)

	if myMap == nil {
		fmt.Println("myMap = ", myMap)
	} else {
		fmt.Println("khac nil")
	}
	fmt.Println()
	var mayMap1 map[string]int
	fmt.Println("in map1 = ", mayMap1)
	if mayMap1 == nil {
		fmt.Println("map1 = ", mayMap1)
	} else {
		fmt.Println("khac nil")
	}

	fmt.Println()

	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	fmt.Println("maymap 2 = ", myMap2)

	myMap2["key2"] = 20

	// reference type
	mymap3 := myMap2
	mymap3["key1"] = 1000
	fmt.Println("maymap 2 again before delete key = ", myMap2)
	fmt.Println("maymap 3 new = ", mymap3)

	delete(myMap2, "key1")
	fmt.Println("my map 2 againg = ", myMap2)

	fmt.Println()

	value, found := myMap2["key2"]
	if found {
		fmt.Println("value = ", value)
	} else {
		fmt.Println("key not found")
	}

}
