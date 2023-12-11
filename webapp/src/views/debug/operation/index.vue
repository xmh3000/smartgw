<template>
  <div class="container">
    <el-card class="box-card">
      <div slot="header">
        操作说明
      </div>
      <el-form label-position="left" label-width="120px">
        <!--        <el-form-item label="label1:">-->
        <!--          {{ data.name }}-->
        <!--        </el-form-item>-->
        <div>
          <h2>一、基础操作</h2>
          <h3>1、一般操作</h3>
          <p>首先设置采集接口</p>
          <p>然后设置设备类型</p>
          <p>设备列表一般不在这里维护，由运维平台下发设备列表</p>
          <h3>2、操作依赖</h3>
          <p>在进行基础数据（采集接口、设备类型）设置时，请关闭采集任务、上报任务</p>
          <p>下发设备列表操作，请关闭采集任务、上报任务</p>
          <h2>二、特殊操作</h2>
          <h3>1. 一个485口接入不同波特率的设备。</h3>
          <p>首先，在配置文件config.yml中设置【openeverytime: true】</p>
          <p>然后在串口设置里，设置波特率最多的设备</p>
          <p>采集周期估算，2.5 * n 秒，n 表示每个485口最多串联设备数，例如：2个485口，1个串接12个表具，1个串接30个表具，则n为30</p>
          <h3>2. 采集队列</h3>
          <p>一个485串口最多连接设备不超过128，因为采集队列大小为128</p>
          <p>另外，采集周期过短，也会造成采集队列已满错误提示（通过smartgw.log日志文件），出现这个错误，可以适当调整采集周期大小</p>
          <p>采集周期估算，2.5 * n 秒，n 表示每个485口最多串联设备数，例如：2个485口，1个串接12个表具，1个串接30个表具，则n为30</p>
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
