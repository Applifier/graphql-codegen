package main

import (
	"encoding/json"

	graphql "github.com/neelance/graphql-go"
)

// PageInfo Information for paginating this connection
type PageInfo struct {
	StartCursor *graphql.ID `json:"startCursor"`
	EndCursor   *graphql.ID `json:"endCursor"`
	HasNextPage bool        `json:"hasNextPage"`
}

// PageInfoResolver resolver for PageInfo
type PageInfoResolver struct {
	PageInfo
}

func (r *PageInfoResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.PageInfo)
}

func (r *PageInfoResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PageInfo)
}

func (r *PageInfoResolver) StartCursor() *graphql.ID {
	return r.PageInfo.StartCursor
}

func (r *PageInfoResolver) EndCursor() *graphql.ID {
	return r.PageInfo.EndCursor
}

func (r *PageInfoResolver) HasNextPage() bool {
	return r.PageInfo.HasNextPage
}
