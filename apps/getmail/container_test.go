package main

import (
	"testing"

	helpers "github.com/bykaj/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/bykaj/getmail:rolling")
	helpers.RequireCommandSucceeds(t, image, nil, "/usr/local/bin/getmail", "--version")
}
