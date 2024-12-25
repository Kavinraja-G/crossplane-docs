package pkg

var markdownGenericTemplate = `
{{- range . }}
### {{ .CompositionName }}
| XR Kind | XR Version |
|---------|-------------|
| {{ .XRKind }} | {{ .XRAPIVersion }} |
#### XRD
| Name | Claim |
|------|-------|
| {{ .LinkedXRD.Name }} | {{ .LinkedXRD.ClaimNameKind }} |
#### XRD Spec
{{- range .LinkedXRD.Versions }}
##### Version: {{ .Version }}
| Field | Path | Type | Description | Default | Required |
|------|-------|------|-------|-------|-------|
{{- range .XRDSpec }}
| {{ .FieldName }} | {{ .Path }} | {{ .Type }} | {{ .Description }} | {{ .Default }} | {{ .Required }} |
{{- end }}
{{- end }}
#### Resources
| Name | Kind | API Version |
|------|------|-------------|
{{- range .Resources }}
| {{ .Name }} | {{ .Kind }} | {{ .APIVersion }} |
{{- end }}
{{- end }}
`

var markdownXRDOnlyTemplate = `
{{- range . }}
### {{ .Name }}
| Name | Claim |
|------|-------|
| {{ .Name }} | {{ .ClaimNameKind }} |
#### Specs
{{- range .Versions }}
##### Version: {{ .Version }}
| Field | Path | Type | Description | Default |  Required |
|------|-------|------|-------|-------|-------|
{{- range .XRDSpec }}
| {{ .FieldName }} | {{ .Path }} | {{ .Type }} | {{ .Description }} | {{ .Default }} | {{ .Required }} |
{{- end }}
{{- end }}
{{- end }}
`
