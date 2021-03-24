package aqywx

import (
    "github.com/asktop/gotools/aqywx/base"
    "github.com/asktop/gotools/aqywx/cache"
    "github.com/asktop/gotools/aqywx/context"
    "github.com/asktop/gotools/aqywx/crm"
    "github.com/asktop/gotools/aqywx/media"
)

// Client struct
type Client struct {
    Context *context.Context
}

// NewClient init
func NewClient(corpid string, corpsecret string, cache ...cache.Cache) *Client {
    client := new(Client)
    client.Context = context.NewContext(corpid, corpsecret, cache...)
    return client
}

// GetBase 通讯录管理接口
func (c *Client) GetBase() *base.Base {
    return base.NewBase(c.Context)
}

// GetMedia 素材管理接口
func (c *Client) GetMedia() *media.Media {
    return media.NewMedia(c.Context)
}

// GetCrm 外部联系人接口
func (c *Client) GetCrm() *crm.Crm {
    return crm.NewCrm(c.Context)
}
