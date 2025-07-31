package yml

import (
	"testing"
)

//func TestMerge(t *testing.T) {
//	files, err := GetYAMLFiles("./test-files")
//	t.Log(files)
//	if err != nil {
//		panic(err)
//	}
//
//	var contents []string
//	for _, file := range files {
//		content, err := os.ReadFile(file)
//		if err != nil {
//			t.Fatal(err)
//		}
//		contents = append(contents, string(content))
//	}
//
//	composite := strings.Join(contents, "\n---\n")
//
//	t.Log(DeepMerge(composite))
//}

const man = `
---
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations: {}
  labels:
    name: sealed-secrets-controller
  name: sealed-secrets-controller
  namespace: kube-system
`

func TestLabel(t *testing.T) {
	t.Log(Label(man, "a", "b"))
	//os.WriteFile("test.yaml", []byte(Label(man)), 0777)
}
