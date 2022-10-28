package ts

import (
	"io.protobj/protobjc"
	"testing"
)

func TestAddImportMessage(t *testing.T) {
	AddImportMessage(protobjc.NewCodeBuilder(), "io.protobj.ProtobjInput")
}
