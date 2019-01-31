// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/graph-gophers/graphql-go"
)

// Human A humanoid creature from the Star Wars universe
type Human struct {
	// ID The ID of the human
	ID graphql.ID `json:"id"`
	// Name What this human calls themselves
	Name string `json:"name"`
	// Height Height in the preferred unit, default is meters
	Height float64 `json:"height"`
	// Mass Mass in kilograms, or null if unknown
	Mass *float64 `json:"mass"`
	// Friends This human's friends, or an empty list if they have none
	Friends *[]*CharacterResolver `json:"friends"`
	// FriendsConnection The friends of the human exposed as a connection with edges
	FriendsConnection *FriendsConnectionResolver `json:"friendsConnection"`
	// AppearsIn The movies this human appears in
	AppearsIn []Episode `json:"appearsIn"`
	// Starships A list of starships this person has piloted, or an empty list if none
	Starships *[]*StarshipResolver `json:"starships"`
}

// HumanResolver resolver for Human
type HumanResolver struct {
	Human
}

// ID The ID of the human
func (r *HumanResolver) ID() graphql.ID {
	return r.Human.ID
}

// Name What this human calls themselves
func (r *HumanResolver) Name() string {
	return r.Human.Name
}

// Height Height in the preferred unit, default is meters
func (r *HumanResolver) Height(args *struct {
	Unit *LengthUnit
}) float64 {
	return r.Human.Height
}

// Mass Mass in kilograms, or null if unknown
func (r *HumanResolver) Mass() *float64 {
	return r.Human.Mass
}

// Friends This human's friends, or an empty list if they have none
func (r *HumanResolver) Friends() *[]*CharacterResolver {
	return r.Human.Friends
}

// FriendsConnection The friends of the human exposed as a connection with edges
func (r *HumanResolver) FriendsConnection(args *struct {
	First *int32
	After *graphql.ID
}) *FriendsConnectionResolver {
	return r.Human.FriendsConnection
}

// AppearsIn The movies this human appears in
func (r *HumanResolver) AppearsIn() []Episode {
	return r.Human.AppearsIn
}

// Starships A list of starships this person has piloted, or an empty list if none
func (r *HumanResolver) Starships() *[]*StarshipResolver {
	return r.Human.Starships
}

func (r *HumanResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Human)
}

func (r *HumanResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Human)
}
