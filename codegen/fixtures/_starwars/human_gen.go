// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/neelance/graphql-go"
)

// Human A humanoid creature from the Star Wars universe
type Human struct {
	Id                graphql.ID                `json:"id"`
	Name              string                    `json:"name"`
	Height            float64                   `json:"height"`
	Mass              *float64                  `json:"mass"`
	Friends           *[]*CharacterResolver     `json:"friends"`
	FriendsConnection FriendsConnectionResolver `json:"friendsConnection"`
	AppearsIn         []EpisodeResolver         `json:"appearsIn"`
	Starships         *[]*StarshipResolver      `json:"starships"`
}

// HumanResolver resolver for Human
type HumanResolver struct {
	Human
}

func (r *HumanResolver) Id() graphql.ID {
	return r.Human.Id
}

func (r *HumanResolver) Name() string {
	return r.Human.Name
}

func (r *HumanResolver) Height() float64 {
	return r.Human.Height
}

func (r *HumanResolver) Mass() *float64 {
	return r.Human.Mass
}

func (r *HumanResolver) Friends() *[]*CharacterResolver {
	return r.Human.Friends
}

func (r *HumanResolver) FriendsConnection() FriendsConnectionResolver {
	return r.Human.FriendsConnection
}

func (r *HumanResolver) AppearsIn() []EpisodeResolver {
	return r.Human.AppearsIn
}

func (r *HumanResolver) Starships() *[]*StarshipResolver {
	return r.Human.Starships
}

func (r *HumanResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Human)
}

func (r *HumanResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Human)
}