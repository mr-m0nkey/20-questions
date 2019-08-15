package main

import "fmt"

type vector struct {
   x, y int
}

func main() {

   var firstName, lastName string;
   var age int;

   fmt.Println("What is your first name?")
   //TODO get first name

   fmt.Println("What is your last name?")
   //TODO get last name

   fmt.Println("How old are you?")
   //TODO get age


   //TODO print result
   fmt.Printf("Welcome, %v %v (%v)", firstName,lastName, age)
}
