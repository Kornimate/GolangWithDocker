package main

import "fmt"

func main() {

	var testString string = "Hello World!"

	var testFloat = 0.12

	var testArray = [3]string{"Hello", "World", "!"}

	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	testInt := 12

	var person User

	var testMap = map[string]int{"Mate": 10, "Petra": 20, "Gabor": 30, "Judit": 40}

	referenceMap := testMap //maps are references

	fastMap := map[int]int{}

	var fastMap2 map[int]int

	makeMap := make(map[int]int)

	person.id = 10
	person.name = "Mate"
	person.password = "nopass"

	fmt.Println(testString)
	fmt.Println(testFloat)
	fmt.Println(testInt)
	fmt.Println(testArray)
	fmt.Println(testSlice)
	fmt.Println(cap(testSlice))
	fmt.Println(len(testSlice))

	testSlice = append(testSlice, 1, 2, 3, 4)

	sliceFromArray := testArray[:len(testArray)-2]

	sliceWithMake := make([]string, len(sliceFromArray))

	copy(sliceWithMake, sliceFromArray)

	fmt.Println(testSlice)
	fmt.Println(sliceFromArray)
	fmt.Println(sliceWithMake)

	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}

		fmt.Printf("%d:%d\n", i, i)

		if i == 8 {
			break
		}
	}

	for index, value := range testArray { // - stands for throwing the vriable value
		fmt.Printf("%d : %v\n", index, value)
	}

	TestFunction("test", 30)

	fmt.Println(ReturnValue(100))
	fmt.Println(ReturnNamedValue(100))

	newName, newInt := MultipleReturns("Mate", 10)

	fmt.Println(newName)
	fmt.Println(newInt)
	fmt.Println(Factorial(5))

	fmt.Printf("(%d) %v %v\n", person.id, person.name, person.password)

	fmt.Println(testMap)

	fastMap[3] = 4

	fmt.Println(fastMap)
	fmt.Println(fastMap2)
	fmt.Println(makeMap)

	delete(testMap, "Mate")

	fmt.Println(testMap)

	val, exists := testMap["notexisting"]

	fmt.Printf("%v %v\n", val, exists)

	delete(referenceMap, "Petra") //maps are references

	fmt.Println(testMap)

	for k, v := range testMap {
		fmt.Printf("%v : %v\n", k, v)
	}
}

func TestFunction(text string, numOfIterations int) {
	for i := 0; i < numOfIterations; i += 2 {
		fmt.Printf("%d : %v\n", i, text)
	}

	fmt.Println()
}

func ReturnValue(num int) int {
	return num * 10
}

func ReturnNamedValue(num int) (result int) {
	result = num * 10
	return
}

func MultipleReturns(name string, occurrence int) (resString string, resInt int) {
	for i := 0; i < occurrence; i++ {
		resString += name
		resString += " "
	}

	resInt = 10 * occurrence

	return
}

func Factorial(num int) (result int) {
	if num <= 1 {
		result = 1

		return
	}

	result = num * Factorial(num-1)

	return
}

type User struct {
	id       int64
	name     string
	password string
}
