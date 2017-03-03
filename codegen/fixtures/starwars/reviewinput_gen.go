// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

// ReviewInput The input object sent when someone is creating a new review
type ReviewInput struct {
	// Stars 0-5 stars
	Stars int32 `json:"stars"`
	// Commentary Comment about the movie, optional
	Commentary *string `json:"commentary"`
}
