// This code is genereated by graphql-codegen
// DO NOT EDIT!

package httpget

import (
	"encoding/json"

	graphql "github.com/graph-gophers/graphql-go"
)

// User
type User struct {
	// ID
	ID graphql.ID `json:"id"`
	// Username
	Username string `json:"username"`
	// Avatar_url
	Avatar_url string `json:"avatar_url"`
}

// UserResolver resolver for User
type UserResolver struct {
	User
}

// ID
func (r *UserResolver) ID() graphql.ID {
	return r.User.ID
}

// Username
func (r *UserResolver) Username() string {
	return r.User.Username
}

// Avatar_url
func (r *UserResolver) Avatar_url() string {
	return r.User.Avatar_url
}

func (r *UserResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.User)
}

func (r *UserResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.User)
}
