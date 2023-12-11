<template>
  <div class="container">
    <el-card class="box-card">
      <div slot="header">
        更新日志
      </div>
      <el-form label-position="left" label-width="120px">
        <div>
          <h4>版本：V 1.0.7 (2023-02-08)</h4>
          <ul>
            <li> 新增
              <ol>
                <li>【设备列表】新增导入导出功能</li>
                <li>【设备类型->属性】新增导出功能</li>
                <li>设备驱动压缩包中如果包含properties.json文件，则自动导入到属性列表中</li>
              </ol>
            </li>
            <li> 接口
              <ol>
                <li>新增GetDevices2接口</li>
                <li>新增SetDevices2接口</li>
                <li>和版本1相比，版本2新增是否独立设置波特率等属性</li>
              </ol>
            </li>
            <li> 优化
              <ol>
                <li>【设备类型->属性】数据类型支持double,int,uint,string四种类型</li>
              </ol>
            </li>
          </ul>
        </div>
        <div>
          <h4>版本：V 1.0.6 (2023-02-02)</h4>
          <ul>
            <li> 优化
              <ol>
                <li>调整ntp服务调用顺序</li>
              </ol>
            </li>
          </ul>
        </div>
        <div>
          <h4>版本：V 1.0.5 (2023-02-01)</h4>
          <ul>
            <li> 新增
              <ol>
                <li>新增【更新日志】模块</li>
              </ol>
            </li>
            <li> 修复
              <ol>
                <li>修改了采集超时有可能破坏上次采集数据的错误</li>
              </ol>
            </li>
          </ul>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script>

import {getToken} from "@/utils/auth";

export default {
  data() {
    return {
      token: getToken(),
      data: {
        name: '测试'
      }
    }
  },
  methods: {}
}
</script>

<style lang="scss" scoped>

</style>
