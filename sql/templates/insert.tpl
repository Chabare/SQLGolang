INSERT INTO `{{.Table}}`({{range $index, $column := .Columns}}{{if $index}}, {{end}}`{{$column}}`{{end}}) VALUES ({{range $index, $value := .Values}}{{if $index}}, {{end}}'{{$value}}'{{end}})
