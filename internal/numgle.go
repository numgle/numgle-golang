package internal

import "log"

func Convert(data Dataset, input string) string {
	result := ""
	for _, letter := range input {
		result += toNumgle(data, letter) + "\n"
	}
	return result
}

func toNumgle(data Dataset, letter rune) string {
	token := Tokenizer(data.TokenRanges, letter)
	switch token {
	case Empty:
		return ""
	case CompleteHangul:
		cho, jung, jong := separateHangul(letter)
		if !isCodePointConvertible(data, cho, jung, jong) {
			return ""
		}
		if jung >= 8 && jung != 20 {
			return data.Jong[jong] + data.Jung[jung-8] + data.Cho[cho]
		}
		return data.Jong[jong] + data.ChoseongUndJungseong[Min(8, jung)][cho]
	case NotCompleteHangul:
		return data.Jamo[letter-data.TokenRanges.NotCompleteHangul.Start]
	case EnglishUpper:
		return data.EnglishUpper[letter-data.TokenRanges.EnglishUpper.Start]
	case EnglishLower:
		return data.EnglishLower[letter-data.TokenRanges.EnglishLower.Start]
	case Number:
		return data.Number[letter-data.TokenRanges.Number.Start]
	case SpecialLetter:
		for index, char := range data.TokenRanges.SpecialLetter {
			if char == letter {
				return data.SpecialLetter[index]
			}
		}
	}
	log.Fatal("Not supported rune")
	return ""
}

func separateHangul(hangul rune) (rune, rune, rune) {
	offset := 44032
	index := hangul - rune(offset)
	cho := index / 28 / 21
	jung := index / 28 % 21
	jong := index % 28
	return cho, jung, jong
}

func isCodePointConvertible(data Dataset, cho rune, jung rune, jong rune) bool {
	switch {
	case jong != 0 && data.Cho[cho] == "":
		return false
	case jung >= 8 && jung != 20:
		return data.Jung[jung-8] != ""
	default:
		return data.ChoseongUndJungseong[Min(jung, 8)][cho] != ""
	}
}
