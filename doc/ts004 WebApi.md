## 一、Web Api
### 约定
返回值:
```
{
  "code": 0 | 1,
  "message": "成功！" | "失败！",
  "data": nil | {
    ...
  }
}
```
### （一）采集接口
1. 实体
```
  user : {
    "id": 0,
    "name": "485-1",
    "type": "serial" | "tcpClient",
    "serial": {
      "name": "本地串口",
      "deviceName“: "/dev/ttymxc6",
      "baudRate": 9600,
      "dataBit": 8,
      "stopBit": 1,
      "check": "N"|"E"|"O"
    },
    "tcpClient": {
      "name": "tcp21",
      "ip": "192.168.100.21",
      "port" 7000,
    },
    "timeout": 3000,
    "interval": 200
  }
```
2. 新增
```
url: /api/collector
方法: POST
  data: user
返回:（主要是id）
  {
    "code": 0 | 1,
    "message": "",
    "data": user
  }
```
2. 编辑（修改）
```
url: /api/collector
方法: PUT
  data: user
返回:
  {
    "code": 0 | 1,
    "message": "修改采集接口成功！",
    "data": user
  }
```
3. 删除
```
url: /api/collector/1
方法: DELETE
返回:（主要是id）
  {
    "code": 0 | 1,
    "message": "删除数据成功！" | "删除数据失败@！",
  }
```
4. 详情（查看）
```
url: /api/collector/1
方法: GET
返回:
  {
    "code": 0,
    "message": "读取数据成功！",
    "data": user
  }
```
5. 获取所有数据
```
url: /api/collectors
方法: GET
返回:
  {
    "code": 0,
    "message": "读取数据成功！",
    "data": [{
    ......
    },{
    ......
    },{
    ......
    }]
  }
```
6. 不提供分页查询
```
nothing to do
```
### （二）设备类型
#### 1 设备类型
实体：
```
deviceType:
{
    "id": 0,
    "name": "DDST422",
    "driver": "DDST422",
    "properties": [
        {
            "name": "电压",
            "description": "电表电压",
            "type": "double",
            "length": 13,
            "decimal": 2,
            "unit": "V",
            "va;ue": null
        },
        {
            "name": "电流",
            "description": "电表电流",
            "type": "double",
            "length": 13,
            "decimal": 2,
            "unit": "A",
            "va;ue": null
        }
    ]
}

deviceProperty:
{
            "name": "电流",
            "description": "电表电流",
            "type": "double",
            "length": 13,
            "decimal": 2,
            "unit": "A",
            "va;ue": null
}
```
1. 新增
```
url: /api/device-type
method: POST
返回值：
{
    "code": 0,
    "data": deviceType,
    "message": "新增设备类型成功！"
}
```
3. 编辑（修改）
```
url: /api/device/type
method: PUT
返回值:
{
    "code": 0,
    "data": deviceType,
    "message": "修改设备类型成功！"
}
```
4. 删除
```
url: /api/device-type/:id
method: DELETE
返回值:
{
    "code": 0,
    "data": null,
    "message": "删除设备类型成功！"
}
```
5. 详情（查看）
```
url: /api/device-type/:id
method: GET
返回值:
{
    "code": 0,
    "data": deviceType,
    "message": "查找设备类型成功！"
}
```
6. 获取所有数据
```
url: /api/device-types
method: GET
返回值:
{
    "code": 0,
    "data": [
        deviceType,
        deviceType
    ],
    "message": "获取所有设备类型数据成功！"
}
```
8. 上传驱动程序
```
url: /device-type/upload/:id
method: POST
上传文件：FileName
返回值:
{
    "code": 0,
    "data": nil,
    "message": "上传文件成功"
}
```
7. 不提供分页查询
#### 2 设备属性
1. 新增
```
url: /api/device-property/:id
method: POST
{
    "code": 0,
    "data": deviceProperty,
    "message": "新增设备属性成功！"
}
```
2. 编辑
```
url: /api/device-poperty/:id/:propertyid
method: PUT
{
    "code": 0,
    "data": deviceProperty,
    "message": "修改设备属性成功！"
}
```
3. 删除
```
url: /api/device-property/:id/:propertyid
method: DELETE
返回值:
{
    "code": 0,
    "data": null,
    "message": "删除设备属性成功！"
}
```
4. 详情（查看）
```
url: /api/device-type/:id/:propertyid
method: GET
返回值：
{
    "code": 0,
    "data": deviceproperty,
    "message": "查找设备属性成功！"
}
```
5. 获取所有数据
```
url: /api/device-type/:id
method: GET
返回值：
{
    "code": 0,
    "data":[
            deviceProperty,
            deviceProperty,
    ],
    "message": "查找设备类型成功！"
}
```
7. 不提供分页查询
### （三）设备清单
```
实体：
device:
{
    "id": 0,
    "name": "电表01",
    "type": {
        "id": 0,
        "name": "DDST422",
        "driver": "DDST422",
        "properties": [
            {
                "name": "电压",
                "description": "电表电压",
                "type": "double",
                "length": 13,
                "decimal": 2,
                "unit": "V",
                "va;ue": null
            },
            {
                "name": "电流",
                "description": "电表电流",
                "type": "double",
                "length": 13,
                "decimal": 2,
                "unit": "A",
                "va;ue": null
            }
        ]    
    },
    "adress": "01",
    "collector":{
        "id": 0,
        "name": "485-1",
        "type": "serial",
        "serial": {
        "name": "本地串口",
        "deviceName": "/dev/ttymxc6",
        "baudRate": 9600,
        "dataBit": 8,
        "stopBit": 1,
        "check": "N"
        },
        "TcpClient":null,
        "timeout": 3000,
        "interval": 200
    },
    "alone": true,
    "serial":{
        "name": "本地串口",
        "deviceName": "/dev/ttymxc6",
        "baudRate": 1200,
        "dataBit": 8,
        "stopBit": 1,
        "check": "N"
    },
    "collectTime": null,
    "collectTotal": 1,
    "collectSuccess": 1,
    "reportTime": null,
    "reportTotal": 1,
    "reportSuccess": 1
}
```
1. 新增
```
url: /api/device
method: POST
返回值:
{
    "code": 0,
    "data": device,
    "message": "新增设备列表成功！"
}
```
2. 编辑
```
url: /api/device
method: PUT
返回值:
{
    "code": 0,
    "data": device,
    "message": "修改设备列表成功！"
}
```
3. 删除
```
url: /api/device/1
method: DELETE
返回值:
{
    "code": 0,
    "data": null,
    "message": "删除设备列表成功！"
}
```
4. 详情（查看）
```
url: /api/device/:id
method: GET
返回值:
{
    "code": 0,
    "data": device,
    "message": "查找设备列表成功！"
}
```
5. 获取所有数据
```
url: /api/devices
method: GET
返回值:
{
    "code": 0,
    "data": [device,device]
    "message": "获取所有设备列表成功！"
}
```
6. 不提供分页查询
### （四）采集任务
```
实体：
collecttask:
{
    "id": 0,
    "name": "电表采集",
    "collector":
    {
    "id": 0,
    "name": "485-1",
    "type": "serial",
    "serial": {
      "name": "本地串口",
      "deviceName": "/dev/ttymxc6",
      "baudRate": 9600,
      "dataBit": 8,
      "stopBit": 1,
      "check": "N"
    },
    "TcpClient":null,
    "timeout": 3000,
    "interval": 200
},
    "cron": "定时策略",
    "status": 0
}
```
1. 新增
```
url: /api/device-task
method: POST
返回值:
{
    "code": 0,
    "data": collecttask,
    "message": "新增采集任务成功！"
}
```
2. 编辑
```
url: /api/collect-task
method: PUT
返回值:
{
    "code": 0,
    "data": {
        collecttask,
    },
    "message": "修改采集任务成功！"
}
```
3. 删除
```
url: /api/collect-task/:id
method: DELETE
返回值:
{
    "code": 0,
    "data": null,
    "message": "删除采集任务成功！"
}
```
4. 详情（查看）
```
url: /api/collect-task/:id
method: GET
返回值:
{
    "code": 0,
    "data": {
        collecttask,
    },
    "message": "查找采集任务成功！"
}
```
5. 获取所有数据
```
url: /api/collect-task
method: GET
返回值:
{
    "code": 0,
    "data": [
        collecttask,
        collecttask
    ],
    "message": "获取所有采集任务数据成功！"
}
```
6. 启动（任务）
7. 停止（任务）
8. 不提供分页查询
### （五）上报任务
1. 数据实体
```
    reportTask: {
        "id": 1,
        "name": "上报任务",
        "reportName": "thingsped",
        "ip": "121.23.22.39",
        "port": 1883,
        "username": admin,
        "password": admin,
        "cron": "1 * * * * *",
        "status": 0
    }
```
2. 新增
```
url: /api/report-task
method: post
返回值：
  {
    "code": 0 | 1,
    "data": reportTask,
    "message": ""
  }  
```
2. 编辑
```
url: /api/report-task
method: PUT
data: reportTask
返回值:
{
    "code": 0,
    "data": reportTask,
    "message": "修改上报服务成功！"
}
```
3. 删除
```
url: /api/report-task/id
method: DELETE
返回值:
{
    "code": 0,
    "data": null,
    "message": "删除上报服务成功！"
}
```
4. 详情（查看）
```
url: /api/report-task/id
method: GET
返回值:
{
    "code": 0,
    "data": reportTask,
    "message": "查找上报服务成功！"
}
```
5. 获取所有数据
```
url: /api/report-tasks
method: GET
返回值:
{
    "code": 0,
    "data": [
        reportTask,
        reportTask
    ],
    "message": "获取所有上报服务数据成功！"
}
```
6. 启动（任务）
```
url: /api/report-task/start/:id
method: GET
返回值:
{
    "code": 0 | 1,
    "data": reportTask,
    "message": "启动任务成功！" | "启动任务失败！"
}
```
7. 停止（任务）
```
url: /api/report-task/stop/:id
method: GET
返回值:
{
    "code": 0 | 1,
    "data": reportTask,
    "message": "停止任务成功！" | "停止任务失败！"
}
```
8. 不提供分页查询
### （五）调试
```
url: /api/debug
method: post
请求：{
    "collectorName": "485-1"
    "data": "37 48 23 44 a3 f4 00 85"
}
响应:{
    "code": 0 | 1
    "data": “49 43 43 ..."
    "message": "成功" | "超时"
}
```

