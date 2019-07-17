package kernel

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var (
	cacheKr *redis.Client
	redisMX sync.Mutex
	redisMC map[string]interface{}
)

func init() {
	if redisMC == nil {
		initRedisMC()
	}
}

type Kr interface {
	Get(key string) string
	Start() *r
	Engine() *redis.Client
	Close()
}

type r struct {
	kr *redis.Client
}

var _ Kr = &r{}

func NewRedis() *r {
	var object *r = &r{}
	if cacheKr != nil {
		object.kr = cacheKr
	}

	return object
}

func (thisKr *r) Start() *r {
	if thisKr.kr == nil {
		thisKr.instanceMaster()
	}

	return thisKr
}

func (thisKr *r) Engine() *redis.Client {
	if thisKr.kr == nil {
		thisKr.instanceMaster()
	}

	return thisKr.kr
}

func (thisKr *r) instanceMaster() *r {
	redisMX.Lock()
	defer redisMX.Unlock()

	if cacheKr != nil {
		thisKr.kr = cacheKr
		return thisKr
	}

	clientKr := redis.NewClient(&redis.Options{
		Addr:     redisMC["Addr"].(string),
		Password: redisMC["Password"].(string),
		DB:       redisMC["DB"].(int),
		PoolSize: redisMC["PoolSize"].(int),
	})

	_, err := clientKr.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("ping error[%s]\n", err.Error()))
	}

	if cacheKr == nil {
		cacheKr = clientKr
	}

	thisKr.kr = clientKr

	return thisKr
}

func (thisKr *r) Get(key string) string {
	v, err := thisKr.kr.Get(key).Result()
	if err != nil {
		panic(fmt.Sprintf("ping error[%s]\n", err.Error()))
	}

	return v
}

func (thisKr *r) Close() {
	if thisKr.kr != nil {
		thisKr.kr.Close()
	}
}

func initRedisMC() {
	var c INI = NewIni().LoadByFN(ConfRedis)

	addr 			:= c.K(
		MapConfLists[ConfRedis][0],
		MapConfParam[MapConfLists[ConfRedis][0]][0],
	).String()
	if addr == "" {addr = KRedisAddr}

	password 		:= c.K(
		MapConfLists[ConfRedis][0],
		MapConfParam[MapConfLists[ConfRedis][0]][1],
	).String()

	db, err 		:= c.K(
		MapConfLists[ConfRedis][0],
		MapConfParam[MapConfLists[ConfRedis][0]][2],
	).Int()
	if err != nil {db = KRedisDB}

	intPoolSize, err:= c.K(
		MapConfLists[ConfRedis][0],
		MapConfParam[MapConfLists[ConfRedis][0]][3],
	).Int()
	if err != nil {intPoolSize = KRedisPoolSize}

	redisMC = map[string]interface{}{
		"Addr":     addr,
		"Password": password,
		"DB":       db,
		"PoolSize": intPoolSize,
	}
}
