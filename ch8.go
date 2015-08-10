package main

import "fmt"

func square(x *float64) {
    *x = *x * *x
}

func swap(x *int, y *int) {
    temp := *x
    *x = *y
    *y = temp
}

func main() {

    // Question 1: guess is 2.25
    fmt.Println("Question 1:")
    x := 1.5
    square(&x)
    fmt.Printf("%f\n", x)
    fmt.Println("")

    // Question 2
    fmt.Println("Question 2:")
    y := 1
    z := 2
    fmt.Printf("y was %d, z was %d\n", y, z)
    swap(&y, &z)
    fmt.Printf("y is now %d, z is now %d\n", y, z)

}
