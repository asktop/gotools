package cache

import (
    "fmt"
    "testing"
)

func TInit() {
    err := StartRedis(Config{Host: "127.0.0.1", Password: "kf123456", Select: 0})
    if err != nil {
        panic(err)
    }
}

//测试存储
func TestSet(t *testing.T) {
    TInit()
    old, err := NewRedis().Set("test:k1", "v1")
    fmt.Println(err)
    fmt.Println(old)
}

//测试获取
func TestGet(t *testing.T) {
    TInit()
    rs, ex, err := NewRedis().Get("test:k1").String()
    fmt.Println(err)
    fmt.Println(ex)
    fmt.Println(rs)
}

//测试删除
func TestDel(t *testing.T) {
    TInit()
    err := NewRedis().Del("test:k1")
    fmt.Println(err)
}

//测试增量
func TestIncr(t *testing.T) {
    TInit()
    rs, err := NewRedis().IncrBy("test:k2", 2)
    fmt.Println(err)
    fmt.Println(rs)
}

//插入/修改哈希表的字段，若哈希表不存在会新建；新建字段返回 1，覆盖字段返回 0
func TestHSet(t *testing.T) {
    TInit()
    err := NewRedis().HSet("test:h1", "k1", "v1")
    fmt.Println(err)
}

//插入/修改哈希表的一个或多个字段，若哈希表不存在会新建；成功返回 OK
func TestHMSet(t *testing.T) {
    TInit()
    mset := map[string]interface{}{}
    mset["a"] = "abc"
    mset["b"] = 123.23542
    mset["c"] = 123
    err := NewRedis().HMSet("test:h1", mset)
    fmt.Println(err)
}

//获取哈希表的指定字段；不存在时返回 nil
func TestHGet(t *testing.T) {
    TInit()
    rs, ex, err := NewRedis().HGet("test:h1", "b").String()
    fmt.Println(err)
    fmt.Println(ex)
    fmt.Println(rs)
}

//获取哈希表的一个或多个字段；返回对应值列表，不存在的key的值为nil；可用string数组来接收
func TestHMGet(t *testing.T) {
    TInit()
    rs, err := NewRedis().HMGet("test:h1", "a", "c", "d", "b")
    fmt.Println(err)
    fmt.Println(rs)
}

//获取哈希表的所有字段名和值；返回先field后紧跟其value的偶数个数的列表，不存在时返回空列表；可用map来接收
func TestHGetAll(t *testing.T) {
    TInit()
    rs, err := NewRedis().HGetAll("test:h2")
    fmt.Println(err)
    fmt.Println(rs)
    fmt.Println(len(rs) == 0)
    fmt.Println(rs == nil)
}

//获取哈希表的所有字段名和值；返回先field后紧跟其value的偶数个数的列表，不存在时返回空列表；可用map来接收
func TestHGetAllStruct(t *testing.T) {
    TInit()
    abc := ABC{}
    ex, err := NewRedis().HGetAllStruct("test:h1", &abc)
    fmt.Println(err)
    fmt.Println(ex)
    fmt.Println(abc)
}

type ABC struct {
    A string `redis:"a"`
    B string `redis:"b"`
    C string `redis:"c"`
}

//将哈希表 key 中的指定字段的整数值加上增量 n；返回修改后的值
func TestHIncrBy(t *testing.T) {
    TInit()
    rs, err := NewRedis().HIncrBy("test:h1", "c", -2)
    fmt.Println(err)
    fmt.Println(rs)
}

//将哈希表 key 中的指定字段的浮点数值加上增量 n；返回修改后的值
func TestHIncrByFloat(t *testing.T) {
    TInit()
    rs, err := NewRedis().HIncrByFloat("test:h1", "b", -3)
    fmt.Println(err)
    fmt.Println(rs)
}

//插入元素和其分数，若元素已存在，则更新其分数并重新排序；返回新增元素的数量
func TestZAdd(t *testing.T) {
    TInit()
    err := NewRedis().ZAdd("test:z1", "b", 2)
    fmt.Println(err)
}

//插入元素和其分数，若元素已存在，则更新其分数并重新排序；返回新增元素的数量
func TestZMAdd(t *testing.T) {
    TInit()
    a := map[string]int64{}
    a["a"] = 1
    a["b"] = 2
    a["c"] = 2
    a["d"] = 3
    a["e"] = 4
    err := NewRedis().ZMAdd("test:z1", a)
    fmt.Println(err)
}

//获取指定元素的分数
func TestZScore(t *testing.T) {
    TInit()
    rs, ex, err := NewRedis().ZScore("test:z1", "bbb")
    fmt.Println(err)
    fmt.Println(ex)
    fmt.Println(rs)
}

//获取指定分数范围的元素（WITHSCORES先元素后紧跟其分数）；返回列表，可以string数组（map）接收
func TestZRangeByScore(t *testing.T) {
    TInit()
    rs, err := NewRedis().ZRangeByScore("test:z1", 2, 3)
    fmt.Println(err)
    fmt.Println(rs)
}

//删除集合中的一个或多个元素；返回成功删除的数量
func TestZRem(t *testing.T) {
    TInit()
    err := NewRedis().ZRem("test:z1", "aaa")
    fmt.Println(err)
}

//删除指定分数区间的元素；返回成功删除的数量
func TestZRemRangeByScore(t *testing.T) {
    TInit()
    err := NewRedis().ZRemRangeByScore("test:z1", 2, 3)
    fmt.Println(err)
}
