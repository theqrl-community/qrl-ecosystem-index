package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	xdraw "golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	_ "golang.org/x/image/webp"
	"gopkg.in/yaml.v3"
)

type Project struct {
	ID              string        `yaml:"id"`
	Name            string        `yaml:"name"`
	ProjectType     string        `yaml:"project_type"`
	Status          string        `yaml:"status"`
	Description     string        `yaml:"description"`
	Category        string        `yaml:"category"`
	Tags            []string      `yaml:"tags"`
	Author          string        `yaml:"author"`
	License         string        `yaml:"license"`
	Created         string        `yaml:"created"`
	Updated         string        `yaml:"updated"`
	URL             string        `yaml:"url"`
	GitHub          string        `yaml:"github"`
	Docs            string        `yaml:"docs"`
	Discord         string        `yaml:"discord"`
	Twitter         string        `yaml:"twitter"`
	OpenSource      bool          `yaml:"open_source"`
	Audited         bool          `yaml:"audited"`
	Audits          []Audit       `yaml:"audits"`
	Clients         []Client      `yaml:"clients"`
	Logo            string        `yaml:"logo"`
	Logos           []Logo        `yaml:"logos"`
	Gallery         []GalleryItem `yaml:"gallery"`
	Features        []string      `yaml:"features"`
	LongDescription string        `yaml:"long_description"`
	// Type-specific blocks
	Dapp           *DappBlock           `yaml:"dapp,omitempty"`
	Application    *ApplicationBlock    `yaml:"application,omitempty"`
	Infrastructure *InfrastructureBlock `yaml:"infrastructure,omitempty"`
	Tooling        *ToolingBlock        `yaml:"tooling,omitempty"`
	Community      *CommunityBlock      `yaml:"community,omitempty"`
}

type Logo struct {
	Path        string `yaml:"path"`
	Description string `yaml:"description"`
}

type GalleryItem struct {
	Type    string `yaml:"type"`
	Path    string `yaml:"path,omitempty"`
	ID      string `yaml:"id,omitempty"`
	Caption string `yaml:"caption"`
}

type Audit struct {
	Auditor  string `yaml:"auditor"`
	AuditURL string `yaml:"audit_url"`
}

type Client struct {
	Platform string `yaml:"platform" json:"platform"`
	URL      string `yaml:"url,omitempty" json:"url,omitempty"`
	GitHub   string `yaml:"github,omitempty" json:"github,omitempty"`
	Default  bool   `yaml:"default,omitempty" json:"default,omitempty"`
}

type DappBlock struct {
	Network         string `yaml:"network"`
	ContractAddress string `yaml:"contract_address"`
	Token           string `yaml:"token"`
}

type ApplicationBlock struct {
	Platforms         []string `yaml:"platforms"`
	SupportedNetworks []string `yaml:"supported_networks"`
}

type InfrastructureBlock struct {
	SupportedNetworks []string `yaml:"supported_networks"`
	Endpoints         []string `yaml:"endpoints"`
}

type ToolingBlock struct {
	Languages []string `yaml:"languages"`
}

type CommunityBlock struct {
	Platforms []string `yaml:"platforms"`
	Language  string   `yaml:"language"`
}

func main() {
	// Ensure generated content and static asset directories exist.
	os.MkdirAll("website/content/projects", 0755)
	os.MkdirAll("website/static", 0755)

	var projects []Project

	// Process active projects
	processDir("projects/active", &projects)
	// Process archived projects
	processDir("projects/archived", &projects)

	removeStaleProjectPages(projects)

	// Generate individual project pages
	for _, p := range projects {
		generateProjectPage(p)
	}
	if err := generateSocialCards(projects, "images", filepath.Join("website", "static", "images", "og")); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating social preview cards: %v\n", err)
		os.Exit(1)
	}

	// Generate JSON index
	generateJSONIndex(projects)

	// Copy local project media into Hugo's published static tree.
	assetTrees := []struct {
		name        string
		source      string
		destination string
	}{
		{"logo", filepath.Join("images", "logos"), filepath.Join("website", "static", "images", "logos")},
		{"screenshot", filepath.Join("images", "screenshots"), filepath.Join("website", "static", "images", "screenshots")},
	}
	for _, assetTree := range assetTrees {
		if err := copyAssetTree(assetTree.source, assetTree.destination); err != nil {
			fmt.Fprintf(os.Stderr, "Error copying %s assets: %v\n", assetTree.name, err)
			os.Exit(1)
		}
	}

	fmt.Printf("Generated %d project pages\n", len(projects))
}

