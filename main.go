package main

import (
	"encoding/json"
	"fmt"
	"github.com/kubescape/storage/pkg/apis/softwarecomposition/v1beta1"
	"github.com/spf13/afero"
	"io/fs"
	"path/filepath"
)

func main() {
	appFs := afero.NewOsFs()
	dataDir := "data/spdx.softwarecomposition.kubescape.io/sbomsyft/"
	err := filepath.Walk(dataDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		}
		if filepath.Ext(path) != ".j" || info.Size() == 0 {
			return nil
		}
		fmt.Printf("opening file: %q\n", path)
		file, err := appFs.Open(path)
		if err != nil {
			return fmt.Errorf("error opening file %q: %v", path, err)
		}
		decoder := json.NewDecoder(file)
		var objPtr v1beta1.SBOMSyft
		err = decoder.Decode(&objPtr)
		if err != nil {
			return fmt.Errorf("error decoding file %q: %v", path, err)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dataDir, err)
		return
	}
}
