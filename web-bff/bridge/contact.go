package bridge

import (
	"encoding/json"
	"github.com/yahya077/portfolio-microservice-app/web-bff/schema"
)

type ContactBridge struct {
}

func (b ContactBridge) Normalize(data []byte) schema.Contact {
	s := schema.ContactResponse{}

	_ = json.Unmarshal(data, &s)

	return s.Data
}

func Contact() ContactBridge {
	return ContactBridge{}
}