func processDir(dir string, projects *[]Project) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", dir, err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".yaml") {
			continue
		}

		filepath := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filepath, err)
			continue
		}

		var project Project
		if err := yaml.Unmarshal(data, &project); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing YAML %s: %v\n", filepath, err)
			continue
		}
		if project.Logo != "" && len(project.Logos) == 0 {
			project.Logos = []Logo{{Path: project.Logo}}
		}

		*projects = append(*projects, project)
	}
}

func copyAssetTree(sourceDir, destinationDir string) error {
	if err := os.RemoveAll(destinationDir); err != nil {
		return err
	}
	if _, err := os.Stat(sourceDir); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return filepath.WalkDir(sourceDir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		destinationPath := filepath.Join(destinationDir, relativePath)

		if entry.IsDir() {
			return os.MkdirAll(destinationPath, 0755)
		}
		if entry.Name() == ".gitkeep" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		return os.WriteFile(destinationPath, data, info.Mode().Perm())
	})
}

func removeStaleProjectPages(projects []Project) {
	expected := make(map[string]bool, len(projects))
	for _, p := range projects {
		expected[projectOutputPath(p)] = true
	}

	for _, dir := range []string{
		filepath.Join("website", "content", "projects", "active"),
		filepath.Join("website", "content", "projects", "archived"),
	} {
		if err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading generated project path %s: %v\n", path, err)
				return nil
			}
			if entry.IsDir() || filepath.Ext(path) != ".md" {
				return nil
			}
			if expected[path] {
				return nil
			}
			if err := os.Remove(path); err != nil {
				fmt.Fprintf(os.Stderr, "Error removing stale project page %s: %v\n", path, err)
			}
			return nil
		}); err != nil && !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error scanning generated project pages in %s: %v\n", dir, err)
		}
	}
}

func generateProjectPage(p Project) {
	outputPath := projectOutputPath(p)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory for %s: %v\n", outputPath, err)
		return
	}

	params := map[string]interface{}{
		"url":          projectPermalink(p),
		"aliases":      legacyProjectAliases(p),
		"title":        p.Name,
		"status":       p.Status,
		"category":     p.Category,
		"categories":   []string{p.Category},
		"tags":         p.Tags,
		"project_type": p.ProjectType,
		"project-types": []string{
			projectTypeSlug(p.ProjectType),
		},
		"author":      p.Author,
		"license":     p.License,
		"created":     p.Created,
		"updated":     p.Updated,
		"description": strings.TrimSpace(p.Description),
		"project_url": p.URL,
		"github":      p.GitHub,
		"docs":        p.Docs,
		"discord":     p.Discord,
		"twitter":     p.Twitter,
		"open_source": p.OpenSource,
		"audited":     p.Audited,
		"audits":      p.Audits,
		"clients":     p.Clients,
		"logos":       p.Logos,
		"features":    p.Features,
	}
	if len(p.Gallery) > 0 {
		params["gallery"] = p.Gallery
	}

	if client, ok := defaultClient(p); ok {
		params["default_client"] = client
	}
	if url := defaultProjectURL(p); url != "" {
		params["default_client_url"] = url
	}
	if github := defaultProjectGitHub(p); github != "" {
		params["default_client_github"] = github
	}

	if p.Dapp != nil {
		params["ecosystem_type"] = "dapp"
		params["network"] = p.Dapp.Network
		params["contract_address"] = p.Dapp.ContractAddress
		params["token"] = p.Dapp.Token
	} else if p.Application != nil {
		params["ecosystem_type"] = "application"
		params["platforms"] = p.Application.Platforms
		params["supported_networks"] = p.Application.SupportedNetworks
	} else if p.Infrastructure != nil {
		params["ecosystem_type"] = "infrastructure"
		params["supported_networks"] = p.Infrastructure.SupportedNetworks
		params["endpoints"] = p.Infrastructure.Endpoints
	} else if p.Tooling != nil {
		params["ecosystem_type"] = "tooling"
		params["languages"] = p.Tooling.Languages
	} else if p.Community != nil {
		params["ecosystem_type"] = "community"
		params["platforms"] = p.Community.Platforms
		params["language"] = p.Community.Language
	}

	content, err := projectPageContent(params, p.LongDescription)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating front matter for %s: %v\n", p.ID, err)
		return
	}

	if err := os.WriteFile(outputPath, content, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", outputPath, err)
	}
}

