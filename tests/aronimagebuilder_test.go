package tests

import (
	"aronimage/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AronImageBuilderSuite struct {
	builder *lib.AronImageBuilder
}

var aronImageBuilderSuite AronImageBuilderSuite

func init() {
	aronImageBuilderSuite.builder = &lib.AronImageBuilder{}
}

func TestWithBase64(t *testing.T) {
	base64String := "xxx"
	aronImageBuilderSuite.builder.WithBase64(base64String)

	assert.Equal(t, aronImageBuilderSuite.builder.Base64, base64String)
}

func TestWithFileName(t *testing.T) {
	fileName := "xxx"
	aronImageBuilderSuite.builder.WithFileName(fileName)

	assert.Equal(t, aronImageBuilderSuite.builder.FileName, fileName)
}

func TestWithExtension(t *testing.T) {
	extensionFile := "xxxxxx"
	aronImageBuilderSuite.builder.WithExtension(extensionFile)

	assert.Equal(t, aronImageBuilderSuite.builder.Extension, extensionFile)
}

func TestBuilAronImage(t *testing.T) {
	aronImage := aronImageBuilderSuite.builder.Build("xxxx")

	assert.IsType(t, aronImage, &lib.Aronimage{})
}
