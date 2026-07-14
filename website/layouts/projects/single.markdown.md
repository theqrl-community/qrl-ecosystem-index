# {{ .Title }}

> {{ .Params.description | plainify }}

{{ with .OutputFormats.Get "html" -}}
- Canonical page: [{{ $.Title }}]({{ .Permalink }})
{{- end }}
- Status: {{ .Params.status }}
- Project type: {{ .Params.project_type }}
- Category: {{ .Params.category }}
{{- with .Params.tags }}
- Tags: {{ delimit . ", " }}
{{- end }}
- Maintainer: {{ .Params.author }}
- License: {{ .Params.license }}
- Open source: {{ if .Params.open_source }}yes{{ else }}no{{ end }}
- Listed as audited: {{ if .Params.audited }}yes{{ else }}no{{ end }}
- Created: {{ .Params.created }}
- Updated: {{ .Params.updated }}
{{- with .Params.network }}
- Network: {{ . }}
{{- end }}
{{- with .Params.contract_address }}
- Contract address: `{{ . }}`
{{- end }}
{{- with .Params.token }}
- Token: {{ . }}
{{- end }}
{{- with .Params.platforms }}
- Platforms: {{ delimit . ", " }}
{{- end }}
{{- with .Params.supported_networks }}
- Supported networks: {{ delimit . ", " }}
{{- end }}
{{- with .Params.languages }}
- Languages: {{ delimit . ", " }}
{{- end }}
{{- with .Params.language }}
- Language: {{ . }}
{{- end }}

## Links

{{ with .Params.project_url -}}
- [Project website]({{ . }})
{{- end }}
{{- with .Params.docs }}
- [Documentation]({{ . }})
{{- end }}
{{- with .Params.github }}
- [Source repository]({{ . }})
{{- end }}
{{- with .Params.discord }}
- [Discord]({{ . }})
{{- end }}
{{- with .Params.twitter }}
- [Social profile]({{ . }})
{{- end }}
{{- with strings.TrimSpace .RawContent }}

## Overview

{{ . }}
{{- end }}
{{- with .Params.features }}

## Capabilities

{{ range . -}}
- {{ . }}
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
{{- with .Params.clients }}

## Clients

{{ range . -}}
{{ $client := . -}}
- **{{ $client.platform }}**{{ if $client.default }} (default){{ end }}{{ with $client.url }}: [Open client]({{ . }}){{ end }}{{ with $client.github }}{{ if $client.url }}; {{ else }}: {{ end }}[Source]({{ . }}){{ end }}
{{ end -}}
{{- end }}
{{- with .Params.audits }}

## Audits

{{ range . -}}
- [{{ .auditor }}]({{ .audit_url }})
{{ end -}}
{{- end }}
{{- with .Params.endpoints }}

## Endpoints

{{ range . -}}
- {{ . }}
{{ end -}}
{{- end }}
