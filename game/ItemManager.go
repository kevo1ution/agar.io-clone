package game

import "github.com/google/uuid"

type itemManager struct {
	Items map[uuid.UUID]event
}

func MakeItemManager() itemManager {
	return itemManager{
		Items: make(map[uuid.UUID]event),
	}
}

func (manager itemManager) Add(id uuid.UUID, item event) {
	manager.Items[id] = item
}
