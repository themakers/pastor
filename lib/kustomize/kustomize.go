package kustomize

import (
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

func draft() string {
	kustomizeDir := "./overlays/dev"

	var (
		fs      = filesys.MakeFsOnDisk()
		options = krusty.MakeDefaultOptions()
		kust    = krusty.MakeKustomizer(options)
	)

	if resMap, err := kust.Run(fs, kustomizeDir); err != nil {
		panic(err)
	} else if yamlOut, err := resMap.AsYaml(); err != nil {
		panic(err)
	} else {
		return string(yamlOut)
	}
}
