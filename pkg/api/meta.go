package api

type Metadata struct {
	// Progress return a channel of numbers from 0.0 ... 1.0 used to inform the % completed by a task
	Progress chan float32

	// Messages return relevant information about the operator step
	Messages chan string

	// Done inform if an operator fully executed and the next is ready for the next stage
	State chan uint
}

func NewMetadata() Metadata {
	return Metadata{
		Progress: make(chan float32),
		Messages: make(chan string),
		State:    make(chan uint),
	}
}
