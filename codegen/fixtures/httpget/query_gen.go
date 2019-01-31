// This code is genereated by graphql-codegen
// DO NOT EDIT!

package httpget

import (
	graphql "github.com/graph-gophers/graphql-go"

	"net/http"

	"encoding/json"

	"fmt"
)

// User
func (r *Resolver) User(args *struct {
	ID graphql.ID
}) (*UserResolver, error) {
	var result *UserResolver
	resp, err := http.Get(fmt.Sprintf("https://static.everyplay.com/developer-quiz/data/users/%s", args.ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// Conversation
func (r *Resolver) Conversation(args *struct {
	ID graphql.ID
}) (*ConversationResolver, error) {
	var result *ConversationResolver
	resp, err := http.Get(fmt.Sprintf("https://static.everyplay.com/developer-quiz/data/conversations/%s", args.ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// Conversations
func (r *Resolver) Conversations() ([]*ConversationResolver, error) {
	var result []*ConversationResolver
	resp, err := http.Get("https://static.everyplay.com/developer-quiz/data/conversations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
