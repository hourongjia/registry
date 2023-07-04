package registry

import (
	"errors"
	"github.com/patrickmn/go-cache"
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
	Cache *cache.Cache
}

// NewRegistryCache 创建注册本地cache
func NewRegistryCache(bucketCnt uint16) (RegistryCache, error) {
	Cache := cache.New(cache.NoExpiration, cache.NoExpiration)
	return RegistryCache{
		Cache: Cache,
	}, nil
}

type CommonRegistry struct {
	Subject Subject
	Cache   *RegistryCache
}

func (receiver *CommonRegistry) AddData(key string, value string) error {
	receiver.Cache.Cache.Add(key, value, cache.NoExpiration)
	// 注册到注册中新
	event := Event{
		Async:     false,
		EventType: Add,
		Data:      EventData{},
	}
	receiver.Subject.NotifyObservers(event)
	return nil
}

func (receiver *CommonRegistry) DeleteData(key string) error {
	// 从注册中心删除
	receiver.Cache.Cache.Delete(key)
	// 注册到注册中新
	event := Event{
		Async:     false,
		EventType: Delete,
		Data:      EventData{},
	}
	receiver.Subject.NotifyObservers(event)
	return nil
}

func (receiver *CommonRegistry) GetData(key string) (interface{}, error) {
	// 注册到注册中新
	event := Event{
		Async:     false,
		EventType: Select,
		Data:      EventData{},
	}
	receiver.Subject.NotifyObservers(event)
	return receiver.Cache.Cache.Get(key), nil
}

func (receiver *CommonRegistry) AddDataWithDeadLine(key string, value string, deadline time.Duration) error {
	// 注册到注册中新
	event := Event{
		Async:     false,
		EventType: Add,
		Data:      EventData{},
	}
	receiver.Subject.NotifyObservers(event)
	// 设置deadline到注册中心
	receiver.Cache.Cache.Set(key, value, deadline)
	return nil
}

func (receiver *CommonRegistry) UpdateDataTime(key string, expire time.Duration) error {
	// 注册到注册中新
	event := Event{
		Async:     false,
		EventType: Update,
		Data:      EventData{},
	}
	receiver.Subject.NotifyObservers(event)
	// 更新数据完成时间戳
	vs, ok := receiver.Cache.Cache.Get(key)
	if !ok {
		return errors.New("can not find data")
	}
	receiver.Cache.Cache.Set(key, vs, expire)
	return nil
}
