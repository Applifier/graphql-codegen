// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Hero
func (r *Resolver) Hero(args *struct {
	Episode *Episode
}) *CharacterResolver {
	return nil
}

// Reviews
func (r *Resolver) Reviews(args *struct {
	Episode Episode
}) []*ReviewResolver {
	return nil
}

// Search
func (r *Resolver) Search(args *struct {
	Text string
}) []*SearchResultResolver {
	return nil
}

// Character
func (r *Resolver) Character(args *struct {
	ID graphql.ID
}) *CharacterResolver {
	return nil
}

// Droid
func (r *Resolver) Droid(args *struct {
	ID graphql.ID
}) *DroidResolver {
	return nil
}

// Human
func (r *Resolver) Human(args *struct {
	ID graphql.ID
}) *HumanResolver {
	return nil
}

// Starship
func (r *Resolver) Starship(args *struct {
	ID graphql.ID
}) *StarshipResolver {
	return nil
}
