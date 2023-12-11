<template>
  <div class="container">
    <el-button disabled style="margin-bottom: 20px" type="primary" icon="el-icon-plus" size="small" @click="add()">新增
    </el-button>
    <el-table :data="data" max-height="700" border style="width: 100%">
      <el-table-column prop="name" label="任务名称"/>
<!--      <el-table-column prop="collector.name" label="采集接口">-->
<!--        <template slot-scope="scope">-->
<!--          <span class="sp1" @click="lookRow(scope.row)">{{ scope.row.collector.name }}</span>-->
<!--        </template>-->
<!--      </el-table-column>-->
      <el-table-column prop="cron" label="定时策略"/>
      <el-table-column prop="status" label="状态">
        <template slot-scope="scope">
          <span style="display: flex;align-items: center">
            <i v-if="scope.row.status === 1" class="el-icon-video-play"/>
            <i v-else class="el-icon-video-pause"/>{{ scope.row.status === 0 ? '已停止' : '运行中' }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template slot-scope="scope">
          <el-tag style="cursor: pointer" size="medium" class="rightMargin" type="success" @click="edit(scope.row)">编辑
          </el-tag>
          <el-popconfirm title="确定删除吗？" @onConfirm="cancel(scope.row)" :disabled="true">
            <el-tag slot="reference" style="cursor: pointer" size="medium" class="rightMargin" type="danger">删除
            </el-tag>
          </el-popconfirm>
          <el-popconfirm :title="`确定${scope.row.status === 0?'启动':'停止'}吗？`" @onConfirm="handleClick(scope.row)">
            <el-tag slot="reference" style="cursor: pointer" size="medium">{{ scope.row.status === 0 ? '启动' : '停止' }}
            </el-tag>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog :title="title" :visible.sync="dialogVisible" width="30%" :before-close="handleClose">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="80px">
        <el-form-item label="服务名称" prop="name">
          <el-input v-model="ruleForm.name" :disabled="title === '编辑' || title === '查看'"/>
        </el-form-item>
<!--        <el-form-item label="采集接口" prop="collector.name">-->
<!--          <el-select v-model="ruleForm.collector.name" :disabled="title === '查看'" placeholder="">-->
<!--            <el-option v-for="item in dataCollector" :key="item.name" :label="item.name" :value="item.name"/>-->
<!--          </el-select>-->
<!--        </el-form-item>-->
        <el-form-item label="定时策略" prop="cron">
          <el-input v-model="ruleForm.cron"/>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="handleOk">确 定</el-button>
      </span>
    </el-dialog>
<!--    <el-dialog title="采集接口" :visible.sync="dialogVisible2" width="30%" :before-close="handleClose2">-->
<!--      <el-form ref="form" :model="row" label-width="80px">-->
<!--        <el-form-item label="接口名称">-->
<!--          <el-input v-model="row.name" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="接口类型">-->
<!--          <el-input v-model="row.type==='Serial'?'串口':'Tcp客户端'" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type==='Serial'" label="串口名称">-->
<!--          <el-input v-model="row.serial.name" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type==='Serial'" label="设备名称">-->
<!--          <el-input v-model="row.serial.deviceName" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type==='Serial'" label="波特率">-->
<!--          <el-input v-model="row.serial.baudRate" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type==='Serial'" label="数据位">-->
<!--          <el-input v-model="row.serial.dataBit" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type==='Serial'" label="停止位">-->
<!--          <el-input v-model="row.serial.stopBit" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type==='Serial'" label="检验">-->
<!--          <el-input v-model="row.serial.check" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type === 'TcpClient'" label="名称">-->
<!--          <el-input v-model="row.tcpClient.name" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type === 'TcpClient'" label="Ip地址">-->
<!--          <el-input v-model="row.tcpClient.ip" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item v-if="row.type ==='TcpClient'" label="端口号">-->
<!--          <el-input v-model="row.tcpClient.port" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="超时时间">-->
<!--          <el-input v-model="row.timeout" disabled/>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="通讯间隔">-->
<!--          <el-input v-model="row.timeout" disabled/>-->
<!--        </el-form-item>-->
<!--      </el-form>-->
<!--    </el-dialog>-->
  </div>
</template>

<script>
import {
  addCollectTask,
  updateCollectTask,
  deleteCollectTask,
  findCollectTasks,
  startCollectTask,
  stopCollectTask
} from '@/api/collect-task'
// import {findCollectors} from '@/api/collector'

export default {

  data() {
    return {
      data: [],
      // dataCollector: [],
      dialogVisible: false,
      dialogVisible2: false,
      row: {},
      title: '',
      ruleForm: {
        name: null,
        // collector: {
        //   name: null,
        // },
        cron: null,
        status: null,
      },
      rules: {
        name: [{required: true, message: '请输入服务名称', trigger: 'blur'},],
        // collector: {
        //   name: [{required: true, message: '请输入服务名称', trigger: 'blur'},]
        // },
        cron: [{required: true, message: '请输入定时策略', trigger: 'blur'},],
      }
    }
  },
  methods: {
    lookRow(row) {
      // this.row = JSON.parse(JSON.stringify(row.collector))
      this.dialogVisible2 = true
    },
    // 启动/停止
    handleClick(row) {
      if (row.status === 0) {
        startCollectTask(row.name).then(
            res => {
              this.$message.info(res.message)
              this.retrieve()
            }
        )
      } else {
        stopCollectTask(row.name).then(
            res => {
              this.$message.info(res.message)
              this.retrieve()
            }
        )
      }
    },
    // 编辑
    edit(row) {
      this.ruleForm = JSON.parse(JSON.stringify(row))
      this.title = '编辑'
      this.dialogVisible = true
    },
    // 新增
    add() {
      this.ruleForm = {
        name: '数据采集',
        cron: '@every 1m30s',
        status: 0,
        // collector: {
        //   name: null,
        // },
      }
      this.title = '新增'
      this.dialogVisible = true
    },
    // 删除
    cancel(row) {
      deleteCollectTask(row.name).then(
          res => {
            this.$message.info(res.message)
            this.retrieve()
          }
      ).catch(
          err => {
            this.$message.error('删除失败')
          }
      )
    },
    //查询数据
    retrieve() {
      findCollectTasks().then(
          res => {
            this.data = res.data
          }
      )
    },
    // 取消
    handleClose() {
      this.$refs['ruleForm'].resetFields();
      this.dialogVisible = false
    },
    handleClose2() {
      this.dialogVisible2 = false
    },
    //保存（新增/编辑）
    handleOk() {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          this.ruleForm.status = Number(this.ruleForm.status)
          // this.ruleForm.collector = this.dataCollector.filter(item => item.name === this.ruleForm.collector.name)[0]
          if (this.title === '新增') {
            addCollectTask(this.ruleForm).then(
                res => {
                  this.retrieve()
                  this.$message.info(res.message)
                  this.dialogVisible = false
                }
            ).catch(
                err => {
                  this.$message.error(err.data.message)
                  this.dialogVisible = false
                }
            )

          } else if (this.title === '编辑') {
            updateCollectTask(this.ruleForm).then(
                res => {
                  this.retrieve()
                  this.$message.info(res.message)
                  this.dialogVisible = false
                }
            ).catch(
                err => {
                  this.$message.error('编辑失败')
                  this.dialogVisible = false
                }
            )
          }
        }
      });

    }
  },
  mounted() {
    this.retrieve()
    // findCollectors().then(
    //     res => {
    //       this.dataCollector = res.data
    //       console.log(666, this.dataCollector)
    //     }
    // )
  }
}
</script>
<style lang="scss" scoped>
.rightMargin {
  margin-right: 10px;
}

.el-icon-video-play {
  font-size: 18px;
  margin-right: 5px;
  color: green;
}

.el-icon-video-pause {
  font-size: 18px;
  margin-right: 5px;
  color: red;
}

.sp1 {
  color: #20a0ff;
  cursor: pointer;
}
</style>
