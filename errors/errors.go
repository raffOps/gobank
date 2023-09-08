package errors

import "fmt"

type InsufficientBalance struct {
	Value float64
}

func (e *InsufficientBalance) Error() string {
	return fmt.Sprintf("Insufficient balance for value %.2f", e.Value)
}

type InvalidValue struct {
	Value float64
}

func (e *InvalidValue) Error() string {
	return fmt.Sprintf("Invalid Value: %f. The Value must be greater than 0", e.Value)
}
