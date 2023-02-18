package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	APIVersion1          = "api/v1"
	PrefixGetSocialMedia = "get-social-media"
	PrefixGetContact     = "get-contact"
	PrefixGetSkill       = "get-skill"
	MessageSuccess       = "Success"
	MessageError         = "Error"
	AddrSkill            = "SKILL_SERVICE_ADDR"
	AddrContact          = "CONTACT_SERVICE_ADDR"
	AddrSocial           = "SOCIAL_SERVICE_ADDR"
)

type Client struct {
	Host       string
	ApiVersion string
}

func (c *Client) Get(endpoint string) ([]byte, error) {
	response, err := http.Get(c.getUrl(endpoint))

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) getUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.getBaseUrl(), endpoint)
}

func (c *Client) getBaseUrl() string {
	return fmt.Sprintf("%s/%s", c.Host, c.ApiVersion)
}

type Option func(f *Client)

func New(addr string, opts ...Option) *Client {
	c := &Client{Host: fmt.Sprintf("http://%s", addr)}

	for _, opt := range opts {
		opt(c)
	}

	if c.ApiVersion == "" {
		c.ApiVersion = APIVersion1
	}

	return c
}

func WithApiV1() Option {
	return func(f *Client) {
		f.ApiVersion = APIVersion1
	}
}

func NewSkill(opts ...Option) *Client {
	return New(os.Getenv(AddrSkill))
}

func NewContact(opts ...Option) *Client {
	return New(os.Getenv(AddrContact))
}

func NewSocial(opts ...Option) *Client {
	return New(os.Getenv(AddrSocial))
}
