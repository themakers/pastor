package yml

import (
	"os"

	"github.com/mikefarah/yq/v4/pkg/yqlib"
	"gopkg.in/op/go-logging.v1"
)

func init() {
	var backend1 = logging.NewLogBackend(os.Stderr, "", 0)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")
	yqlib.GetLogger().SetBackend(backend1Leveled)
	yqlib.InitExpressionParser()
}
