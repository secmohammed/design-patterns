package main

import "fmt"

type Shape interface {
    Render() string
}
type Circle struct {
    Radius float32
}
type Square struct {
    Length float32
}

func (c *Circle) Render() string {
    return fmt.Sprintf("circle is rendered with radius %f", c.Radius)
}
func (s *Square) Render() string {
    return fmt.Sprintf("square is rendered with length %f", s.Length)
}

// Instead of adding color to each struct, which violates OCP.
type ColoredShape struct {
    Shape Shape
    Color string
}

func (c *ColoredShape) Render() string {
    return fmt.Sprintf("%s and with color of %s", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
    Shape        Shape
    Transparency float32
}

func (t *TransparentShape) Render() string {
    return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}
func main() {
    circle := Circle{11}
    fmt.Println(circle.Render())
    coloredCircle := ColoredShape{&circle, "Red"}
    fmt.Println(coloredCircle.Render())
    // we lost the ability of doing this.
    // coloredCircle.Resize()
    // as this is only available for the circle, but isn't available for square.
    // and to avoid this issue, we can compose decorators.
    rhsCircle := TransparentShape{&coloredCircle, 0.5}
    fmt.Println(rhsCircle.Render())
}
