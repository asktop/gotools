package atx

import (
    "github.com/asktop/dbr"
    "github.com/asktop/gotools/db"
)

/* ******************** */
/* 数据库事务 公共部分 */
/* ******************** */

type Tx struct {
    tx dbr.SessionRunner //session 或 tx
}

//设置tx和是否锁表（tx的情况下默认锁表）
func (m *Tx) SetTx(tx dbr.SessionRunner) dbr.SessionRunner {
    if tx == nil {
        m.tx = db.Session()
    } else {
        m.tx = tx
    }
    return m.GetTx()
}

//获取tx
func (m *Tx) GetTx() (tx dbr.SessionRunner) {
    if m.tx == nil {
        m.tx = db.Session()
    }
    return m.tx
}
