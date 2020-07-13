package aredis

import (
    "encoding/json"
    "errors"
    big "github.com/asktop/decimal"
    "github.com/asktop/gotools/alog"
    "github.com/gomodule/redigo/redis"
    "github.com/shopspring/decimal"
    "net"
    "time"
)

/*
import "github.com/gomodule/redigo/redis"
import "github.com/garyburd/redigo/redis" （已废弃）
 */

var redisPool *redis.Pool

//redis配置
type Config struct {
    Host        string `json:"host" yaml:"host"`
    Port        string `json:"port" yaml:"port"`
    Password    string `json:"password" yaml:"password"`
    Select      int    `json:"select" yaml:"select"`
    MaxIdle     int    `json:"maxidle" yaml:"maxidle"`
    MaxActive   int    `json:"maxactive" yaml:"maxactive"`
    IdleTimeout int    `json:"idletimeout" yaml:"idletimeout"`
}

//初始化连接redis
func StartRedis(config Config) error {
    //获取redis配置
    if config.Host == "" {
        config.Host = "127.0.0.1"
    }
    if config.Port == "" {
        config.Port = "6379"
    }
    address := net.JoinHostPort(config.Host, config.Port)
    password := config.Password
    db := config.Select
    if config.MaxActive == 0 {
        config.MaxActive = 500
    }
    maxActive := config.MaxActive
    if config.MaxIdle == 0 {
        config.MaxIdle = 300
    }
    maxIdle := config.MaxIdle
    if config.IdleTimeout == 0 {
        config.IdleTimeout = 180
    }
    idleTimeout := time.Duration(config.IdleTimeout) * time.Second //时间单位：秒
    alog.Info("--- 连接 redis ---", "address:", address, "password:", password, "db:", db)
    //连接redis并创建连接池
    redisPool = &redis.Pool{
        Dial: func() (redis.Conn, error) {
            conn, err := redis.Dial("tcp", address)
            if err != nil {
                return nil, err
            }
            //输入密码
            if password != "" {
                if _, err := conn.Do("AUTH", password); err != nil {
                    conn.Close()
                    return nil, err
                }
            }
            //选择存贮库
            if db != 0 {
                if _, err := conn.Do("SELECT", db); err != nil {
                    conn.Close()
                    return nil, err
                }
            }
            return conn, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if time.Since(t) < time.Minute {
                return nil
            }
            _, err := c.Do("PING")
            return err
        },
        MaxActive:   maxActive,
        MaxIdle:     maxIdle,
        IdleTimeout: idleTimeout,
    }

    //redis连接测试
    conn := redisPool.Get()
    defer conn.Close()
    _, err := conn.Do("PING")
    if err != nil {
        alog.Error("--- 连接 redis 出错 ---", "err:", err)
        return err
    }
    return nil
}

type Redis struct {
    conn      redis.Conn
    autoClose bool
}

func NewRedis(noAutoClose ...bool) *Redis {
    autoClose := true
    if len(noAutoClose) > 0 && noAutoClose[0] {
        autoClose = false
    }
    return &Redis{
        conn:      redisPool.Get(),
        autoClose: autoClose,
    }
}

func (r *Redis) Close() {
    r.conn.Close()
}

//执行指定命令
func (r *Redis) Do(command string, args ...interface{}) (reply interface{}, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    return r.conn.Do(command, args...)
}

//获取所有匹配的key
func (r *Redis) Keys(pattern string) (keys []string, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    return redis.Strings(r.conn.Do("KEYS", pattern))
}

//设置|取消过期时间
// @param seconds > 0 设置过期时间
// @param seconds <= 0 取消过期时间
func (r *Redis) Expire(key string, seconds int64) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    //设置过期时间
    if seconds > 0 {
        _, err := redis.Bool(r.conn.Do("EXPIRE", key, seconds))
        if err != nil {
            return err
        }
    }
    //取消过期时间
    if seconds <= 0 {
        _, err := redis.Bool(r.conn.Do("PERSIST", key))
        if err != nil {
            return err
        }
    }
    return nil
}

//判断是否存在
func (r *Redis) Exists(key string) (exists bool, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    return redis.Bool(r.conn.Do("EXISTS", key))
}

//删除
func (r *Redis) Del(key string) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    _, err := redis.Bool(r.conn.Do("DEL", key))
    if err != nil {
        return err
    }
    return nil
}

//设置 string
func (r *Redis) Set(key string, value interface{}, seconds ...int64) (old string, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    old, err = redis.String(r.conn.Do("GETSET", key, value))
    _, err = IfExistsOfErr(err)
    if err == nil && len(seconds) > 0 {
        err = r.Expire(key, seconds[0])
    }
    return
}

