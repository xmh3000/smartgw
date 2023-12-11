<template>
  <div class="container">
    <div class="box">
      <div>
        <el-tooltip>
          <div slot="content">内存总量<br/>{{ data.MemTotal }}M</div>
          <div class="card">
            <div class="left">
              <div class="img-wrapper">
                <div class="img img0"></div>
              </div>
            </div>
            <div class="right">
              <div class="name">内存使用率</div>
              <div class="rate">{{ data.MemUse }}%</div>
            </div>
          </div>
        </el-tooltip>
      </div>
      <div>
        <el-tooltip>
          <div slot="content">磁盘总量<br/>{{ data.DiskTotal }}M</div>
          <div class="card">
            <div class="left">
              <div class="img-wrapper">
                <div class="img img1"></div>
              </div>
            </div>
            <div class="right">
              <div class="name">磁盘使用率</div>
              <div class="rate">{{ data.DiskUse }}%</div>
            </div>
          </div>
        </el-tooltip>
      </div>
      <div>
        <el-tooltip>
          <div slot="content">设备总数<br/>{{ data.DeviceTotal }}个</div>
          <div class="card">
            <div class="left">
              <div class="img-wrapper">
                <div class="img img2"></div>
              </div>
            </div>
            <div class="right">
              <div class="name">设备在线率</div>
              <div class="rate">{{ data.DeviceOnline }}%</div>
            </div>
          </div>
        </el-tooltip>
      </div>
    </div>
    <div class="content">
      <el-card class="box-card">
        <div slot="header" style="display: flex;justify-content: space-between;align-items: center">
          <span>网关信息</span>
          <div>
            <el-popconfirm
              title="确定重启吗？"
              @onConfirm="restart()"
            >
              <el-button type="primary" slot="reference">重启</el-button>
            </el-popconfirm>
          </div>
        </div>
        <el-form label-position="left" label-width="120px">
          <el-form-item label="网关编号:">
            {{ data.Name }}
          </el-form-item>
          <el-form-item label="硬件版本:">
            {{ data.HardVer }}
          </el-form-item>
          <el-form-item label="软件版本:">
            {{ data.SoftVer }}
          </el-form-item>
          <el-form-item label="系统时间:">
            {{ data.SystemRTC }}
            <el-tag class="leftMargin" @click="timing()" style="cursor: pointer" size="medium">校时
            </el-tag>
          </el-form-item>
          <el-form-item label="系统运行:">
            {{ data.RunTime }}
          </el-form-item>
          <el-form-item label="串口功能:">
            {{ data.OpenEveryTime }}
          </el-form-item>
          <el-form-item label="软件更新:">
            {{ data.SN }}
          </el-form-item>
        </el-form>
<!--        <div class="content-1">-->
<!--          <el-popconfirm-->
<!--            title="确定重启吗？"-->
<!--            @onConfirm="restart()"-->
<!--          >-->
<!--            <el-button type="primary" slot="reference">重启</el-button>-->
<!--          </el-popconfirm>-->
<!--        </div>-->
      </el-card>
<!--      <div class="box2">-->
<!--                <dashboardLine :data="data"></dashboardLine>-->
<!--      </div>-->
    </div>
  </div>
</template>

<script>
// 引入组件
import dashboardLine from "@/components/dashboard-line";
import {systemNtp, systemReboot, systemStatus} from "@/api/debug"

export default {
  components: {
    dashboardLine,
  },
  data() {
    return {
      data: {}
    }
  },
  methods: {
    restart() {
      systemReboot().then(
        res => {
          this.retrieve()
        }
      )
    },
    timing() {
      systemNtp().then(
        res => {
          this.retrieve()
        }
      )
    },
    retrieve() {
      systemStatus().then(
        res => {
          this.data = res.data
        }
      )
    }
  },
  mounted() {
    this.retrieve()
  }
}
</script>

<style lang="scss" scoped>
.box {
  display: flex;
}

.box > div {
  margin-right: 20px;
}

.card {
  display: flex;
  justify-content: space-between;
  padding: 20px;
  background: #fff;
  border-radius: 10px;
  border: 1px solid #5a5e66;
}

.name {
  line-height: 18px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 16px;
  margin-bottom: 12px;
  font-weight: bold;
}

.rate {
  font-size: 20px;
  font-weight: bold;
}

.img-wrapper {
  width: 80px;
  height: 80px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.img {
  width: 72px;
  height: 72px;
  background-size: 100%;
  background-repeat: no-repeat;
}

.img0 {
  background-image: url("~@/assets/image/memory.png");
}

.img1 {
  background-image: url("~@/assets/image/disk.png");
}

.img2 {
  background-image: url("~@/assets/image/dashboard.png");
}

.content {
  margin-top: 30px;
  display: flex;
  background-color: #fff;
}

.box-card {
  width: 500px;
}

.box2 {
  flex: 1;
  margin-left: 20px;
}

.content-1 {
  margin-bottom: 20px;
}

.leftMargin {
  margin-left: 10px;
}
</style>
