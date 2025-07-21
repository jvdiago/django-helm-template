package djangoapp

import (
	"embed"
	"io/fs"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

//go:embed Chart.yaml
//go:embed values.yaml
//go:embed templates
//go:embed templates/_helpers.tpl
var chartFS embed.FS

// Chart loads the embedded chart from chartFS.
func Chart() (*chart.Chart, error) {
	var files []*loader.BufferedFile

	// Walk the embedded FS and buffer every file
	if err := fs.WalkDir(chartFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := chartFS.ReadFile(path)
		if err != nil {
			return err
		}
		files = append(files, &loader.BufferedFile{
			Name: path,
			Data: data,
		})
		return nil
	}); err != nil {
		return nil, err
	}

	// Load everything as a chart
	return loader.LoadFiles(files)
}
