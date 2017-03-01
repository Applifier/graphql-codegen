// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"
)

// FriendsConnection A connection object for a character's friends
type FriendsConnection struct {
	TotalCount int64                `json:"totalCount"`
	Edges      *FriendsEdgeResolver `json:"edges"`
	Friends    *CharacterResolver   `json:"friends"`
	PageInfo   PageInfoResolver     `json:"pageInfo"`
}

// FriendsConnectionResolver resolver for FriendsConnection
type FriendsConnectionResolver struct {
	FriendsConnection
}

func (r *FriendsConnectionResolver) TotalCount() int64 {
	return r.FriendsConnection.TotalCount
}

func (r *FriendsConnectionResolver) Edges() *FriendsEdgeResolver {
	return r.FriendsConnection.Edges
}

func (r *FriendsConnectionResolver) Friends() *CharacterResolver {
	return r.FriendsConnection.Friends
}

func (r *FriendsConnectionResolver) PageInfo() PageInfoResolver {
	return r.FriendsConnection.PageInfo
}

func (r *FriendsConnectionResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.FriendsConnection)
}

func (r *FriendsConnectionResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FriendsConnection)
}
