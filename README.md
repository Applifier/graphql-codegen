[![Build Status](https://travis-ci.com/Applifier/graphql-codegen.svg?token=jeWt6weUpeDp6aNSSaST&branch=master)](https://travis-ci.com/Applifier/graphql-codegen)

# graphql-codegen

This tool generates boilerplate code for [graphql-go](https://github.com/graph-gophers/graphql-go) based on a provided graphql.schema file. Additional config file can be provided to customize output.

```sh
go install github.com/Applifier/graphql-codegen
```

## TODO
- [ ] Allow custom templates (without rebuilding)

## example

Example below generates code for [schema.graphql](https://github.com/Applifier/graphql-codegen/blob/master/codegen/fixtures/httpget/schema.graphql). Optional config [config.hcl](https://github.com/Applifier/graphql-codegen/blob/master/codegen/fixtures/httpget/config.hcl) file is provided to tell the codegen where and how to fetch the data for the resolvers.

```sh
graphql-codegen generate -s=codegen/fixtures/httpget/schema.graphql -c=codegen/fixtures/httpget/config.hcl -p=httpget -o=test_output/
```

Example of the generated code (_gen.go files) can be found under [/codegen/fixtures/httpget](https://github.com/Applifier/graphql-codegen/tree/master/codegen/fixtures/httpget)

- More examples under [codegen/fixtures](https://github.com/Applifier/graphql-codegen/tree/master/codegen/fixtures)

## templates

### default
Resolve field based on a struct property with the same name as the schema type field

### custom
Skips code generation for the field

Check example [config](https://github.com/Applifier/graphql-codegen/blob/master/codegen/fixtures/httpget/config.hcl#L35) and [conversation.go](https://github.com/Applifier/graphql-codegen/blob/master/codegen/fixtures/httpget/coversation.go#L12)

### http_resolver
Resolve field with a http GET request to a server
```hcl
type "Query" {
  field "user" {
    imports = ["\"fmt\""]
    template "http_resolver" {
      url = "fmt.Sprintf(\"https://static.everyplay.com/developer-quiz/data/users/%s\", args.ID)"
    }
  }
}
```
