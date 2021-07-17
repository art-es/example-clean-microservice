package calculator

type Calculator struct{}

// Doubling is a method of multiplying the specified number by 2.
func (*Calculator) Doubling(num float64) float64 {
	return num * 2
}
