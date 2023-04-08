package main
// import ("time")
import "fmt"

//interfaces
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main() {
	dog := Dog {
		Name: "Samson",
		Breed: "German Shepherd",
	}
	// gorilla := Gorilla {
	// 	Name: "Peter",
	// 	Color: "Black",
	// 	NumberOfTeeth: 29,
	// }
	PrintInfo(&dog)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), " and has ", a.NumberOfLegs(), " legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

//looping
// func main(){
// 	var firstLine = "Once upon a time"
// 	for i, l := range firstLine {
// 		fmt.Println(i, l)
// 	}
// 	// animals := make(map[string]string)
// 	// animals["dog"] = "Jack"
// 	// animals["cat"] = "Jill"
// 	// for animalType, animal := range animals {
// 	// 	fmt.Println(animalType, animal)
// 	// }
// 	// animals := []string{"dog", "horse", "cat", "pig", "antelope"}
// 	// for _, animal := range animals {
// 	// 	fmt.Println(animal)
// 	// }
// }

//decision struct
// func main(){
// 	 myNum := 16
// 	 if myNum >= 16{
// 		myNum = myNum + 5
// 		fmt.Println(myNum)
// 	 } else {
// 		fmt.Println("Can't do anything with this")
// 	 }
// 	// isCat := "dog"
// 	// if isCat == "dog"{
// 	// 	fmt.Println("This is not a cat")
// 	// } else {
// 	// 	fmt.Println("This is a cat")
// 	// }
// 	// var isTrue bool
// 	// isTrue = true
// 	// if isTrue {
// 	// 	fmt.Println("It is true")
// 	// }
// }

//slices

// func main(){
// 	var mySlice []string
// 	mySlice = append(mySlice, "Otonye", "Peace", "Cynthia", "James","Sear")
// 	fmt.Println(mySlice[0:2])
// }

//map

// type User struct {
// 	FirstName string
// }

// func main() {
// 	myMap := make(map[string]User)
// 	me := User {
// 		FirstName: "Otonye",
// 	}
// 	myMap["me"] = me
// 	fmt.Println(myMap["me"].FirstName)
// }

// type myStruct struct {
// 	FirstName string
// }

// func (m *myStruct) printFirstName() string {
// 	return m.FirstName
// }

// func main() {
// 	var myVar myStruct
// 	myVar.FirstName = "John"
// 	fmt.Println(myVar.printFirstName())
// }


// type User struct {
// 	FirstName string
// 	LastName string
// 	PhoneNumber string
// 	Age string
// 	BirthDate time.Time
// }

// func main(){
// 	user := User {
// 		FirstName: "Otonye",
// 		LastName: "Amietubodie",
// 		PhoneNumber: "080020202",
// 		Age: "27",
		
// 	}
// 	fmt.Println(user)
// }

// func main(){
// 	var s2 string = "Six"
// 	fmt.Println(s)
// 	fmt.Println(s2)
// }

// func saySomething(s string) (string, string){
// 	return s, "DOg"
// }

//pointers

// func main(){
// 	var color string = "Green"
// 	fmt.Println(color)
// 	pointerpointer(&color)
// 	fmt.Println(color)
// }

// func pointerpointer(s *string){
// 	newvalue := "Red"
// 	*s = newvalue
// }