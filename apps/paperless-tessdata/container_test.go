package main

import (
	"os"
	"testing"

	"github.com/moby/moby/client"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"

	helpers "github.com/bykaj/containers/tests"
)

// The image is FROM scratch: nothing can run in it, so check its filesystem directly.
func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/bykaj/paperless-tessdata:rolling")

	for _, lang := range []string{"deu", "eng", "fra", "ita", "nld", "osd", "spa"} {
		requireFileMode(t, image, "/tessdata/"+lang+".traineddata", 0o444)
	}

	for _, path := range []string{
		"/tessdata/configs/hocr",
		"/tessdata/configs/pdf",
		"/tessdata/configs/txt",
		"/tessdata/tessconfigs/batch",
		"/tessdata/pdf.ttf",
	} {
		helpers.RequireFileExists(t, image, path)
	}
}

// requireFileMode asserts that a regular file exists in the image with the given permissions.
func requireFileMode(t *testing.T, image, path string, perm os.FileMode) {
	t.Helper()

	ctx := t.Context()

	ctr, err := testcontainers.Run(ctx, image,
		testcontainers.WithNoStart(),
		testcontainers.WithCmd("/"),
	)
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	cli, err := testcontainers.NewDockerClientWithOpts(ctx)
	require.NoError(t, err)
	defer cli.Close()

	res, err := cli.ContainerStatPath(ctx, ctr.GetContainerID(), client.ContainerStatPathOptions{Path: path})
	require.NoError(t, err, "file %q should exist in image %q", path, image)
	require.True(t, res.Stat.Mode.IsRegular(), "%q should be a regular file", path)
	require.Equal(t, perm, res.Stat.Mode.Perm(), "%q permissions", path)
}
