// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/graph-gophers/graphql-go"
)

// Droid An autonomous mechanical character in the Star Wars universe
type Droid struct {
	// ID The ID of the droid
	ID graphql.ID `json:"id"`
	// Name What others call this droid
	Name string `json:"name"`
	// Friends This droid's friends, or an empty list if they have none
	Friends *[]*CharacterResolver `json:"friends"`
	// FriendsConnection The friends of the droid exposed as a connection with edges
	FriendsConnection *FriendsConnectionResolver `json:"friendsConnection"`
	// AppearsIn The movies this droid appears in
	AppearsIn []Episode `json:"appearsIn"`
	// PrimaryFunction This droid's primary function
	PrimaryFunction *string `json:"primaryFunction"`
}

// DroidResolver resolver for Droid
type DroidResolver struct {
	Droid
}

// ID The ID of the droid
func (r *DroidResolver) ID() graphql.ID {
	return r.Droid.ID
}

// Name What others call this droid
func (r *DroidResolver) Name() string {
	return r.Droid.Name
}

// Friends This droid's friends, or an empty list if they have none
func (r *DroidResolver) Friends() *[]*CharacterResolver {
	return r.Droid.Friends
}

// FriendsConnection The friends of the droid exposed as a connection with edges
func (r *DroidResolver) FriendsConnection(args *struct {
	First *int32
	After *graphql.ID
}) *FriendsConnectionResolver {
	return r.Droid.FriendsConnection
}

// AppearsIn The movies this droid appears in
func (r *DroidResolver) AppearsIn() []Episode {
	return r.Droid.AppearsIn
}

// PrimaryFunction This droid's primary function
func (r *DroidResolver) PrimaryFunction() *string {
	return r.Droid.PrimaryFunction
}

func (r *DroidResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Droid)
}

func (r *DroidResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Droid)
}
