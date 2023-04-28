package gogen

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
	"strings"
)

//go:embed sub-main.tpl
var subMainTemplate string

func genSubMain(dir, rootPkg, moduleName string, cfg *config.Config, api *spec.ApiSpec) error {
	arr := strings.Split(rootPkg, "/")
	packageName := arr[len(arr)-1]

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          "",
		filename:        packageName + ".go",
		templateName:    "subMainTemplate",
		category:        category,
		templateFile:    subMainTemplateFile,
		builtinTemplate: subMainTemplate,
		data: map[string]string{
			"importPackages": genSubMainImports(rootPkg, moduleName),
			"packageName":    packageName,
		},
	})
}

func genSubMainImports(parentPkg, moduleName string) string {
	var imports []string
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, handlerDir)))
	imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(moduleName, contextDir)))
	imports = append(imports, fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceURL))
	return strings.Join(imports, "\n\t")
}
