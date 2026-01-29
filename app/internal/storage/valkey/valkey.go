package valkey

import (
	"context"
	"fmt"
	"url-shortener/internal/config"

	"github.com/valkey-io/valkey-go"
)

type ValkeyStore struct {
	client valkey.Client
	ctx    context.Context
}

func NewValkeyStore(c *config.Config) (*ValkeyStore, error) {
	connString := fmt.Sprintf("%s:%s", c.ValkeyHost, c.ValkeyPort)
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{connString}})
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	return &ValkeyStore{
		client: client,
		ctx:    ctx,
	}, nil
}

func (s *ValkeyStore) Get(code string) (string, error) {
	linkKey := fmt.Sprintf("links:%s", code)
	link, err := s.client.Do(s.ctx, s.client.B().Get().Key(linkKey).Build()).ToString()
	if err != nil {
		return "", err
	}
	return link, err
}

func (s *ValkeyStore) Save(code, link string) (string, error) {
	linkKey := fmt.Sprintf("links:%s", code)
	_, err := s.client.Do(s.ctx, s.client.B().Set().Key(linkKey).Value(link).Build()).AsBool()
	return code, err
}
