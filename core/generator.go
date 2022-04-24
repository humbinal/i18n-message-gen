package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/fs"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Gen 扫描指定目录下的指定文件后缀的go源文件,解析注释
// 参考文章: https://www.dazhuanlan.com/zhangzw/topics/1445991
func Gen(src string, dest string) {
	outputs, _ := ParseFile(src + "/errorcode.go")
	var firstBuffer bytes.Buffer
	e := json.NewEncoder(&firstBuffer)
	e.SetIndent("", "  ")
	err := e.Encode(outputs)
	if err != nil {
		log.Fatal(err)
	}
	println(firstBuffer.String())
	writeError := ioutil.WriteFile(dest, firstBuffer.Bytes(), fs.ModePerm)
	if writeError != nil {
		log.Fatal(writeError)
	}
}

func ParseFile(fileName string) (map[string]string, error) {
	fileSet := token.NewFileSet()
	// 解析文件，主要解析token
	// TODO 使用ParseDir
	f, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// 类型检查, 得到常量的值
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("", fileSet, []*ast.File{f}, nil)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for _, v := range f.Scope.Objects {
		if v.Kind == ast.Con {
			d := v.Decl.(*ast.ValueSpec)
			var label = pkg.Scope().Lookup(v.Name).(*types.Const).Val().String()
			num, _ := strconv.ParseInt(label, 10, 64)
			label = fmt.Sprintf("0x%09x", num)
			value := d.Doc.Text()
			value = strings.TrimSpace(value)
			if value == "ignore" {
				continue
			}
			m[label] = value
		}
	}
	return m, nil
}
