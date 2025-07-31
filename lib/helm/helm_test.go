package helm

import (
	"testing"
)

func TestRenderChart(t *testing.T) {
	var manifest = RenderChartFromRemoteRepo(
		"https://topolvm.github.io/topolvm",
		"topolvm",
		"15.5.6",
		"topolvm-release",
		"topolvm",
		map[string]interface{}{},
	)

	t.Log(manifest)
}
