package utils

import (
	"embed"
	"errors"
	"time"

	"github.com/dgraph-io/ristretto/v2"
	"github.com/mojocn/base64Captcha"
)

//go:embed fonts/JetBrainsMono-Bold.ttf
var defaultEmbeddedFontsFS embed.FS
var DefaultEmbeddedFonts = base64Captcha.NewEmbeddedFontsStorage(defaultEmbeddedFontsFS)

type CaptchaCache struct {
	cache *ristretto.Cache[string, string]
}

func NewCaptchaCache() (*CaptchaCache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config[string, string]{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &CaptchaCache{
		cache: cache,
	}, nil
}

func (c *CaptchaCache) Set(key, value string) error {
	if value, ok := c.cache.Get(key); ok && value != "" {
		return errors.New("key already exists")
	}
	c.cache.SetWithTTL(key, value, 1, 120*time.Second)
	return nil
}

func (c *CaptchaCache) Get(key string, clear bool) string {
	value, ok := c.cache.Get(key)
	if ok && clear {
		c.cache.Del(key)
	}
	return value
}

func (c *CaptchaCache) Verify(id, answer string, clear bool) (match bool) {
	match = c.Get(id, clear) == answer
	return
}
