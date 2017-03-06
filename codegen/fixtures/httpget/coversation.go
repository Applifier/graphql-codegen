package httpget

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// This file contains custom definitions for ConversationResolver

// Title returns a title for the conversation
func (r *ConversationResolver) Title() (string, error) {
	var user *UserResolver
	resp, err := http.Get(fmt.Sprintf("https://static.everyplay.com/developer-quiz/data/users/%s", r.Conversation.With_user_id))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", err
	}

	return fmt.Sprintf("Conversation between you and %s", user.Username()), err
}
