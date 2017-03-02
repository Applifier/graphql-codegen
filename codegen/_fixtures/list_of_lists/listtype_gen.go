// This code is genereated by graphql-codegen
// DO NOT EDIT!

package list_of_lists

import (
	"encoding/json"
)

// ListType Information for paginating this connection
type ListType struct {
	// List <nil>
	List *[]*string `json:"list"`
	// ListOfList <nil>
	ListOfList *[]*[]*string `json:"listOfList"`
	// NonNullListOfList <nil>
	NonNullListOfList [][]string `json:"nonNullListOfList"`
}

// ListTypeResolver resolver for ListType
type ListTypeResolver struct {
	ListType
}

// List <nil>
func (r *ListTypeResolver) List() *[]*string {
	return r.ListType.List
}

// ListOfList <nil>
func (r *ListTypeResolver) ListOfList() *[]*[]*string {
	return r.ListType.ListOfList
}

// NonNullListOfList <nil>
func (r *ListTypeResolver) NonNullListOfList() [][]string {
	return r.ListType.NonNullListOfList
}

func (r *ListTypeResolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.ListType)
}

func (r *ListTypeResolver) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ListType)
}
