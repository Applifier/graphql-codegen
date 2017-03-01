// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"
)

// Review Represents a review for a movie
type Review struct {
	Stars      int32   `json:"stars"`
	Commentary *string `json:"commentary"`
}

// ReviewResolver resolver for Review
type ReviewResolver struct {
	Review
}

func (r *ReviewResolver) Stars() int32 {
	return r.Review.Stars
}

func (r *ReviewResolver) Commentary() *string {
	return r.Review.Commentary
}

func (r *ReviewResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Review)
}

func (r *ReviewResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Review)
}
