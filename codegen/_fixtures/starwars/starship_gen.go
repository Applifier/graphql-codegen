// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/neelance/graphql-go"
)

// Starship
type Starship struct {
	ID     graphql.ID `json:"id"`
	Name   string     `json:"name"`
	Length float64    `json:"length"`
}

// StarshipResolver resolver for Starship
type StarshipResolver struct {
	Starship
}

func (r *StarshipResolver) ID() graphql.ID {
	return r.Starship.ID
}

func (r *StarshipResolver) Name() string {
	return r.Starship.Name
}

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
