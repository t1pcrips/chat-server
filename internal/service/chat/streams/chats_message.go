package streams

import (
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/pkg/errs"
	"sync"
)

// создаем карту ключ - чат йд - значение - канал из меседжей
type ChatMessageChannels struct {
	channels map[int64]chan *model.Message
	rwMutex  sync.RWMutex
}

func NewChatMessageChannels() *ChatMessageChannels {
	return &ChatMessageChannels{
		channels: make(map[int64]chan *model.Message),
	}
}

func (c *ChatMessageChannels) InitChannelForChat(chatId, buffersize int64) chan *model.Message {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	chanel := make(chan *model.Message, buffersize)
	c.channels[chatId] = chanel

	return chanel
}

func (c *ChatMessageChannels) GetChanelForChat(chatId int64) (chan *model.Message, bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	channel, ok := c.channels[chatId]

	return channel, ok
}

func (c *ChatMessageChannels) SendMessageForChat(chatId int64, message *model.Message) error {
	chanel, ok := c.GetChanelForChat(chatId)
	if !ok {
		return errs.ErrChatChanel
	}

	chanel <- message

	return nil
}
