// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	graphql "github.com/neelance/graphql-go"
)

// Character A character from the Star Wars universe
type Character interface {
	ID() graphql.ID

	Name() string

	Friends() *[]*CharacterResolver

	FriendsConnection(args *struct {
		First *int32
		After *graphql.ID
	}) FriendsConnectionResolver

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