//批量设置 string
func (r *Redis) MSet(mset map[string]interface{}) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    if len(mset) > 0 {
        var args []interface{}
        for k, v := range mset {
            args = append(args, k, v)
        }
        rs, err := redis.String(r.conn.Do("MSET", args...))
        if err != nil {
            return err
        }
        if rs != "OK" {
            return errors.New(rs)
        }
    }
    return nil
}

type value struct {
    reply interface{}
    err   error
}

func (r *Redis) GetBytes(key string) (reply []byte, exists bool, err error) {
    return r.Get(key).Bytes()
}

//获取 string
func (r *Redis) Get(key string) *value {
    if r.autoClose {
        defer r.conn.Close()
    }
    reply, err := r.conn.Do("GET", key)
    return &value{
        reply: reply,
        err:   err,
    }
}

func (v *value) Bytes() (reply []byte, exists bool, err error) {
    reply, err = redis.Bytes(v.reply, v.err)
    exists, err = IfExistsOfErr(err)
    return
}

func (v *value) String() (reply string, exists bool, err error) {
    reply, err = redis.String(v.reply, v.err)
    exists, err = IfExistsOfErr(err)
    return
}

func (v *value) Bool() (reply bool, exists bool, err error) {
    reply, err = redis.Bool(v.reply, v.err)
    exists, err = IfExistsOfErr(err)
    return
}

func (v *value) Int() (reply int, exists bool, err error) {
    reply, err = redis.Int(v.reply, v.err)
    exists, err = IfExistsOfErr(err)
    return
}

func (v *value) Int64() (reply int64, exists bool, err error) {
    reply, err = redis.Int64(v.reply, v.err)
    exists, err = IfExistsOfErr(err)
    return
}

func (v *value) Float64() (reply float64, exists bool, err error) {
    reply, err = redis.Float64(v.reply, v.err)
    exists, err = IfExistsOfErr(err)
    return
}

func (v *value) Big() (reply *big.Big, exists bool, err error) {
    reply = new(big.Big)
    data, e := redis.String(v.reply, v.err)
    exists, err = IfExistsOfErr(e)
    if exists {
        if _, ok := reply.SetString(data); !ok {
            err = errors.New(data + " to decimal.Big err")
        }
    }
    return
}

func (v *value) Decimal() (reply decimal.Decimal, exists bool, err error) {
    data, e := redis.String(v.reply, v.err)
    exists, err = IfExistsOfErr(e)
    if exists {
        err = reply.Scan(data)
    }
    return
}

func (v *value) JsonUnmarshal(reply interface{}) (exists bool, err error) {
    data, e := redis.Bytes(v.reply, v.err)
    exists, err = IfExistsOfErr(e)
    if exists {
        err = json.Unmarshal(data, reply)
    }
    return
}

//分析error判断是否存在
func IfExistsOfErr(err error) (bool, error) {
    if err != nil {
        if err == redis.ErrNil {
            return false, nil
        } else {
            return false, err
        }
    }
    return true, nil
}

//将存储的数字加n；返回修改后的值
func (r *Redis) IncrBy(key string, n int64) (new int64, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    if n >= 0 {
        return redis.Int64(r.conn.Do("INCRBY", key, n))
    } else {
        return redis.Int64(r.conn.Do("DECRBY", key, -n))
    }
}

//判断哈希表的字段是否存在；存在返回 1，否则返回 0
func (r *Redis) HExists(key string, field string) (exists bool, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    return redis.Bool(r.conn.Do("HEXISTS", key, field))
}

//删除哈希表的一个或多个字段；返回删除数量
func (r *Redis) HDel(key string, fields ...string) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    if len(fields) > 0 {
        var args []interface{}
        args = append(args, key)
        for _, f := range fields {
            args = append(args, f)
        }
        _, err := redis.Int64(r.conn.Do("HDEL", args...))
        if err != nil {
            return err
        }
    }
    return nil
}

//插入/修改哈希表的字段，若哈希表不存在会新建；新建字段返回 1，覆盖字段返回 0
func (r *Redis) HSet(key string, field string, value interface{}) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    _, err := redis.Int64(r.conn.Do("HSET", key, field, value))
    if err != nil {
        return err
    }
    return nil
}