func projectPageContent(params map[string]interface{}, longDescription string) ([]byte, error) {
	frontMatter, err := yaml.Marshal(params)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("---\n%s---\n\n%s\n", frontMatter, escapeMarkdown(longDescription))), nil
}

func projectOutputPath(p Project) string {
	// Determine section based on status
	section := "projects/active"
	if p.Status == "archived" {
		section = "projects/archived"
	}

	return filepath.Join("website", "content", section, p.ID+".md")
}

func projectPermalink(p Project) string {
	return "/projects/" + p.ID + "/"
}

func legacyProjectAliases(p Project) []string {
	return []string{
		"/projects/active/" + p.ID + "/",
		"/projects/archived/" + p.ID + "/",
	}
}

func generateJSONIndex(projects []Project) {
	type IndexProject struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Status      string   `json:"status"`
		Category    string   `json:"category"`
		Description string   `json:"description"`
		URL         string   `json:"url"`
		GitHub      string   `json:"github"`
		Logo        string   `json:"logo,omitempty"`
		Clients     []Client `json:"clients,omitempty"`
	}

	var index []IndexProject
	for _, p := range projects {
		index = append(index, IndexProject{
			ID:          p.ID,
			Name:        p.Name,
			Status:      p.Status,
			Category:    p.Category,
			Description: strings.TrimSpace(p.Description),
			URL:         defaultProjectURL(p),
			GitHub:      defaultProjectGitHub(p),
			Logo:        defaultLogoURL(p),
			Clients:     p.Clients,
		})
	}

	data, _ := json.MarshalIndent(map[string]interface{}{
		"generated": time.Now().UTC().Format(time.RFC3339),
		"count":     len(projects),
		"projects":  index,
	}, "", "  ")

	if err := os.WriteFile("website/static/index.json", append(data, '\n'), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing JSON index: %v\n", err)
	}
}

func escapeMarkdown(s string) string {
	return s
}

func defaultClient(p Project) (Client, bool) {
	for _, client := range p.Clients {
		if client.Default {
			return client, true
		}
	}
	for _, client := range p.Clients {
		if client.URL != "" {
			return client, true
		}
	}
	if len(p.Clients) > 0 {
		return p.Clients[0], true
	}
	return Client{}, false
}

func defaultProjectURL(p Project) string {
	if client, ok := defaultClient(p); ok && client.URL != "" {
		return client.URL
	}
	return p.URL
}

func defaultProjectGitHub(p Project) string {
	if client, ok := defaultClient(p); ok && client.GitHub != "" {
		return client.GitHub
	}
	return p.GitHub
}

func defaultLogoURL(p Project) string {
	if len(p.Logos) == 0 || p.Logos[0].Path == "" {
		return ""
	}
	return "/images/logos/" + filepath.ToSlash(strings.TrimPrefix(p.Logos[0].Path, "/"))
}

func projectTypeSlug(projectType string) string {
	switch projectType {
	case "dapp":
		return "dapps"
	case "application":
		return "applications"
	default:
		return projectType
	}
}

const (
	socialCardWidth  = 1200
	socialCardHeight = 630
)

