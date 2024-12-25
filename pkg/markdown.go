package pkg

import (
	"encoding/json"
	"fmt"
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

const xrdOutputFieldName = "status"

// GenMarkdownDocs driver function for markdown sub-command
func GenMarkdownDocs(cmd *cobra.Command, searchPath string) error {
	outputOpts, err := getOutputOpts(cmd)
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

	// now iterate all discovered XRDs and build XRD output map with XR & API spec details
	var xrdOutputs []CompResourceDefinitionData
	for _, xrd := range discoveredXRDs {
		var xrdVersions []XRDVersion
		for _, version := range xrd.Spec.Versions {
			openAPIV3Schema, err := getOpenAPIV3Schema(version.Schema.OpenAPIV3Schema.Raw)
			if err != nil {
				return err
			}

			var xrdInputData []XRDSpecData
			var xrdOutputData []XRDSpecData
			getXRDSpecData(openAPIV3Schema, &xrdInputData, &xrdOutputData, []string{}, []string{})
			xrdVersions = append(xrdVersions, XRDVersion{
				Version:       version.Name,
				XRDInputSpec:  xrdInputData,
				XRDOutputSpec: xrdOutputData,
			})
		}

		xrdOutputs = append(xrdOutputs, CompResourceDefinitionData{
			Name:                  xrd.Name,
			CompositeResourceKind: xrd.Spec.Names.Kind,
			ClaimNameKind:         xrd.Spec.ClaimNames.Kind,
			Versions:              xrdVersions,
		})
	}

	// handle only XRDs if set to true
	if outputOpts.PrintXRDOnly {
		if err = outputMarkdownDocs(outputOpts, xrdOutputs); err != nil {
			return err
		}
		return nil
	}

	// now iterate all discovered compositions, link with respective XRDs and build the output data
	var mdOutputData []GenericOutputData
	for _, comp := range discoveredCompositions {
		resources, err := genericCompositionMode(comp.Spec.Resources)
		if err != nil {
			return err
		}

		linkedXRD, found := findXRDByKind(xrdOutputs, comp.Spec.CompositeTypeRef.Kind)
		if !found {
			fmt.Errorf("could not find XRD for composite kind %s", comp.Spec.CompositeTypeRef.Kind)
		}

		mdOutputData = append(mdOutputData, GenericOutputData{
			CompositionName: comp.Name,
			XRAPIVersion:    comp.Spec.CompositeTypeRef.APIVersion,
			XRKind:          comp.Spec.CompositeTypeRef.Kind,
			Resources:       resources,
			LinkedXRD:       linkedXRD,
		})
	}

	// output generic markdown docs
	if err = outputMarkdownDocs(outputOpts, mdOutputData); err != nil {
		return err
	}

	return nil
}

// getOutputOpts builds the output options for the generator
func getOutputOpts(cmd *cobra.Command) (OutputOpts, error) {
	outFile, err := cmd.Flags().GetString("output-file")
	if err != nil {
		return OutputOpts{}, err
	}
	xrdOnly, err := cmd.Flags().GetBool("xrd-only")
	if err != nil {
		return OutputOpts{}, err
	}

	outputTemplate := markdownGenericTemplate
	if xrdOnly {
		outputTemplate = markdownXRDOnlyTemplate
	}

	return OutputOpts{OutputFileName: outFile, OutputTemplate: outputTemplate, PrintXRDOnly: xrdOnly}, nil
}

// findXRDByKind checks if the given slice contains XRDs mapped to the specific kind
func findXRDByKind(xrdOutputs []CompResourceDefinitionData, kind string) (CompResourceDefinitionData, bool) {
	for _, xrd := range xrdOutputs {
		if xrd.CompositeResourceKind == kind {
			return xrd, true
		}
	}
	return CompResourceDefinitionData{}, false
}

// getOpenAPIV3Schema parses the XRDs raw API schema as JSONSchemaProps
func getOpenAPIV3Schema(rawSchemaData []byte) (extv1.JSONSchemaProps, error) {
	openAPIV3Schema := extv1.JSONSchemaProps{}
	err := json.Unmarshal(rawSchemaData, &openAPIV3Schema)
	if err != nil {
		return openAPIV3Schema, err
	}

	return openAPIV3Schema, nil
}

// getXRDSpecData returns the specifications for the given XRD API
func getXRDSpecData(schema extv1.JSONSchemaProps, xrdInputData *[]XRDSpecData, xrdOutputData *[]XRDSpecData, schemaPath []string, requiredFields []string) {
	for propName, propValue := range schema.Properties {
		fullPath := append(schemaPath, propName)

		specData := XRDSpecData{
			FieldName:   propName,
			Path:        strings.Join(schemaPath, "."),
			Type:        propValue.Type,
			Description: propValue.Description,
			Required:    slices.Contains(requiredFields, propName),
			Default: func() string {
				if propValue.Default != nil {
					return string(propValue.Default.Raw)
				}
				return "n/a"
			}(),
		}

		// store XRD output specs separately
		if propName == xrdOutputFieldName || slices.Contains(schemaPath, xrdOutputFieldName) {
			*xrdOutputData = append(*xrdOutputData, specData)
		} else {
			*xrdInputData = append(*xrdInputData, specData)
		}

		// check if this is an array of objects
		if propValue.Type == "array" && propValue.Items != nil {
			// handle items recursively if they are objects
			itemSchema := propValue.Items.Schema
			if itemSchema.Type == "object" {
				getXRDSpecData(*itemSchema, xrdInputData, xrdOutputData, fullPath, itemSchema.Required)
			}
			continue
		}

		if propValue.Type == "object" && propValue.Properties != nil {
			// recursively iterate all the nested properties
			getXRDSpecData(propValue, xrdInputData, xrdOutputData, fullPath, propValue.Required)
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
func outputMarkdownDocs(opts OutputOpts, mdOutputData interface{}) error {
	tmpl, err := template.New("md_docs").Parse(opts.OutputTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(opts.OutputFileName)
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
