package context

import (
    "github.com/asktop/gotools/aqywx/cache"
    "net/http"
    "sync"
)

var (
    defaultMemory     = cache.NewMemory()
    acccessTokenLocks = map[string]*sync.RWMutex{}
)

// Context struct
type Context struct {
    CorpID     string // 企业ID
    CorpSecret string // 应用的凭证密钥

    AccessTokenLock *sync.RWMutex
    Cache           cache.Cache

    Writer  http.ResponseWriter
    Request *http.Request
}

// NewClient init
func NewContext(corpid string, corpsecret string, cache ...cache.Cache) *Context {
    context := new(Context)
    context.CorpID = corpid
    context.CorpSecret = corpsecret

    cacheKey := context.GetAccessTokenCacheKey()
    if lock, ok := acccessTokenLocks[cacheKey]; !ok {
        lock = new(sync.RWMutex)
        acccessTokenLocks[cacheKey] = lock
        context.AccessTokenLock = lock
    } else {
        context.AccessTokenLock = lock
    }

    if len(cache) > 0 && cache[0] != nil {
        context.Cache = cache[0]
    } else {
        context.Cache = defaultMemory
    }
    return context
}

// Query returns the keyed url query value if it exists
func (ctx *Context) Query(key string) string {
    value, _ := ctx.GetQuery(key)
    return value
}

// GetQuery is like Query(), it returns the keyed url query value
func (ctx *Context) GetQuery(key string) (string, bool) {
    req := ctx.Request
    if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
        return values[0], true
    }
    return "", false
}
