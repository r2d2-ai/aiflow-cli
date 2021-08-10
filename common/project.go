package common

import "github.com/r2d2-ai/aiflow-cli/util"

type AppProject interface {
	Validate() error
	Name() string
	Dir() string
	BinDir() string
	SrcDir() string
	Executable() string
	AddImports(ignoreError bool, addToJson bool, imports ...util.Import) error
	RemoveImports(imports ...string) error
	GetPath(AIflowImport util.Import) (string, error)
	DepManager() util.DepManager

	GetGoImports(withVersion bool) ([]util.Import, error)
}
