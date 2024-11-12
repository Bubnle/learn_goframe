package main

import (
	_ "learn_goframe/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"learn_goframe/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
