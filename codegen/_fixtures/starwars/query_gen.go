// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	graphql "github.com/neelance/graphql-go"
)

// Hero <nil>
func (r *Resolver) Hero(args *struct {
	Episode *Episode
}) *CharacterResolver {
	return nil
}

// Reviews <nil>
func (r *Resolver) Reviews(args *struct {
	Episode Episode
}) []*ReviewResolver {
	return nil
}

// Search <nil>
func (r *Resolver) Search(args *struct {
	Text string
}) []*SearchResultResolver {
	return nil
}

// Character <nil>
func (r *Resolver) Character(args *struct {
	ID graphql.ID
}) *CharacterResolver {
	return nil
}

// Droid <nil>
func (r *Resolver) Droid(args *struct {
	ID graphql.ID
}) *DroidResolver {
	return nil
}

// Human <nil>
func (r *Resolver) Human(args *struct {
	ID graphql.ID
}) *HumanResolver {
	return nil
}

// Starship <nil>
func (r *Resolver) Starship(args *struct {
	ID graphql.ID
}) *StarshipResolver {
	return nil
}
