{{- define "opstack.name" -}}
opstack
{{- end }}

{{- define "opstack.fullname" -}}
{{ include "opstack.name" . }}-{{ .Release.Name }}
{{- end }}
