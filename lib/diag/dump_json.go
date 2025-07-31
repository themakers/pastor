package diag

// TODO: make deterministic: "encoding/json/v2"
	
import (
	"encoding/json"
)

func DumpJSON(v any) string {
	if data, err := json.MarshalIndent(v, "", "  "); err != nil {
		panic(err)
	} else {
		return string(data)
	}
}
