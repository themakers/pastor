package yml

import (
	"github.com/mikefarah/yq/v4/pkg/yqlib"
)

func DeepMerge(composite string) string {
	var (
		prefs = yqlib.NewDefaultYamlPreferences()
	)

	prefs.EvaluateTogether = false
	prefs.ColorsEnabled = false

	var (
		evaluator  = yqlib.NewStringEvaluator()
		encoder    = yqlib.NewYamlEncoder(prefs)
		decoder    = yqlib.NewYamlDecoder(prefs)
		expression = ". as $item ireduce ({}; . *+d $item )"
	)

	if res, err := evaluator.EvaluateAll(expression, composite, encoder, decoder); err != nil {
		panic(err)
	} else {
		return res
	}
}
