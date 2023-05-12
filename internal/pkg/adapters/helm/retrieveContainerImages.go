package helm

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"helm.sh/helm/v3/pkg/chart/loader"
)

func (a *adapter) RetrieveContainerImages(repositoryDirectory string) (*[]string, error) {
	var images []string
	err := filepath.Walk(repositoryDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("path: %v\n", path)

		chart, err := loader.LoadDir(path)
		if err != nil {
			return err
		}

		fmt.Printf("chart: %v\n", chart)
		if info.IsDir() || filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
			fmt.Println("Skipping file: " + path)
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return err
		}

		pattern := `(image|value):\s*("[^"]+"|'[^']+'|[^#'" ]+)`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(string(content), -1)

		for _, match := range matches {
			image := strings.Trim(match[len(match)-1], "\"'")
			images = append(images, image)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &images, nil
}
