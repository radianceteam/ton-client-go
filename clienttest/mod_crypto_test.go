package clienttest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"

	"github.com/radianceteam/ton-client-go/client"
)

func newTestClient() *client.Client {
	c, err := client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: "net.ton.dev"},
	})
	if err != nil {
		panic(err)
	}
	return c
}

func TestModCrypto(t *testing.T) {
	a := assert.New(t)
	c := newTestClient()
	defer c.Close()

	keys, err := c.CryptoGenerateRandomSignKeys()
	a.NoError(err, "call Client.version")
	a.Len(keys.Public, 64, "hex len")
	a.Len(keys.Secret, 64, "hex len")
}

func TestModCryptoMnemonicFromRandom(t *testing.T) {
	a := assert.New(t)
	c := newTestClient()
	defer c.Close()

	r, err := c.CryptoMnemonicFromRandom(&client.ParamsOfMnemonicFromRandom{})
	a.NoError(err, "call crypto.mnemonic_from_random")
	a.Len(strings.Split(r.Phrase, " "), 12, "default phrase size")
	r, err = c.CryptoMnemonicFromRandom(&client.ParamsOfMnemonicFromRandom{WordCount: null.IntFrom(24)})
	a.NoError(err, "call crypto.mnemonic_from_random")
	r, err = c.CryptoMnemonicFromRandom(&client.ParamsOfMnemonicFromRandom{WordCount: null.IntFrom(13)})
	a.Error(err, "bip39 invalid wc")
}

func TestModCryptoMnemonicWords(t *testing.T) {
	a := assert.New(t)
	c := newTestClient()
	defer c.Close()

	r, err := c.CryptoMnemonicWords(&client.ParamsOfMnemonicWords{})
	a.NoError(err, "call crypto.mnemonic_words")
	a.Len(strings.Split(r.Words, " "), 2048, "default dictionary size")
	r, err = c.CryptoMnemonicWords(&client.ParamsOfMnemonicWords{Dictionary: null.IntFrom(1)})
	a.NoError(err, "call crypto.mnemonic_words")
	a.Len(strings.Split(r.Words, " "), 2048, "default dictionary size")
}