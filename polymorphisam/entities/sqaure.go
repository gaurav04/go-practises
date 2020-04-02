package entities

type Square struct {
	Name string
	A    float64
}

func (square Square) GetName() string {
	return square.Name
}

func (square Square) Area() float64 {
	return square.A * square.A
}

func (square Square) Perimeter() float64 {
	return square.A * 4
}
