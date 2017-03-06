[![Build Status](https://travis-ci.com/Applifier/graphql-codegen.svg?token=jeWt6weUpeDp6aNSSaST&branch=master)](https://travis-ci.com/Applifier/graphql-codegen)

# graphql-codegen

This tool generates boilerplate code for [graphql-go](https://github.com/neelance/graphql-go) based on a provided graphql.schema file. Additional config file can be provided to customize output.

```sh
go install github.com/Applifier/graphql-codegen
```

## example

Example below generates code for [schema.graphql](https://github.com/Applifier/graphql-codegen/blob/master/codegen/fixtures/httpget/schema.graphql). Optional config [config.hcl](https://github.com/Applifier/graphql-codegen/blob/master/codegen/fixtures/httpget/config.hcl) file is provided to tell the codegen where and how to fetch the data for the resolvers.

```sh
graphql-codegen generate -s=codegen/fixtures/httpget/schema.graphql -c=codegen/fixtures/httpget/config.hcl -p=httpget -o=test_output/
```
