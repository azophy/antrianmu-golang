// adapted from https://echo.labstack.com/guide/templates/
package common

import (
  "io"
  "os"
  "bytes"
  "strings"
  "path/filepath"
  "html/template"

	"github.com/labstack/echo/v4"
  "antrianmu-golang/web/config"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
  baseTemplate string
}

// crreate customized template renderer that could handle template in root dir
// adapted from:
// - https://stackoverflow.com/a/50581032/2496217
// - https://gist.github.com/logrusorgru/abd846adb521a6fb39c7405f32fec0cf
func NewTemplateRenderer(rootDir, extension, baseTemplate string) *TemplateRenderer {
    cleanRoot := filepath.Clean(rootDir)
    pfx := len(cleanRoot)+1
    root := template.New("")

    _ = filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
        if !info.IsDir() && strings.HasSuffix(path, extension) {
            if e1 != nil {
                return e1
            }

            b, e2 := os.ReadFile(path)
            if e2 != nil {
                return e2
            }

            name := path[pfx:]
            t := root.New(name)
            _, e2 = t.Parse(string(b))
            if e2 != nil {
                return e2
            }
        }

        return nil
    })

  return &TemplateRenderer{
      templates: root,
      baseTemplate: baseTemplate,
  }
}

// Render renders a template document with base tempplate
// this involves 2-stage rendering: render the blocks, that putting it to baseTemplate
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

  // render contenct block
  var blockContent bytes.Buffer
  if err := t.templates.ExecuteTemplate(&blockContent, name, data); err != nil {
      return err
  }

  // render base template
	return t.templates.ExecuteTemplate(w, t.baseTemplate, map[string]interface{}{
    "APP_URL": config.ConfAppUrl,
    "page_title": data.(map[string]interface{})["title"],
    "content_block": template.HTML(blockContent.String()),
  })
}

