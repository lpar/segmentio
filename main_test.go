package main

import (
	"compress/gzip"
	_ "embed"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
	segment "github.com/segmentio/encoding/json"
)

var testData []byte

func TestMain(m *testing.M) {
	testData = loadTestData("testdata/code.json.gz")
	code := m.Run()
	os.Exit(code)
}

func loadTestData(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(gz)
	if err != nil {
		panic(err)
	}
	return data
}

func BenchmarkUnmarshalStdlib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var obj interface{}
		err := json.Unmarshal(testData, &obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalJSONiter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var obj interface{}
		err := jsoniter.Unmarshal(testData, &obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalSegmentIO(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var obj interface{}
		err := segment.Unmarshal(testData, &obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalStdlib(b *testing.B) {
	var obj interface{}
	err := json.Unmarshal(testData, &obj)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalJSONiter(b *testing.B) {
	var obj interface{}
	err := jsoniter.Unmarshal(testData, &obj)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := jsoniter.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalSegmentIO(b *testing.B) {
	var obj interface{}
	err := jsoniter.Unmarshal(testData, &obj)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := segment.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}
