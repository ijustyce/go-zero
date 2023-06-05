package generator

import (
	_ "embed"
	"fmt"
	"path/filepath"

	conf "github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed client.tpl
var clientTemplateText string

func (g *Generator) GenClient(ctx DirContext, proto parser.Proto, cfg *conf.Config,
	c *ZRpcContext) error {
	clientFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	fileName := filepath.Join(ctx.GetMain().Filename, fmt.Sprintf("%v_client.go", clientFilename))

	text, err := pathx.LoadTemplate(category, clientTemplateFile, clientTemplateText)
	if err != nil {
		return err
	}

	return util.With("client").GoFmt(false).Parse(text).SaveTo(map[string]any{}, fileName, false)
}
