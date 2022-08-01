package files

import (
	"api-builder/templates"
)

const SharedCacheStore templates.Template = `
	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
		DB:      0,
	}))
`
