package process

type ActivationFunction interface {
	Activate(input float64) float64
}
