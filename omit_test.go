package omit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	ID   Omit[int]     `json:"id,omitzero"`
	Name Omit[*string] `json:"name,omitzero"`
}

func TestOptional_MarshalJSON(t *testing.T) {
	data := []struct {
		name     string
		value    user
		expected string
	}{
		{
			name: "present",
			value: user{
				ID:   New(1),
				Name: NewPtr("john"),
			},
			expected: `{"id":1,"name":"john"}`,
		},
		{
			name: "optional",
			value: user{
				ID:   NewZero[int](),
				Name: NewZero[*string](),
			},
			expected: `{}`,
		},
		{
			name: "optional null",
			value: user{
				ID:   NewZero[int](),
				Name: NewNilPtr[string](),
			},
			expected: `{"name":null}`,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			b, err := json.Marshal(d.value)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, d.expected, string(b))
		})
	}
}

func TestOptional_UnmarshalJSON(t *testing.T) {
	data := []struct {
		name     string
		json     string
		expected user
	}{
		{
			name: "present",
			json: `{"id":1,"name":"john"}`,
			expected: user{
				ID:   New(1),
				Name: NewPtr("john"),
			},
		},
		{
			name: "optional",
			json: "{}",
			expected: user{
				ID:   NewZero[int](),
				Name: NewZero[*string](),
			},
		},
		{
			name: "optional null",
			json: `{"name":null}`,
			expected: user{
				ID:   NewZero[int](),
				Name: NewNilPtr[string](),
			},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			var u user
			if err := json.Unmarshal([]byte(d.json), &u); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, d.expected, u)
		})
	}
}
