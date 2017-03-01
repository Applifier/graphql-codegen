// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/neelance/graphql-go"
)

// FriendsEdge An edge object for a character's friends
type FriendsEdge struct {
	Cursor graphql.ID         `json:"cursor"`
	Node   *CharacterResolver `json:"node"`
}

// FriendsEdgeResolver resolver for FriendsEdge
type FriendsEdgeResolver struct {
	FriendsEdge
}

func (r *FriendsEdgeResolver) Cursor() graphql.ID {
	return r.FriendsEdge.Cursor
}

func (r *FriendsEdgeResolver) Node() *CharacterResolver {
	return r.FriendsEdge.Node
}

func (r *FriendsEdgeResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.FriendsEdge)
}

func (r *FriendsEdgeResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FriendsEdge)
}
