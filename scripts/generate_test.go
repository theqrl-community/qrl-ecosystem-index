package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestGalleryUnmarshalPreservesOrder(t *testing.T) {
	data := []byte(`
id: example-project
gallery:
  - type: youtube
    id: M7lc1UVf-VE
    caption: Project demonstration
  - type: image
    path: example-project/details.png
    caption: Detail screen
`)

	var project Project
	if err := yaml.Unmarshal(data, &project); err != nil {
		t.Fatalf("unmarshal project: %v", err)
	}
	if len(project.Gallery) != 2 {
		t.Fatalf("got %d gallery items, want 2", len(project.Gallery))
	}
	if project.Gallery[0].Type != "youtube" || project.Gallery[0].ID != "M7lc1UVf-VE" || project.Gallery[1].Path != "example-project/details.png" {
		t.Fatalf("gallery order or fields changed: %#v", project.Gallery)
	}
}

func TestProjectPageContentIncludesGallery(t *testing.T) {
	params := map[string]interface{}{
		"title": "Example project",
		"gallery": []GalleryItem{
			{Type: "youtube", ID: "M7lc1UVf-VE", Caption: "Project demonstration"},
			{Type: "image", Path: "example-project/details.png", Caption: "Detail screen"},
		},
	}

	content, err := projectPageContent(params, "Long description")
	if err != nil {
		t.Fatalf("generate project page content: %v", err)
	}
	text := string(content)
	if !strings.Contains(text, "id: M7lc1UVf-VE") || !strings.Contains(text, "path: example-project/details.png") || !strings.Contains(text, "caption: Detail screen") {
		t.Fatalf("generated content is missing gallery items:\n%s", text)
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

func TestGenerateSocialCardsWithAndWithoutProjectMedia(t *testing.T) {
	root := t.TempDir()
	assetRoot := filepath.Join(root, "images")
	outputRoot := filepath.Join(root, "og")
	if err := os.MkdirAll(filepath.Join(assetRoot, "screenshots", "with-media"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(assetRoot, "logos", "with-media"), 0755); err != nil {
		t.Fatal(err)
	}

	screenshotPath := filepath.Join(assetRoot, "screenshots", "with-media", "screen.png")
	screenshotFile, err := os.Create(screenshotPath)
	if err != nil {
		t.Fatal(err)
	}
	screenshot := image.NewRGBA(image.Rect(0, 0, 320, 180))
	for y := 0; y < 180; y++ {
		for x := 0; x < 320; x++ {
			screenshot.Set(x, y, color.RGBA{R: uint8(x % 255), G: uint8(y % 255), B: 80, A: 255})
		}
	}
	if err := png.Encode(screenshotFile, screenshot); err != nil {
		t.Fatal(err)
	}
	if err := screenshotFile.Close(); err != nil {
		t.Fatal(err)
	}

	logo := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><circle cx="50" cy="50" r="45" fill="#2a8e9a"/></svg>`
	if err := os.WriteFile(filepath.Join(assetRoot, "logos", "with-media", "icon.svg"), []byte(logo), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(outputRoot, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(outputRoot, "stale.png"), []byte("stale"), 0644); err != nil {
		t.Fatal(err)
	}

	projects := []Project{
		{
			ID:          "with-media",
			Name:        "Project With Media",
			ProjectType: "dapp",
			Description: "A project whose generated card uses its first gallery image and SVG logo.",
			Logos:       []Logo{{Path: "with-media/icon.svg"}},
			Gallery: []GalleryItem{
				{Type: "youtube", ID: "M7lc1UVf-VE", Caption: "Project video"},
				{Type: "image", Path: "with-media/screen.png", Caption: "Project screen"},
			},
		},
		{
			ID:          "video-only",
			Name:        "Video Only Project",
			ProjectType: "community",
			Description: "A video-only project whose generated card uses the branded initials treatment.",
			Gallery:     []GalleryItem{{Type: "youtube", ID: "M7lc1UVf-VE", Caption: "Project video"}},
		},
	}

	if err := generateSocialCards(projects, assetRoot, outputRoot); err != nil {
		t.Fatalf("generate social cards: %v", err)
	}
	for _, path := range []string{
		filepath.Join(outputRoot, "default.png"),
		filepath.Join(outputRoot, "projects", "with-media.png"),
		filepath.Join(outputRoot, "projects", "video-only.png"),
	} {
		file, err := os.Open(path)
		if err != nil {
			t.Fatalf("generated card missing: %v", err)
		}
		config, err := png.DecodeConfig(file)
		file.Close()
		if err != nil {
			t.Fatalf("decode generated card %s: %v", path, err)
		}
		if config.Width != socialCardWidth || config.Height != socialCardHeight {
			t.Fatalf("generated card %s is %dx%d, want %dx%d", path, config.Width, config.Height, socialCardWidth, socialCardHeight)
		}
	}
	if _, err := os.Stat(filepath.Join(outputRoot, "stale.png")); !os.IsNotExist(err) {
		t.Fatalf("stale social card should be removed: %v", err)
	}
}

func TestFitImageDimensionsPreservesGalleryImageAspectRatio(t *testing.T) {
	tests := []struct {
		name       string
		source     image.Rectangle
		wantWidth  int
		wantHeight int
	}{
		{name: "landscape uses maximum width", source: image.Rect(0, 0, 200, 100), wantWidth: 420, wantHeight: 210},
		{name: "portrait uses maximum height", source: image.Rect(0, 0, 100, 200), wantWidth: 225, wantHeight: 450},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			width, height := fitImageDimensions(test.source, 420, 450)
			if width != test.wantWidth || height != test.wantHeight {
				t.Fatalf("fit dimensions = %dx%d, want %dx%d", width, height, test.wantWidth, test.wantHeight)
			}
		})
	}
}
