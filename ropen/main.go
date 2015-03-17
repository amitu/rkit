/*
ropen command is used to create a temporary html file, and launch browser to
open that html file.

Since gopher compiler generates a .js file, we need a html file to embed that
js, and have to debug the code by opening that html file, and since we may have
to change the name of js file in html file often, and since this grows tiring
fast, this command exists.

Usage:

	$ gopherjs build && ropen

Gopherjs will build a js file named after directory, ropen will follow the same.

Alternately:

	$ gopherjs build test.go && ropen test

This will lead to test.js file being generated, and ropen will take test on
command line and create html page with test.js.

TODO:

	- support linux(xdg-open), windows?
	- command line flag to specify which browser to use, currently chrome only

*/
package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var (
	HTMLTemplate *template.Template
)

func TempFile(prefix, suffix string) (*os.File, error) {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return os.Create(
		filepath.Join(
			os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix,
		),
	)
}

func init() {
	var err error
	HTMLTemplate, err = template.New("html").Parse(`
		<html>
			<head>
				<title>foo</title>
			</head>
			<body>
				<script src="file://{{.}}"></script>
			</body>
		</html>
	`,
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	js := flag.Arg(0)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	if js == "" {
		js = path.Base(dir)
	}

	if !strings.HasSuffix(js, ".js") {
		js = js + ".js"
	}

	if !strings.HasPrefix(js, "/") {
		js = filepath.Join(dir, js)
	}

	if _, err := os.Stat(js); os.IsNotExist(err) {
		fmt.Println("no such file:", js)
		return
	}

	tmp, err := TempFile("rkit-", ".html")
	if err != nil {
		fmt.Println(err)
		return
	}

	HTMLTemplate.Execute(tmp, js)

	cmd := exec.Command("open", "-a", "Google Chrome", tmp.Name())
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
		return
	}
}
