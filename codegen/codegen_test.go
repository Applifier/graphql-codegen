package codegen

import (
	"testing"

	"github.com/Applifier/graphql-codegen/config"
)

func TestCodegen(t *testing.T) {
	tests := []struct {
		schema   string
		config   config.Config
		expected map[string]string
	}{
		{
			config: config.Config{
				Package: "main",
			},
			schema: `
    # Information for paginating this connection
    type PageInfo {
      startCursor: ID
      endCursor: ID
      hasNextPage: Boolean!
    }
    `,
			expected: map[string]string{
				"pageinfo_gen.go": `
        package main;

        import (
          graphql "github.com/neelance/graphql-go"
        )

        // pageInfoResolver Information for paginating this connection
        type pageInfoResolver struct {
          startCursor        *graphql.ID
          endCursor          *graphql.ID
          hasNextPage        bool
        }

        func (r *pageInfoResolver) StartCursor() *graphql.ID {
	        return r.startCursor
        }

        func (r *pageInfoResolver) EndCursor() *graphql.ID {
        	return r.endCursor
        }

        func (r *pageInfoResolver) HasNextPage() bool {
        	return r.hasNextPage
        }
        `,
			},
		},
	}

	for _, test := range tests {
		RunTest(test.schema, test.config, test.expected, t)
	}
}
