package morse

var EncodeTestCases = []struct {
	description string
	decoded     string
	encoded     string
	ok          error
}{
	{
		"Invalid empty input",
		"",
		"",
		errInvalidLength,
	},
	{
		"Invalid input containing only one space",
		" ",
		"",
		errInvalidLength,
	},
	{
		"Valid morse input",
		"HELLO WORLD",
		".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
		nil,
	},
	{
		"Valid input without any spaces and lower case characters",
		"HELLoWORLD",
		".... . .-.. .-.. --- .-- --- .-. .-.. -..",
		nil,
	},
	{
		"Valid input with punctuations and space",
		"HELLO! WORLD(?)",
		".... . .-.. .-.. --- -.-.-- / .-- --- .-. .-.. -.. -.--. ..--.. -.--.-",
		nil,
	},
	{
		"Invalid input without any characters and only spaces",
		"     \t\n",
		"",
		errInvalidLength,
	},
	{
		"Valid input with numbers, spaces and punctuations and lowercase letters",
		"MY1 DoG BiT ME! I NeED hELP ASAP! 123",
		"-- -.-- .---- / -.. --- --. / -... .. - / -- . -.-.-- / .. / -. . . -.. / .... . .-.. .--. / .- ... .- .--. -.-.-- / .---- ..--- ...--",
		nil,
	},
	{
		"Complete gibberish input",
		"FGASDKFJGHJKSDJHFG3246.,?!/()&:=+-@",
		"..-. --. .- ... -.. -.- ..-. .--- --. .... .--- -.- ... -.. .--- .... ..-. --. ...-- ..--- ....- -.... .-.-.- --..-- ..--.. -.-.-- -..-. -.--. -.--.- .-... ---... -...- .-.-. -....- .--.-.",
		nil,
	},
	{
		"Invalid input containing invalid characters",
		"I_dont%know",
		"",
		errInvalidChar,
	},
}

var DecodeTestCases = []struct {
	description string
	encoded     string
	decoded     string
	ok          error
}{
	{
		"Invalid morse code with empty string",
		"",
		"",
		errInvalidLength,
	},
	{
		"Invalid morse code with only one space",
		" ",
		"",
		errInvalidLength,
	},
	{
		"Valid morse code",
		".... . .-.. .-.. --- / .-- --- .-. .-.. -..",
		"HELLO WORLD",
		nil,
	},
	{
		"Valid morse code without any spaces and lower case characters",
		".... . .-.. .-.. --- .-- --- .-. .-.. -..",
		"HELLOWORLD",
		nil,
	},
	{
		"Valid morse code with punctuations and space",
		".... . .-.. .-.. --- -.-.-- / .-- --- .-. .-.. -.. -.--. ..--.. -.--.-",
		"HELLO! WORLD(?)",
		nil,
	},
	{
		"Valid morse code with numbers, spaces and punctuations",
		"-- -.-- .---- / -.. --- --. / -... .. - / -- . -.-.-- / .. / -. . . -.. / .... . .-.. .--. / .- ... .- .--. -.-.-- / .---- ..--- ...--",
		"MY1 DOG BIT ME! I NEED HELP ASAP! 123",
		nil,
	},
	{
		"Complete gibberish morse code",
		"..-. --. .- ... -.. -.- ..-. .--- --. .... .--- -.- ... -.. .--- .... ..-. --. ...-- ..--- ....- -.... .-.-.- --..-- ..--.. -.-.-- -..-. -.--. -.--.- .-... ---... -...- .-.-. -....- .--.-.",
		"FGASDKFJGHJKSDJHFG3246.,?!/()&:=+-@",
		nil,
	},
	{
		"Invalid morse code containing invalid characters",
		".._. __. ?? 123I:C",
		"",
		errInvalidChar,
	},
}
