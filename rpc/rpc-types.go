package rpc

// import (
//     "fmt"
//     "math"
// )

// How we communicate pressed keys to the math manager script
type KeyDownCommand struct {
	KeyGlyph rune
}

type KeyUpCommand struct {
	KeyGlyph rune
}

type Point3D struct {
	X, Y, Z float32
}

type Point2D struct {
	X, Y     float32
	DrawSize uint8
}

type Line2D struct {
	Point1    Point2D
	Point2    Point2D
	DrawWidth uint8
}

type Polygon2D struct {
	Edges []Line2D
}

type Color struct {
	R, G, B uint8
}

type GuiDrawCommand struct {
	Lines  []Line2D  // List of endpoint pairs
	Points []Point2D // List of 3d points
	// Polygons []Polygon // List of polygons
}

// func (dc *DrawCommand) k
