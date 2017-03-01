// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"
)

// Mutation The mutation type, represents all updates we can make to our data
type Mutation struct {
	CreateReview *ReviewResolver `json:"createReview"`
}

// MutationResolver resolver for Mutation
type MutationResolver struct {
	Mutation
}

func (r *MutationResolver) CreateReview() *ReviewResolver {
	return r.Mutation.CreateReview
}

func (r *MutationResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Mutation)
}

func (r *MutationResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Mutation)
}
