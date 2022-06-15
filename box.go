package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) > b.shapesCapacity {
		return errors.New("out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	var shape Shape
	shape, ok := b.shapes[i].(Shape)
	if !ok {
		return shape, errors.New("shape by index doesn't exist")
	} else {
		return shape, nil
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	//return append(slice[:s], slice[s+1:]...)
	var shape Shape
	shape, ok := b.shapes[i].(Shape)
	if !ok {
		return shape, errors.New("shape by index doesn't exist")
	} else {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return shape, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var removedShape Shape
	removedShape, ok := b.shapes[i].(Shape)
	if !ok {
		return removedShape, errors.New("shape by index doesn't exist")
	} else {
		b.shapes[i] = shape
		return removedShape, nil
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumP float64
	for i, _ := range b.shapes {
		sumP += b.shapes[i].CalcPerimeter()
	}
	return sumP
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumA float64
	for i, _ := range b.shapes {
		sumA += b.shapes[i].CalcArea()
	}
	return sumA
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	for i, _ := range b.shapes {
		circ, ok := b.shapes[i].(Circle)
		fmt.Println(circ)
		if !ok {
			return errors.New("circles are not exist in the list")
		} else {
			b.shapes = nil
			return nil
		}
	}
	return nil
}
