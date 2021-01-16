package alru

import (
    "container/list"
    "sync"
)

type Node struct {
    Key interface{}
    Val interface{}
}

type Lru struct {
    mu    sync.Mutex
    size  int
    vlist *list.List
    cache map[interface{}]*list.Element
}

func NewLru(size int) *Lru {
    return &Lru{
        size:  size,
        vlist: list.New(),
        cache: make(map[interface{}]*list.Element),
    }
}

func (l *Lru) init() {
    if l.vlist == nil {
        l.vlist = list.New()
    }
    if l.cache == nil {
        l.cache = make(map[interface{}]*list.Element)
    }
}

func (l *Lru) Add(key interface{}, val interface{}) *Node {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.init()
    //已经存在，更新
    if e, ok := l.cache[key]; ok {
        e.Value.(*Node).Val = val
        l.vlist.MoveToFront(e)
        return nil
    }
    //不存在，新增
    newe := l.vlist.PushFront(&Node{
        Key: key,
        Val: val,
    })
    l.cache[key] = newe

    if l.size > 0 && l.vlist.Len() > l.size {
        if e := l.vlist.Back(); e != nil {
            node := e.Value.(*Node)
            delete(l.cache, node.Key)
            l.vlist.Remove(e)
            return node
        }
    }
    return nil
}

func (l *Lru) Get(key interface{}) (val interface{}, ok bool) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.init()

    if e, ok := l.cache[key]; ok {
        l.vlist.MoveToFront(e)
        return e.Value.(*Node).Val, true
    }

    return nil, false
}

func (l *Lru) Del(key interface{}) *Node {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.init()

    if e, ok := l.cache[key]; ok {
        delete(l.cache, key)
        l.vlist.Remove(e)
        node := e.Value.(*Node)
        return node
    }

    return nil
}

func (l *Lru) GetAll() []*Node {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.init()

    var data []*Node
    for e := l.vlist.Front(); e != nil; e = e.Next() {
        data = append(data, e.Value.(*Node))
    }
    return data
}
