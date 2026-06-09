package main

import(
	"html/template"
	"fmt"
	"chatapp.new.net/ui"
	"path"
	"bytes"
	"net/http"
)
const tplDir = "html"
const tplExt = ".html"

func mustParseTemplates(pages ...string) *template.Template {

	files := pages

	files = normalize(tplDir, tplExt, files...)

	ts, err := template.ParseFS(ui.Files, files...)
	if err != nil {
		panic(fmt.Sprintf(
			"failed to parse templates (%q): %s",
			pages[0],
			err.Error(),
		))
	}
	return ts
}

func normalize(baseDir, ext string, files ...string) []string {

	for i, f := range files {
		if path.Ext(f) == "" {
			files[i] += tplExt
		}
		files[i] = path.Join(baseDir, files[i])
	}

	return files
}



func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	cache["room.html"] = mustParseTemplates("pages/room")

	return cache, nil
}
func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data any) error {

	ts, ok := app.templateCache[page]
    if !ok {
        err := fmt.Errorf("the template %s does not exist", page)
        app.logger.Error(err.Error())
        return err
    }

    buf := new(bytes.Buffer)
    err := ts.Execute(buf, data)
    if err != nil {
        return err 
    }

    w.WriteHeader(status)
    buf.WriteTo(w)

    return nil
}

