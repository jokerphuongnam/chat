module chat-call

go 1.23.4

require (
	chat-config v0.0.0
	chat-database v0.0.0
	chat-logs v0.0.0
)

replace chat-config => ../internal/chat-config

replace chat-database => ../internal/chat-database

replace chat-logs => ../internal/chat-logs