//插入/修改哈希表的一个或多个字段，若哈希表不存在会新建；成功返回 OK
func (r *Redis) HMSet(key string, hset map[string]interface{}) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    if len(hset) > 0 {
        var args []interface{}
        args = append(args, key)
        for k, v := range hset {
            args = append(args, k, v)
        }
        rs, err := redis.String(r.conn.Do("HMSET", args...))
        if err != nil {
            return err
        }
        if rs != "OK" {
            return errors.New(rs)
        }
    }
    return nil
}

//获取哈希表的指定字段；不存在时返回 nil
func (r *Redis) HGet(key string, field string) *value {
    if r.autoClose {
        defer r.conn.Close()
    }
    reply, err := r.conn.Do("HGET", key, field)
    return &value{
        reply: reply,
        err:   err,
    }
}

//获取哈希表的一个或多个字段；返回对应值列表，不存在的key的值为nil；可用string数组来接收
func (r *Redis) HMGet(key string, fields ...string) (fv map[string]string, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    fv = map[string]string{}
    if len(fields) > 0 {
        var args []interface{}
        args = append(args, key)
        for _, f := range fields {
            args = append(args, f)
        }
        reply, e := redis.Strings(r.conn.Do("HMGET", args...))
        if e != nil {
            return fv, e
        }
        for i, f := range fields {
            fv[f] = reply[i]
        }
    }
    return fv, nil
}

//获取哈希表的所有字段名和值；返回先field后紧跟其value的偶数个数的列表，不存在时返回空列表；可用map来接收
func (r *Redis) HGetAll(key string) (fv map[string]string, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    return redis.StringMap(r.conn.Do("HGETALL", key))
}

//获取哈希表的所有字段名和值；返回先field后紧跟其value的偶数个数的列表，不存在时返回空列表；可用map来接收
func (r *Redis) HGetAllStruct(key string, v interface{}) (exists bool, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    values, err := redis.Values(r.conn.Do("HGETALL", key))
    if err != nil {
        return false, err
    }
    if len(values) == 0 {
        return false, nil
    }
    return true, redis.ScanStruct(values, v)
}

//将哈希表 key 中的指定字段的整数值加上增量 n；返回修改后的值
func (r *Redis) HIncrBy(key string, field string, n int64) (new int64, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    return redis.Int64(r.conn.Do("HINCRBY", key, field, n))
}

//将哈希表 key 中的指定字段的浮点数值加上增量 n；返回修改后的值
func (r *Redis) HIncrByFloat(key string, field string, n float64) (new float64, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    return redis.Float64(r.conn.Do("HINCRBYFLOAT", key, field, n))
}

//插入元素和其分数，若元素已存在，则更新其分数并重新排序；返回新增元素的数量
func (r *Redis) ZAdd(key string, value string, score int64) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    _, err := redis.Int64(r.conn.Do("ZADD", key, score, value))
    return err
}

//插入元素和其分数，若元素已存在，则更新其分数并重新排序；返回新增元素的数量
func (r *Redis) ZMAdd(key string, mset map[string]int64) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    if len(mset) > 0 {
        var args []interface{}
        args = append(args, key)
        for value, score := range mset {
            args = append(args, score, value)
        }
        _, err := redis.Int64(r.conn.Do("ZADD", args...))
        return err
    }
    return nil
}

//获取指定元素的分数
func (r *Redis) ZScore(key string, value string) (score int64, exists bool, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    score, err = redis.Int64(r.conn.Do("ZSCORE", key, value))
    exists, err = IfExistsOfErr(err)
    return
}

//获取指定分数范围的元素（WITHSCORES先元素后紧跟其分数）；返回列表，可以string数组（map[value]score）接收
func (r *Redis) ZRangeByScore(key string, minScore int64, maxScore int64) (values map[string]string, err error) {
    if r.autoClose {
        defer r.conn.Close()
    }
    return redis.StringMap(r.conn.Do("ZRANGEBYSCORE", key, minScore, maxScore, "WITHSCORES"))
}

//删除集合中的一个或多个元素；返回成功删除的数量
func (r *Redis) ZRem(key string, values ...string) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    if len(values) > 0 {
        var args []interface{}
        args = append(args, key)
        for _, value := range values {
            args = append(args, value)
        }
        _, err := redis.Int64(r.conn.Do("ZREM", args...))
        return err
    }
    return nil
}

//删除指定分数区间的元素；返回成功删除的数量
func (r *Redis) ZRemRangeByScore(key string, minScore int64, maxScore int64) error {
    if r.autoClose {
        defer r.conn.Close()
    }
    _, err := redis.Int64(r.conn.Do("ZREMRANGEBYSCORE", key, minScore, maxScore))
    return err
}
