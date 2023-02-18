package bridge

import (
	"encoding/json"
	"github.com/yahya077/portfolio-microservice-app/web-bff/schema"
)

type SocialBridge struct {
}

func (b SocialBridge) Normalize(data []byte) []schema.Social {
	var s schema.SocialResponse

	_ = json.Unmarshal(data, &s)

	return s.Data
}

func Social() SocialBridge {
	return SocialBridge{}
}
