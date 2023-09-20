package sms

import (
	"sync"
)

var (
	_tencentStatus    bool // 腾讯云启用状态
	_volcengineStatus bool // 火山引擎启用状态
	_mutex            sync.RWMutex
)

// Client define sms method
type Client interface {
	Send(mobile string, code string) error // 发送验证码短信
}

type ServerType int8

const (
	ServerTypeTencent    ServerType = 0
	ServerTypeVolcengine ServerType = 1
)

// NewClient
func NewClient(conf TencentConfig, volcConf VolcConfig, server ServerType) Client {
	// 若存在服务被禁用，返回启用的服务商client
	_mutex.RLock()
	if _tencentStatus && !_volcengineStatus {
		server = ServerTypeTencent
	} else if !_tencentStatus && _volcengineStatus {
		server = ServerTypeVolcengine
	}
	_mutex.RUnlock()

	// new
	switch server {
	case ServerTypeTencent:
		return NewTencentSmsClient(conf)
	case ServerTypeVolcengine:
		return NewVolcSmsClient(volcConf)
	default:
		// 默认走腾讯
		return NewTencentSmsClient(conf)
	}
}

func UpdateServerStatus(server ServerType, status bool) {
	_mutex.Lock()
	defer _mutex.Unlock()

	switch server {
	case ServerTypeTencent:
		_tencentStatus = status
	case ServerTypeVolcengine:
		_volcengineStatus = status
	}
}
