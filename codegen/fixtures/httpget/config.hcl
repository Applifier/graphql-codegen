package = "httpget"

type "User" {
}

type "Message" {
  field "from" {
    imports = ["\"fmt\""]
    template "http_resolver" {
      fields = [
        {
          field_name = "from_user_id"
          field_type = "string"
        }
      ]
      url = "fmt.Sprintf(\"https://static.everyplay.com/developer-quiz/data/users/%s\", r.{{.TypeName}}.From_user_id)"
    }
  }

  field "conversation" {
    imports = ["\"fmt\""]
    template "http_resolver" {
      fields = [
        {
          field_name = "conversation_id"
          field_type = "string"
        }
      ]
      url = "fmt.Sprintf(\"https://static.everyplay.com/developer-quiz/data/conversations/%s\", r.{{.TypeName}}.Conversation_id)"
    }
  }
}

type "Conversation" {
  field "with_user" {
    imports = ["\"fmt\""]
    template "http_resolver" {
      fields = [
        {
          field_name = "with_user_id"
          field_type = "string"
        }
      ]
      url = "fmt.Sprintf(\"https://static.everyplay.com/developer-quiz/data/users/%s\", r.{{.TypeName}}.With_user_id)"
    }
  }

  field "messages" {
    imports = ["\"fmt\""]
    template "http_resolver" {
      url = "fmt.Sprintf(\"https://static.everyplay.com/developer-quiz/data/conversations/%s/messages\", r.{{.TypeName}}.ID)"
    }
  }
}

type "Query" {
  field "user" {
    imports = ["\"fmt\""]
    template "http_resolver" {
      url = "fmt.Sprintf(\"https://static.everyplay.com/developer-quiz/data/users/%s\", args.ID)"
    }
  }

  field "conversation" {
    imports = ["\"fmt\""]
    template "http_resolver" {
      url = "fmt.Sprintf(\"https://static.everyplay.com/developer-quiz/data/conversations/%s\", args.ID)"
    }
  }

  field "conversations" {
    imports = ["\"fmt\""]
    template "http_resolver" {
      url = "\"https://static.everyplay.com/developer-quiz/data/conversations\""
    }
  }
}
