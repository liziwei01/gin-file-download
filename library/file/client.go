/*
 * @Author: liziwei01
 * @Date: 2022-03-09 19:26:04
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-04-06 16:43:49
 * @Description: file content
 */
package file

import (
	"sync"
)

var (
	// 初始化互斥锁
	mu sync.Mutex
)

type Client interface {
	name() string
	Location() string
}

type client struct {
	conf *Config
}

func New(config *Config) Client {
	c := &client{
		conf: config,
	}
	return c
}

func NewDefault() Client {
	c := &client{
		conf: nil,
	}
	return c
}

func (c *client) name() string {
	return c.conf.Name
}

func (c *client) Location() string {
	return c.conf.Location
}
