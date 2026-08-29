package main

import (
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestClassificationMatchesSchema(t *testing.T) {
	classification, err := loadClassification(filepath.Join("..", "website", "data", "classification.yaml"))
	if err != nil {
		t.Fatalf("load classification: %v", err)
	}

	type schemaProperty struct {
		Enum []string `json:"enum"`
	}
	var schema struct {
		Version     int                       `json:"version"`
		Properties  map[string]schemaProperty `json:"properties"`
		Definitions map[string]schemaProperty `json:"$defs"`
	}
	schemaData, err := os.ReadFile(filepath.Join("..", "schema", "project.schema.json"))
	if err != nil {
		t.Fatalf("read schema: %v", err)
	}
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		t.Fatalf("unmarshal schema: %v", err)
	}
	if schema.Version != 6 {
		t.Fatalf("schema version = %d, want 6", schema.Version)
	}

	var projectTypeIDs []string
	var categoryIDs []string
	var capabilityIDs []string
	var platformIDs []string
	for _, projectType := range classification.ProjectTypes {
		projectTypeIDs = append(projectTypeIDs, projectType.ID)
	}
	for _, category := range classification.Categories {
		categoryIDs = append(categoryIDs, category.ID)
	}
	for _, capability := range classification.Capabilities {
		capabilityIDs = append(capabilityIDs, capability.ID)
	}
	for _, platform := range classification.Platforms {
		platformIDs = append(platformIDs, platform.ID)
	}
	if !reflect.DeepEqual(schema.Properties["project_type"].Enum, projectTypeIDs) {
		t.Fatalf("schema project types = %#v, classification project types = %#v", schema.Properties["project_type"].Enum, projectTypeIDs)
	}
	if !reflect.DeepEqual(schema.Definitions["category"].Enum, categoryIDs) {
		t.Fatalf("schema categories = %#v, classification categories = %#v", schema.Definitions["category"].Enum, categoryIDs)
	}
	if !reflect.DeepEqual(schema.Definitions["capability"].Enum, capabilityIDs) {
		t.Fatalf("schema capabilities = %#v, classification capabilities = %#v", schema.Definitions["capability"].Enum, capabilityIDs)
	}
	if !reflect.DeepEqual(schema.Definitions["platform"].Enum, platformIDs) {
		t.Fatalf("schema platforms = %#v, classification platforms = %#v", schema.Definitions["platform"].Enum, platformIDs)
	}
}

func TestValidateProjectClassifications(t *testing.T) {
	classification := Classification{
		ProjectTypes: []ProjectTypeDefinition{{ID: "application", TaxonomySlug: "applications", Label: "Applications", Description: "User software"}},
		Categories:   []CategoryDefinition{{ID: "security", Label: "Security", Description: "Security projects"}, {ID: "payments", Label: "Payments", Description: "Payment projects"}},
		Capabilities: []CapabilityDefinition{{ID: "wallet", Label: "Wallets", Description: "Interactive wallet software"}},
		Platforms:    []PlatformDefinition{{ID: "web", Label: "Web", Description: "Runs in a browser"}},
		Networks:     []NetworkDefinition{{ID: "qrl-1-mainnet", Label: "QRL 1.x Mainnet", Generation: "1.x", Environment: "mainnet"}},
	}
	projects := []Project{
		{ID: "wallet", ProjectType: "application", PrimaryCategory: "security", SecondaryCategories: []string{"payments"}, Capabilities: []string{"wallet"}, Platforms: []string{"web"}},
	}
	if err := validateClassification(classification); err != nil {
		t.Fatalf("valid classification rejected: %v", err)
	}
	if err := validateProjectClassifications(projects, classification); err != nil {
		t.Fatalf("valid projects rejected: %v", err)
	}

	projects[0].SecondaryCategories = []string{"security"}
	if err := validateProjectClassifications(projects, classification); err == nil || !strings.Contains(err.Error(), "repeats category") {
		t.Fatalf("duplicate category error = %v", err)
	}
}

func TestValidateClassificationRejectsDuplicateCategory(t *testing.T) {
	duplicate := Classification{
		ProjectTypes: []ProjectTypeDefinition{{ID: "application", TaxonomySlug: "applications", Label: "Applications", Description: "User software"}},
		Categories:   []CategoryDefinition{{ID: "security", Label: "Security", Description: "One"}, {ID: "security", Label: "Security", Description: "Two"}},
		Capabilities: []CapabilityDefinition{{ID: "wallet", Label: "Wallets", Description: "Interactive wallet software"}},
		Platforms:    []PlatformDefinition{{ID: "web", Label: "Web", Description: "Runs in a browser"}},
		Networks:     []NetworkDefinition{{ID: "qrl-1-mainnet", Label: "QRL 1.x Mainnet", Generation: "1.x", Environment: "mainnet"}},
	}
	if err := validateClassification(duplicate); err == nil || !strings.Contains(err.Error(), "duplicate category") {
		t.Fatalf("duplicate category error = %v", err)
	}
}

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

