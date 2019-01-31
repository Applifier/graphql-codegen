// This code is genereated by graphql-codegen
// DO NOT EDIT!

package httpget

import (
	graphql "github.com/graph-gophers/graphql-go"

	"net/http"

	"encoding/json"

	"fmt"
)

// Conversation
type Conversation struct {
	// ID
	ID graphql.ID `json:"id"`

	With_user_id string `json:"with_user_id"`
}

// ConversationResolver resolver for Conversation
type ConversationResolver struct {
	Conversation
}

// ID
func (r *ConversationResolver) ID() graphql.ID {
	return r.Conversation.ID
}

// With_user
func (r *ConversationResolver) With_user() (*UserResolver, error) {
	var result *UserResolver
	resp, err := http.Get(fmt.Sprintf("https://static.everyplay.com/developer-quiz/data/users/%s", r.Conversation.With_user_id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// Messages
func (r *ConversationResolver) Messages() ([]*MessageResolver, error) {
	var result []*MessageResolver
	resp, err := http.Get(fmt.Sprintf("https://static.everyplay.com/developer-quiz/data/conversations/%s/messages", r.Conversation.ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func (r *ConversationResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Conversation)
}

func (r *ConversationResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Conversation)
}
