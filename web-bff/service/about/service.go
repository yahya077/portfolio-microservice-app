package about

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yahya077/portfolio-microservice-app/web-bff/bridge"
	"github.com/yahya077/portfolio-microservice-app/web-bff/contract"
	"github.com/yahya077/portfolio-microservice-app/web-bff/module/client"
	"github.com/yahya077/portfolio-microservice-app/web-bff/module/service"
	"github.com/yahya077/portfolio-microservice-app/web-bff/schema"
)

const (
	DefinitionAboutSchema = "about_schema"
)

type Service struct {
}

func (s Service) HandleGet(ctx *fiber.Ctx) contract.IServiceContext {
	skillData, err := client.NewSkill().Get(client.PrefixGetSkill)

	if err != nil {
		return service.Context().SetCode(service.CodeInternalError).SetSchema(err)
	}

	contactData, err := client.NewContact().Get(client.PrefixGetContact)

	if err != nil {
		return service.Context().SetCode(service.CodeInternalError).SetSchema(err)
	}

	socialData, err := client.NewSocial().Get(client.PrefixGetSocialMedia)

	if err != nil {
		return service.Context().SetCode(service.CodeInternalError).SetSchema(err)
	}

	d := schema.About{
		Skills:  bridge.Skill().Normalize(skillData),
		Contact: bridge.Contact().Normalize(contactData),
		Socials: bridge.Social().Normalize(socialData),
	}

	return service.Context().SetCode(service.CodeSuccess).SetSchema(d)
}

func New() Service {
	return Service{}
}
