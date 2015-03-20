package rkit

import (
	"io/ioutil"
	"net/http"
)

// LoadFile loads a file from data folder of an app. All resorces should be kept
// in data folder.
func LoadFile(filename string) ([]byte, error) {
	resp, err := http.Get("http://127.0.0.1:8877/data" + filename)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
