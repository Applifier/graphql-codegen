// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"

	graphql "github.com/graph-gophers/graphql-go"
)

// PageInfo Information for paginating this connection
type PageInfo struct {
	// StartCursor
	StartCursor *graphql.ID `json:"startCursor"`
	// EndCursor
	EndCursor *graphql.ID `json:"endCursor"`
	// HasNextPage
	HasNextPage bool `json:"hasNextPage"`
}

// PageInfoResolver resolver for PageInfo
type PageInfoResolver struct {
	PageInfo
}

// StartCursor
func (r *PageInfoResolver) StartCursor() *graphql.ID {
	return r.PageInfo.StartCursor
}

// EndCursor
func (r *PageInfoResolver) EndCursor() *graphql.ID {
	return r.PageInfo.EndCursor
}

// HasNextPage
func (r *PageInfoResolver) HasNextPage() bool {
	return r.PageInfo.HasNextPage
}

func (r *PageInfoResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.PageInfo)
}

func (r *PageInfoResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PageInfo)
}
