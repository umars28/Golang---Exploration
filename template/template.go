package template

import (
	"html/template"
	"path"
)

func TemplateHtml() {
	template.ParseFiles(
		path.Join("views/mahasiswa", "mahasiswa.html"),
		path.Join("views/template", "main.html"),
		path.Join("views/template", "header.html"),
		path.Join("views/template", "sidebar.html"),
		path.Join("views/template", "footer.html"),
	)
}
