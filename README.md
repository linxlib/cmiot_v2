# cmiot_v2
 中移v2接口sdk go版（非官方）
 
# 适用

- https://api.iot.10086.cn/v2/ ， 移动旧平台， [官方接口文档地址](https://api.iot.10086.cn/api/index.html#/resourceCenter/docDownload)
 
# 版本

[![GitHub version](https://badge.fury.io/gh/linxlib%2Fcmiot_v2.svg)](https://badge.fury.io/gh/linxlib%2Fcmiot_v2)

# 使用方法

```go
package main

import (
    cmiot "github.com/linxlib/cmiot_v2"
)

const (
	APPID    = "xxxx"
	PASSWORD = "xxxx"
	ECID     = "xxxx"
)

func main()  {
    c := cmiot.NewCMIOTClient(APPID, PASSWORD, ECID)
 	c.SetForwarderId("ip")
 	c.SetDebug(false)
}

```

- `SetForwarderId` 官方后台指定的白名单服务器的ip地址，不需要把程序跑在白名单服务上，也不需要使用代理服务器
- `SetDebug` 设为true则在终端输出http请求内容和返回的数据

# 进度
目前暂只实现了部分单卡查询的接口， 批量查询的尚未实现

# 注意
- sdk未经过完全测试
- 如果开多个协程进行调用， 需要控制请求频率， 否则容易无法返回正常的数据。 可能在多台电脑上开多个程序（手动分布式）运行会比较好。

# 引用
sdk引用了部分第三方库的代码

- `https://github.com/guonaihong/gout` 中的Query Encode代码
- `https://github.com/NETkiddy/cmq-go` 中的http请求的部分代码



