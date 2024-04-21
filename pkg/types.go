package pkg

type CompResourceData struct {
	Name       string
	Kind       string
	APIVersion string
}

// XRDSpecData struct for API specs in the output
type XRDSpecData struct {
	FieldName   string
	Path        string
	Type        string
	Description string
	Required    bool
}

// XRDVersion used for representing individual API versions
type XRDVersion struct {
	Version string
	XRDSpec []XRDSpecData
}

// CompResourceDefinitionData defines the XRD data and its XR, XRC details if any
type CompResourceDefinitionData struct {
	Name                  string
	CompositeResourceKind string
	ClaimNameKind         string
	Versions              []XRDVersion
}

// MarkdownOutputData output data struct used by markdown generator
type MarkdownOutputData struct {
	CompositionName string
	XRAPIVersion    string
	XRKind          string
	Resources       []CompResourceData
	LinkedXRD       CompResourceDefinitionData
}
