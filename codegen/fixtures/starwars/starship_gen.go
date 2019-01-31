// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/graph-gophers/graphql-go"
)

// Starship
type Starship struct {
	// ID The ID of the starship
	ID graphql.ID `json:"id"`
	// Name The name of the starship
	Name string `json:"name"`
	// Length Length of the starship, along the longest axis
	Length float64 `json:"length"`
}

// StarshipResolver resolver for Starship
type StarshipResolver struct {
	Starship
}

// ID The ID of the starship
func (r *StarshipResolver) ID() graphql.ID {
	return r.Starship.ID
}

// Name The name of the starship
func (r *StarshipResolver) Name() string {
	return r.Starship.Name
}

// Length Length of the starship, along the longest axis
func (r *StarshipResolver) Length(args *struct {
	Unit *LengthUnit
}) float64 {
	return r.Starship.Length
}

func (r *StarshipResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Starship)
}

func (r *StarshipResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Starship)
}
