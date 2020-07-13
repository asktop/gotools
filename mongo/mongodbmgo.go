package mongo

import (
    "fmt"
    "github.com/asktop/gotools/log"
    "gopkg.in/mgo.v2"
)

var (
    mgoSession *mgo.Session
    database   string
)

type Config struct {
    Host        string `json:"host" yaml:"host"`
    Port        int    `json:"port" yaml:"port"`
    Username    string `json:"username" yaml:"username"`
    Password    string `json:"password" yaml:"password"`
    Database    string `json:"database" yaml:"database"`
    MaxPoolSize int    `json:"maxpoolsize" yaml:"maxpoolsize"`
}

func (c *Config) GetConfig() string {
    if c.Host == "" {
        c.Host = "127.0.0.1"
    }
    if c.Port == 0 {
        c.Port = 27017
    }
    if c.MaxPoolSize == 0 {
        c.MaxPoolSize = 500
    }
    return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authMechanism=SCRAM-SHA-1",
        c.Username,
        c.Password,
        c.Host,
        c.Port,
        c.Database,
    )
}

//初始化连接mongodb
func StartMongoDbMgo(config Config) error {
    //获取配置
    addr := config.GetConfig()
    log.Info("--- 连接 mongodb ---", "addr:", addr)
    //创建连接池
    session, err := mgo.Dial(addr)
    if err != nil {
        log.Error("--- 连接 mongodb 出错 ---", "err:", err)
        return err
    }
    //session设置的模式分别为:
    //Strong：
    //session 的读写一直向主服务器发起并使用一个唯一的连接，因此所有的读写操作完全的一致。
    //Monotonic：
    //session 的读操作开始是向其他服务器发起（且通过一个唯一的连接），只要出现了一次写操作，session 的连接就会切换至主服务器。由此可见此模式下，能够分散一些读操作到其他服务器，但是读操作不一定能够获得最新的数据。
    //Eventual：
    //session 的读操作会向任意的其他服务器发起，多次读操作并不一定使用相同的连接，也就是读操作不一定有序。session 的写操作总是向主服务器发起，但是可能使用不同的连接，也就是写操作也不一定有序。
    session.SetMode(mgo.Monotonic, true)
    session.SetPoolLimit(config.MaxPoolSize)
    mgoSession = session
    database = config.Database
    return nil
}

func NewMgoSession(typ ...int) *mgo.Session {
    if len(typ) > 0 && typ[0] == 1 {
        return mgoSession.Clone()
    }
    return mgoSession.Copy()
}

type mgoDb struct {
    session *mgo.Session
    db      string
}

func NewMgoDb(db ...string) *mgoDb {
    dbName := database
    if len(db) > 0 && db[0] != "" {
        dbName = db[0]
    }
    return &mgoDb{
        session: mgoSession.Copy(),
        db:      dbName,
    }
}

func (m *mgoDb) C(collection string) *mgo.Collection {
    return m.session.DB(m.db).C(collection)
}

func (m *mgoDb) Close() {
    m.session.Close()
}
