package main

import (
   "fmt"
)

// Embedded struct
type Person struct {
   Name string
   Age  int
}

// Embedding struct
type Employee struct {
   Person      // Embedding the Person struct
   CompanyName string
}

func main() {
   // Creating an Employee object
   emp := Employee{
      Person: Person{
         Name: "John Doe",
         Age:  35,
      },
      CompanyName: "ACME Corp",
   }

   // Accessing the embedded Person object
   fmt.Println("Employee Name:", emp.Name)
   fmt.Println("Employee Age:", emp.Age)

   // Accessing the fields of the embedding object
   fmt.Println("Employee Company:", emp.CompanyName)
}