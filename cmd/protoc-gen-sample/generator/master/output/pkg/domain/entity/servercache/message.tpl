{{- define "message" }}
// {{ .GoName }} {{ .Comment }}
type {{ .GoName }} struct {
{{- range .Columns }}
    // {{ .Comment }}
    {{ .GoName }} {{ .Type }}
{{- end }}
}

type {{ .GoName }}Slice []*{{ .GoName }}

{{ range .Messages -}}
  {{- template "message" . }}
{{- end -}}
{{ end -}}
