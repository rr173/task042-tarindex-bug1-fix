package tarindex

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"testing"
)

func TestProbeEmptySearchItemsArray(t *testing.T) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}

	svc := NewService()
	summary, err := svc.Create(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := svc.Search(summary.ID, Filters{Name: "missing-*"})
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(data, []byte(`"items":[]`)) {
		t.Fatalf("empty search must encode items as [], got %s", data)
	}
}
