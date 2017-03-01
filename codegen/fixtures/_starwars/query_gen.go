// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"
)

// Query The query type, represents all of the entry points into our object graph
type Query struct {
	Hero      *CharacterResolver   `json:"hero"`
	Reviews   ReviewResolver       `json:"reviews"`
	Search    SearchResultResolver `json:"search"`
	Character *CharacterResolver   `json:"character"`
	Droid     *DroidResolver       `json:"droid"`
	Human     *HumanResolver       `json:"human"`
	Starship  *StarshipResolver    `json:"starship"`
}

// QueryResolver resolver for Query
type QueryResolver struct {
	Query
}

func (r *QueryResolver) Hero() *CharacterResolver {
	return r.Query.Hero
}

func (r *QueryResolver) Reviews() ReviewResolver {
	return r.Query.Reviews
}

func (r *QueryResolver) Search() SearchResultResolver {
	return r.Query.Search
}

func (r *QueryResolver) Character() *CharacterResolver {
	return r.Query.Character
}

func (r *QueryResolver) Droid() *DroidResolver {
	return r.Query.Droid
}

func (r *QueryResolver) Human() *HumanResolver {
	return r.Query.Human
}

func (r *QueryResolver) Starship() *StarshipResolver {
	return r.Query.Starship
}

func (r *QueryResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Query)
}

func (r *QueryResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Query)
}
