package speller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const baseUrl = "https://speller.yandex.net"
const serviceUri = baseUrl + "/services/spellservice.json/checkText"

type SpellCheckResult struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func CheckText(text string) ([]SpellCheckResult, error) {
	data := url.Values{}
	data.Set("text", text)

	resp, err := http.PostForm(serviceUri, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []SpellCheckResult
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	return results, nil
}