var (
	cardPaper       = color.RGBA{R: 244, G: 242, B: 235, A: 255}
	cardInk         = color.RGBA{R: 20, G: 38, B: 49, A: 255}
	cardMuted       = color.RGBA{R: 76, G: 92, B: 101, A: 255}
	cardLine        = color.RGBA{R: 205, G: 210, B: 207, A: 255}
	cardAccent      = color.RGBA{R: 42, G: 142, B: 154, A: 255}
	cardAccentLight = color.RGBA{R: 215, G: 235, B: 234, A: 255}
	cardWhite       = color.RGBA{R: 255, G: 255, B: 255, A: 255}
)

type socialCardFonts struct {
	label   font.Face
	title   font.Face
	titleSm font.Face
	body    font.Face
	initial font.Face
}

func generateSocialCards(projects []Project, assetRoot, outputRoot string) error {
	if err := os.RemoveAll(outputRoot); err != nil {
		return err
	}
	projectOutputRoot := filepath.Join(outputRoot, "projects")
	if err := os.MkdirAll(projectOutputRoot, 0755); err != nil {
		return err
	}

	fonts, err := newSocialCardFonts()
	if err != nil {
		return err
	}
	if err := writeSocialCard(filepath.Join(outputRoot, "default.png"), renderDefaultSocialCard(fonts)); err != nil {
		return err
	}

	for _, project := range projects {
		card, err := renderProjectSocialCard(project, assetRoot, fonts)
		if err != nil {
			return fmt.Errorf("%s: %w", project.ID, err)
		}
		if err := writeSocialCard(filepath.Join(projectOutputRoot, project.ID+".png"), card); err != nil {
			return fmt.Errorf("%s: %w", project.ID, err)
		}
	}
	return nil
}

func newSocialCardFonts() (socialCardFonts, error) {
	regular, err := opentype.Parse(goregular.TTF)
	if err != nil {
		return socialCardFonts{}, err
	}
	bold, err := opentype.Parse(gobold.TTF)
	if err != nil {
		return socialCardFonts{}, err
	}
	makeFace := func(parsed *opentype.Font, size float64) (font.Face, error) {
		return opentype.NewFace(parsed, &opentype.FaceOptions{Size: size, DPI: 72, Hinting: font.HintingFull})
	}

	label, err := makeFace(bold, 18)
	if err != nil {
		return socialCardFonts{}, err
	}
	title, err := makeFace(bold, 66)
	if err != nil {
		return socialCardFonts{}, err
	}
	titleSm, err := makeFace(bold, 54)
	if err != nil {
		return socialCardFonts{}, err
	}
	body, err := makeFace(regular, 25)
	if err != nil {
		return socialCardFonts{}, err
	}
	initial, err := makeFace(bold, 108)
	if err != nil {
		return socialCardFonts{}, err
	}
	return socialCardFonts{
		label:   label,
		title:   title,
		titleSm: titleSm,
		body:    body,
		initial: initial,
	}, nil
}

func renderDefaultSocialCard(fonts socialCardFonts) image.Image {
	card := image.NewRGBA(image.Rect(0, 0, socialCardWidth, socialCardHeight))
	draw.Draw(card, card.Bounds(), &image.Uniform{C: cardPaper}, image.Point{}, draw.Src)
	drawCardBackground(card)
	drawLabel(card, fonts.label, "QRL / COMMUNITY INDEX", 72, 64, cardAccent)
	drawWrappedText(card, fonts.title, "QRL Ecosystem\nIndex", 72, 190, 600, 2, 76, cardInk)
	drawWrappedText(card, fonts.body, "A community-maintained view of projects, tools, services, and resources connected to QRL 2.0.", 72, 410, 580, 3, 36, cardMuted)
	drawCardMotif(card, "QI", fonts, 770, 90, 350, 410)
	drawFooter(card, fonts)
	return card
}

