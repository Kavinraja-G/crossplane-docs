package pkg

var markdownGenericTemplate = `
{{- range . }}
### {{ .CompositionName }}
| XR-Kind | XR-APIVersion |
|---------|-------------|
| {{ .XRKind }} | {{ .XRAPIVersion }} |

#### Resources
| Name | Kind | API Version |
|------|------|-------------|
{{- range .Resources }}
| {{ .Name }} | {{ .Kind }} | {{ .APIVersion }} |
{{- end }}
{{- end }}
`
