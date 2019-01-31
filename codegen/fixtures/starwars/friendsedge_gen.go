// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/graph-gophers/graphql-go"
)

// FriendsEdge An edge object for a character's friends
type FriendsEdge struct {
	// Cursor A cursor used for pagination
	Cursor graphql.ID `json:"cursor"`
	// Node The character represented by this friendship edge
	Node *CharacterResolver `json:"node"`
}

// FriendsEdgeResolver resolver for FriendsEdge
type FriendsEdgeResolver struct {
	FriendsEdge
}

// Cursor A cursor used for pagination
func (r *FriendsEdgeResolver) Cursor() graphql.ID {
	return r.FriendsEdge.Cursor
}

// Node The character represented by this friendship edge
func (r *FriendsEdgeResolver) Node() *CharacterResolver {
	return r.FriendsEdge.Node
}

func (r *FriendsEdgeResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.FriendsEdge)
}

func (r *FriendsEdgeResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FriendsEdge)
}
