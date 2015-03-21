package rkit

import (
	"io/ioutil"
	"net/http"
)

// LoadFile loads a file from data folder of an app. All resorces should be kept
// in data folder.
//
// Each file, eg icon would be loaded separately. which for http like scenario
// is suboptimal, so init() of this module loads a special resource called
// resources.rbndl, which is generated using github.com/amitu/rkit/rbndl tool.
//
// If the file is present in resources.rbndl, then it will be returned from
// there else an http request would be made. to fetch it.
//
// evey call to LoadFile() in later case would make a http request, with cache
// busting technique if second flag is set to false. in case the resource is
// available in rbndl file, cache flag would be ignored.
//
// rbndl will contain images, fonts, any other static resource one may need.
func LoadFile(filename string, cache bool) ([]byte, error) {
	resp, err := http.Get("http://127.0.0.1:8877/data" + filename)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// LoadBundle can be used to load other bundles. By default rkit will download
// and cache a bundle names resources.rbndl. Some programs may want to load a
// separate bundle on demand, eg when showing a not often used UI.
//
// The loaded bundle would be merged into existing bundles, and further calls
// to LoadFile() would have access to files loaded in this bundle.
//
// This is a blocking, synchronous call.
func LoadBundle(bundlename string) error {
	return nil
}
