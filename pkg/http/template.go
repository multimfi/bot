package http

const DefaultTemplate = `
{{- if eq .A.Status "firing"}}A{{else}}r{{end}}
{{index .A.Labels "instance"}}:
{{index .A.Labels "job"}},
{{index .A.Labels "alertname" -}}
{{with $s := index .A.Annotations "summary"}}: {{$s}}{{end}}
{{if eq .A.Status "firing"}}(since {{.A.Since}})
{{- if ne .R.Name ""}} - {{.R.Name}}{{else}} - {{.A.Responders}}{{end -}}
{{- else}}(lasted {{.A.Lasted}})
{{- end}}`
