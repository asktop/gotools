package mongo

import (
    "fmt"
    "github.com/asktop/decimal"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "testing"
)

func TInit() {
    err := StartMongoDbMgo(Config{Host: "127.0.0.1", Port: 27017, Username: "", Password: "", Database: "test"})
    if err != nil {
        panic(err)
    }
}

type persionMgo struct {
    Id   bson.ObjectId `bson:"_id"`
    Name string        `bson:"name"`
    Age  int           `bson:"age"`
}

//插入
func TestMgoInsert(t *testing.T) {
    TInit()
    conn := NewMgoDb()
    defer conn.Close()
    err := conn.C("test").Insert(bson.M{"name": "test1", "age": 21})
    fmt.Println(err)
    err = conn.C("test").Insert(bson.M{"name": "test2", "age": 22})
    fmt.Println(err)
    err = conn.C("test").Insert(bson.M{"name": "test3", "age": 23})
    fmt.Println(err)
}

//插入多条
func TestMgoInsertMany(t *testing.T) {
    TInit()
    conn := NewMgoDb()
    defer conn.Close()
    var rows []interface{}
    rows = append(rows, bson.M{"name": "test4", "age": 11}, bson.M{"name": "test5", "age": 12}, bson.M{"name": "test6", "age": 13})
    err := conn.C("test").Insert(rows...)
    fmt.Println(err)
}

//插入多条
type User struct {
    Id       int64  `json:"id" bson:"id"`
    UserName string `json:"user_name" bson:"user_name"`
    Age      string `json:"age" bson:"age"`
}

func TestMgoInsertMany2(t *testing.T) {
    TInit()
    conn := NewMgoDb()
    defer conn.Close()
    rows := []interface{}{
        User{Id: 7, UserName: "test7", Age: decimal.New(185469685145454, 5).String()},
        User{Id: 8, UserName: "test8", Age: decimal.New(185469685145454, 8).String()},
        User{Id: 8, UserName: "test9", Age: decimal.New(185469685145454, 10).String()},
    }
    err := conn.C("test").Insert(rows...)
    fmt.Println(err)
}

//查询全部
func TestMgoFindAll(t *testing.T) {
    TInit()
    findMgoAll()
}

func findMgoAll() {
    //datas := []bson.M{}
    datas := []persionMgo{}
    conn := NewMgoDb()
    defer conn.Close()
    err := conn.C("test").Find(nil).Sort("age").All(&datas)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, data := range datas {
            fmt.Println(data.Id.String(), data)
        }
    }
}

//查询条件
func TestMgoFind(t *testing.T) {
    TInit()
    datas := []bson.M{}
    conn := NewMgoDb()
    defer conn.Close()
    err := conn.C("test").Find(bson.M{"name": "test1", "age": bson.M{"$lt": 12}}).All(&datas)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, data := range datas {
            fmt.Println(data)
        }
    }
    fmt.Println("--------------")
    err = conn.C("test").Find(nil).Select(bson.M{"_id": 0, "name": 1, "age": 1}).All(&datas)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, data := range datas {
            fmt.Println(data)
        }
    }
}

type User2 struct {
    MId      bson.ObjectId `json:"_id" bson:"_id"`
    Id       int64         `json:"id" bson:"id"`
    Name     string        `json:"name" bson:"name"`
    UserName string        `json:"user_name" bson:"user_name"`
    Age      string        `json:"age" bson:"age"`
}

func TestMgoFind2(t *testing.T) {
    TInit()
    datas := []User2{}
    conn := NewMgoDb()
    defer conn.Close()
    err := conn.C("test").Find(nil).All(&datas)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, data := range datas {
            fmt.Println(data)
        }
    }
}

//查询条件
func TestMgoFind3(t *testing.T) {
    TInit()
    datas := []bson.M{}
    conn := NewMgoDb()
    defer conn.Close()
    err := conn.C("test").Pipe([]bson.M{
        bson.M{"$match": bson.M{"name": bson.M{"$in": []string{"test1", "test2", "test4"}}}},
        bson.M{"$group": bson.M{"_id": "$name"}},
    }).All(&datas)
    if err != nil {
        fmt.Println(err)
    } else {
        for _, data := range datas {
            fmt.Println(data)
        }
    }
}

//更新
func TestMgoUpdate(t *testing.T) {
    TInit()
    conn := NewMgoDb()
    defer conn.Close()
    info, err := conn.C("test").UpdateAll(bson.M{"name": "test1", "age": bson.M{"$lte": 10}}, bson.M{"$set": bson.M{"age": 9}})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(info)
    }
    findMgoAll()
}

//删除
func TestMgoRemove(t *testing.T) {
    TInit()
    conn := NewMgoDb()
    defer conn.Close()
    data := bson.M{}
    err := conn.C("test").Find(bson.M{"name": "test3"}).One(&data)
    if err != nil {
        fmt.Println(err)
        fmt.Println(err == mgo.ErrNotFound)
    } else {
        id := data["_id"]
        fmt.Println(id)
        err = conn.C("test").RemoveId(id)
        fmt.Println(err)
    }
    findMgoAll()
}

//批量删除
func TestMgoRemoveAll(t *testing.T) {
    TInit()
    conn := NewMgoDb()
    defer conn.Close()
    var names []string
    names = append(names, "test4", "test5", "test6")
    _, err := conn.C("test").RemoveAll(bson.M{"name": bson.M{"$in": names}})
    if err != nil {
        fmt.Println(err)
    }
    findMgoAll()
}
