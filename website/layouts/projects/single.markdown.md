{{- $clientLinks := where .Params.links "type" "application" -}}
{{- $projectLinks := where .Params.links "type" "ne" "application" -}}
# {{ .Title }}

> {{ .Params.description | plainify }}

{{ with .OutputFormats.Get "html" -}}
- Canonical page: [{{ $.Title }}]({{ .Permalink }})
{{- end }}
- Status: {{ .Params.display_status }}
- Type: {{ partial "project-type-label.html" .Params.project_type }} (`{{ .Params.project_type }}`)
- Primary category: {{ partial "category-label.html" .Params.primary_category }} (`{{ .Params.primary_category }}`)
{{- with .Params.secondary_categories }}
- Secondary categories: {{ range $index, $category := . }}{{ if $index }}, {{ end }}{{ partial "category-label.html" $category }}{{ end }}
{{- end }}
- Capabilities: {{ range $index, $capability := .Params.capabilities }}{{ if $index }}, {{ end }}{{ partial "capability-label.html" $capability }}{{ end }}
{{- with .Params.platforms }}
- Available on: {{ range $index, $platform := . }}{{ if $index }}, {{ end }}{{ partial "platform-label.html" $platform }}{{ end }}
{{- end }}
- QRL relationship: {{ .Params.qrl_relationship | humanize }}
- QRL support: {{ range $index, $support := .Params.qrl_support }}{{ if $index }}; {{ end }}QRL {{ $support.generation }}{{ with $support.environments }} ({{ delimit . ", " }}){{ end }}{{ end }}
- Publisher: {{ with .GetTerms "publishers" }}{{ range first 1 . }}[{{ $.Params.publisher.name }}]({{ with .OutputFormats.Get "markdown" }}{{ .Permalink }}{{ else }}{{ .Permalink }}{{ end }}){{ end }}{{ else }}{{ .Params.publisher.name }}{{ end }}
- Source availability: {{ .Params.source_availability | humanize }}
- Listed: {{ .Params.listed_at }}
- Data updated: {{ .Params.data_updated_at }}
{{- with .Params.last_verified_at }}
- Last verified: {{ . }}
{{- end }}
{{- with .Params.keywords }}
- Search keywords: {{ delimit . ", " }}
{{- end }}

{{- with $projectLinks }}

## Project links

{{ range . -}}
- [{{ with .label }}{{ . }}{{ else }}{{ .type | humanize | title }}{{ with .platform }} · {{ . }}{{ end }}{{ end }}]({{ .url }})
{{ end -}}
{{- end }}
{{- with $clientLinks }}

## Clients

{{ range . -}}
- [{{ with .label }}{{ . }}{{ else }}{{ with .platform }}{{ . | humanize | title }}{{ else }}Client{{ end }}{{ end }}]({{ .url }}){{ if .primary }} — default{{ end }}
{{ end -}}
{{- end }}
{{- with .Params.repositories }}

## Repositories

{{ range . -}}
- [{{ .role | humanize | title }} repository]({{ .url }}) — license: {{ .license }}
{{ end -}}
{{- end }}
{{- with strings.TrimSpace .RawContent }}

## Overview

{{ . }}
{{- end }}
{{- with .Params.features }}

## Features

{{ range . -}}
- {{ . }}
{{ end -}}
{{- end }}
{{- with .Params.deployments }}

## Deployments

{{ range . -}}
- **{{ partial "network-label.html" .network }}** — {{ .operational_state | humanize }}; source verification: {{ .source_verification | humanize }}
{{ range .identifiers }}  - {{ .type | humanize | title }}{{ with .role }} ({{ . }}){{ end }}: `{{ .value }}`
{{ end }}{{ range .evidence }}  - [Deployment evidence]({{ . }})
{{ end }}{{ end -}}
{{- end }}
{{- with .Params.security_reviews }}

## Security review reports

Security review reports are provenance records, not safety ratings.

{{ range . -}}
- [{{ .auditor }} report]({{ .report_url }}){{ with .report_date }} — {{ . }}{{ end }}. {{ .scope }} Remediation: {{ .remediation_status | humanize }}.
{{ end -}}
{{- end }}
{{- with .Params.assets }}

## Assets

{{ range . -}}
- **{{ .name }}** — {{ .type | humanize }}{{ with .symbol }} (`{{ . }}`){{ end }}{{ with .deployment_id }}; deployment: `{{ . }}`{{ end }}{{ with .identifier }}; identifier: `{{ . }}`{{ end }}{{ with .evidence_url }}; [evidence]({{ . }}){{ end }}
{{ end -}}
{{- end }}
{{- if or .Params.relationships .Params.previous_names }}

## Project history

{{ range .Params.relationships -}}
- {{ .type | humanize | title }} [{{ .project_id }}]({{ printf "/projects/%s/" .project_id | absURL }})
{{ end -}}
{{ with .Params.previous_names }}- Previous names: {{ delimit . ", " }}{{ end }}
{{- end }}
{{- with .Params.evidence }}

## Evidence references

These submitted references provide context and do not create a trust or ownership label.

{{ range . -}}
- [{{ .type | humanize | title }} evidence]({{ .url }}){{ with .checked_at }} — checked {{ . }}{{ end }}{{ with .note }}. {{ . }}{{ end }}
{{ end -}}
{{- end }}
{{- with .Params.gallery }}

## Gallery

{{ range . -}}
{{ if eq .type "youtube" -}}
- [{{ .caption }}](https://www.youtube.com/watch?v={{ .id }}) — YouTube video
{{ else -}}
- [{{ .caption }}]({{ printf "/images/screenshots/%s" .path | absURL }})
{{ end -}}
{{ end -}}
{{- end }}