func TestProjectPageContentOmitsEmptyBody(t *testing.T) {
	content, err := projectPageContent(map[string]interface{}{
		"url": "/projects/qrc20-factory/",
	}, "\n")
	if err != nil {
		t.Fatalf("generate project page content: %v", err)
	}

	want := "---\nurl: /projects/qrc20-factory/\n---\n"
	if string(content) != want {
		t.Fatalf("generated empty project page = %q, want %q", content, want)
	}
}

func TestGenerateProjectPageAddsExactAttributionTaxonomies(t *testing.T) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}
	temporaryDirectory := t.TempDir()
	t.Cleanup(func() {
		if err := os.Chdir(workingDirectory); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})

	if err := os.Chdir(temporaryDirectory); err != nil {
		t.Fatalf("change to temporary directory: %v", err)
	}

	project := Project{
		ID:                  "example-project",
		Name:                "Example project",
		ProjectType:         "application",
		PrimaryCategory:     "finance",
		SecondaryCategories: []string{"payments-commerce"},
		Capabilities:        []string{"wallet"},
		Platforms:           []string{"web"},
		Maturity:            "beta",
		Availability:        "live",
		QRLSupport:          []QRLSupport{{Generation: "2.0", Environments: []string{"testnet"}}},
		Publisher:           Publisher{Name: "The QRL", URL: "https://example.com"},
		Maintainers:         []Maintainer{{Name: "The QRL", Contact: "https://example.com/contact"}},
		SecurityReviews: []SecurityReview{{
			Auditor: "Example Security", ReportURL: "https://example.com/report", Scope: "Example scope", RemediationStatus: "not-reported",
		}},
	}
	generateProjectPage(project)

	content, err := os.ReadFile(projectOutputPath(project))
	if err != nil {
		t.Fatalf("read generated project page: %v", err)
	}
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) != 3 {
		t.Fatalf("generated project page is missing front matter: %s", content)
	}

	var metadata struct {
		ID              string           `yaml:"id"`
		ProjectType     string           `yaml:"project_type"`
		ProjectTypes    []string         `yaml:"project-types"`
		Categories      []string         `yaml:"categories"`
		Capabilities    []string         `yaml:"capabilities"`
		Platforms       []string         `yaml:"platforms"`
		DisplayStatus   string           `yaml:"display_status"`
		Publisher       Publisher        `yaml:"publisher"`
		Publishers      []string         `yaml:"publishers"`
		Maintainers     []string         `yaml:"maintainers"`
		SecurityReviews []SecurityReview `yaml:"security_reviews"`
	}
	if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err != nil {
		t.Fatalf("unmarshal generated front matter: %v", err)
	}
	if metadata.Publisher.Name != "The QRL" {
		t.Fatalf("generated publisher = %q, want %q", metadata.Publisher.Name, "The QRL")
	}
	if !reflect.DeepEqual(metadata.Publishers, []string{"The QRL"}) {
		t.Fatalf("generated publisher taxonomy = %#v, want [The QRL]", metadata.Publishers)
	}
	if len(metadata.Maintainers) != 1 || metadata.Maintainers[0] != "The QRL" {
		t.Fatalf("generated maintainers = %#v, want [The QRL]", metadata.Maintainers)
	}
	if metadata.ID != "example-project" || metadata.ProjectType != "application" {
		t.Fatalf("generated v6 identity/type = %q/%q", metadata.ID, metadata.ProjectType)
	}
	if !reflect.DeepEqual(metadata.ProjectTypes, []string{"applications"}) {
		t.Fatalf("generated project type taxonomy = %#v", metadata.ProjectTypes)
	}
	if !reflect.DeepEqual(metadata.Categories, []string{"finance", "payments-commerce"}) {
		t.Fatalf("generated category taxonomy = %#v", metadata.Categories)
	}
	if !reflect.DeepEqual(metadata.Capabilities, []string{"wallet"}) {
		t.Fatalf("generated capability taxonomy = %#v", metadata.Capabilities)
	}
	if !reflect.DeepEqual(metadata.Platforms, []string{"web"}) {
		t.Fatalf("generated platforms = %#v", metadata.Platforms)
	}
	if metadata.DisplayStatus != "Beta · Testnet" {
		t.Fatalf("generated display status = %q", metadata.DisplayStatus)
	}
	if len(metadata.SecurityReviews) != 1 || metadata.SecurityReviews[0].Auditor != "Example Security" {
		t.Fatalf("generated security reviews = %#v", metadata.SecurityReviews)
	}
}

