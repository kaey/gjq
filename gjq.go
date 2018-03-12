package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"
)

var (
	inputF  = flag.String("i", "-", "file to read from (default stdin)")
	formatF = flag.String("f", "", "format file")
)

func main() {
	flag.Parse()

	outfmt := flag.Arg(0)
	var tmpl *template.Template
	if outfmt != "" {
		t, err := template.New("main").Funcs(funcMap).Parse(outfmt)
		if err != nil {
			fmt.Println("format error:", err)
			os.Exit(2)
		}
		tmpl = t
	}

	w := os.Stdout
	r := os.Stdin
	if *inputF != "-" {
		f, err := os.Open(*inputF)
		if err != nil {
			fmt.Println("open file error:", err)
			os.Exit(1)
		}
		r = f
	}

	d := json.NewDecoder(r)
	d.UseNumber()

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")

	for {
		m := make(map[string]interface{}, 20)
		if err := d.Decode(&m); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			os.Exit(1)
		}

		if tmpl == nil {
			if err := e.Encode(m); err != nil {
				fmt.Println("write error:", err)
				os.Exit(1)
			}

			continue
		}

		if err := tmpl.Execute(w, m); err != nil {
			fmt.Println("write error:", err)
			os.Exit(1)
		}
		//w.Write([]byte("\n"))
	}
}
