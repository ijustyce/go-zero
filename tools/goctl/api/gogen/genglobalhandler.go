package gogen

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

//go:embed global-handler.tpl
var globalHandlerTemplate string

func genGlobalHandler(dir string, cfg *config.Config, api *spec.ApiSpec) error {

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "rest",
		filename:        "global_handler.go",
		templateName:    "globalHandlerTemplate",
		category:        category,
		templateFile:    globalHandlerTemplateFile,
		builtinTemplate: globalHandlerTemplate,
		data: map[string]string{
			"importPackages": genGlobalHandlerImports(),
		},
	})
}

func genGlobalHandlerImports() string {
	var imports []string
	imports = append(imports, fmt.Sprintf("\"%s/rest/httpx\"", vars.ProjectOpenSourceURL))
	return strings.Join(imports, "\n\t")
}
