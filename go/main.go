package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/GeertJohan/go.rice"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"

	pdf "github.com/adrg/go-wkhtmltopdf"
	"github.com/voxelbrain/goptions"
)

func checkDir(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	return nil
}

func md2html(input, output string) error {
	var Temp struct {
		Title string
		Body  template.HTML
	}

	if err := checkDir(filepath.Dir(output)); err != nil {
		return err
	}

	r, err := os.Open(input)

	if err != nil {
		return err
	}

	defer r.Close()

	content, err := ioutil.ReadAll(r)

	if err != nil {
		return err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.HardLineBreak | parser.MathJax | parser.Tables | parser.FencedCode
	pars := parser.NewWithExtensions(extensions)

	Temp.Title = filepath.Base(input)
	Temp.Body = template.HTML(string(markdown.ToHTML(content, pars, nil)))

	box, err := rice.FindBox("views")
	if err != nil {
		return err
	}

	htmlContent, err := box.String("md2html.html")
	if err != nil {
		return err
	}

	tmpl, err := template.New(Temp.Title).Parse(htmlContent)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	return tmpl.Execute(f, Temp)
}

func html2pdf(input, output string) error {
	pdf.Init()
	defer pdf.Destroy()

	// Create object from file
	object, err := pdf.NewObject(input)
	if err != nil {
		return err
	}

	if err := object.SetOption("footer.right", "[page]"); err != nil {
		return err
	}

	// Create converter
	converter := pdf.NewConverter()
	defer converter.Destroy()

	// Add created objects to the converter
	converter.AddObject(object)

	// Add converter options
	if err := converter.SetOption("margin.left", "10mm"); err != nil {
		return err
	}
	if err := converter.SetOption("margin.right", "10mm"); err != nil {
		return err
	}
	if err := converter.SetOption("margin.top", "10mm"); err != nil {
		return err
	}
	if err := converter.SetOption("margin.bottom", "10mm"); err != nil {
		return err
	}

	// Convert the objects and get the output PDF document
	content, err := converter.Convert()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(output, content, 0755)
}

func main() {

	options := struct {
		Input  string `goptions:"-i, --input, description='Path to input markdown file', obligatory"`
		Output string `goptions:"-o, --output, description='Path to output html file'"`
		Pdf    string `goptions:"-p, --pdf, description='Path to pdf, using Python with pdfkit module '"`
	}{}

	goptions.ParseAndFail(&options)

	if len(os.Args) <= 1 {
		goptions.PrintHelp()
		os.Exit(0)
	}

	re := regexp.MustCompile(`(md|markdown)$`)

	output := options.Output
	if options.Output == "" {
		output = re.ReplaceAllString(options.Input, "html")
	}

	if err := md2html(options.Input, options.Output); err != nil {
		log.Fatal(err)
	} else {
		output = options.Output
	}

	if options.Pdf != "" {
		if err := html2pdf(output, options.Pdf); err != nil {
			log.Fatal(err)
		}
	}
}
