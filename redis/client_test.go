package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	err := Set("redis:test", "123")
	assert.NoError(t, err)
	value, err := Get("redis:test")
	assert.NoError(t, err)
	assert.Equal(t, "123", value)
	_, err = Get("jinxing.liu")
	assert.Error(t, err)
	fmt.Println(err)
}

func TestSet(t *testing.T) {
	err := Set("redis:test", "123")
	assert.NoError(t, err)

	err = redisClient.HMSet(ctx, "redis:user", map[string]interface{}{
		"username": "jinxing.liu",
		"age": 12,
		"gender": 1,
		"status": 2,
		"created_at": time.Now(),
	}).Err()

	fmt.Println(err)
}
