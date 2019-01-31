// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Character A character from the Star Wars universe
type Character interface {

	// ID The ID of the character
	ID() graphql.ID

	// Name The name of the character
	Name() string

	// Friends The friends of the character, or an empty list if they have none
	Friends() *[]*CharacterResolver

	// FriendsConnection The friends of the character exposed as a connection with edges
	FriendsConnection(args *struct {
		First *int32
		After *graphql.ID
	}) *FriendsConnectionResolver

	// AppearsIn The movies this character appears in
	AppearsIn() []Episode
}

// CharacterResolver resolver for Character
type CharacterResolver struct {
	Character
}

func (r *CharacterResolver) ToHuman() (*HumanResolver, bool) {
	c, ok := r.Character.(*HumanResolver)
	return c, ok
}

func (r *CharacterResolver) ToDroid() (*DroidResolver, bool) {
	c, ok := r.Character.(*DroidResolver)
	return c, ok
}
