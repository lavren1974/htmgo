package run

import (
	"github.com/maddalax/htmgo/cli/tasks/astgen"
	"github.com/maddalax/htmgo/cli/tasks/css"
	"github.com/maddalax/htmgo/cli/tasks/process"
)

func Build() {
	astgen.GenAst(process.ExitOnError)
	css.GenerateCss(process.ExitOnError)
	process.RunOrExit("rm -rf ./dist")
	process.RunOrExit("mkdir -p ./dist/assets/dist")
	process.RunOrExit("cp -r ./assets/dist/* ./dist/assets/dist/")
	process.RunOrExit("go build -o \"./dist\" .")
	process.RunOrExit("echo \"Build successful\"")
}
