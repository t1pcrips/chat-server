package streams

import (
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/service"
	"sync"
)

// создам стрим - ключ юзер - сервис Стрим чатов из меседжей
type Chat struct {
	stream  map[*model.User]service.StreamChatMessages
	rwMutex sync.RWMutex
}

// инициализируем чат
func NewChat() *Chat {
	return &Chat{
		stream: make(map[*model.User]service.StreamChatMessages),
	}
}

func (c *Chat) SetStreamForUser(user *model.User, stream service.StreamChatMessages) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	c.stream[user] = stream
}

func (c *Chat) GetStreamForUser(user *model.User) (service.StreamChatMessages, bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	stream, ok := c.stream[user]
	return stream, ok
}

func (c *Chat) GetStreamForUsers() map[*model.User]service.StreamChatMessages {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	streams := make(map[*model.User]service.StreamChatMessages, len(c.stream))

	for user, stream := range c.stream {
		streams[user] = stream
	}

	return streams
}

func (c *Chat) DeleteStreamForUser(user *model.User) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	delete(c.stream, user)
}
