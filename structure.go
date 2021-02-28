package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {

    emp1 := Employee{
        firstName: "Nandhini",
        age:       24,
        salary:    30000,
        lastName:  "Murugaiyan",
    }

    emp2 := Employee{"Raja", "Lakshmi", 20, 15000}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}