package http

const DefaultTemplate = `
{{- if eq .A.Status "firing"}}A{{else}}r{{end}}
{{index .A.Labels "instance"}}:
{{index .A.Labels "job"}},
{{index .A.Labels "alertname" -}}
{{with $s := index .A.Annotations "summary"}}: {{$s}}{{end}}
{{if eq .A.Status "firing"}}(since {{.A.Since}})
{{- if ne .R.Name ""}}{{if not .R.Active}} - {{.R.Name}}{{end}}{{else}}{{if .H}} - {{.A.Responders}}{{end}}{{end -}}
{{- else}}(lasted {{.A.Lasted}})
{{- end}}`
