package main

import (
    "fmt"
    "strings"
)

type Line struct {
    X1, Y1, X2, Y2 int
}
type VectorImage struct {
    Lines []Line
}

// given interface.
func NewRectangle(width, height int) *VectorImage {
    // given width or height is 5, we start from 0 to 4.
    width--
    height--
    return &VectorImage{
        []Line{
            Line{0, 0, width, 0},
            Line{0, 0, 0, height},
            Line{width, 0, width, height},
            Line{0, height, width, height},
        },
    }
}

type Point struct {
    X, Y int
}
type RasterImage interface {
    GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
    maxX, maxY := 0, 0
    points := owner.GetPoints()
    for _, pixel := range points {
        if pixel.X > maxX {
            maxX = pixel.X
        }
        if pixel.Y > maxY {
            maxY = pixel.Y
        }
    }
    maxX++
    maxY++
    data := make([][]rune, maxY)
    for i := 0; i < maxY; i++ {
        data[i] = make([]rune, maxX)
        for j := range data[i] {
            data[i][j] = ' '
        }
    }
    for _, point := range points {
        data[point.Y][point.X] = '*'
    }
    b := strings.Builder{}
    for _, line := range data {
        b.WriteString(string(line))
        b.WriteRune('\n')
    }
    return b.String()
}

// solution.
type vectorToRasterAdapter struct {
    points []Point
}

func minmax(a, b int) (int, int) {
    if a < b {
        return a, b
    }
    return b, a
}
func (v vectorToRasterAdapter) GetPoints() []Point {
    return v.points
}
func (v *vectorToRasterAdapter) addLine(line Line) {
    left, right := minmax(line.X1, line.X2)
    top, bottom := minmax(line.Y1, line.Y2)
    dx := right - left
    dy := line.Y2 - line.Y1
    if dx == 0 {
        for y := top; y <= bottom; y++ {
            v.points = append(v.points, Point{left, y})
        }
    } else if dy == 0 {
        for x := left; x <= right; x++ {
            v.points = append(v.points, Point{x, top})
        }
    }
    fmt.Println("we have", len(v.points), "points")
}

//Factory as we are going to construct the raster interface/implement
func VectorToRaster(vi *VectorImage) RasterImage {
    adapter := vectorToRasterAdapter{}
    for _, line := range vi.Lines {
        adapter.addLine(line)
    }
    return adapter // as a RasterImage
}
func main() {
    rc := NewRectangle(6, 4)
    // convert rectangle to a raster image.
    a := VectorToRaster(rc)
    fmt.Print(DrawPoints(a))
}
