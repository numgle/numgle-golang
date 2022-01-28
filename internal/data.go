package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Range struct {
	Start rune `json:"start"`
	End   rune `json:"end"`
}

type TokenRange struct {
	CompleteHangul    Range  `json:"completeHangul"`
	NotCompleteHangul Range  `json:"notCompleteHangul"`
	EnglishUpper      Range  `json:"EnglishUpper"`
	EnglishLower      Range  `json:"lowercase"`
	Number            Range  `json:"number"`
	SpecialLetter     []rune `json:"special"`
}

type Dataset struct {
	Cho                  []string   `json:"cho"`
	Jung                 []string   `json:"jung"`
	Jong                 []string   `json:"jong"`
	ChoseongUndJungseong [][]string `json:"cj"`
	Jamo                 []string   `json:"han"`
	EnglishUpper         []string   `json:"englishUpper"`
	EnglishLower         []string   `json:"englishLower"`
	Number               []string   `json:"number"`
	SpecialLetter        []string   `json:"special"`
	TokenRanges          TokenRange `json:"range"`
}

func GetDataset() Dataset {
	datasetUrl := "https://raw.githubusercontent.com/numgle/dataset/main/src/data.json"
	res, err := http.Get(datasetUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data Dataset
	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	return data
}
