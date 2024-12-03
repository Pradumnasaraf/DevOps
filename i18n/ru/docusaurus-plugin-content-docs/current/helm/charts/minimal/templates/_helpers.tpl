{{- define "envShort" -}}
{{- if eq .Values.environment "production" -}}
prod
{{- else -}}
non-prod
{{- end -}}
{{- end -}}
