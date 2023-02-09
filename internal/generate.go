package internal

import (
	"fmt"
	"unicode"
)

type GeneratorOptions struct {
	Length    int
	Sets      int
	Numbers   int
	Caps      int
	Dashes    bool
	Ambiguous bool
}

func Generate(opt *GeneratorOptions) (res string, err error) {

	totalLength := opt.Length * opt.Sets

	if (opt.Numbers + opt.Caps) > totalLength {
		fmt.Printf("the number of numbers and capitals cannot be greater than the total length of the string")
		return "", ErrSilent
	}

	var letters string
	if opt.Ambiguous {
		letters, err = GenerateRandomString(totalLength)
		if err != nil {
			return "", err
		}
	} else {
		letters, err = GenerateRandomUnambiguousString(totalLength)
		if err != nil {
			return "", err
		}
	}

	numberIndices, err := GetDifferentRandomNumbers(totalLength, opt.Numbers, []int{})
	if err != nil {
		fmt.Printf("error getting random numbers: %v", err)
		return "", err
	}

	capitalIndices, err := GetDifferentRandomNumbers(totalLength, opt.Caps, numberIndices)
	if err != nil {
		fmt.Printf("error getting random numbers: %v", err)
		return "", err
	}

	fmt.Println(numberIndices, capitalIndices)

	for _, numberIndex := range numberIndices {
		// generate a random number between 0 and 9
		number, err := GenerateRandomNumberChar()
		if err != nil {
			return "", err
		}

		// replace the letter at numberIndex with the number
		letters = letters[:numberIndex] + number + letters[numberIndex+1:]
	}

	for _, capitalIndex := range capitalIndices {
		// capitalize the letter at the random index
		r := rune(letters[capitalIndex])
		letters = letters[:capitalIndex] + string(unicode.ToUpper(r)) + letters[capitalIndex+1:]
	}

	if !opt.Dashes {
		return letters, nil
	}

	// place a dash every 6 characters
	for i := 0; i < len(letters); i++ {
		if i%opt.Length == 0 && i != 0 {
			res += "-"
		}
		res += string(letters[i])
	}

	return res, nil
}
