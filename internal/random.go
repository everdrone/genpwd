package internal

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// from: https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func GenerateRandomString(n int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenerateRandomUnambiguousString(n int) (string, error) {
	const letters = "abcdefghjkmnpqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenerateRandomNumberChar() (string, error) {
	const letters = "0123456789"
	var ret byte
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
	if err != nil {
		return "", err
	}
	ret = letters[num.Int64()]

	return string(ret), nil
}

func GetDifferentRandomNumbers(max int, n int, avoid []int) ([]int, error) {
	if n > max {
		return nil, errors.New("n cannot be greater than max")
	}

	if len(avoid) > max {
		return nil, errors.New("avoid cannot be greater than max")
	}

	avoidSeen := make(map[int]bool)
	for _, a := range avoid {
		avoidSeen[a] = true
	}

	uniq := &UniqueRand{seen: avoidSeen}

	var ret []int
	for i := 0; i < n; i++ {
		num, err := uniq.Int(max)
		if err != nil {
			return nil, err
		}

		ret = append(ret, int(num))
	}

	return ret, nil
}
