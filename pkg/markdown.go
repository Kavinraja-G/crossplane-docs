package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/Kavinraja-G/crossplane-docs/utils"

	"github.com/spf13/cobra"

	xv1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func GenMarkdownDocs(cmd *cobra.Command, searchPath string) error {
	outputFileName, err := cmd.Flags().GetString("output-file")
	if err != nil {
		return err
	}

	// find and discover compositions in the provided search path
	discoveredCompositions, err := findAndDiscoverCompositions(searchPath)
	if err != nil {
		return err
	}

	// now iterate all discovered compositions and build the output data
	// TODO: currently, we skip if the composition is based on pipeline. Need to enable support for them
	var mdOutputData []MarkdownOutputData
	for _, comp := range discoveredCompositions {
		compMode := comp.Spec.Mode
		var resources []CompResourceData
		if compMode != nil && *compMode == xv1.CompositionModePipeline {
			fmt.Printf("Composition: %s is in pipeline mode. Skipping...\n", comp.Name)
		} else {
			resources, err = genericCompositionMode(comp.Spec.Resources)
			if err != nil {
				return err
			}
		}
		mdOutputData = append(mdOutputData, MarkdownOutputData{
			CompositionName: comp.Name,
			XRAPIVersion:    comp.Spec.CompositeTypeRef.APIVersion,
			XRKind:          comp.Spec.CompositeTypeRef.Kind,
			Resources:       resources,
		})
	}

	// generate markdown docs
	if err = outputMarkdownDocs(outputFileName, mdOutputData); err != nil {
		return err
	}

	return nil
}

// findAndDiscoverCompositions returns compositions fetched from the target path
func findAndDiscoverCompositions(searchPath string) ([]xv1.Composition, error) {
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
		discoveredCompositions = append(discoveredCompositions, comp)
	}

	return discoveredCompositions, nil
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
