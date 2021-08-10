package api

import (
	"github.com/r2d2-ai/ai-box-cli/util"
)

var verbose = false

func SetVerbose(enable bool) {
	verbose = enable
	util.SetVerbose(enable)
}

func Verbose() bool {
	return verbose
}

//TODO use a logger like struct for API that can be used to log or console output
