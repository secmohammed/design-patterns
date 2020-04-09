package main

import "fmt"

type Image interface {
    Draw()
}
type Bitmap struct {
    filename string
}

func NewBitmap(filename string) *Bitmap {
    fmt.Println("Loading image from", filename)
    return &Bitmap{filename}
}

func (b *Bitmap) Draw() {
    fmt.Println("Drawing image", b.filename)
}
func DrawImage(image Image) {
    fmt.Println("about to draw the image")
    image.Draw()
    fmt.Println("Done drawing the image")
}

// lazy load the image, so it gets only rendered when needed via
// proxy pattern.
type LazyBitmap struct {
    filename string
    bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
    return &LazyBitmap{filename: filename}
}

func (l *LazyBitmap) Draw() {
    // virutal because when we create a lazy bitmap, it hasn't been materialized yet.
    // it hasn't been construted yet until somebody asks for it.

    if l.bitmap == nil {
        l.bitmap = NewBitmap(l.filename)
    }
    l.bitmap.Draw()
}
func main() {
    bmp := NewBitmap("demo.png")
    DrawImage(bmp)
    //if we instantiate the bitmap without using it, it will print loading, which isn't lazy
    // _ := NewBitmap("demo.png") this will print laoding from image, even though we haven't drown it yet.
    lbmp := NewLazyBitmap("demo.png")
    DrawImage(lbmp)
}
