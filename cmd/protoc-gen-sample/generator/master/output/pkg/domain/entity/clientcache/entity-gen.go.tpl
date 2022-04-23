{{ template "autogen_comment" }}
package clientcache

import (
	"github.com/hikyaru-suzuki/go-con-2022-spring-sample/pkg/domain/proto/client/enums"
)

// {{ .Comment }}
type {{ .CamelName }} struct {
	{{- range .Columns }}
	// {{ .Comment }}
	{{ .CamelName }} {{ .Type }} `json:"{{ .SnakeName }}"`
	{{- end }}
}

type {{ .CamelName }}Slice []*{{ .CamelName }}
