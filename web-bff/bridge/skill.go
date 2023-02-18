package bridge

import (
	"encoding/json"
	"github.com/yahya077/portfolio-microservice-app/web-bff/schema"
)

type SkillBridge struct {
}

func (b SkillBridge) Normalize(data []byte) []schema.Skill {
	var s schema.SkillResponse

	_ = json.Unmarshal(data, &s)

	return s.Data
}

func Skill() SkillBridge {
	return SkillBridge{}
}
