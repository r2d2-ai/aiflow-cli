package util

import (
	"errors"
	"fmt"
	"path"
	"regexp"
)

/* util.Import struct defines the different fields which can be extracted from a AIflow import
these imports are stored in AIflow.json in the "imports" array, for instance:

 "imports": [
   "github.com/r2d2-ai/ai-box/common@v1.0.0:/activity/log",
   "github.com/r2d2-ai/ai-box/common/activity/rest@v1.0.0"
   "rest_activity github.com/r2d2-ai/ai-box/common@v1.0.0:/activity/rest",
   "rest_trigger github.com/r2d2-ai/ai-box/common:/trigger/rest",
   "github.com/r2d2-ai/ai-box/flow"
 ]

*/

type AIflowImport struct {
	modulePath         string
	relativeImportPath string
	version            string
	alias              string
}

func NewAIflowImportFromPath(flowImportPath string) (Import, error) {
	flowImport, err := ParseImport(flowImportPath)
	if err != nil {
		return nil, err
	}
	return flowImport, nil
}

func NewAIflowImport(modulePath, relativeImportPath, version, alias string) Import {
	return &AIflowImport{modulePath: modulePath, relativeImportPath: relativeImportPath, version: version, alias: alias}
}

func NewAIflowImportWithVersion(imp Import, version string) Import {
	return &AIflowImport{modulePath: imp.ModulePath(), relativeImportPath: imp.RelativeImportPath(), version: version, alias: imp.Alias()}
}

type Import interface {
	fmt.Stringer

	ModulePath() string
	RelativeImportPath() string
	Version() string
	Alias() string

	CanonicalImport() string // canonical import is used in AIflow.json imports array and to check for equality of imports
	GoImportPath() string    // the import path used in .go files
	GoGetImportPath() string // the import path used by "go get" command
	GoModImportPath() string // the import path used by "go mod edit" command
	IsClassic() bool         // an import is "classic" if it has no : character separator, hence no relative import path
	CanonicalAlias() string  // canonical alias is the alias used in the AIflow.json
}

type Imports []Import

func (flowImport *AIflowImport) ModulePath() string {
	return flowImport.modulePath
}

func (flowImport *AIflowImport) RelativeImportPath() string {
	return flowImport.relativeImportPath
}

func (flowImport *AIflowImport) Version() string {
	return flowImport.version
}

func (flowImport *AIflowImport) Alias() string {
	return flowImport.alias
}

func (flowImport *AIflowImport) CanonicalImport() string {
	alias := ""
	if flowImport.alias != "" {
		alias = flowImport.alias + " "
	}
	version := ""
	if flowImport.version != "" {
		version = "@" + flowImport.version
	}
	relativeImportPath := ""
	if flowImport.relativeImportPath != "" {
		relativeImportPath = ":" + flowImport.relativeImportPath
	}

	return alias + flowImport.modulePath + version + relativeImportPath
}

func (flowImport *AIflowImport) CanonicalAlias() string {
	if flowImport.alias != "" {
		return flowImport.alias
	} else {
		return path.Base(flowImport.GoImportPath())
	}
}

func (flowImport *AIflowImport) GoImportPath() string {
	return flowImport.modulePath + flowImport.relativeImportPath
}

func (flowImport *AIflowImport) GoGetImportPath() string {
	version := "@latest"
	if flowImport.version != "" {
		version = "@" + flowImport.version
	}
	return flowImport.modulePath + flowImport.relativeImportPath + version
}

func (flowImport *AIflowImport) GoModImportPath() string {
	version := "@latest"
	if flowImport.version != "" {
		version = "@" + flowImport.version
	}
	return flowImport.modulePath + version
}

func (flowImport *AIflowImport) IsClassic() bool {
	return flowImport.relativeImportPath == ""
}

func (flowImport *AIflowImport) String() string {
	version := ""
	if flowImport.version != "" {
		version = " " + flowImport.version
	}
	relativeImportPath := ""
	if flowImport.relativeImportPath != "" {
		relativeImportPath = flowImport.relativeImportPath
	}

	return flowImport.modulePath + relativeImportPath + version
}

var flowImportPattern = regexp.MustCompile(`^(([^ ]*)[ ]+)?([^@:]*)@?([^:]*)?:?(.*)?$`) // extract import path even if there is an alias and/or a version

func ParseImports(flowImports []string) (Imports, error) {
	var result Imports

	for _, flowImportPath := range flowImports {
		flowImport, err := ParseImport(flowImportPath)
		if err != nil {
			return nil, err
		}
		result = append(result, flowImport)
	}

	return result, nil
}

func ParseImport(flowImport string) (Import, error) {
	if !flowImportPattern.MatchString(flowImport) {
		return nil, errors.New(fmt.Sprintf("The AIflow import '%s' cannot be parsed.", flowImport))
	}

	matches := flowImportPattern.FindStringSubmatch(flowImport)

	result := &AIflowImport{modulePath: matches[3], relativeImportPath: matches[5], version: matches[4], alias: matches[2]}

	return result, nil
}
