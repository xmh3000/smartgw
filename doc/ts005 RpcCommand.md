# RPC命令
> ## 命令总览
> v1/devices/me/rpc/request/***RequestID***

> **面向网关**
> 
> - 重启
> - 校时
> - 升级
> - 打开看门狗
> - 关闭看门狗
> - 获取网关表具列表
> - 下发网关表具列表
> - ping命令
> 
> **面向表具**
> 
> - 充值
> - 开合闸
> - 设置单价
> - 读取单价
> - 保电
> - 读取金额
> - 读取电量
> - 路灯校时
> - 常规策略下发
> - 临时策略下发
> - 读路灯状态

> ## 面向网关

> ### 重启
> **REQUEST**
> ```
> {
>   "method": "gateway",
>   "requestID": "1",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "reboot",
>       }
>   ]
> }
> ```
> **RESPONSE**
> ```
> 无回复
> ```

> ### 校时
> **REQUEST**
> ```
> {
>   "method": "gateway",
>   "requestID": "2",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "ntp",
>       }
>   ]
> }
> ```
> **RESPONSE**
> ```
> {
>   "method": "gateway",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "npt",
>           "cmdStatus": 0 | 1,
>       }
>   ]
> }
> ```

> ### 升级
> **REQUEST**
> ```
> {
>   "method": "gateway",
>   "requestID": "3",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "upgrade",
>           "cmdParams": {
>               "url": "http://xxxxxxxxx"
>           }
>       }
>   ]
> }
> ```
> **RESPONSE**
> ```
> 无回复
> ```


> ### 打开看门狗
> **REQUEST**
> ```
> {
>   "method": "gateway",
>   "requestID": "4",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "watchdogstart",
>       }
>   ]
> }
> ```
> **RESPONSE**
> ```
> {
>   "method": "gateway",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "watchdogstart",
>           "cmdStatus": 0 | 1,
>       }
>   ]
> }
> ```


> ### 关闭看门狗
> **REQUEST**
> ```
> {
>   "method": "gateway",
>   "requestID": "5",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "watchdogstop",
>       }
>   ]
> }
> ```
> **RESPONSE**
> ```
> {
>   "method": "gateway",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "watchdogstop",
>           "cmdStatus": 0 | 1,
>           "cmdResult":
>       }
>   ]
> }
> ```

> ### 获取网关表具列表
> **REQUEST**
> ```
> {
>   "method": "gateway",
>   "requestID": "6",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "getDevices",
>       }
>   ]
> }
> ```
> **RESPONSE**
> ```
> {
>   "method": "gateway",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "getDevices",
>           "cmdStatus": 0 | 1,
>           "cmdResult":
>       }
>   ]
> }
> ```

> ### 下发网关表具列表
> **REQUEST**
> ```
> {
>   "method": "gateway",
>   "requestID": "7",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "setDevices",
>       }
>   ]
> }
> ```
> **RESPONSE**
> ```
> {
>   "method": "gateway",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "setDevices",
>           "cmdStatus": 0 | 1,
>           "cmdResult": [
>                           {
>                               "name": "XAWL001",
>                               "type": {
>                                           "name": "DDSY1946A",
>                                           "driver": "/home/plugin",
>                                           "properties": [{ 
>                                               "name": "DDSY1946A",
>                                               "description": "null",
>                                               "type": "double",
>                                               "length": 8,
>                                               "decimal": 2,
>                                               "unit": "kwh",
>                                               "value": 36.21,  任意类型
>                                               "reported": true | false,
>                                           },
>                                           ……
>                                       }
>                               "address": "001",
>                               "collector": {
>                                               "name": "rs485-1",
>                                               "type": "serial" | "tcpClient", 二选一
>                                               "serial": {
>                                                           "name": "/dev/tyy",
>                                                           "deviceName": "DDSY1946A",
>                                                           "baudRate": 9600,
>                                                           "dataBit": 8,
>                                                           "stopBit": "1.5",
>                                                           "check": "o",
>                                                           }
>                                               "tcpClient": {
>                                                           "name": "tcp1",
>                                                           "ip": "192.168.100.13",
>                                                           "port": 1000,
>                                                           }
>                                             },
>                               "timeout": 0,
>                               "interval": 120,
>                               "alone": true | false,   如果是true，需要增加下一个成员serial
>                               "serial": ,
>                               "online": true | false,
>                               "collectTime": 1666666666,    uint64类型
>                               "collectTotal": 100，
>                               "collectSuccess": 99,
>                               "reportTime": 1666666666,    uint64类型
>                               "reportTotal": 10,
>                               "reportSuccess": 9,
>                           },
>                           ……
>                       ]
>       }
>   ]
> }
> ```

> ### 下发网关表具列表
> **REQUEST**
> ```
> {
>   "method": "gateway",
>   "requestID": "8",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "ping",
>       }
>   ]
> }
> ```
> **RESPONSE**
> ```
> {
>   "method": "gateway",
>   "params": [
>       {
>           "clientID": "YBD001",
>           "cmdName": "ping",
>           "cmdStatus": 0,
>           "cmdResult": 
>       }
>   ]
> }
> ```