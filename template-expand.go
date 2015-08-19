// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// Take a template and json data file and execute the template.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	templateFileName := flag.String("t", "", "template file")
	inputFileName := flag.String("i", "", "json data file name")

	flag.Parse()

	if *templateFileName == "" || *inputFileName == "" {
		fmt.Fprintf(os.Stderr, "See http://godoc.org/text/template for template help.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	df, err := os.Open(*inputFileName)
	e("Failed to open input file", err)
	defer df.Close()

	tt, err := ioutil.ReadFile(*templateFileName)
	e("Failed to read template file", err)
	t, err := template.New("").Parse(string(tt))
	e("Failed template file", err)

	coder := json.NewDecoder(df)
	obj := make(map[string]interface{})
	err = coder.Decode(&obj)
	e("Failed decode json data file", err)
	err = t.Execute(os.Stdout, obj)
	e("Failed executing template", err)
}

func e(ctx string, err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "See http://godoc.org/text/template for template help.\n")
	fmt.Fprintf(os.Stderr, "%s: %v\n", ctx, err)
	os.Exit(2)
}
