package redis

import (
	"fmt"
	"log"
	"os"
	"testing"

	r "github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
)

var redisAddr = fmt.Sprintf("%s:%s", os.Getenv("REDIS_MASTER_SERVICE_HOST"), os.Getenv("REDIS_MASTER_SERVICE_PORT"))

const (
	testHash = "testHash"
)

// CheckRedisEnv checks for the presence of the following
// environment variables
//   REDIS_MASTER_SERVICE_HOST
//   REDIS_MASTER_SERVICE_PORT
func CheckRedisEnv() error {
	envs := []string{
		"REDIS_MASTER_SERVICE_HOST",
		"REDIS_MASTER_SERVICE_PORT",
	}
	for _, e := range envs {
		if len(os.Getenv(e)) == 0 {
			return fmt.Errorf("env %s is not set", e)
		}
	}
	return nil
}

type TestRedis struct {
	client *r.Client
}

// NewTestRedisFromEnv is a constructor for a TestRedis instance.
func NewTestRedisFromEnv() (*TestRedis, error) {
	tr := new(TestRedis)
	if err := CheckRedisEnv(); err != nil {
		return tr, err
	}
	tr.client = r.NewClient(&r.Options{
		Addr: redisAddr,
	})
	err := tr.client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("error pinging redis %s", err)
	}
	return tr, nil
}

func TestMain(m *testing.M) {
	_, err := NewTestRedisFromEnv()
	if err != nil {
		log.Fatalf("unable to construct new TestRedisFromEnv instance %s", err)
	}
	os.Exit(m.Run())
}

func TestSet(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	repo, err := NewCache(redisAddr)
	assert.NoError(err, "error connecting to redis")
	err = repo.Set("art", "vandelay")
	assert.NoError(err, "error in setting key")
}

func TestGet(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	repo, err := NewCache(redisAddr)
	assert.NoError(err, "error connecting to redis")
	err = repo.Set("art", "vandelay")
	assert.NoError(err, "error in setting key")
	key, err := repo.Get("art")
	assert.NoError(err, "error getting key")
	assert.Equal(key, "vandelay", "should retrieve correct value")
}

func TestExists(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	repo, err := NewCache(redisAddr)
	assert.NoError(err, "error connecting to redis")
	err = repo.Set("art", "vandelay")
	assert.NoError(err, "error in setting key")
	lookup, err := repo.Exists("art")
	assert.NoError(err, "error finding key")
	assert.True(lookup, "should find previously set key")
	badLookup, err := repo.Exists("obrien-murphy")
	assert.NoError(err, "error finding key")
	assert.False(badLookup, "should not find random key")
}

func TestHSet(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	repo, err := NewCache(redisAddr)
	assert.NoError(err, "error connecting to redis")
	err = repo.HSet(testHash, "art", "vandelay")
	assert.NoError(err, "error in setting key")
}

func TestHGet(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	repo, err := NewCache(redisAddr)
	assert.NoError(err, "error connecting to redis")
	err = repo.HSet(testHash, "art", "vandelay")
	assert.NoError(err, "error in setting key")
	key, err := repo.HGet(testHash, "art")
	assert.NoError(err, "error getting key")
	assert.Equal(key, "vandelay", "should retrieve correct value")
}

func TestHExists(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	repo, err := NewCache(redisAddr)
	assert.NoError(err, "error connecting to redis")
	err = repo.HSet(testHash, "art", "vandelay")
	assert.NoError(err, "error in setting key")
	lookup, err := repo.HExists(testHash, "art")
	assert.NoError(err, "error finding key")
	assert.True(lookup, "should find previously set key")
	badLookup, err := repo.HExists(testHash, "obrien-murphy")
	assert.NoError(err, "error finding key")
	assert.False(badLookup, "should not find random key")
}
