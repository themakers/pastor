package cluster_config

import (
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/themakers/pastor/lib/yml"
)

func LoadClusterConfigFromDir[T any](dir string) T {
	if files, err := yml.GetYAMLFiles(dir); err != nil {
		panic(err)
	} else if composite, err := yml.MergeYAMLFiles(files); err != nil {
		panic(err)
	} else {
		var merged = yml.DeepMerge(composite)

		dec := yaml.NewDecoder(strings.NewReader(merged))
		dec.KnownFields(true)

		var cfg T
		if err := dec.Decode(&cfg); err != nil {
			panic(err)
		} else {
			return cfg
		}
	}
}
