package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Project struct {
	ID              string   `yaml:"id"`
	Name            string   `yaml:"name"`
	ProjectType     string   `yaml:"project_type"`
	Status          string   `yaml:"status"`
	Description     string   `yaml:"description"`
	Category        string   `yaml:"category"`
	Tags            []string `yaml:"tags"`
	Author          string   `yaml:"author"`
	License         string   `yaml:"license"`
	Created         string   `yaml:"created"`
	Updated         string   `yaml:"updated"`
	URL             string   `yaml:"url"`
	GitHub          string   `yaml:"github"`
	Docs            string   `yaml:"docs"`
	Discord         string   `yaml:"discord"`
	Twitter         string   `yaml:"twitter"`
	OpenSource      bool     `yaml:"open_source"`
	Audited         bool     `yaml:"audited"`
	AuditURL        string   `yaml:"audit_url"`
	Logos           []Logo   `yaml:"logos"`
	Features        []string `yaml:"features"`
	LongDescription string   `yaml:"long_description"`
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

	// Generate individual project pages
	for _, p := range projects {
		generateProjectPage(p)
	}

	// Generate JSON index
	generateJSONIndex(projects)

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

		*projects = append(*projects, project)
	}
}

func generateProjectPage(p Project) {
	// Determine section based on status
	section := "projects/active"
	if p.Status == "archived" {
		section = "projects/archived"
	}

	outputPath := filepath.Join("website/content", section, p.ID+".md")
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory for %s: %v\n", outputPath, err)
		return
	}

	params := map[string]interface{}{
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
		"audit_url":   p.AuditURL,
		"logos":       p.Logos,
		"features":    p.Features,
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

	frontMatter, err := yaml.Marshal(params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating front matter for %s: %v\n", p.ID, err)
		return
	}

	content := fmt.Sprintf("---\n%s---\n\n%s\n", frontMatter, escapeMarkdown(p.LongDescription))

	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", outputPath, err)
	}
}

func generateJSONIndex(projects []Project) {
	type IndexProject struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Status      string `json:"status"`
		Category    string `json:"category"`
		Description string `json:"description"`
		URL         string `json:"url"`
		GitHub      string `json:"github"`
	}

	var index []IndexProject
	for _, p := range projects {
		index = append(index, IndexProject{
			ID:          p.ID,
			Name:        p.Name,
			Status:      p.Status,
			Category:    p.Category,
			Description: strings.TrimSpace(p.Description),
			URL:         p.URL,
			GitHub:      p.GitHub,
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
