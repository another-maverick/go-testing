package main

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func BenchmarkTesting(b *testing.B) {
	fmt.Printf("the number of iterations is %d \n", b.N)
	myTemplate := "Warriors Guard is  {{.Name}}"
	templateData := &map[string]string{
		"Name": "Steph Curry",
	}
	var buf bytes.Buffer

	for i:=0; i< b.N; i++ {
		parsedTemplate, _ := template.New("test").Parse(myTemplate)
		parsedTemplate.Execute(&buf, templateData)
		buf.Reset()
	}
}
