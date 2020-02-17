package main

import (
	"bytes"
	"testing"
	"text/template"
)

// To run each of the following on 1,2,3,4 CPU's execute --  "go test -bench . -cpu=1,2,3,4"

// Here the template gets compiled with every iteration. Takes longer
// Here you would NOT see difference when the code is run with more processors because we are not making use of parallel benchmark
func BenchmarkTesting(b *testing.B) {
	b.Logf("the number of iterations is %d \n", b.N)
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

// Here template is compiled ahead of time and not for every iteration. Runs much faster
// Here you would NOT see difference when the code is run with more processors because we are not making use of parallel benchmark
func BenchmarkOptimizedTesting(b *testing.B) {
	b.Logf("the number of iterations is %d \n", b.N)
	myTemplate := "Warriors Guard is  {{.Name}}"
	parsedTemplate, _ := template.New("test").Parse(myTemplate)
	templateData := &map[string]string{
		"Name": "Steph Curry",
	}
	var buf bytes.Buffer

	for i:=0; i< b.N; i++ {
		parsedTemplate.Execute(&buf, templateData)
		buf.Reset()
	}
}

// This runs several goroutines in parallel. Here you would see difference when the code is run with more processors because we are making use of parallel benchmark
func BenchmarkParallelTesting(b *testing.B) {
	b.Logf("the number of iterations is %d \n", b.N)
	myTemplate := "Warriors Guard is  {{.Name}}"
	parsedTemplate, _ := template.New("test").Parse(myTemplate)
	templateData := &map[string]string{
		"Name": "Steph Curry",
	}
	b.RunParallel(func(pb *testing.PB){
		var buf bytes.Buffer
		for pb.Next(){
			parsedTemplate.Execute(&buf, templateData)
			buf.Reset()
		}
	})
}

