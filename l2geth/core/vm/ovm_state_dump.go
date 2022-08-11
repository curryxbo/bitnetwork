package vm

import (
	"os"
)

// UsingBvm is used to enable or disable functionality necessary for the Bvm.
var UsingBvm bool

func init() {
	UsingBvm = os.Getenv("USING_Bvm") == "true"
}
