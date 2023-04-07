package main
// import ("time")
import "fmt"

//slices

func main(){
	var mySlice []string
	mySlice = append(mySlice, "Otonye", "Peace", "Cynthia", "James","Sear")
	fmt.Println(mySlice[0:2])
}

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