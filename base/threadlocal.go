package base

import (
	"github.com/funny/goid"
	"sync"
)

/*
*为goruntine增加本地存储,每个goruntine 都有自己的私有数据，他们都存储在
GoruntineStore 结构内, 可以根据GoruntineId 进行检索!!!
* */

type GoruntineId int

type GoruntineStoreData interface{}

//type GoruntineStoreData []interface{}

type GoruntineStore struct {
	c        chan int
	gsVector map[GoruntineId]GoruntineStoreData
	lock     *sync.Mutex
}

var goruntineStore *GoruntineStore

func init() {
	CreateGoruntineStore()
}

var p int64

func CreateGoruntineStore() {
	if goruntineStore != nil {
		return
	}
	goruntineStore = new(GoruntineStore)
	goruntineStore.c = make(chan int)
	goruntineStore.gsVector = make(map[GoruntineId]GoruntineStoreData)
	goruntineStore.lock = &sync.Mutex{}
}

func GoruntineSetSpecific(value GoruntineStoreData) {
	defer func() {
		goruntineStore.lock.Unlock()
	}()
	goruntineStore.lock.Lock()
	goruntineStore.gsVector[GoruntineId(goid.Get())] = value
}

//func GoruntineSetSpecific(goruntineStore *GoruntineStore, value interface{}) {
//append(goruntineStore.GSVector[runtime.GetGoId], value)
//}

func GoruntineGetSpecific() GoruntineStoreData {
	defer func() {
		goruntineStore.lock.Unlock()
	}()
	goruntineStore.lock.Lock()
	return goruntineStore.gsVector[GoruntineId(goid.Get())]
}

/*func GoruntineGetSpecific(goruntineStore *GoruntineStore) {*/
//return goruntineStore.GSVector[runtime.GetGoId()]
/*}*/
