// Command ropen is used to create a temporary html file containing a js file, and
// launch browser to open that html file.
//
// Since gopher compiler generates a .js file, we need a html file to embed that
// js, and have to debug the code by opening that html file, and since we may have
// to change the name of js file in html file often, and since this grows tiring
// fast, this command exists.
//
// Usage:
//
// 	$ gopherjs build && ropen
//
// Gopherjs will build a js file named after directory, ropen will follow the same.
//
// Alternately:
//
// 	$ gopherjs build test.go && ropen test
//
// This will lead to test.js file being generated, and ropen will take test on
// command line and create html page with test.js.
//
// TODO:
//
// 	- support linux(xdg-open), windows?
// 	- command line flag to specify which browser to use, currently chrome only
//
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var (
	addr = flag.String(
		"http", "127.0.0.1:8877", "HTTP Server host:port.",
	)
	js string
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(js)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, `
		<html>
			<head>
				<title>foo</title>
			</head>
			<body>
				<script>%s</script>
			</body>
		</html>
	`, string(data))
}

func main() {
	flag.Parse()

	js = flag.Arg(0)
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

	go func() {
		time.Sleep(time.Millisecond * 100)
		cmd := exec.Command("open", "-a", "Google Chrome", "http://"+*addr)
		err = cmd.Run()

		if err != nil {
			panic(err)
			return
		}
	}()

	http.HandleFunc("/", serveIndex)
	http.ListenAndServe(*addr, nil)
}
