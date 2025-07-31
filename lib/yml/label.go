package yml

import (
	"fmt"
	
	"github.com/mikefarah/yq/v4/pkg/yqlib"
)

func Label(doc, label, value string) string {
	var (
		prefs = yqlib.NewDefaultYamlPreferences()
	)

	prefs.EvaluateTogether = false
	prefs.ColorsEnabled = false

	var (
		evaluator  = yqlib.NewStringEvaluator()
		encoder    = yqlib.NewYamlEncoder(prefs)
		decoder    = yqlib.NewYamlDecoder(prefs)
		expression = fmt.Sprintf(
			`(.metadata.labels // {} ) as $labels | .metadata.labels = ($labels + {"%s": "%s"}) | select(.)`,
			label, value,
		)
	)

	if res, err := evaluator.EvaluateAll(expression, doc, encoder, decoder); err != nil {
		panic(err)
	} else {
		return res
	}
}
