package udp

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseToUDPAddr(t *testing.T) {
	address, err := ParseToUDPAddr("127.0.0.1:9098")
	assert.NoError(t, err)
	assert.Equal(t, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 9098}, address)

	_, err = ParseToUDPAddr("127.0.0.1")
	fmt.Println(err)
	assert.Error(t, err)

	_, err = ParseToUDPAddr("127.0.0:1")
	fmt.Println(err)
	assert.Error(t, err)

	_, err = ParseToUDPAddr("127.0.0.1:abc")
	fmt.Println(err)
	assert.Error(t, err)

	_, err = ParseToUDPAddr("127.0.abc.1:100")
	fmt.Println(err)
	assert.Error(t, err)
}
