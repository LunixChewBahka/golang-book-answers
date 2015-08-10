package main

import "fmt"
import "math"

/**
 * Line struct, used mainly by rectangle but not a Shape
 */
type Line struct {
    x1, y1, x2, y2 float64
}
func (ln *Line) distance() float64 {
    a := ln.x2 - ln.x1
    b := ln.y2 - ln.y1
    return math.Sqrt(a * a + b * b)
}

/**
 * Generic shape interface for things that have perimeters and areas.
 */
type Shape interface {
    area() float64
    perimeter() float64
}

/**
 * Circle struct
 */
type Circle struct {
    x, y, r float64
}
func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}

/**
 * Rectangle struct
 */
type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
    ln_w := Line{x1: r.x1, y1: r.y1, x2: r.x2, y2: r.y1}
    ln_h := Line{x1: r.x1, y1: r.y1, x2: r.x1, y2: r.y2}
    w := ln_w.distance()
    h := ln_h.distance()
    return w * h
}
func (r *Rectangle) perimeter() float64 {
    ln_w := Line{x1: r.x1, y1: r.y1, x2: r.x2, y2: r.y1}
    ln_h := Line{x1: r.x1, y1: r.y1, x2: r.x1, y2: r.y2}
    w := ln_w.distance()
    h := ln_h.distance()
    return 2 * w + 2 * h
}

/**
 * A collector of shapes that adds up perimeters and areas.
 */
type MultiShape struct {
    shapes []Shape
}
func (ms *MultiShape) addShapes(shapes ...Shape) {
    for _, val := range shapes {
        ms.shapes = append(ms.shapes, val)
        fmt.Printf("Adding shape with area of %f and perimeter of %f.\n", val.area(), val.perimeter())
    }
}
func (ms *MultiShape) calculateTotalArea() float64 {
    var total float64
    for _, val := range ms.shapes {
        total += val.area()
    }
    return total
}
func (ms *MultiShape) calculateTotalPerimeter() float64 {
    var total float64
    for _, val := range ms.shapes {
        total += val.perimeter()
    }
    return total
}

func main() {
    rect := Rectangle{x1: 0, y1: 0, x2: 2, y2: 2}
    circ := Circle{x: 0, y: 0, r: 8}
    ms := new(MultiShape)
    ms.addShapes(&rect, &circ)
    fmt.Printf("Area of shapes is: %f\n", ms.calculateTotalArea())
    fmt.Printf("Perimeter of shapes is: %f\n", ms.calculateTotalPerimeter())
}
