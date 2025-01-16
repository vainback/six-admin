package request

import (
	"bytes"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

func Get(uri string, values *URLParam, result any) error {
	resp, err := http.Get(uri + "?" + values.Builder())
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	resultBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return jsoniter.Unmarshal(resultBytes, result)
}

func Post(uri string, data any, result any) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(uri, "application/json", bytes.NewReader(dataBytes))
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	resultBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return jsoniter.Unmarshal(resultBytes, result)
}

func POSTForm(uri string, data *URLParam, result any) error {
	resp, err := http.PostForm(uri, data.values)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	resultBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(resultBytes, result)
}
