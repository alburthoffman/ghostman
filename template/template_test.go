package template_test

import (
	"fmt"
	"github.com/newly12/ghostman/template"
)

func ExampleParseTemplate() {
	tmpl := template.ParseTemplate("../examples/templates/get/root.yml")
	fmt.Println(tmpl.Usage)
	fmt.Println(tmpl.Url)
	fmt.Println(tmpl.Protc)
	fmt.Println(tmpl.Port)
	fmt.Println(tmpl.Body)

	// Output:
	// access /
	// /
	// http
	// 80
	//
}