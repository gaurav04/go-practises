package entities

type Geometry interface {
	GetName() string
	Area() float64
	Perimeter() float64
}
