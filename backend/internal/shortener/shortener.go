package shortener

import (
	"crypto/rand"
	"math/big"
)

type Generator struct {
    alphabet string
    length   int
}

func NewGen() *Generator {
    return &Generator{
        alphabet: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_",
        length:   10,
    }
}

func (g *Generator) Generate() (string, error) {
	result := []byte{}
	lenth := int64(len(g.alphabet))
	for i:= 0; i < g.length; i++ {
		temp, err := rand.Int(rand.Reader, big.NewInt(lenth))
		if err != nil {
			return "", err
		}
		idx := temp.Int64()
		result = append(result, g.alphabet[idx])
	}

	return string(result), nil
}