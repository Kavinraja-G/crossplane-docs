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
| {{ .LinkedXRDData.Name }} | {{ .LinkedXRDData.ClaimNameKind }} |
#### XRD Spec
| Field | Path | Type | Description | Required |
|------|-------|------|-------|-------|
{{- range .LinkedXRDData.XRDSpec }}
| {{ .FieldName }} | {{ .Path }} | {{ .Type }} | {{ .Description }} | {{ .Required }} |
{{- end }}
#### Resources
| Name | Kind | API Version |
|------|------|-------------|
{{- range .Resources }}
| {{ .Name }} | {{ .Kind }} | {{ .APIVersion }} |
{{- end }}
{{- end }}
`
