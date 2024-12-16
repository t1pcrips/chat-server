package streams

import "sync"

// создаем карту ключ - чат йд - значение структура чата, где мы имеем канал передачи данных и мютекс
type Chats struct {
	chats   map[int64]*Chat
	rwMutex sync.RWMutex
}

func NewChats() *Chats {
	return &Chats{
		chats: make(map[int64]*Chat),
	}
}

func (c *Chats) CreateOrGetChat(chatId int64) *Chat {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	chat, ok := c.chats[chatId]
	if !ok {
		chat = NewChat()
		c.chats[chatId] = chat
	}

	return chat
}
