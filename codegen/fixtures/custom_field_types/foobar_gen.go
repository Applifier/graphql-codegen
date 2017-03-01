package custom_field_types

import (
	"encoding/json"
)

// FooBar
type FooBar struct {
	BarFoo *BarFooResolver `json:"barFoo"`
}

// FooBarResolver resolver for FooBar
type FooBarResolver struct {
	FooBar
}

func (r *FooBarResolver) BarFoo() *BarFooResolver {
	return r.FooBar.BarFoo
}

func (r *FooBarResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.FooBar)
}

func (r *FooBarResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FooBar)
}
