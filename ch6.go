package main

import "fmt"

func main() {

    fmt.Println("Question 1:")
    var numbers [4]int
    numbers[3] = 2
    fmt.Printf("%d\n", numbers[3])

    fmt.Println("Question 2:")
    x := make([]int, 3, 9) // 3 is length of slice, 9 is underlying array length
    fmt.Printf("%d\n", len(x))

    fmt.Println("Question 3:")
    y := [6]string{"a","b","c","d","e","f"}
    fmt.Printf("%v\n", y[2:5])

    fmt.Println("Question 4:")
    z := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }
    var smallest int = 1000;
    for i := 0; i < len(z); i++ {
        if smallest > z[i] {
            smallest = z[i]
        }
    }
    fmt.Printf("%d\n", smallest)
    
}
