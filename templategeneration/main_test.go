package templategeneration

import (
	"os"
	"testing"

	"github.com/anz-bank/sysl/pkg/parse"

	"github.com/sirupsen/logrus"

	"github.com/spf13/afero"
)

func TestDataModel(t *testing.T) {
	t.Parallel()
	fs := afero.NewOsFs()
	filename := "simple.sysl"
	plantumlService := os.Getenv("SYSL_PLANTUML")
	if plantumlService == "" {
		panic("Error: Set SYSL_PLANTUML env variable")
	}
	m, err := parse.NewParser().Parse(filename, fs)
	if err != nil {
		panic(err)
	}
	NewProject(filename, "docs", plantumlService, logrus.New(), fs, m).
		ExecuteTemplateAndDiagrams()

}