package common

import (
	"github.com/r2d2-ai/ai-box-cli/util"
)

var verbose = false
var appProject AppProject

func SetVerbose(enable bool) {
	verbose = enable
	util.SetVerbose(enable)
}

func Verbose() bool {
	return verbose
}

func CurrentProject() AppProject {
	return appProject
}

func SetCurrentProject(project AppProject) {
	appProject = project
}
