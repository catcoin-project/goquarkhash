package quarkhash

import (
	"encoding/hex"
	"fmt"
	"testing"
)

type Test struct {
	Data     string
	Expected string
}

func TestQuarkHash(t *testing.T) {
	tests := []Test{
		{
			Data:     "0000000000000000000000000000000000000000000000000000000000000000",
			Expected: "72e0ba1220a18502a00038c4a70c9ab35ed51d73c9953adafae3b76150ae77a6",
		},
		{
			Data:     "0101010101010101010101010101010101010101010101010101010101010101",
			Expected: "9f0ef8df8191ef0d7dddd09f7df71eb5e2ad9b4feb1dd0fd0328f50d19fab702",
		},
		{
			Data:     "ff65a6be6d14fce72ad4aff559f7579100f295cfae91a64ec6bd1939602a4eb6",
			Expected: "870e9d5b6cb5831fcbd21a644f3b7c164614115160b3f523f126715eb78ea0ac",
		},
		{
			Data:     "9ac322dca9b55db8f4579dc9bf6df51334a194da9aac86b08dd4c32c72cc8718",
			Expected: "ca9f80d482d59f12d897c2b706b55fcc2fa77ba915739656cbede8cdc197cd8b",
		},
	}

	for _, test := range tests {
		decoded_len := hex.DecodedLen(len(test.Data))
		decoded := make([]byte, decoded_len)
		hex.Decode(decoded, []byte(test.Data))

		hash := QuarkHash(decoded)

		hash_str := fmt.Sprintf("%x", hash)

		if hash_str != test.Expected {
			t.Errorf("Invalid hash %s != %s", hash_str, test.Expected)
			return
		}
	}
}
