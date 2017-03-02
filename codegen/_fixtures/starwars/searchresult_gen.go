// This code is genereated by graphql-codegen
// DO NOT EDIT!

package starwars

// SearchResultResolver resolver for SearchResult
type SearchResultResolver struct {
	searchResult interface{}
}

func (r *SearchResultResolver) ToHuman() (*HumanResolver, bool) {
	c, ok := r.searchResult.(*HumanResolver)
	return c, ok
}

func (r *SearchResultResolver) ToDroid() (*DroidResolver, bool) {
	c, ok := r.searchResult.(*DroidResolver)
	return c, ok
}

func (r *SearchResultResolver) ToStarship() (*StarshipResolver, bool) {
	c, ok := r.searchResult.(*StarshipResolver)
	return c, ok
}
