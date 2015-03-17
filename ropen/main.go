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

	cmd := exec.Command("open", tmp.Name())
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
		return
	}
}
