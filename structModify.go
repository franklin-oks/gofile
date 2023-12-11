package main

import "fmt"
//  define the structure

type student struct {
   Name string
   Age int
}
// method to modify the struct
func (s *student) modifyStruct(updatedName string, updatedAge int){
	s.Name = updatedName
	s.Age = updatedAge
}

// method to access the struct without modifying

func(s student) accessStruct(){
	fmt.Printf("name: %s, age: %d\n", s.Name, s.Age)
}

func main(){
	
// create an instance of the struct
myInstance := student{Name: "obinna", Age: 24}

// initial struct values
fmt.Println("initial student details")
myInstance.accessStruct()

// to modify the struct using method
myInstance.modifyStruct("franklin", 26)

// printing the modify values
fmt.Println("updated details:")
myInstance.accessStruct()

}