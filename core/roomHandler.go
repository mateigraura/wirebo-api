package core

import (
	"github.com/mateigraura/wirebo-api/domain"
)

type Room struct {
	broadcast chan *domain.Message
}
