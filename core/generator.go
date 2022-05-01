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
	"os"
	"path"
	"strconv"
	"strings"
)

func Generate(src string, dest string, fileName string) {
	outputs, err := ParseFile(src)
	if err != nil {
		log.Fatal(err)
	}
	var firstBuffer bytes.Buffer
	e := json.NewEncoder(&firstBuffer)
	e.SetIndent("", "  ")
	err = e.Encode(outputs)
	if err != nil {
		log.Fatal(err)
	}
	if verbose {
		println("generate output:")
		println(firstBuffer.String())
	}
	dest = dest + "/" + fileName
	destDir := path.Dir(dest)
	_ = os.MkdirAll(destDir, fs.ModePerm)
	writeError := ioutil.WriteFile(dest, firstBuffer.Bytes(), fs.ModePerm)
	if writeError != nil {
		log.Fatal(writeError)
	}
}

// ParseFile 扫描指定目录下的指定文件后缀的go源文件,解析注释
// 参考文章: https://www.dazhuanlan.com/zhangzw/topics/1445991
func ParseFile(src string) (map[string]string, error) {
	fileSet := token.NewFileSet()
	//解析结果map
	result := make(map[string]string)
	//根据src类型分别处理文件和目录
	fileInfo, _ := os.Stat(src)
	if fileInfo.IsDir() {
		parsedPackages, err := parser.ParseDir(fileSet, src, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		for _, packageAST := range parsedPackages {
			for _, astFile := range packageAST.Files {
				err = parseSingleFile(fileSet, astFile, result)
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		astFile, err := parser.ParseFile(fileSet, src, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		err = parseSingleFile(fileSet, astFile, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func parseSingleFile(fileSet *token.FileSet, astFile *ast.File, result map[string]string) error {
	// 类型检查, 得到常量的值
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("", fileSet, []*ast.File{astFile}, nil)
	if err != nil {
		return err
	}
	for _, v := range astFile.Scope.Objects {
		if v.Kind == ast.Con {
			d := v.Decl.(*ast.ValueSpec)
			var label = pkg.Scope().Lookup(v.Name).(*types.Const).Val().String()
			num, _ := strconv.ParseInt(label, 10, 64)
			label = fmt.Sprintf("0x%09x", num)
			value := strings.TrimSpace(d.Doc.Text())
			// go推荐注释前面带变量名, 去除
			if strings.HasPrefix(value, v.Name) {
				value = value[len(v.Name)+1:]
			}
			value = strings.TrimSpace(value)
			if value == "ignore" {
				continue
			}
			result[label] = value
		}
	}
	return nil
}
