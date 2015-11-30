UPDATE `{{.Table}}` SET {{range $colIndex, $column := .Columns}}{{if $colIndex}}, {{end}}`{{$column}}`='{{index $.Values $colIndex}}'{{end}}
