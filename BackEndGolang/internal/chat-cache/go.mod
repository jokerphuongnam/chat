module chat-cache

go 1.23.2

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/uuid v1.6.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

require (
	chat-logs v0.0.0
	github.com/robfig/cron/v3 v3.0.1
)

replace chat-logs => ../chat-logs
