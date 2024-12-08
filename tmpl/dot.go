package tmpl

var DotTmpl = `digraph tepl{
{{range $pkgver,$node := .NodeMap -}}
	{{printf "	%#p[label=\"%s\", shape=\"box\"]\n" $node $pkgver}}
{{- end}}

{{range $x := .Paths -}}
	{{printf "	%s\n" $x}}
{{- end}}
}
`
