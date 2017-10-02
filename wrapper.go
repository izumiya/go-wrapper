package wrapper

import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"regexp"
)

type Wrapper struct {
	Body   string
	Labels []string
}

var diff *diffmatchpatch.DiffMatchPatch = diffmatchpatch.New()

var label *regexp.Regexp = regexp.MustCompile(`{{(.+?)}}`)

var placeholder string = "\x01"

func New(body string) *Wrapper {
	return new(Wrapper).load(body, label)
}

func (w *Wrapper) load(body string, re *regexp.Regexp) *Wrapper {
	w.Body = re.ReplaceAllString(body, placeholder)
	w.Labels = []string{}
	for _, v := range re.FindAllStringSubmatch(body, -1) {
		w.Labels = append(w.Labels, v[1])
	}
	return w
}

type Map struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (w *Wrapper) Extract(src string) []Map {
	i := 0
	e := []Map{}
	m := Map{}
	for _, v := range diff.DiffMain(w.Body, src, false) {
		if v.Type == 0 {
			continue
		}
		if v.Type == -1 {
			m.Key = w.Labels[i]
			i++
			continue
		}
		if v.Type == 1 {
			m.Value = v.Text
			e = append(e, m)
			m = Map{}
		}
	}
	return e
}
