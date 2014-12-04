# EventLoop

## DataStructure

- eventloop

```
type EventLoop struct {
    // atomic , 是否处于事件循环
    looping_ bool
    //当前对象所属的ID
    threadid_ int32

    t_looInThisThread *EventLoop
    this              *EventLoop
}
```

每个线程最多只能有一个`EventLoop` 对象，当一个`eventloop`成功被创建时，会以`{gid, gid}`键值对写到`goruntineStore`中。


- GoroutineStore
线程本地化存储
```
type GoruntineStoreData interface{}

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
```

这之事一个单例模式。

