// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	graphql "github.com/neelance/graphql-go"
)

func (r *Resolver) Hero(args *struct {
	Episode *Episode
}) *CharacterResolver {
	return nil
}

func (r *Resolver) Reviews(args *struct {
	Episode Episode
}) []*ReviewResolver {
	return nil
}

func (r *Resolver) Search(args *struct {
	Text string
}) []*SearchResultResolver {
	return nil
}

func (r *Resolver) Character(args *struct {
	ID graphql.ID
}) *CharacterResolver {
	return nil
}

func (r *Resolver) Droid(args *struct {
	ID graphql.ID
}) *DroidResolver {
	return nil
}

func (r *Resolver) Human(args *struct {
	ID graphql.ID
}) *HumanResolver {
	return nil
}

func (r *Resolver) Starship(args *struct {
	ID graphql.ID
}) *StarshipResolver {
	return nil
}
