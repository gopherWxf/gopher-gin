package goft

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	file := dir + "/application.yaml"
	b, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}

/*
	自动生成一个 funcmap.go
	goft.GenTplFunc("src/funcs")
*/
// GenTplFunc 根据 函数生成 funcMap   ast 语法树
func GenTplFunc(path string) {
	path = strings.Replace(path, "\\", "/", -1)
	pList := strings.Split(path, "/")
	pkgName := pList[len(pList)-1]
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	funcMap := make(map[string]string)
	impList := make(map[string]struct{})
	for _, f := range dir {
		fp := path + "/" + f.Name()
		if f.IsDir() || filepath.Ext(fp) != ".go" || f.Name() == "funcmap.go" {
			continue
		} else {
			fset := token.NewFileSet()
			ast_file, err := parser.ParseFile(fset, fp, nil, 0|parser.ParseComments)
			if err != nil {
				log.Fatal(err)
			}
			for _, imp := range ast_file.Imports {
				impName := ""
				if imp.Name != nil {
					impName = imp.Name.Name
				}
				impList[fmt.Sprintf("%s %s", impName, imp.Path.Value)] = struct{}{}
			}
			for _, dec := range ast_file.Decls {
				if fn, ok := dec.(*ast.FuncDecl); ok {
					var output []byte
					buffer := bytes.NewBuffer(output)
					newf := &ast.FuncDecl{
						Doc:  fn.Doc,
						Name: &ast.Ident{Name: ""},
						Body: fn.Body,
						Recv: fn.Recv,
						Type: fn.Type,
					}
					format.Node(buffer, fset, newf)
					funcMap[fn.Name.String()] = buffer.String()
				}
			}
		}
	}
	tpl := `
// Code generated  DO NOT EDIT
package {{.pkg}}
import (
 {{range $k, $v := .import }}
    {{$k}}
 {{end}}
 )
var FuncMap=map[string]interface{}{
 {{range $key, $value := .map }}
   "{{$key}}" :{{$value}},
  {{end}}
}
			`
	buf := bytes.Buffer{}
	tmpl, _ := template.New("funcMap").Parse(tpl)
	_ = tmpl.Execute(&buf, gin.H{"map": funcMap, "pkg": pkgName, "import": impList})
	autocode_file := path + "/funcmap.go"
	autocode, err := os.OpenFile(autocode_file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Fprint(autocode, buf.String()))
}
