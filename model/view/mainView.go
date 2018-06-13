package view

import "html/template"

type MainViewTemplate struct {
	Body       template.HTML
	Css        []string
	JavaScript []string
}
