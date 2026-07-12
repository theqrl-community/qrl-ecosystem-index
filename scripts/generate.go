package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Project struct {
	ID              string       `yaml:"id"`
	Name            string       `yaml:"name"`
	ProjectType     string       `yaml:"project_type"`
	Status          string       `yaml:"status"`
	Description     string       `yaml:"description"`
	Category        string       `yaml:"category"`
	Tags            []string     `yaml:"tags"`
	Author          string       `yaml:"author"`
	License         string       `yaml:"license"`
	Created         string       `yaml:"created"`
	Updated         string       `yaml:"updated"`
	URL             string       `yaml:"url"`
	GitHub          string       `yaml:"github"`
	Docs            string       `yaml:"docs"`
	Discord         string       `yaml:"discord"`
	Twitter         string       `yaml:"twitter"`
	OpenSource      bool         `yaml:"open_source"`
	Audited         bool         `yaml:"audited"`
	Audits          []Audit      `yaml:"audits"`
	Clients         []Client     `yaml:"clients"`
	Logo            string       `yaml:"logo"`
	Logos           []Logo       `yaml:"logos"`
	Screenshots     []Screenshot `yaml:"screenshots"`
	Features        []string     `yaml:"features"`
	LongDescription string       `yaml:"long_description"`
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

type Screenshot struct {
	Path    string `yaml:"path"`
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
	if len(p.Screenshots) > 0 {
		params["screenshots"] = p.Screenshots
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

	if err := os.WriteFile("website/static/index.json", data, 0644); err != nil {
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
