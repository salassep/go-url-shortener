package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://medium.com/write-a-catalyst/good-vs-bad-apis-for-developers-12f52b5ac69c"
	shortUrl := "goodvsbad"

	SaveUrlMapping(shortUrl, initialLink)

	retrievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initialLink, retrievedUrl)
}
