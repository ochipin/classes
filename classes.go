package classes

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// New : Classes 構造体を生成する
func New(dirname string) *Classes {
	return &Classes{
		Dirname:    dirname,
		Func:       func(f os.FileInfo) bool { return true },
		Mode:       0644,
		ExportOnly: true,
	}
}

// Classes : ソースコード内で定義されている構造体名一覧を取得する構造体
type Classes struct {
	Dirname    string
	Func       func(os.FileInfo) bool
	Mode       parser.Mode
	ExportOnly bool
	fset       *token.FileSet
}

// Classlist : 構造体一覧を返却する
func (class *Classes) Classlist() ([]string, error) {
	var clsnames []string

	class.fset = token.NewFileSet()
	pkgs, err := parser.ParseDir(class.fset, class.Dirname, class.Func, class.Mode)
	if err != nil {
		return nil, err
	}

	// app.go, base.go, ... などのパッケージ内のソースをすべて網羅する
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			// ソースコード上で定義されている構造体一覧を取得
			result := class.analyze(file)
			// 構造体一覧を結合
			if result != nil {
				clsnames = append(clsnames, result...)
			}
		}
	}

	// 構造体一覧を返却する
	return clsnames, nil
}

func (class *Classes) analyze(file *ast.File) []string {
	var result []string

	for _, decl := range file.Decls {
		switch td := decl.(type) {
		// 構造体の場合
		case *ast.GenDecl:
			if td.Tok != token.TYPE {
				continue
			}
			for _, spec := range td.Specs {
				name := class.spec(spec)
				if name != "" {
					result = append(result, name)
				}
			}
		}
	}
	return result
}

func (class *Classes) spec(spec ast.Spec) string {
	var result string
	s := spec.(*ast.TypeSpec)

	switch s.Type.(type) {
	// 構造体の場合
	case *ast.StructType:
		// エクスポートされていない構造体は非対象とする
		if class.ExportOnly && s.Name.IsExported() == false {
			return ""
		}
		result = s.Name.String()
	}
	return result
}
