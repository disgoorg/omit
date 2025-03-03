package omit

import (
	"encoding/json"
	"fmt"
)

var _ json.Marshaler = (*Omit[int])(nil)
var _ json.Unmarshaler = (*Omit[int])(nil)

// Ptr returns a pointer to the given value.
func Ptr[T any](v T) *T {
	return &v
}

// New creates a new Omit with the given value set.
func New[T any](v T) Omit[T] {
	return Omit[T]{
		Value: v,
		OK:    true,
	}
}

// NewPtr creates a new Omit with the given pointer value set.
func NewPtr[T any](v T) Omit[*T] {
	return Omit[*T]{
		Value: &v,
		OK:    true,
	}
}

// NewNilPtr creates a new Omit with the value set to nil.
func NewNilPtr[T any]() Omit[*T] {
	return Omit[*T]{
		OK: true,
	}
}

// NewZero creates a new Omit with the value not set.
func NewZero[T any]() Omit[T] {
	return Omit[T]{
		OK: false,
	}
}

// Omit is a type that can be used to represent a value which may or may not be set.
// This is useful for omitting the value in JSON. The zero value of Omit is not set.
type Omit[T any] struct {
	Value T
	OK    bool
}

// String returns the string representation of the value if it is set, otherwise it returns "<omitted>".
func (o Omit[T]) String() string {
	if !o.OK {
		return "<omitted>"
	}
	return fmt.Sprint(o.Value)
}

// IsZero returns true if the value is not set. This is useful for omitting the value in JSON.
func (o Omit[T]) IsZero() bool {
	return !o.OK
}

// MarshalJSON marshals the value if it is set, otherwise it returns nil.
func (o Omit[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals the value if it is set, otherwise it returns nil.
func (o *Omit[T]) UnmarshalJSON(data []byte) error {
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	o.Value = v
	o.OK = true
	return nil
}

// Or returns the value if it is set, otherwise it returns the default value.
func (o Omit[T]) Or(def T) T {
	if !o.OK {
		return def
	}
	return o.Value
}
