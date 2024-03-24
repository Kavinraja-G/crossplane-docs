package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"

	xv1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

type compResourceData struct {
	Name       string
	Kind       string
	APIVersion string
}

type markdownOutputData struct {
	Title     string
	Overview  string
	Resources []compResourceData
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

	compMode := comp.Spec.Mode
	var resources []compResourceData
	if compMode != nil && *compMode == xv1.CompositionModePipeline {
		fmt.Println("Composition mode is Pipeline. Skipping...")
	} else {
		resources = genericCompositionMode(comp.Spec.Resources)
	}

	outputData := markdownOutputData{
		Title:     "My Crossplane Composition Documentation",
		Overview:  "This document provides an overview of my Crossplane composition.",
		Resources: resources,
	}

	tmpl := template.New("xdocs.tmpl")
	tmpl, err = tmpl.ParseFiles("xdocs.tmpl")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("output.md")
	if err != nil {
		fmt.Println("Error while creating markdown output", err)
	}

	err = tmpl.Execute(file, outputData)
	if err != nil {
		fmt.Println("Error while writing markdown output", err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func genericCompositionMode(resources []xv1.ComposedTemplate) []compResourceData {
	var resource xv1.ComposedTemplate
	var docOutputs []compResourceData
	for _, resource = range resources {
		rawBase, err := extractRawBaseFromRawExtension(resource.Base)
		if err != nil {
			fmt.Println("Error extracting raw Base", err)
		}

		baseAPIVersion, _ := rawBase["apiVersion"].(string)
		baseKind, _ := rawBase["kind"].(string)

		docOutputs = append(docOutputs, compResourceData{
			Name:       *resource.Name,
			APIVersion: baseAPIVersion,
			Kind:       baseKind,
		})
	}

	return docOutputs
}

func extractRawBaseFromRawExtension(rawExt runtime.RawExtension) (map[string]interface{}, error) {
	var rawMap map[string]interface{}
	if err := json.Unmarshal(rawExt.Raw, &rawMap); err != nil {
		return nil, err
	}

	return rawMap, nil
}

// markdownEscape escapes markdown special characters in a string.
func markdownEscape(input string) string {
	// List of markdown characters to escape
	// You might need to expand this list based on your requirements
	specialChars := []string{"|", "\\"}

	for _, char := range specialChars {
		input = strings.ReplaceAll(input, char, "\\"+char)
	}

	return input
}
