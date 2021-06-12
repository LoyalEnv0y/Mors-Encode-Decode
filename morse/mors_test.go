package morse

import (
	"testing"
)

func TestEncode(t *testing.T) {
	for _, test := range EncodeTestCases {
		if encoded, ok := Encode(test.decoded); ok != test.ok {
			t.Fatalf("FAIL => [%s] Decoded: %s, Expected: %t, Got: %t", test.description, test.decoded, test.ok, ok)
		} else if encoded != test.encoded {
			t.Fatalf("FAIL =>[%s] Decoded: %s, Expected: %s, Got: %s", test.description, test.decoded, test.encoded, encoded)
		} else {
			t.Logf("PASSED => [%s]", test.description)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range DecodeTestCases {
		if decoded, ok := Decode(test.encoded); ok != test.ok {
			t.Fatalf("[%s]: Encoded %s, Expected: %t, Got: %t", test.description, test.encoded, test.ok, ok)
		} else if decoded != test.decoded {
			t.Fatalf("[%s]: Encoded %s, Expected: %s, Got %s", test.description, test.encoded, test.decoded, decoded)
		} else {
			t.Logf("PASSED => [%s]", test.description)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range EncodeTestCases {
			Encode(v.decoded)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range DecodeTestCases {
			Decode(v.encoded)
		}
	}
}
