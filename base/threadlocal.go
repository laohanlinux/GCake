package base

import (
	"github.com/laohanlinux/go-logger/logger"
	"sync"
)

/*
*为goruntine增加本地存储,每个goruntine 都有自己的私有数据，他们都存储在
GoruntineStore 结构内, 可以根据GoruntineId 进行检索!!!
* */

type GoruntineId int32

type GoruntineStoreData interface{}

//type GoruntineStoreData []interface{}

type GoruntineStore struct {
	c        chan int32
	gsVector map[GoruntineId]GoruntineStoreData
	lock     *sync.Mutex
}

var goruntineStore *GoruntineStore

func init() {
	CreateGoruntineStore()
	go func() {
		if goruntineStore != nil {
			for {
				GoruntineDeleteSpecific(<-goruntineStore.c)
			}
		}
	}()
}

func CreateGoruntineStore() {
	if goruntineStore != nil {
		return
	}
	goruntineStore = new(GoruntineStore)
	goruntineStore.c = make(chan int32)
	goruntineStore.gsVector = make(map[GoruntineId]GoruntineStoreData)
	goruntineStore.lock = &sync.Mutex{}
}

func GoruntineSetSpecific(value GoruntineStoreData) {
	defer func() {
		goruntineStore.lock.Unlock()
	}()
	goruntineStore.lock.Lock()
	goruntineStore.gsVector[GoruntineId(CurrentGoroutineId())] = value
}

//func GoruntineSetSpecific(goruntineStore *GoruntineStore, value interface{}) {
//append(goruntineStore.GSVector[runtime.GetGoId], value)
//}

func GoruntineGetSpecific() GoruntineStoreData {
	defer func() {
		goruntineStore.lock.Unlock()
	}()
	goruntineStore.lock.Lock()
	return goruntineStore.gsVector[GoruntineId(CurrentGoroutineId())]
}

func GoruntineDeleteSpecific(gid int32) {
	logger.Info("delete a thread local store, gid: ", gid)
	delete(goruntineStore.gsVector, GoruntineId(gid))
}

/*func GoruntineGetSpecific(goruntineStore *GoruntineStore) {*/
//return goruntineStore.GSVector[runtime.GetGoId()]
/*}*/
