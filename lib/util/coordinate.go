package util

type Coordinate[T any] struct {
	X, Y T
}

func NewCoordinate[T any](x, y T) *Coordinate[T] {
	return &Coordinate[T]{
		X: x,
		Y: y,
	}
}