func TestGenerateJSONIndexV6Contract(t *testing.T) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}
	temporaryDirectory := t.TempDir()
	t.Cleanup(func() {
		if err := os.Chdir(workingDirectory); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})
	if err := os.Chdir(temporaryDirectory); err != nil {
		t.Fatalf("change to temporary directory: %v", err)
	}
	if err := os.MkdirAll(filepath.Join("website", "static"), 0755); err != nil {
		t.Fatal(err)
	}

	project := Project{
		ID:                  "example-project",
		Name:                "Example project",
		ProjectType:         "protocol",
		PrimaryCategory:     "finance",
		SecondaryCategories: []string{"payments-commerce"},
		Capabilities:        []string{"dex"},
		Platforms:           []string{},
		Keywords:            []string{"example"},
		Maturity:            "beta",
		Availability:        "live",
		Maintenance:         "active",
		QRLRelationship:     "deployed",
		QRLSupport:          []QRLSupport{{Generation: "2.0", Environments: []string{"testnet"}}},
		Deployments: []Deployment{{
			ID: "qrl-2-testnet-v2", Network: "qrl-2-testnet-v2", OperationalState: "live",
			Identifiers: []DeploymentIdentifier{}, Evidence: []string{}, SourceVerification: "unknown",
		}},
		Description:        "Example project description.",
		SourceAvailability: "full",
		Repositories:       []Repository{{ID: "main", Role: "contracts", URL: "https://example.com/source", License: "MIT"}},
		Links: []Link{
			{Type: "website", URL: "https://example.com/fallback"},
			{Type: "application", URL: "https://example.com/app", Primary: true},
		},
		Logos: []Logo{{Path: "example-project/icon.png"}},
	}
	generateJSONIndex([]Project{project})

	data, err := os.ReadFile(filepath.Join("website", "static", "index.json"))
	if err != nil {
		t.Fatalf("read generated index: %v", err)
	}
	var document map[string]interface{}
	if err := json.Unmarshal(data, &document); err != nil {
		t.Fatalf("unmarshal generated index: %v", err)
	}
	expectedTopLevel := map[string]bool{"schema_version": true, "generated_at": true, "count": true, "projects": true}
	if len(document) != len(expectedTopLevel) {
		t.Fatalf("top-level JSON fields = %#v", document)
	}
	for field := range document {
		if !expectedTopLevel[field] {
			t.Fatalf("unexpected top-level JSON field %q", field)
		}
	}
	if document["schema_version"] != float64(6) || document["count"] != float64(1) {
		t.Fatalf("schema/count = %#v/%#v", document["schema_version"], document["count"])
	}
	projects, ok := document["projects"].([]interface{})
	if !ok || len(projects) != 1 {
		t.Fatalf("projects = %#v", document["projects"])
	}
	entry := projects[0].(map[string]interface{})
	requiredFields := []string{
		"id", "name", "project_type", "primary_category", "secondary_categories", "capabilities", "platforms", "keywords",
		"maturity", "availability", "maintenance", "display_status", "qrl_relationship", "qrl_support", "deployments",
		"description", "primary_url", "source_availability", "repositories", "links", "logo",
	}
	for _, field := range requiredFields {
		if _, ok := entry[field]; !ok {
			t.Errorf("generated JSON project is missing %q", field)
		}
	}
	for _, legacy := range []string{"category", "status", "qrl_versions", "github", "open_source", "audited", "author", "tags"} {
		if _, ok := entry[legacy]; ok {
			t.Errorf("generated JSON contains legacy field %q", legacy)
		}
	}
	if entry["primary_url"] != "https://example.com/app" {
		t.Errorf("primary_url = %#v", entry["primary_url"])
	}
	if entry["display_status"] != "Beta · Testnet" {
		t.Errorf("display_status = %#v", entry["display_status"])
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
	if err := os.WriteFile(filepath.Join(source, ".DS_Store"), []byte("metadata"), 0644); err != nil {
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
	if _, err := os.Stat(filepath.Join(destination, ".DS_Store")); !os.IsNotExist(err) {
		t.Fatalf("hidden metadata should not be copied: %v", err)
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
			ProjectType: "protocol",
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
			ProjectType: "resource",
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
