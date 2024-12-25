package pkg

type OutputOpts struct {
	OutputFileName string
	OutputTemplate string
	PrintXRDOnly   bool
}

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
	Default     string
	Required    bool
}

// XRDVersion used for representing individual API versions
type XRDVersion struct {
	Version       string
	XRDInputSpec  []XRDSpecData
	XRDOutputSpec []XRDSpecData
}

// CompResourceDefinitionData defines the XRD data and its XR, XRC details if any
type CompResourceDefinitionData struct {
	Name                  string
	CompositeResourceKind string
	ClaimNameKind         string
	Versions              []XRDVersion
}

// GenericOutputData output data struct used by markdown generator
type GenericOutputData struct {
	CompositionName string
	XRAPIVersion    string
	XRKind          string
	Resources       []CompResourceData
	LinkedXRD       CompResourceDefinitionData
}
