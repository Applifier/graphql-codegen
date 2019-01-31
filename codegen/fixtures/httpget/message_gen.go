// This code is genereated by graphql-codegen
// DO NOT EDIT!

package httpget

import (
	graphql "github.com/graph-gophers/graphql-go"

	"net/http"

	"encoding/json"

	"fmt"
)

// Message
type Message struct {
	// ID
	ID graphql.ID `json:"id"`

	Conversation_id string `json:"conversation_id"`

	From_user_id string `json:"from_user_id"`

	// Body dasadssadsad
	Body string `json:"body"`
	// Created_at
	Created_at string `json:"created_at"`
}

// MessageResolver resolver for Message
type MessageResolver struct {
	Message
}

// ID
func (r *MessageResolver) ID() graphql.ID {
	return r.Message.ID
}

// Conversation
func (r *MessageResolver) Conversation() (*ConversationResolver, error) {
	var result *ConversationResolver
	resp, err := http.Get(fmt.Sprintf("https://static.everyplay.com/developer-quiz/data/conversations/%s", r.Message.Conversation_id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// From user blalala
func (r *MessageResolver) From() (*UserResolver, error) {
	var result *UserResolver
	resp, err := http.Get(fmt.Sprintf("https://static.everyplay.com/developer-quiz/data/users/%s", r.Message.From_user_id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// Body dasadssadsad
func (r *MessageResolver) Body() string {
	return r.Message.Body
}

// Created_at
func (r *MessageResolver) Created_at() string {
	return r.Message.Created_at
}

func (r *MessageResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.Message)
}

func (r *MessageResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Message)
}
