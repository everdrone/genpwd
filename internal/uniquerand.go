package internal

import (
	"crypto/rand"
	"math/big"
)

// from: https://stackoverflow.com/a/39925864

type UniqueRand struct {
	seen map[int]bool
}

func (u *UniqueRand) Int(max int) (int, error) {
	for {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			return 0, err
		}

		i := int(num.Int64())
		if !u.seen[i] {
			u.seen[i] = true
			return i, nil
		}
	}
}
