package main

import "fmt"

type Buffer struct {
    widht, height int
    buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
    return &Buffer{width, height, make([]rune, width*height)}
}
func (b *Buffer) At(index int) rune {
    return b.buffer[index]
}

type Viewport struct {
    buffer *Buffer
    offset int
}

func NewViewport(buffer *Buffer) *Viewport {
    return &Viewport{buffer: buffer}
}
func (v *Viewport) GetCharacterAt(index int) rune {
    return v.buffer.At(v.offset + index)
}

type Console struct {
    buffer    []*Buffer
    viewports []*Viewport
    offset    int
}

func NewConsole(width, height, offset int) *Console {
    b := NewBuffer(width, height)
    v := NewViewport(b)
    return &Console{[]*Buffer{b}, []*Viewport{v}, offset}
}
func (c *Console) GetCharacterAt(index int) rune {
    return c.viewports[0].GetCharacterAt(index)
}
func main() {
    c := NewConsole(10, 200, 0)
    u := c.GetCharacterAt(1)
    fmt.Println(u)
}
