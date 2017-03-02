{{ define "arguments" }}args *struct{
  {{range .}}{{.Name | capitalize}} {{.Type}}
  {{end}}
}{{ end }}
{{define "receiver"}}{{if is_entry . }}Resolver{{else}}{{.}}Resolver{{end}}{{end}}
{{if eq .TypeKind "OBJECT"}}
{{$hasArguments := gt (.MethodArguments | len) 0}}
// {{capitalize .MethodName}} {{.MethodDescription}}
func (r *{{template "receiver" .TypeName}}) {{capitalize .MethodName}}({{if $hasArguments}}{{template "arguments" .MethodArguments}}{{end}}) {{.MethodReturnType}} {
  {{if is_entry .TypeName}}return nil{{else}}return r.{{.TypeName}}.{{capitalize .MethodReturn}}{{end}}
}
{{end}}
{{if eq .TypeKind "INTERFACE"}}
{{$hasArguments := gt (.MethodArguments | len) 0}}
// {{capitalize .MethodName}} {{.MethodDescription}}
{{capitalize .MethodName}}({{if $hasArguments}}{{template "arguments" .MethodArguments}}{{end}}) {{.MethodReturnType}}
{{end}}
