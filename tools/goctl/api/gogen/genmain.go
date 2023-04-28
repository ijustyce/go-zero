package gogen

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

//go:embed main.tpl
var mainTemplate string

func genMain(dir, rootPkg, moduleName string, cfg *config.Config, api *spec.ApiSpec) error {
	_, serviceName := filepath.Split(rootPkg)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        strings.ReplaceAll(moduleName, "-", "_") + ".go",
		templateName:    "mainTemplate",
		category:        category,
		templateFile:    mainTemplateFile,
		builtinTemplate: mainTemplate,
		data: map[string]string{
			"importPackages": genMainImports(rootPkg, moduleName),
			"packageName":    "main",
			"serviceName":    serviceName,
		},
	})
}

func genMainImports(parentPkg, moduleName string) string {
	var imports []string
	imports = append(imports, fmt.Sprintf("\"%s\"", parentPkg))
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(moduleName, configDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(moduleName, contextDir)))
	imports = append(imports, fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceURL))
	imports = append(imports, fmt.Sprintf("\"%s/core/conf\"", vars.ProjectOpenSourceURL))
	return strings.Join(imports, "\n\t")
}
