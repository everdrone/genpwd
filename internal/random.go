package internal

import (
	"crypto/rand"
	"errors"
	"math/big"

	"golang.org/x/exp/slices"
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

	var ret []int
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			return nil, err
		}

		for slices.Contains(avoid, int(num.Int64())) {
			num, err = rand.Int(rand.Reader, big.NewInt(int64(max)))
		}

		ret = append(ret, int(num.Int64()))
	}

	return ret, nil
}

func GetNonCoincidentNumbers(max int) (int, int, error) {
	// generate two non equal random numbers between 0 and max
	num1, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, 0, err
	}

	num2, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil {
		return 0, 0, err
	}

	for num1.Int64() == num2.Int64() {
		num2, err = rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			return 0, 0, err
		}
	}

	return int(num1.Int64()), int(num2.Int64()), nil
}
