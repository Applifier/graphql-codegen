// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

import (
	"encoding/json"
)

// Review Represents a review for a movie
type Review struct {
	// Stars The number of stars this review gave, 1-5
	Stars int32 `json:"stars"`
	// Commentary Comment about the movie
	Commentary *string `json:"commentary"`
}

// ReviewResolver resolver for Review
type ReviewResolver struct {
	Review
}

// Stars The number of stars this review gave, 1-5
func (r *ReviewResolver) Stars() int32 {
	return r.Review.Stars
}

// Commentary Comment about the movie
func (r *ReviewResolver) Commentary() *string {
	return r.Review.Commentary
}

func (r *ReviewResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Review)
}

func (r *ReviewResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Review)
}
