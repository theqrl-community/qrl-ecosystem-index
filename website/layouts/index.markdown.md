# {{ site.Title }}

> {{ site.Params.description }}

The QRL Ecosystem Index is a community-maintained directory of projects, tools, services, and resources connected to QRL 2.0. Listings are informational and do not imply endorsement, audit, or affiliation.

## Core Resources

{{ with site.GetPage "/projects" }}{{ with .OutputFormats.Get "markdown" -}}
- [Project directory]({{ .Permalink }})
{{- end }}{{ end }}
{{- with site.GetPage "/about" }}{{ with .OutputFormats.Get "markdown" }}
- [About the index]({{ .Permalink }})
{{- end }}{{ end }}
{{- with site.GetPage "/ideas" }}{{ with .OutputFormats.Get "markdown" }}
- [Builder ideas]({{ .Permalink }})
{{- end }}{{ end }}
- [Structured project index]({{ "index.json" | absURL }})

## Active Projects

{{ $projects := sort (where (where site.RegularPages "Section" "projects") "Params.status" "!=" "archived") "Title" -}}
{{- range $projects }}
{{- $project := . }}
{{- with $project.OutputFormats.Get "markdown" }}
- [{{ $project.Title }}]({{ .Permalink }}): {{ $project.Params.description | plainify }}
{{- end }}
{{- end }}