func renderProjectSocialCard(project Project, assetRoot string, fonts socialCardFonts) (image.Image, error) {
	card := image.NewRGBA(image.Rect(0, 0, socialCardWidth, socialCardHeight))
	draw.Draw(card, card.Bounds(), &image.Uniform{C: cardPaper}, image.Point{}, draw.Src)
	drawCardBackground(card)

	markRect := image.Rect(72, 74, 158, 160)
	drawRoundedRect(card, markRect, 18, cardWhite)
	drawRoundedBorder(card, markRect, 18, cardLine, 2)
	logoDrawn := false
	if len(project.Logos) > 0 && project.Logos[0].Path != "" {
		logoPath := filepath.Join(assetRoot, "logos", filepath.FromSlash(project.Logos[0].Path))
		logo, err := loadProjectLogo(logoPath, 66, 66)
		if err != nil {
			return nil, fmt.Errorf("load logo: %w", err)
		}
		drawImageContain(card, logo, image.Rect(82, 84, 148, 150))
		logoDrawn = true
	}
	if !logoDrawn {
		initials := projectInitials(project.Name)
		drawCenteredText(card, fonts.label, initials, markRect, cardInk)
	}

	drawLabel(card, fonts.label, "QRL / ECOSYSTEM INDEX", 176, 105, cardAccent)
	drawLabel(card, fonts.label, strings.ToUpper(projectTypeLabel(project.ProjectType)), 176, 143, cardMuted)

	titleFace := fonts.title
	if len([]rune(project.Name)) > 18 {
		titleFace = fonts.titleSm
	}
	titleLines := wrapText(titleFace, project.Name, 570)
	if len(titleLines) > 2 {
		titleFace = fonts.titleSm
	}
	drawWrappedText(card, titleFace, project.Name, 72, 238, 570, 3, 68, cardInk)
	drawWrappedText(card, fonts.body, strings.TrimSpace(project.Description), 72, 410, 570, 3, 35, cardMuted)

	if galleryImage, ok := firstGalleryImage(project.Gallery); ok {
		screenshotPath := filepath.Join(assetRoot, "screenshots", filepath.FromSlash(galleryImage.Path))
		screenshot, err := loadRasterImage(screenshotPath)
		if err != nil {
			return nil, fmt.Errorf("load first screenshot: %w", err)
		}
		drawScreenshotPanel(card, screenshot)
	} else {
		drawCardMotif(card, projectInitials(project.Name), fonts, 728, 58, 402, 476)
	}

	drawFooter(card, fonts)
	return card, nil
}

func firstGalleryImage(gallery []GalleryItem) (GalleryItem, bool) {
	for _, item := range gallery {
		if item.Type == "image" && item.Path != "" {
			return item, true
		}
	}
	return GalleryItem{}, false
}

func drawCardBackground(card *image.RGBA) {
	draw.Draw(card, image.Rect(0, 0, 14, socialCardHeight), &image.Uniform{C: cardAccent}, image.Point{}, draw.Src)
}

func drawScreenshotPanel(card *image.RGBA, screenshot image.Image) {
	contentWidth, contentHeight := fitImageDimensions(screenshot.Bounds(), 420, 450)
	frameWidth := contentWidth + 36
	frameHeight := contentHeight + 36
	available := image.Rect(684, 38, 1140, 524)
	left := available.Min.X + (available.Dx()-frameWidth)/2
	top := available.Min.Y + (available.Dy()-frameHeight)/2
	panelRect := image.Rect(left, top, left+frameWidth, top+frameHeight)
	shadowRect := panelRect.Add(image.Pt(11, 11))
	drawRoundedRect(card, shadowRect, 24, color.RGBA{R: 20, G: 38, B: 49, A: 35})
	drawRoundedRect(card, panelRect, 24, cardWhite)
	drawRoundedBorder(card, panelRect, 24, cardLine, 2)

	imageRect := panelRect.Inset(18)
	fitted := resizeImage(screenshot, imageRect.Dx(), imageRect.Dy())
	mask := roundedMask(imageRect.Dx(), imageRect.Dy(), 14)
	draw.DrawMask(card, imageRect, fitted, image.Point{}, mask, image.Point{}, draw.Over)
}

