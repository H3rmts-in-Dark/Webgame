package serve

import (
	"regexp"

	"Server/logging"
)

var mimeTypes map[*regexp.Regexp]string

func LoadMime() {
	typers := logging.LoadMimeTypes()
	mimeTypes = make(map[*regexp.Regexp]string, len(typers))
	for ext, typ := range typers {
		mimeTypes[regexp.MustCompile(ext)] = typ
	}
	logging.Log(logging.SERVE, "Loaded Mime Types")
}

func getMime(filename string) (string, bool) {
	for r, s := range mimeTypes {
		if r.Match([]byte(filename)) {
			return s, true
		}
	}
	return "", false
}
