package main

import (
	"testing"

	helpers "github.com/bykaj/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/bykaj/k8s-crd-extractor:rolling")
	helpers.RequireFileExists(t, image, "/app/crd-extractor.sh")
	helpers.RequireFileExists(t, image, "/app/openapi2jsonschema.py")
}
