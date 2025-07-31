package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(url, path string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	checkResponseStatus(resp)

	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		panic(err)
	}
}

func Get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	checkResponseStatus(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func checkResponseStatus(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		if body, err := io.ReadAll(resp.Body); err != nil {
			panic(err)
		} else {
			panic(fmt.Sprintf("%d - %s => %s", resp.StatusCode, resp.Status, string(body)))
		}
	}
}
