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
	XRDSpec               []XRDSpecData
}

type XRDSpecData struct {
	FieldName   string
	Path        string
	Type        string
	Description string
	Required    bool
}

// MarkdownOutputData output data struct used by markdown generator
type MarkdownOutputData struct {
	CompositionName string
	XRAPIVersion    string
	XRKind          string
	Resources       []CompResourceData
	LinkedXRDData   CompResourceDefinitionData
}
