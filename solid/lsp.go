package main

import "fmt"

// ---
// the following example is breaking the LSP
// the behavior of Square is not compatible with Rectangle

type CanSize interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	w, h int
}

func (r *Rectangle) GetWidth() int {
	return r.w
}

func (r *Rectangle) SetWidth(width int) {
	r.w = width
}

func (r *Rectangle) GetHeight() int {
	return r.h
}

func (r *Rectangle) SetHeight(height int) {
	r.h = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.w = size
	sq.h = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.w = width
	s.h = width
}

func (s *Square) SetHeight(height int) {
	s.w = height
	s.h = height
}

func CalcArea(sized CanSize) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

func main() {
	rc := &Rectangle{2, 3}
	CalcArea(rc)

	sq := NewSquare(5)
	CalcArea(sq)
}
