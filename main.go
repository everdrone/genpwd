package main

import (
	"crypto/rand"
	"fmt"
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

func getNonCoincidentNumbers(max int) (int, int, error) {
	// generate two non equal random numbers between 0 and strLen
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

func main() {
	letters, err := GenerateRandomString(6 * 3)
	if err != nil {
		panic(err)
	}

	// generate a random number between 0 and 9
	number, err := GenerateRandomNumberChar()
	if err != nil {
		panic(err)
	}

	capitalIndex, numberIndex, err := getNonCoincidentNumbers(6 * 3)
	if err != nil {
		panic(err)
	}

	// capitalize the letter at the random index
	letters = letters[:capitalIndex] + string(letters[capitalIndex]-32) + letters[capitalIndex+1:]

	// replace the letter at numberIndex with the number
	letters = letters[:numberIndex] + number + letters[numberIndex+1:]

	// fmt.Println(letters)

	// place a dash every 6 characters
	var dashLetters string
	for i := 0; i < len(letters); i++ {
		if i%6 == 0 && i != 0 {
			dashLetters += "-"
		}
		dashLetters += string(letters[i])
	}

	fmt.Println(dashLetters)
}
