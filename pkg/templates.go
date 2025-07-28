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
##### Inputs
| Field | Path | Type | Description | Default | Required |
|------|-------|------|-------|-------|-------|
{{- range .XRDInputSpec }}
| {{ .FieldName }} | {{ .Path }} | {{ .Type }} | {{ .Description }} | {{ .Default }} | {{ .Required }} |
{{- end }}
##### Outputs
| Field | Path | Type | Description | Default | Required |
|------|-------|------|-------|-------|-------|
{{- range .XRDOutputSpec }}
| {{ .FieldName }} | {{ .Path }} | {{ .Type }} | {{ .Description }} | {{ .Default }} | {{ .Required }} |
{{- end }}
{{- end }}
{{- if .Resources }}
#### Resources
| Name | Kind | API Version |
|------|------|-------------|
{{- range .Resources }}
| {{ .Name }} | {{ .Kind }} | {{ .APIVersion }} |
{{- end }}
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
##### Inputs
| Field | Path | Type | Description | Default |  Required |
|------|-------|------|-------|-------|-------|
{{- range .XRDInputSpec }}
| {{ .FieldName }} | {{ .Path }} | {{ .Type }} | {{ .Description }} | {{ .Default }} | {{ .Required }} |
{{- end }}
##### Outputs
| Field | Path | Type | Description | Default |  Required |
|------|-------|------|-------|-------|-------|
{{- range .XRDOutputSpec }}
| {{ .FieldName }} | {{ .Path }} | {{ .Type }} | {{ .Description }} | {{ .Default }} | {{ .Required }} |
{{- end }}
{{- end }}
{{- end }}
`
