// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/neelance/graphql-go"
)

// Droid An autonomous mechanical character in the Star Wars universe
type Droid struct {
	Id                graphql.ID                `json:"id"`
	Name              string                    `json:"name"`
	Friends           *[]*CharacterResolver     `json:"friends"`
	FriendsConnection FriendsConnectionResolver `json:"friendsConnection"`
	AppearsIn         []EpisodeResolver         `json:"appearsIn"`
	PrimaryFunction   *string                   `json:"primaryFunction"`
}

// DroidResolver resolver for Droid
type DroidResolver struct {
	Droid
}

func (r *DroidResolver) Id() graphql.ID {
	return r.Droid.Id
}

func (r *DroidResolver) Name() string {
	return r.Droid.Name
}

func (r *DroidResolver) Friends() *[]*CharacterResolver {
	return r.Droid.Friends
}

func (r *DroidResolver) FriendsConnection() FriendsConnectionResolver {
	return r.Droid.FriendsConnection
}

func (r *DroidResolver) AppearsIn() []EpisodeResolver {
	return r.Droid.AppearsIn
}

func (r *DroidResolver) PrimaryFunction() *string {
	return r.Droid.PrimaryFunction
}

func (r *DroidResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Droid)
}

func (r *DroidResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Droid)
}