func drawCardMotif(card *image.RGBA, initials string, fonts socialCardFonts, x, y, width, height int) {
	rect := image.Rect(x, y, x+width, y+height)
	drawRoundedRect(card, rect, 28, cardInk)
	circleCenter := image.Pt(rect.Min.X+width/2, rect.Min.Y+height/2)
	drawCircle(card, circleCenter, minInt(width, height)/3, cardAccent)
	drawCircle(card, circleCenter, minInt(width, height)/3-12, cardAccentLight)
	drawCenteredText(card, fonts.initial, initials, image.Rect(circleCenter.X-130, circleCenter.Y-100, circleCenter.X+130, circleCenter.Y+100), cardInk)
}

func drawFooter(card *image.RGBA, fonts socialCardFonts) {
	draw.Draw(card, image.Rect(72, 579, 1128, 581), &image.Uniform{C: cardLine}, image.Point{}, draw.Src)
	drawLabel(card, fonts.label, "QRLECOSYSTEM.COM", 72, 610, cardMuted)
}

func loadProjectLogo(path string, width, height int) (image.Image, error) {
	if strings.EqualFold(filepath.Ext(path), ".svg") {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		icon, err := oksvg.ReadIconStream(file)
		if err != nil {
			return nil, err
		}
		icon.SetTarget(0, 0, float64(width), float64(height))
		canvas := image.NewRGBA(image.Rect(0, 0, width, height))
		scanner := rasterx.NewScannerGV(width, height, canvas, canvas.Bounds())
		dasher := rasterx.NewDasher(width, height, scanner)
		icon.Draw(dasher, 1)
		return canvas, nil
	}
	return loadRasterImage(path)
}

func loadRasterImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoded, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func drawImageContain(destination *image.RGBA, source image.Image, bounds image.Rectangle) {
	sourceBounds := source.Bounds()
	if sourceBounds.Dx() == 0 || sourceBounds.Dy() == 0 {
		return
	}
	scale := math.Min(float64(bounds.Dx())/float64(sourceBounds.Dx()), float64(bounds.Dy())/float64(sourceBounds.Dy()))
	width := maxInt(1, int(math.Round(float64(sourceBounds.Dx())*scale)))
	height := maxInt(1, int(math.Round(float64(sourceBounds.Dy())*scale)))
	x := bounds.Min.X + (bounds.Dx()-width)/2
	y := bounds.Min.Y + (bounds.Dy()-height)/2
	xdraw.CatmullRom.Scale(destination, image.Rect(x, y, x+width, y+height), source, sourceBounds, draw.Over, nil)
}

func fitImageDimensions(source image.Rectangle, maxWidth, maxHeight int) (int, int) {
	if source.Dx() == 0 || source.Dy() == 0 {
		return maxWidth, maxHeight
	}
	scale := math.Min(float64(maxWidth)/float64(source.Dx()), float64(maxHeight)/float64(source.Dy()))
	return maxInt(1, int(math.Round(float64(source.Dx())*scale))), maxInt(1, int(math.Round(float64(source.Dy())*scale)))
}

func resizeImage(source image.Image, width, height int) image.Image {
	sourceBounds := source.Bounds()
	destination := image.NewRGBA(image.Rect(0, 0, width, height))
	if sourceBounds.Dx() == 0 || sourceBounds.Dy() == 0 {
		return destination
	}
	xdraw.CatmullRom.Scale(destination, destination.Bounds(), source, sourceBounds, draw.Src, nil)
	return destination
}

func writeSocialCard(path string, card image.Image) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	return encoder.Encode(file, card)
}

func drawLabel(destination draw.Image, face font.Face, text string, x, baseline int, ink color.Color) {
	drawer := font.Drawer{Dst: destination, Src: &image.Uniform{C: ink}, Face: face, Dot: fixedPoint(x, baseline)}
	drawer.DrawString(text)
}

func drawWrappedText(destination draw.Image, face font.Face, text string, x, baseline, maxWidth, maxLines, lineHeight int, ink color.Color) {
	lines := wrapText(face, text, maxWidth)
	if len(lines) > maxLines {
		lines = lines[:maxLines]
		lines[maxLines-1] = ellipsize(face, lines[maxLines-1], maxWidth)
	}
	for index, line := range lines {
		drawLabel(destination, face, line, x, baseline+index*lineHeight, ink)
	}
}

