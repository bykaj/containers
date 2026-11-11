package main

import (
	"testing"

	helpers "github.com/bykaj/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/bykaj/radicale:rolling")
	helpers.RequireHTTPEndpoint(t, image, helpers.HTTPTestConfig{Port: "5232"}, nil)
}
