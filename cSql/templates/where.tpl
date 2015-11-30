 {{define "where"}} WHERE {{with .Conditions}}{{range $condIndex, $cond := .}}{{if $condIndex}} AND {{end}}`{{$cond.Column}}`='{{$cond.Value}}'{{end}}{{end}}{{end}}
