package httpget

import (
	"io/ioutil"
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestHTTPSchema(t *testing.T) {
	schema, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		t.Fatal(err)
	}

	resolver := &Resolver{}

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema: graphql.MustParseSchema(string(schema), resolver),
			Query: `
				{
					user(id: "1") {
            username,
						avatar_url
          }
				}
			`,
			ExpectedResult: `
				{
					"user": {
					  "username": "John",
						"avatar_url": "http://placekitten.com/g/300/300"
					}
				}
			`,
		},
		{
			Schema: graphql.MustParseSchema(string(schema), resolver),
			Query: `
				{
					conversations {
            id,
						with_user {
							id,
							username,
							avatar_url,
						},
          }
				}
			`,
			ExpectedResult: `
			{
				 "conversations":[
						{
							 "id":"1",
							 "with_user":{
									"avatar_url":"http://placekitten.com/g/301/301",
									"id":"2",
									"username":"Amy"
							 }
						},
						{
							 "id":"2",
							 "with_user":{
									"avatar_url":"http://placekitten.com/g/302/302",
									"id":"3",
									"username":"Jeremy"
							 }
						},
						{
							 "id":"3",
							 "with_user":{
									"avatar_url":"http://placekitten.com/g/303/303",
									"id":"4",
									"username":"Hannah"
							 }
						},
						{
							 "id":"4",
							 "with_user":{
									"avatar_url":"http://placekitten.com/g/304/304",
									"id":"5",
									"username":"Charles"
							 }
						},
						{
							 "id":"5",
							 "with_user":{
									"avatar_url":"http://placekitten.com/g/305/305",
									"id":"6",
									"username":"George"
							 }
						}
				 ]
				}
			`,
		},
		{
			Schema: graphql.MustParseSchema(string(schema), resolver),
			Query: `
				{
					conversation(id: "1") {
            id,
						title,
						with_user {
							id,
							username,
							avatar_url,
						},
						messages {
							id,
							body,
							from {
								username
							}
						}
          }
				}
			`,
			ExpectedResult: `
			{
			 "conversation":{
					"id":"1",
					"title": "Conversation between you and Amy",
					"messages":[
						 {
								"body":"Moi!",
								"from":{
									 "username":"John"
								},
								"id":"1"
						 }
					],
					"with_user":{
						 "avatar_url":"http://placekitten.com/g/301/301",
						 "id":"2",
						 "username":"Amy"
					}
			 }
			}
			`,
		},
	})
}
