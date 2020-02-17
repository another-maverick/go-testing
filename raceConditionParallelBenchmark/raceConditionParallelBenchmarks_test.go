package main

import (
	"bytes"
	"testing"
	"text/template"
)

// To run each of the following on 1,2,3,4 CPU's and check for race condition execute --  "go test -race -bench . -cpu=1,2,3,4"

// This runs several goroutines in parallel. Here you would see difference when the code is run with more processors because we are making use of parallel benchmark
func BenchmarkParallelTesting(b *testing.B) {
	b.Logf("the number of iterations is %d \n", b.N)
	myTemplate := "Warriors Guard is  {{.Name}}"
	parsedTemplate, _ := template.New("test").Parse(myTemplate)
	templateData := &map[string]string{
		"Name": "Steph Curry",
	}
	var buf bytes.Buffer
	b.RunParallel(func(pb *testing.PB){
		//Moved the below line outside the loop to force Race condition
		//var buf bytes.Buffer
		for pb.Next(){
			parsedTemplate.Execute(&buf, templateData)
			buf.Reset()
		}
	})
}

