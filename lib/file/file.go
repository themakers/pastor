package file

import "os"

func Read(path string) []byte {
	if data, err := os.ReadFile(path); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			panic(err)
		}
	} else {
		return data
	}
}
