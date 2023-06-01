package service

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/hilaily/lib/cmdx"
	"github.com/urfave/cli/v2"

	"github.com/hilaily/cmds/modtool/common"
)

var (
	count = 0
)

// RenameCommand ...
func RenameCommand() *cli.Command {
	return &cli.Command{
		Name:   "rename",
		Usage:  "rename a mod, modtool rename <old_name> <new_name>",
		Action: rename,
	}
}

func rename(ctx *cli.Context) error {
	oldName := ctx.Args().Get(0)
	newName := ctx.Args().Get(1)

	if oldName == "" || newName == "" {
		cmdx.Throw("you should specify a new mod name")
	}
	if oldName == newName {
		cmdx.Throw("old name and new name are same")
	}

	modFile, err := common.FindModFile()
	cmdx.CheckErr(err)
	// 解析命令行参数
	dir := filepath.Dir(modFile)
	cmdx.CheckErr(err)

	cmdx.Green("find these info:")
	cmdx.Green("  mod directory: %s", dir)
	cmdx.Green("  old mod name: %s", oldName)
	cmdx.Green("  new mod name: %s", newName)

	// 遍历指定目录下的所有 Go 源代码文件
	err = dealDir(dir+"/", oldName, newName)
	cmdx.CheckErr(err)
	cmdx.Green("rename finished, %d files changed", count)
	return nil
}

func dealDir(dir string, oldName, newName string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		return dealOneFile(path, oldName, newName)
	})
	return err
}

func dealOneFile(path string, oldName, newName string) error {
	// 解析 Go 源代码文件为 AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// 遍历 AST 中的所有 ImportSpec 节点，找到需要替换的包名并进行替换
	newImports := make([]*ast.ImportSpec, 0, len(node.Imports))
	for _, vv := range node.Imports {
		imp := vv
		if strings.HasPrefix(imp.Path.Value, "\""+oldName) {
			imp.Path.Value = strings.ReplaceAll(imp.Path.Value, oldName, newName)
		}
		newImports = append(newImports, imp)
	}
	node.Imports = newImports

	// 将修改后的 AST 输出为 Go 源代码文件
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	count++
	return (&printer.Config{Tabwidth: 8, Mode: printer.UseSpaces | printer.TabIndent}).Fprint(out, fset, node)
}