func wrapText(face font.Face, text string, maxWidth int) []string {
	var lines []string
	for _, paragraph := range strings.Split(text, "\n") {
		words := strings.Fields(paragraph)
		if len(words) == 0 {
			lines = append(lines, "")
			continue
		}
		current := words[0]
		for _, word := range words[1:] {
			candidate := current + " " + word
			if measureText(face, candidate) <= maxWidth {
				current = candidate
				continue
			}
			lines = append(lines, current)
			current = word
		}
		lines = append(lines, current)
	}
	return lines
}

func ellipsize(face font.Face, text string, maxWidth int) string {
	text = strings.TrimSpace(text)
	for measureText(face, text+"…") > maxWidth && len(text) > 0 {
		runes := []rune(text)
		text = strings.TrimSpace(string(runes[:len(runes)-1]))
	}
	return text + "…"
}

func measureText(face font.Face, text string) int {
	drawer := font.Drawer{Face: face}
	return drawer.MeasureString(text).Ceil()
}

func drawCenteredText(destination draw.Image, face font.Face, text string, bounds image.Rectangle, ink color.Color) {
	metrics := face.Metrics()
	width := measureText(face, text)
	height := (metrics.Ascent + metrics.Descent).Ceil()
	x := bounds.Min.X + (bounds.Dx()-width)/2
	baseline := bounds.Min.Y + (bounds.Dy()-height)/2 + metrics.Ascent.Ceil()
	drawLabel(destination, face, text, x, baseline, ink)
}

func projectInitials(name string) string {
	words := strings.FieldsFunc(name, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsNumber(r) })
	if len(words) >= 2 {
		return strings.ToUpper(string([]rune(words[0])[:1]) + string([]rune(words[1])[:1]))
	}
	clean := []rune(strings.TrimSpace(name))
	if len(clean) == 0 {
		return "QI"
	}
	if len(clean) == 1 {
		return strings.ToUpper(string(clean))
	}
	return strings.ToUpper(string(clean[:2]))
}

func projectTypeLabel(projectType string) string {
	switch projectType {
	case "dapp":
		return "DApp"
	case "application":
		return "Application"
	case "infrastructure":
		return "Infrastructure"
	case "tooling":
		return "Tooling"
	case "community":
		return "Community"
	default:
		return projectType
	}
}

func drawRoundedRect(destination draw.Image, rect image.Rectangle, radius int, fill color.Color) {
	mask := roundedMask(rect.Dx(), rect.Dy(), radius)
	draw.DrawMask(destination, rect, &image.Uniform{C: fill}, image.Point{}, mask, image.Point{}, draw.Over)
}

func drawRoundedBorder(destination draw.Image, rect image.Rectangle, radius int, border color.Color, thickness int) {
	drawRoundedRect(destination, rect, radius, border)
	inner := rect.Inset(thickness)
	drawRoundedRect(destination, inner, maxInt(0, radius-thickness), cardWhite)
}

func roundedMask(width, height, radius int) *image.Alpha {
	mask := image.NewAlpha(image.Rect(0, 0, width, height))
	radius = minInt(radius, minInt(width/2, height/2))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := maxInt(radius-x, maxInt(0, x-(width-radius-1)))
			dy := maxInt(radius-y, maxInt(0, y-(height-radius-1)))
			if dx == 0 || dy == 0 || dx*dx+dy*dy <= radius*radius {
				mask.SetAlpha(x, y, color.Alpha{A: 255})
			}
		}
	}
	return mask
}

func drawCircle(destination draw.Image, center image.Point, radius int, fill color.Color) {
	for y := -radius; y <= radius; y++ {
		halfWidth := int(math.Sqrt(float64(radius*radius - y*y)))
		draw.Draw(destination, image.Rect(center.X-halfWidth, center.Y+y, center.X+halfWidth+1, center.Y+y+1), &image.Uniform{C: fill}, image.Point{}, draw.Over)
	}
}

func fixedPoint(x, y int) fixed.Point26_6 {
	return fixed.P(x, y)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
