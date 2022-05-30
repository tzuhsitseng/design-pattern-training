package main

import "fmt"

type shapeType int

type shape interface {
	getType() shapeType
	accept(visitor visitor)
}

type visitor interface {
	visitRectangle(rec *rectangle)
	visitSquare(sq *square)
}

type rectangle struct {
	w, h int
}

func (r *rectangle) getType() shapeType {
	return rectangleType
}

func (r *rectangle) accept(visitor visitor) {
	visitor.visitRectangle(r)
}

type square struct {
	side int
}

func (s *square) getType() shapeType {
	return squareType
}

func (s *square) accept(visitor visitor) {
	visitor.visitSquare(s)
}

type areaCalculator struct {
}

func (c *areaCalculator) visitRectangle(rec *rectangle) {
	fmt.Printf("area of rec: %d\n", rec.h*rec.w)
}

func (c *areaCalculator) visitSquare(sq *square) {
	fmt.Printf("area of square: %d\n", sq.side*sq.side)
}

const (
	rectangleType shapeType = iota + 1
	squareType
)

func main() {
	sq := square{side: 10}
	rec := rectangle{w: 50, h: 3}
	ac := areaCalculator{}

	sq.accept(&ac)
	rec.accept(&ac)
}
