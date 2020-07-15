package cache

import (
    big "github.com/asktop/decimal"
    "github.com/asktop/gotools/acast"
    "github.com/asktop/gotools/log"
    "github.com/astaxie/beego/cache"
    _ "github.com/astaxie/beego/cache/redis"
    "github.com/shopspring/decimal"
    "strconv"
)

var beegoCache cache.Cache

func StartCache(config Config) error {
    //获取redis配置
    if config.Host == "" {
        config.Host = "127.0.0.1"
    }
    if config.Port == "" {
        config.Port = "6379"
    }
    cacheConfig := `{"key":"cache","conn":"` + config.Host + `:` + config.Port + `","dbNum":"` + strconv.Itoa(config.Select) + `","password":"` + config.Password + `"}`
    log.Info("--- 连接 redis cache ---", "config:", cacheConfig)
    var err error
    beegoCache, err = cache.NewCache("redis", cacheConfig)
    if err != nil {
        log.Error("--- 连接 redis cache 出错 ---", "err:", err)
        return err
    }
    return nil
}

type Cache struct {
    cache.Cache
}

func NewCache() *Cache {
    c := &Cache{}
    c.Cache = beegoCache
    return c
}

func (c *Cache) GetBool(key string) bool {
    value := c.Get(key)
    return acast.ToBool(value)
}

func (c *Cache) GetInt(key string) int {
    value := c.Get(key)
    return acast.ToInt(value)
}

func (c *Cache) GetInt64(key string) int64 {
    value := c.Get(key)
    return acast.ToInt64(value)
}

func (c *Cache) GetFloat64(key string) float64 {
    value := c.Get(key)
    return acast.ToFloat64(value)
}

func (c *Cache) GetString(key string) string {
    value := c.Get(key)
    return acast.ToString(value)
}

func (c *Cache) GetBig(key string) *big.Big {
    value := c.Get(key)
    return acast.ToBig(value)
}

func (c *Cache) GetDecimal(key string) decimal.Decimal {
    value := c.Get(key)
    return acast.ToDecimal(value)
}
