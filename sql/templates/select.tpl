SELECT {{range $colIndex, $column := .Columns}}{{if $colIndex}}, {{end}}{{if eq $column "*"}}{{$column}}{{else}}`{{$column}}`{{end}}{{end}} FROM `{{.Table}}{{template "where" .}}`
