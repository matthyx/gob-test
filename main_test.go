package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/kubescape/storage/pkg/apis/softwarecomposition/v1beta1"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	huge   = "data/spdx.softwarecomposition.kubescape.io/sbomsyft/kubescape/apachepulsar-pulsar-all-3.0.3-9b50a3"
	big    = "data/spdx.softwarecomposition.kubescape.io/sbomsyft/kubescape/quay.io-argoproj-argocd-v2.10.6-012b98"
	medium = "data/spdx.softwarecomposition.kubescape.io/sbomsyft/kubescape/quay.io-kubescape-kubescape-v3.0.9-4441f0"
	small  = "data/spdx.softwarecomposition.kubescape.io/sbomsyft/kubescape/docker.io-memcached-1.5.17-alpine-eb5f68"
)

var current = small

func TestConvert(t *testing.T) {
	appFs := afero.NewOsFs()
	file, err := appFs.Open(current + ".j")
	require.NoError(t, err)
	decoder := json.NewDecoder(file)
	var objPtr v1beta1.SBOMSyft
	err = decoder.Decode(&objPtr)
	require.NoError(t, err)
	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)
	err = encoder.Encode(objPtr)
	require.NoError(t, err)
	err = afero.WriteFile(appFs, current+".g", b.Bytes(), 0644)
	require.NoError(t, err)
}

func TestJsonDecode(t *testing.T) {
	appFs := afero.NewOsFs()
	file, err := appFs.Open(current + ".j")
	require.NoError(t, err)
	decoder := json.NewDecoder(file)
	var objPtr v1beta1.SBOMSyft
	err = decoder.Decode(&objPtr)
	require.NoError(t, err)
}

func TestGobDecode(t *testing.T) {
	appFs := afero.NewOsFs()
	file, err := appFs.Open(current + ".g")
	require.NoError(t, err)
	decoder := gob.NewDecoder(file)
	var objPtr v1beta1.SBOMSyft
	err = decoder.Decode(&objPtr)
	require.NoError(t, err)
}

func BenchmarkJsonDecode(b *testing.B) {
	appFs := afero.NewOsFs()
	for i := 0; i < b.N; i++ {
		file, err := appFs.Open(current + ".j")
		require.NoError(b, err)
		decoder := json.NewDecoder(file)
		var objPtr v1beta1.SBOMSyft
		err = decoder.Decode(&objPtr)
		require.NoError(b, err)
	}
	b.ReportAllocs()
}

func BenchmarkGobDecode(b *testing.B) {
	appFs := afero.NewOsFs()
	for i := 0; i < b.N; i++ {
		file, err := appFs.Open(current + ".g")
		require.NoError(b, err)
		decoder := gob.NewDecoder(file)
		var objPtr v1beta1.SBOMSyft
		err = decoder.Decode(&objPtr)
		require.NoError(b, err)
	}
	b.ReportAllocs()
}
