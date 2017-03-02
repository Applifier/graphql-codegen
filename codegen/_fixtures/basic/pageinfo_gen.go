// This code is genereated by graphql-codegen
// DO NOT EDIT!

package basic

import (
	"encoding/json"

	graphql "github.com/neelance/graphql-go"
)

// PageInfo Information for paginating this connection
type PageInfo struct {
	// StartCursor <nil>
	StartCursor *graphql.ID `json:"startCursor"`
	// EndCursor <nil>
	EndCursor *graphql.ID `json:"endCursor"`
	// HasNextPage <nil>
	HasNextPage bool `json:"hasNextPage"`
}

// PageInfoResolver resolver for PageInfo
type PageInfoResolver struct {
	PageInfo
}

// StartCursor <nil>
func (r *PageInfoResolver) StartCursor() *graphql.ID {
	return r.PageInfo.StartCursor
}

// EndCursor <nil>
func (r *PageInfoResolver) EndCursor() *graphql.ID {
	return r.PageInfo.EndCursor
}

// HasNextPage <nil>
func (r *PageInfoResolver) HasNextPage() bool {
	return r.PageInfo.HasNextPage
}

func (r *PageInfoResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.PageInfo)
}

func (r *PageInfoResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PageInfo)
}
