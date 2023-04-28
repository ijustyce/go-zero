package gogen

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
	"strings"
)

//go:embed global-handler.tpl
var globalHandlerTemplate string

func genGlobalHandler(dir string, cfg *config.Config, api *spec.ApiSpec) error {

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "rest",
		filename:        "globalHandler.go",
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
	imports = append(imports, "\"net/http\"")
	return strings.Join(imports, "\n\t")
}
