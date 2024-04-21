package pkg

import (
	"encoding/json"
	"os"
	"strings"
	"text/template"

	"github.com/Kavinraja-G/crossplane-docs/utils"

	"github.com/spf13/cobra"

	xv1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/strings/slices"
	"sigs.k8s.io/yaml"
)

// GenMarkdownDocs driver function for markdown sub-command
func GenMarkdownDocs(cmd *cobra.Command, searchPath string) error {
	outputFileName, err := cmd.Flags().GetString("output-file")
	if err != nil {
		return err
	}

	// find and discover XRDs in the provided search path
	discoveredXRDs, err := discoverXRDs(searchPath)
	if err != nil {
		return err
	}

	// find and discover compositions in the provided search path
	discoveredCompositions, err := discoverCompositions(searchPath)
	if err != nil {
		return err
	}

	// now iterate all discovered XRDs and build XRD output map with XRName as Key
	xrdOutputs := make(map[string]CompResourceDefinitionData)
	for _, xrd := range discoveredXRDs {
		openAPIV3Schema := extv1.JSONSchemaProps{}
		err = json.Unmarshal(xrd.Spec.Versions[0].Schema.OpenAPIV3Schema.Raw, &openAPIV3Schema)
		if err != nil {
			return err
		}

		var xrdOutputData []XRDSpecData
		getXRDSpecData(openAPIV3Schema, &xrdOutputData, []string{}, []string{})

		xrdOutputs[xrd.Spec.Names.Kind] = CompResourceDefinitionData{
			Name:                  xrd.Name,
			CompositeResourceKind: xrd.Spec.Names.Kind,
			ClaimNameKind:         xrd.Spec.ClaimNames.Kind,
			XRDSpec:               xrdOutputData,
		}
	}

	// now iterate all discovered compositions, link with respective XRDs and build the output data
	var mdOutputData []MarkdownOutputData
	for _, comp := range discoveredCompositions {
		resources, err := genericCompositionMode(comp.Spec.Resources)
		if err != nil {
			return err
		}

		mdOutputData = append(mdOutputData, MarkdownOutputData{
			CompositionName: comp.Name,
			XRAPIVersion:    comp.Spec.CompositeTypeRef.APIVersion,
			XRKind:          comp.Spec.CompositeTypeRef.Kind,
			Resources:       resources,
			LinkedXRDData:   xrdOutputs[comp.Spec.CompositeTypeRef.Kind],
		})
	}

	// generate markdown docs
	if err = outputMarkdownDocs(outputFileName, mdOutputData); err != nil {
		return err
	}

	return nil
}

// getXRDSpecData returns the specifications for the given XRD API
func getXRDSpecData(schema extv1.JSONSchemaProps, xrcOutputData *[]XRDSpecData, schemaPath []string, requiredFields []string) {
	for propName, propValue := range schema.Properties {
		*xrcOutputData = append(*xrcOutputData, XRDSpecData{
			FieldName:   propName,
			Path:        strings.Join(schemaPath, "."),
			Type:        propValue.Type,
			Description: propValue.Description,
			Required:    slices.Contains(requiredFields, propName),
		})
		if propValue.Properties != nil {
			getXRDSpecData(propValue, xrcOutputData, append(schemaPath, propName), propValue.Required)
		}
	}
}

// discoverCompositions returns compositions fetched from the target path
func discoverCompositions(searchPath string) ([]xv1.Composition, error) {
	yamlFiles, err := utils.FindYAMLFiles(searchPath)
	if err != nil {
		return nil, err
	}

	var discoveredCompositions []xv1.Composition
	for _, file := range yamlFiles {
		// read yamlFiles in the specified directory
		rawComp, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		// unmarshall and store yaml as discovered compositions
		var comp xv1.Composition
		err = yaml.Unmarshal(rawComp, &comp)
		if err != nil {
			return nil, err
		}

		// validate if it is a Composition by checking the Kind
		// support only if the composition mode is Resources
		kind := comp.Kind
		compMode := comp.Spec.Mode
		if (kind != "" && kind == xv1.CompositionKind) && (compMode == nil || *compMode == xv1.CompositionModeResources) {
			discoveredCompositions = append(discoveredCompositions, comp)
		}
	}

	return discoveredCompositions, nil
}

// discoverXRDs returns XRDs fetched from the target path
func discoverXRDs(searchPath string) ([]xv1.CompositeResourceDefinition, error) {
	// read yamlFiles in the specified directory
	yamlFiles, err := utils.FindYAMLFiles(searchPath)
	if err != nil {
		return nil, err
	}

	var discoveredXRDs []xv1.CompositeResourceDefinition
	for _, file := range yamlFiles {
		rawXRD, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		// unmarshall and store yaml as discovered XRDs
		var xrd xv1.CompositeResourceDefinition
		err = yaml.Unmarshal(rawXRD, &xrd)
		if err != nil {
			return nil, err
		}

		// validate if it is a XRD by checking the Kind
		kind := xrd.Kind
		if kind != "" && kind == xv1.CompositeResourceDefinitionKind {
			discoveredXRDs = append(discoveredXRDs, xrd)
		}
	}

	return discoveredXRDs, nil
}

// genericCompositionMode this handles logic for the Compositions without pipeline mode
func genericCompositionMode(resources []xv1.ComposedTemplate) ([]CompResourceData, error) {
	var resource xv1.ComposedTemplate
	var docOutputs []CompResourceData
	for _, resource = range resources {
		rawBase, err := extractRawBaseFromRawExtension(resource.Base)
		if err != nil {
			return nil, err
		}

		baseAPIVersion, _ := rawBase["apiVersion"].(string)
		baseKind, _ := rawBase["kind"].(string)

		docOutputs = append(docOutputs, CompResourceData{
			Name:       *resource.Name,
			APIVersion: baseAPIVersion,
			Kind:       baseKind,
		})
	}

	return docOutputs, nil
}

// extractRawBaseFromRawExtension unmarshalls the raw resources
func extractRawBaseFromRawExtension(rawExt runtime.RawExtension) (map[string]interface{}, error) {
	var rawMap map[string]interface{}
	if err := json.Unmarshal(rawExt.Raw, &rawMap); err != nil {
		return nil, err
	}

	return rawMap, nil
}

// outputMarkdownDocs generates Markdown docs file based on the discovered composition outputs
func outputMarkdownDocs(outputFileName string, mdOutputData []MarkdownOutputData) error {
	tmpl := template.New("md_docs")
	tmpl, err := tmpl.Parse(markdownGenericTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, mdOutputData)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
