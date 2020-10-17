package client

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModClient(t *testing.T) {
	a := assert.New(t)
	c, err := NewClient(Config{
		Network: ConfigNetwork{ServerAddress: "net.ton.dev"},
	})
	if err == nil {
		defer c.Close()
	}
	fmt.Println("client", c, err)
	a.NoError(err, "client created")
	version, err := c.ClientVersion()
	a.NoError(err, "call client.version")
	a.Equal("1.0.0", version, "dll with specified version")
	rawReference, err := c.ClientGetAPIReference()
	a.NoError(err, "call client.get_api_version")
	var jsonDocs interface{}
	err = json.Unmarshal(rawReference, &jsonDocs)
	a.NoError(err, "get_api_version reference unmarshal")
	a.NotNil(jsonDocs, "successfully unmarshalled")
}