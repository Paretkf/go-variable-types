package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 1. Literal types
	// Note: Literal type copy by value
	var theSting2 string
	theSting2 = "This is String 2"
	theString := "This variable type string"
	theInt := 305
	theBool := true
	printString(theString)
	printString(theSting2)
	printInt(theInt)
	printBool(theBool)
	printUnderline()
	// 2. Type in Golang are static type
	// printString(theInt) // Compile error

	// 3. interface{} type can accept all types รับตัวแปรอะไรก็ได้
	var interF interface{}
	interF = theSting2

	// ถ้าจะรับค่า Interface กลับมาต้อง casting interface ก่อน
	string2 := interF.(string)
	printString(string2)

	printInterface(interF)
	printInterface(theString)
	printInterface(theInt)
	printInterface(theBool)
	printUnderline()

	// 4. map type
	// Note: Map variable is the pointer of map, pass by pointer of map
	// เหมือน Dictionary
	// เก็บ [Key]Values
	theMap := map[string]interface{}{}
	// [string] = key เป็น string
	// interface = เป็น Values ของ map
	// {} ตัวสุดท้ายให้ประกาศ instance ให้ด้วย
	// ประกาศอีกวิธี theMap := make(map[string]interface{})
	theMap["firstname"] = "Rachata"
	theMap["lastname"] = "Tongphakdi"
	theMap["citizen_id"] = "1234"
	theMap["no"] = 2345
	printMap(theMap)

	theMapInt := map[string]int{}
	theMapInt["num1"] = 1
	theMapInt["num2"] = 2
	printMapInt(theMapInt)
	// Map จะส่ง Pointer เข้าไป ที่ Function ด้วย ทำให้เมื่อมีการ update map จะให้ค่าเปลี่ยน
	updateMap(theMap) // This will update map that pass in to function
	printMapAsJSON(theMap)
	// Note: We can use 2 variables to get map value
	// ถ้ามี gender ใน theMap ok จะเป็น true
	gender, ok := theMap["gender"]
	if ok {
		fmt.Println("gender is ", gender.(string)) // Casting
	} else {
		fmt.Println("No Gender specify")
	}

	printUnderline()

	// 5. slice type (dynamic array)
	// ควรจะเติม s ด้วย slice = array
	// Note: We should name slice phurally
	// Note: Slices pass by pointer of slice (just like map)
	theSlices := []string{}
	// [] บอกว่าเเป็น array
	// string = type ของ slice
	// {} ให้ประกาศ instance รอไว้
	// theSlices := make([]string, 0)
	theSlices = append(theSlices, "item 1")
	// in JS array.push("item 1")
	// ex [js] theSlices.push("item 1") = [go] theSlices = append(theSlices, "item 1")
	theSlices = append(theSlices, "item 2")
	theSlices = append(theSlices, "item 3")
	printSlice(theSlices)
	// updateSliceIndex0(theSlices) // This will update value of index 0 of slice
	printSliceAsJSON(theSlices)
	printUnderline()

	// 6. Struct type
	// Note: It is my practice to always create struct variable as pointer (*Type)
	//   to avoid confusion in team member, so will can always assume that strut is a pointer
	theCitizen := &Citizen{
		Firstname: "Chaiyapong",
		Lastname:  "Lapliengtrakul",
		CitizenID: "1234",
	}

	theCitizen2 := &Citizen{
		CitizenID: "454",
	}

	printCitizen(theCitizen)
	printCitizen(theCitizen2)
	// Note: JSON marshal will use struct tag ส่วนมาใช้แบบนี้
	printCitizenAsJSON(theCitizen)
	printUnderline()

	// 7. nil value
	// Note: Literal cannot be nil (string, int, bool, ..)
	// theString = nil // Error
	// theInt = nil    // Error
	// theBool = nil   // Error

	// Note: map, slice and struct and be nil
	theMap = nil
	theSlices = nil
	theCitizen = nil
	// nil = null
	var theMapNil = map[string]interface{}{}
	printInterface("Nil map : ")
	printMap(theMapNil)

	// Note: assign value to nil map and struct will cause runtime error
	// theMap["key"] = "value" // Runtime Error
	// theCitizen.Firstname = "Chaiyapong" // Runtime Error

	// Note: append value to nil slice is OK
	theSlices = append(theSlices, "value") // OK

	// 8. Check nil or zero using len()
	if len(theMap) == 0 {
		fmt.Println("theMap is nil")
	}
	if len(theSlices) == 0 {
		fmt.Println("theSlices is nil")
	}
	if theCitizen == nil {
		fmt.Println("theCitizen is nil")
	}
	printUnderline()

	// 9. Enum is just constant in golang
	var theGender GenderType
	theGender = Male
	theGender = Female
	theGender = Unspecify
	switch theGender {
	case Male:
		fmt.Println("Gender is Male")
		break
	case Female:
		fmt.Println("Gender is Female")
		break
	case Unspecify:
		fmt.Println("Gender is Unspecify")
		break
	}
	printGender(theGender)
	printUnderline()
}

func printUnderline() {
	fmt.Println("---")
}

func printString(input string) {
	fmt.Println("string = ", input)
}

func printInt(input int) {
	fmt.Println("int = ", input)
}

func printBool(input bool) {
	fmt.Println("bool = ", input)
}

func printInterface(input interface{}) {
	fmt.Println("interface = ", input)
}

func printMap(input map[string]interface{}) {
	fmt.Println("map = ", input)
}

func printMapInt(input map[string]int) {
	fmt.Println("map int = ", input)
}

func updateMap(input map[string]interface{}) {
	input["updated"] = true
}

func printMapAsJSON(input map[string]interface{}) {
	// NOTE แปลง Map String
	js, _ := json.Marshal(input)
	// fmt.Println("map JSON Not String = ", js)
	fmt.Println("map JSON = ", string(js))
}

func printSlice(input []string) {
	fmt.Println("slice = ", input)
}

func updateSliceIndex0(input []string) {
	input[0] = "Item 0"
}

func printSliceAsJSON(input []string) {
	js, _ := json.Marshal(input)
	fmt.Println("slice JSON = ", string(js))
}

// Citizen is type represent person in country
// json: คือ struct tag
type Citizen struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	CitizenID string `json:"citizen_id"`
}

func printCitizen(input *Citizen) {
	fmt.Println("Citizen = ", input)
}

func printCitizenAsJSON(input *Citizen) {
	js, _ := json.Marshal(input)
	fmt.Println("Citizen JSON = ", string(js))
}

// GenderType is enum for Gender
type GenderType string

const (
	// Unspecify is gender type for Unspecify
	Unspecify GenderType = "UNSPECIFY"
	// Male is gender type for Male
	Male GenderType = "MALE"
	// Female is gender type for Female
	Female GenderType = "FEMALE"
)

func printGender(input GenderType) {
	fmt.Println("Gender = ", input)
}
