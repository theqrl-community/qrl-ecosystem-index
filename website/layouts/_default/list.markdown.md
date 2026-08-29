{{- $heading := .Title -}}
{{- $description := .Description -}}
{{- $categoryMeta := dict -}}
{{- if and .Data.Term (eq .Data.Plural "project-types") -}}
    {{- $heading = partial "project-type-label.html" .Data.Term -}}
{{- else if and .Data.Term (eq .Data.Plural "categories") -}}
    {{- $categoryMeta = partial "category-meta.html" .Data.Term -}}
    {{- $heading = $categoryMeta.category.label -}}
{{- else if and .Data.Term (eq .Data.Plural "capabilities") -}}
    {{- $capabilityMeta := partial "capability-meta.html" .Data.Term -}}
    {{- $heading = $capabilityMeta.label -}}
    {{- $description = $capabilityMeta.description -}}
{{- else if and .Data.Term (eq .Data.Plural "publishers") -}}
    {{- $description = printf "Projects published by %s." $heading -}}
{{- end -}}# {{ $heading }}

{{- with $description }}

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
