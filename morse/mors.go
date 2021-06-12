// Package morse contains useful tools to encode a text to a morse code or decode a morse code text.
package morse

import (
	"errors"
	"strings"
)

var (
	// ToMorse to store the rune values of latin characters with their counterpart morse codes
	ToMorse = map[rune]string{'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
		'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
		'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...",
		'T': "-", 'U': "..-", 'V': "...-", 'Y': "-.--", 'W': ".--", 'X': "-..-", 'Z': "--..",
		'0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.", '.': ".-.-.-", ',': "--..--",
		'?': "..--..", '!': "-.-.--", '/': "-..-.", '(': "-.--.", ')': "-.--.-", '&': ".-...",
		':': "---...", '=': "-...-", '+': ".-.-.", '-': "-....-", '@': ".--.-.", ' ': "/",
	}
	// ToLatin to store the each morse code's counterpart latin character
	ToLatin = map[string]rune{".-": 'A', "-...": 'B', "-.-.": 'C', "-..": 'D', ".": 'E',
		"..-.": 'F', "--.": 'G', "....": 'H', "..": 'I', ".---": 'J', "-.-": 'K', ".-..": 'L',
		"--": 'M', "-.": 'N', "---": 'O', ".--.": 'P', "--.-": 'Q', ".-.": 'R', "...": 'S',
		"-": 'T', "..-": 'U', "...-": 'V', "-.--": 'Y', ".--": 'W', "-..-": 'X', "--..": 'Z',
		"-----": '0', ".----": '1', "..---": '2', "...--": '3', "....-": '4', ".....": '5',
		"-....": '6', "--...": '7', "---..": '8', "----.": '9', ".-.-.-": '.', "--..--": ',',
		"..--..": '?', "-.-.--": '!', "-..-.": '/', "-.--.": '(', "-.--.-": ')', ".-...": '&',
		"---...": ':', "-...-": '=', ".-.-.": '+', "-....-": '-', ".--.-.": '@', "/": ' ',
	}
	errInvalidLength = errors.New("unexpected input length")
	errInvalidChar   = errors.New("input contains unexpected characters")
)

// Encode takes a decoded massage with the length of 1 or higher and returns it's morse coded version
// if the input contains any character that doesn't appear in the ToMorse map it will throw an error
// every space character is being represented with "/". Output will always be capitalized.
func Encode(decoded string) (string, error) {
	// deleting the unnecessary spaces and checking the length of the input
	decoded = strings.TrimSpace(decoded)
	if len(decoded) == 0 {
		return "", errInvalidLength
	}
	decoded = strings.ToUpper(decoded)

	var encoded string

	for _, v := range decoded {
		if _, ok := ToMorse[v]; !ok && v != ' ' {
			return "", errInvalidChar
		}
		encoded += ToMorse[v] + " "
	}

	encoded = strings.TrimSpace(encoded)
	return encoded, nil
}

// Decode will take a string of encoded morse code(s) and translate them to latin alphabet.
// it will throw an error if the given input is empty or if it contains any character other than
// those being stored in ToLatin map. it will put a space character everytime it sees a "/".
func Decode(encoded string) (string, error) {
	encoded = strings.TrimSpace(encoded)
	if len(encoded) == 0 {
		return "", errInvalidLength
	}

	var (
		decoded string
		set     string
	)

	for _, v := range encoded + " " {
		if v == ' ' {
			if _, ok := ToLatin[set]; ok {
				decoded += string(ToLatin[set])
				set = ""
				continue
			}
			return "", errInvalidChar
		}
		set += string(v)
	}
	return decoded, nil
}
