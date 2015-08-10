package main

import "fmt"


func fib(n int) int {
    if n < 0 {
        panic("n must be >= 0")
    } else if n == 0 {
        return 0
    } else if (n == 1) {
        return 1
    } else {
        return fib(n - 1) + fib(n - 2)
    }
}

func main() {

    // Question 1
    fmt.Println("Question 1:")
    sum := func(numbers []int) int {
        total := 0
        for _, val := range numbers {
            total += val
        }
        return total
    }
    fmt.Printf("%d\n\n", sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 1000}))

    // Question 2
    fmt.Println("Question 2:")
    halfParity := func(someint int) (int, bool) {
        half := someint / 2
        parity := someint % 2 == 0
        return half, parity
    }
    numbers := []int{23, 45, 66, 1, -100}
    for _, val := range numbers {
        h, p := halfParity(val)
        fmt.Printf("Half:Parity of %d is %d:%t\n", val, h, p)
    }
    fmt.Println("")

    // Question 3
    fmt.Println("Question 3:")
    max := func(numbers ...int) int {
        var max int;
        for i, v := range numbers {
            if i == 0 || max < v {
                max = v
            }
        }
        return max
    }
    numbers2 := []int{23, 45, 66, 1, -100}
    fmt.Printf("Max of %v is %d\n\n", numbers2, max(numbers2...))

    // Question 4
    fmt.Println("Question 4:")
    makeOddGenerator := func() func() uint {
        i := uint(1)
        return func() (ret uint) {
            ret = i
            i += 2
            return i
        }
    }
    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd()) // 1
    fmt.Println(nextOdd()) // 3
    fmt.Println(nextOdd()) // 5
    fmt.Println("")

    // Question 5
    fmt.Println("Question 5:")
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    for i := 0; i <= 30; i++ {
        fmt.Printf("%d", fib(i))
        if i < 30 {
            fmt.Printf(", ")
        }
    }
    fmt.Println("")
    fmt.Println(fib(-100))

}
