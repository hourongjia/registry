package registry

import (
	"github.com/orca-zhang/ecache"
	"time"
)

// RegistryTableApi 注册表
type RegistryTableApi interface {
	AddData(key string, value string) error
	AddDataWithDeadLine(key string, value string, deadline float64) error
	DeleteData(key string) error
	UpdateDataTime(key string, deadline float64) error
	GetData(key string) (interface{}, error)
}

type RegistryCache struct {
	Cache *ecache.Cache
}

// NewRegistryCache 创建注册本地cache
func NewRegistryCache(bucketCnt uint16) (RegistryCache, error) {
	var c = ecache.NewLRUCache(bucketCnt, 200, 10*time.Second)
	return RegistryCache{
		Cache: c,
	}, nil
}

type CommonRegistry struct {
	Cache *RegistryCache
}

func (receiver *CommonRegistry) AddData(key string, value string) error {
	receiver.Cache.Cache.Put(key, value)
	// 注册到注册中新
	return nil
}

func (receiver *CommonRegistry) DeleteData(key string) error {
	// 从注册中心删除
	receiver.Cache.Cache.Del(key)
	return nil
}

func (receiver *CommonRegistry) GetData(key string) (interface{}, error) {
	// 从注册中心获取数据
	return receiver.Cache.Cache.Get(key), nil
}

func (receiver *CommonRegistry) AddDataWithDeadLine(key string, value string, deadline float64) error {
	// 设置deadline到注册中心
	return nil
}

func (receiver *CommonRegistry) UpdateDataTime(key string, value float64) error {
	// 更新数据完成时间戳
	return nil
}
