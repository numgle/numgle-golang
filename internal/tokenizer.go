package internal

type Token int32

const (
	_           = iota
	Empty Token = iota
	CompleteHangul
	NotCompleteHangul
	EnglishUpper
	EnglishLower
	Number
	SpecialLetter
	Unknown
)

func Tokenizer(data TokenRange, letter rune) Token {
	switch {
	case letter == ' ', letter == '\r', letter == '\n':
		return Empty
	case letter >= data.CompleteHangul.Start && letter <= data.CompleteHangul.End:
		return CompleteHangul
	case letter >= data.NotCompleteHangul.Start && letter <= data.NotCompleteHangul.End:
		return NotCompleteHangul
	case letter >= data.EnglishUpper.Start && letter <= data.EnglishUpper.End:
		return EnglishUpper
	case letter >= data.EnglishLower.Start && letter <= data.EnglishLower.End:
		return EnglishLower
	case letter >= data.Number.Start && letter <= data.Number.End:
		return Number
	case letter == '?', letter == '!', letter == '.', letter == '^', letter == '-':
		// TODO
		return SpecialLetter
	}
	return Unknown
}
