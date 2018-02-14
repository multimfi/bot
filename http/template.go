package http

const DefaultTemplate = `
{{- $since := 60000000000 }}
{{- if eq .A.Status "firing"}}A{{else}}r{{end}}
{{index .A.Labels "instance"}}:
{{index .A.Labels "job"}},
{{index .A.Labels "source"}},
{{index .A.Labels "alertname" -}}
{{with $s := index .A.Annotations "summary"}}: {{$s}}{{end}}
{{if eq .A.Status "firing"}}{{if gt .A.Since $since }}(since {{.A.Since}}){{end}}
{{- if ne .R.Name ""}}{{if not .R.Active}} - {{.R.Name}}{{end}}{{else}}{{if .H}} - {{.A.Responders}}{{end}}{{end -}}
{{- else}}(lasted {{.A.Lasted}})
{{- end}}`
