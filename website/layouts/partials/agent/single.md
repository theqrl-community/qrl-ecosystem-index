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
