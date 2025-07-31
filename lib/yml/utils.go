package yml

import (
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strings"
)

//func IsKubernetesManifest(text string) bool {
//}

func SplitIntoDocuments(text string) []string {
	var docs = strings.Split(text, "\n---")
	docs = lo.Map(docs, func(item string, index int) string {
		item = strings.Trim(item, "\r\n ")
		item = strings.TrimPrefix(item, "---")
		item = strings.TrimSuffix(item, "---")
		item = strings.Trim(item, "\r\n ")
		return item
	})
	docs = lo.Filter(docs, func(item string, index int) bool {
		return item != ""
	})
	return docs
}

func MergeDocuments(docs ...string) string {
	for i := range docs {
		docs[i] = strings.Trim(docs[i], "\r\n")
	}
	return strings.Join(docs, "\n---\n")
}

func GetYAMLFiles(dir string) ([]string, error) {
	pattern := filepath.Join(dir, "*.yaml")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func MergeYAMLFiles(files []string) (string, error) {
	var contents []string

	for _, file := range files {
		if content, err := os.ReadFile(file); err != nil {
			return "", err
		} else {
			contents = append(contents, string(content))
		}
	}

	return strings.Join(contents, "\n---\n"), nil
}
