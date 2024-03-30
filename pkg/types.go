package pkg

type CompResourceData struct {
	Name       string
	Kind       string
	APIVersion string
}

type CompResourceDefinitionData struct {
	Name                  string
	CompositeResourceKind string
	ClaimNameKind         string
}

// MarkdownOutputData output data struct used by markdown generator
type MarkdownOutputData struct {
	CompositionName string
	XRAPIVersion    string
	XRKind          string
	Resources       []CompResourceData
	LinkedXRDData   CompResourceDefinitionData
}
