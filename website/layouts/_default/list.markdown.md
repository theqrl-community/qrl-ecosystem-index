# {{ .Title }}

{{- with .Description }}

> {{ . | plainify }}
{{- end }}
{{- with .OutputFormats.Get "html" }}

- Canonical page: [{{ $.Title }}]({{ .Permalink }})
{{- end }}
{{- with strings.TrimSpace .RawContent }}

{{ . }}
{{- end }}
{{- if .Data.Terms }}

## Terms

{{- range .Data.Terms.Alphabetical }}
{{- $term := .Page }}
{{- $count := .Count }}
{{- with $term.OutputFormats.Get "markdown" }}
- [{{ $term.Title }}]({{ .Permalink }}): {{ $count }} project{{ if ne $count 1 }}s{{ end }}.
{{- end }}
{{- end }}
{{ else }}
{{- $pages := sort .RegularPages "Title" }}
{{- if $pages }}

## Projects

{{- range $pages }}
{{- $page := . }}
{{- with $page.OutputFormats.Get "markdown" }}
- [{{ $page.Title }}]({{ .Permalink }}){{ with $page.Params.description }}: {{ . | plainify }}{{ end }}
{{- end }}
{{- end }}
{{- end }}
{{ end }}
