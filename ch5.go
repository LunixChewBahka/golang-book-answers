package main

import "fmt"

func main() {

    // Question 1
    fmt.Print("\nQuestion 1: ")
    i := 10
    if i > 10 {
        fmt.Println("Big")
    } else {
        fmt.Println("Small")
    }

    // Question 2
    fmt.Print("\nQuestion 2: ")
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 {
            fmt.Printf("%d", i)
            if i < 99 {
                fmt.Printf(", ")
            }
        }
    }

    // Question 3
    fmt.Print("\n\nQuestion 3: ")
    for i := 1; i <= 100; i++ {
        if i % 3 != 0 && i % 5 != 0 {
            fmt.Printf("%d", i)
        } else {
            if i % 3 == 0 {
                fmt.Printf("Fizz")
            }
            if i % 5 == 0 {
                fmt.Printf("Buzz")
            }
        }
        if i < 100 {
            fmt.Printf(", ")
        }
    }

    fmt.Println("")

}
