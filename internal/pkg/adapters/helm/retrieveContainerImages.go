package helm

func (a *adapter) RetrieveContainerImages(repositoryDirectory string) (*[]string, error) {
	return nil, nil
}

// func (a *adapter) RetrieveContainerImages(repositoryDirectory string) (*[]string, error) {
// 	var images []string
// 	err := filepath.Walk(repositoryDirectory, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}

// 		fmt.Printf("path: %v\n", path)

// 		if filepath.Ext(path) == ".tgz" {
// 			a.viewChartFiles(path)
// 			return nil
// 		}

// 		if info.IsDir() || filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
// 			fmt.Println("Skipping file: " + path)
// 			return nil
// 		}

// 		content, err := os.ReadFile(path)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}

// 		pattern := `(image|value):\s*("[^"]+"|'[^']+'|[^#'" ]+)`
// 		re := regexp.MustCompile(pattern)
// 		matches := re.FindAllStringSubmatch(string(content), -1)

// 		for _, match := range matches {
// 			image := strings.Trim(match[len(match)-1], "\"'")
// 			images = append(images, image)
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	return &images, nil
// }

// func (a *adapter) viewChartFiles(chartPath string) {

// 	// Create a temporary directory to unpack the chart
// 	tempDir, err := ioutil.TempDir("", "helm-chart")
// 	if err != nil {
// 		fmt.Printf("Failed to create temporary directory: %v\n", err)
// 		return
// 	}
// 	defer os.RemoveAll(tempDir)

// 	// Unpack the Helm chart into the temporary directory
// 	if err := unpackChart(chartPath, tempDir); err != nil {
// 		fmt.Printf("Failed to unpack Helm chart: %v\n", err)
// 		return
// 	}

// 	// Retrieve container images mentioned in the values file
// 	images, err := retrieveContainerImages(filepath.Join(tempDir, "values.yaml"))
// 	if err != nil {
// 		fmt.Printf("Failed to retrieve container images: %v\n", err)
// 		return
// 	}

// 	// Print the retrieved container images
// 	fmt.Println("Container Images:")
// 	for _, image := range images {
// 		fmt.Println(image)
// 	}
// }

// func unpackChart(chartPath, destinationDir string) error {
// 	cmd := exec.Command("tar", "xzf", chartPath, "-C", destinationDir)
// 	err := cmd.Run()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // retrieveContainerImages retrieves container images mentioned in the values file.
// func retrieveContainerImages(valuesPath string) ([]string, error) {
// 	var images []string

// 	// Read the contents of the values file
// 	content, err := os.ReadFile(valuesPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Find container image references using a regular expression
// 	// This pattern matches strings starting with "image:" followed by a valid container image format
// 	pattern := `image:\s*("[^"]+"|'[^']+'|[^#'" ]+)`
// 	re := regexp.MustCompile(pattern)
// 	matches := re.FindAllStringSubmatch(string(content), -1)

// 	// Extract container images from the matches
// 	for _, match := range matches {
// 		image := strings.Trim(match[len(match)-1], "\"'")
// 		images = append(images, image)
// 	}

// 	return images, nil
// }
