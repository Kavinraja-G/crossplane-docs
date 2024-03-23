package main

import (
	"encoding/json"
	"fmt"
	xv1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"sigs.k8s.io/yaml"
)

type CompositionOutputFormat struct {
	Name       string `json:"name,omitempty"`
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"api_version,omitempty"`
}

func main() {
	fmt.Println("Running Crossplane docs...")

	// read composition file
	rawComp, err := os.ReadFile("examples/sample-composition.yaml")
	if err != nil {
		fmt.Print("Error reading file...", err)
	}

	var comp xv1.Composition
	err = yaml.Unmarshal(rawComp, &comp)
	if err != nil {
		fmt.Print("Error during unmarshall...", err)
	}

	var resource xv1.ComposedTemplate
	var docOutputs []CompositionOutputFormat
	for _, resource = range comp.Spec.Resources {
		rawBase, err := extractRawBaseFromRawExtension(resource.Base)
		if err != nil {
			fmt.Println("Error extracting raw Base", err)
		}

		baseAPIVersion, _ := rawBase["apiVersion"].(string)
		baseKind, _ := rawBase["kind"].(string)

		docOutputs = append(docOutputs, CompositionOutputFormat{
			Name:       *resource.Name,
			APIVersion: baseAPIVersion,
			Kind:       baseKind,
		})
	}

	for _, docOutput := range docOutputs {
		fmt.Printf("Name: %s | APIVersion: %s | Kind: %s\n", docOutput.Name, docOutput.APIVersion, docOutput.Kind)
	}
}

func extractRawBaseFromRawExtension(rawExt runtime.RawExtension) (map[string]interface{}, error) {
	var rawMap map[string]interface{}
	if err := json.Unmarshal(rawExt.Raw, &rawMap); err != nil {
		return nil, err
	}

	return rawMap, nil
}
