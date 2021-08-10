package api

import (
	"fmt"
	"os/exec"

	"github.com/r2d2-ai/aiflow-cli/common"
	"github.com/r2d2-ai/aiflow-cli/util"
)

func UpdatePkg(project common.AppProject, pkg string) error {

	if Verbose() {
		fmt.Printf("Updating Package: %s \n", pkg)
	}

	err := util.ExecCmd(exec.Command("go", "get", "-u", pkg), project.SrcDir())
	return err
}
