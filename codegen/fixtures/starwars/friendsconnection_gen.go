// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"
)

// FriendsConnection A connection object for a character's friends
type FriendsConnection struct {
	// TotalCount The total number of friends
	TotalCount int32 `json:"totalCount"`
	// Edges The edges for each of the character's friends.
	Edges *[]*FriendsEdgeResolver `json:"edges"`
	// Friends A list of the friends, as a convenience when edges are not needed.
	Friends *[]*CharacterResolver `json:"friends"`
	// PageInfo Information for paginating this connection
	PageInfo *PageInfoResolver `json:"pageInfo"`
}

// FriendsConnectionResolver resolver for FriendsConnection
type FriendsConnectionResolver struct {
	FriendsConnection
}

// TotalCount The total number of friends
func (r *FriendsConnectionResolver) TotalCount() int32 {
	return r.FriendsConnection.TotalCount
}

// Edges The edges for each of the character's friends.
func (r *FriendsConnectionResolver) Edges() *[]*FriendsEdgeResolver {
	return r.FriendsConnection.Edges
}

// Friends A list of the friends, as a convenience when edges are not needed.
func (r *FriendsConnectionResolver) Friends() *[]*CharacterResolver {
	return r.FriendsConnection.Friends
}

// PageInfo Information for paginating this connection
func (r *FriendsConnectionResolver) PageInfo() *PageInfoResolver {
	return r.FriendsConnection.PageInfo
}

func (r *FriendsConnectionResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.FriendsConnection)
}

func (r *FriendsConnectionResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FriendsConnection)
}
