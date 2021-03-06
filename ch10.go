package main

import (
    "fmt"
    "time"
)

func sleep(n int) {
    fmt.Printf("Sleeping for %d seconds...\n", n)
    wait := <- time.After(time.Millisecond * time.Duration(n))
    fmt.Printf("I am now awake (%v).\n", wait)
}

func main() {

    // Question 2:
    fmt.Println("Question 2:")
    sleep(1000)
    fmt.Println("")

    // Question 3:
    fmt.Println("Question 3:")
    c := make(chan int, 10)
    go func() {
        for i := 0; i < 20; i++ {
            c <- i
            time.Sleep(time.Millisecond * 10)
        }
    }()
    go func() {
        for i := 0; i < 5; i++ {
            time.Sleep(time.Millisecond * 2000)
            num := <- c
            fmt.Printf("Received %d\n", num)
            num = <- c
            fmt.Printf("Received %d\n", num)
        }
    }()
    var input string
    fmt.Scanln(&input)

}
