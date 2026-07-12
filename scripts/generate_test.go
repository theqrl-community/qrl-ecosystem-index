package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestScreenshotsUnmarshalPreservesOrder(t *testing.T) {
	data := []byte(`
id: example-project
screenshots:
  - path: example-project/overview.webp
    caption: Overview screen
  - path: example-project/details.png
    caption: Detail screen
`)

	var project Project
	if err := yaml.Unmarshal(data, &project); err != nil {
		t.Fatalf("unmarshal project: %v", err)
	}
	if len(project.Screenshots) != 2 {
		t.Fatalf("got %d screenshots, want 2", len(project.Screenshots))
	}
	if project.Screenshots[0].Path != "example-project/overview.webp" || project.Screenshots[1].Caption != "Detail screen" {
		t.Fatalf("screenshot order or fields changed: %#v", project.Screenshots)
	}
}

func TestProjectPageContentIncludesScreenshots(t *testing.T) {
	params := map[string]interface{}{
		"title": "Example project",
		"screenshots": []Screenshot{
			{Path: "example-project/overview.webp", Caption: "Overview screen"},
			{Path: "example-project/details.png", Caption: "Detail screen"},
		},
	}

	content, err := projectPageContent(params, "Long description")
	if err != nil {
		t.Fatalf("generate project page content: %v", err)
	}
	text := string(content)
	if !strings.Contains(text, "path: example-project/overview.webp") || !strings.Contains(text, "caption: Detail screen") {
		t.Fatalf("generated content is missing screenshots:\n%s", text)
	}
	if !strings.Contains(text, "\n---\n\nLong description\n") {
		t.Fatalf("generated content is missing body: %s", text)
	}
}

func TestCopyAssetTreeCopiesFilesAndRemovesStaleOutput(t *testing.T) {
	root := t.TempDir()
	source := filepath.Join(root, "source")
	destination := filepath.Join(root, "destination")
	if err := os.MkdirAll(filepath.Join(source, "example-project"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(destination, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "example-project", "screen.webp"), []byte("image"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, ".gitkeep"), nil, 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(destination, "stale.webp"), []byte("stale"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := copyAssetTree(source, destination); err != nil {
		t.Fatalf("copy asset tree: %v", err)
	}
	if _, err := os.Stat(filepath.Join(destination, "example-project", "screen.webp")); err != nil {
		t.Fatalf("copied file missing: %v", err)
	}
	if _, err := os.Stat(filepath.Join(destination, "stale.webp")); !os.IsNotExist(err) {
		t.Fatalf("stale output was not removed: %v", err)
	}
	if _, err := os.Stat(filepath.Join(destination, ".gitkeep")); !os.IsNotExist(err) {
		t.Fatalf(".gitkeep should not be copied: %v", err)
	}
}

func TestCopyAssetTreeHandlesMissingSource(t *testing.T) {
	root := t.TempDir()
	destination := filepath.Join(root, "destination")
	if err := os.MkdirAll(destination, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(destination, "stale.webp"), []byte("stale"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := copyAssetTree(filepath.Join(root, "missing"), destination); err != nil {
		t.Fatalf("missing source should be accepted: %v", err)
	}
	if _, err := os.Stat(destination); !os.IsNotExist(err) {
		t.Fatalf("stale destination should be removed when source is missing: %v", err)
	}
}