### （六）升级
```
url: /api/debug/grade
上传文件：FileName
返回值:
{
    "code": 0,
    "data": nil,
    "message": "升级成功！"
}
```

### （七）首页
```
1、设备状态
url: /api/debug/system-status
method: POST
返回值:
{
	MemTotal:     "0",                      磁盘总量
	MemUse:       "0",                      磁盘使用率
	DiskTotal:    "0",                      内存总量
	DiskUse:      "0",                      内存使用率
	Name:         "smartgw",              网关名称
	SN:           "22005260001",            产品序列号
	HardVer:      "smartgw-V.A",          硬件版本
	SoftVer:      "V0.9.2",                 软件版本
	SystemRTC:    "2022-06-2 12:00:00",     系统时间
	RunTime:      "0",                      设备运行时间
	DeviceOnline: "0",                      设备在线率
}

2、重启
url: /api/debug/system-reboot
method: POST
返回值：
{
    "code": 0,
    "message": "",
    "data": "",
}
3、校时
url: /api/debug/system-ntp
method: POST
返回值：
{
    "code": 0,
    "message": "2006-01-02 15:04:05",
    "data": "",
}
```
### （八） 网关网口信息
```
1、获取所有网络信息
url: /api/ethernets
method: GET
返回值：
{
    "code": 0,
    "message": "获取所有网口成功！",
    “data”: {
        "name": "eth0",
        "index": 1,
        "MTU": 1500,
        "MAC": "A0A0A0A0A0A0",
        "flags": "up",
        "ip": "192.168.3.31",
        "netmask": "255.255.255.0",
        "gateway": "192.168.3.1",
        "configEnabled": false,
        "DHCPEnabled": false,
        "configIP": "192.168.3.31",
        "configNetMask": "255.255.255.0",
        "configGateway": "192.168.3.1"
    }
}

2、编辑网络信息
url: /api/ethernet
method: PUT
返回值：
{
    “code”: 0,
    "message": "修改网口成功！",
    "data": 同上
}

3、从数据库中删除某网络信息
url: /ethernet/:name
method: DELETE
返回值：
{
    “code”: 0,
    "message": "删除网口成功！",
    "data": nil
}

4、获取某一网络信息
url: /ethernet/:name
method: GET
返回值：
{
    “code”: 0,
    "message": "查找网口成功！",
    "data": 同上
}
    
```
### （九）开启看门狗
```
url: /debug/watchdogstart
method: POST
返回值：
{
    "code": 0 | 1,
    "message": "开启看门狗成功 | 失败",
    "data": "",
}
```

### （十）关闭看门狗
```
url: /debug/watchdogstop
method: POST
{
    "code": 0 | 1,
    "message": "关闭看门狗成功 | 失败",
    "data": "",
}
```

## 二、北向接口
### （一）数据上报

### （二）命令下达

