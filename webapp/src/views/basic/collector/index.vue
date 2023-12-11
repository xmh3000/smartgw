<template>
  <div class="container">
    <el-button style="margin-bottom: 20px" @click="add()" type="primary" icon="el-icon-plus" size="small">新增</el-button>
    <el-table
        :data="data"
        max-height="700"
        border
        style="width: 100%">
      <el-table-column prop="name" label="接口名称" width="180"></el-table-column>
      <el-table-column prop="type" label="接口类型" width="180">
        <template slot-scope="scope">
          {{ scope.row.type === 'Serial' ? '串口' : '网口' }}
        </template>
      </el-table-column>
      <el-table-column width="500" label="参数">
        <template slot-scope="scope">
          <span v-if="scope.row.type === 'Serial' ">
            {{
              `物理接口：${scope.row.serial.deviceName} 波特率：${scope.row.serial.baudRate} 数据位：${scope.row.serial.dataBit} 停止位：${scope.row.serial.stopBit} 校验位：${scope.row.serial.check}`
            }}
          </span>
          <span v-else>
            {{ `目标IP：${scope.row.tcpClient.ip} 端口：${scope.row.tcpClient.port}` }}
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="timeout" label="超时时间"></el-table-column>
      <el-table-column prop="interval" label="通讯间隔"></el-table-column>
      <el-table-column label="操作" width="180">
        <template slot-scope="scope">
          <el-tag class="rightMargin" @click="handleClick(scope.row)" style="cursor: pointer" type="warning"
                  size="medium">查看
          </el-tag>
          <el-tag class="rightMargin" @click="edit(scope.row)" type="success" style="cursor: pointer" size="medium">编辑
          </el-tag>
          <el-popconfirm
              title="确定删除吗？"
              @onConfirm="cancel(scope.row)"
          >
            <el-tag type="danger" style="cursor: pointer" slot="reference" size="medium">删除</el-tag>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
        :title="title"
        :visible.sync="dialogVisible"
        width="30%"
        :before-close="handleClose">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="80px">
        <el-form-item label="接口名称" prop="name">
          <el-input :disabled="title === '编辑' || this.title === '查看'" v-model="ruleForm.name" placeholder="电表、水表、TCP1、TCP2、TCP3..."></el-input>
        </el-form-item>
        <el-form-item label="接口类型" prop="type">
          <el-select :disabled="this.title === '查看'" v-model="ruleForm.type" placeholder="">
            <el-option label="串口" value="Serial"></el-option>
            <el-option label="网口" value="TcpClient"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="ruleForm.type === 'TcpClient'" label="名称" prop="tcpClient.name">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.tcpClient.name"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.type === 'TcpClient'" label="Ip地址" prop="tcpClient.ip">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.tcpClient.ip"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.type === 'TcpClient'" label="端口号" prop="tcpClient.port">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.tcpClient.port"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.type === 'Serial'" label="串口名称" prop="serial.name">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.name" placeholder="电表、水表、TCP1、TCP2、TCP3..."></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.type === 'Serial'" label="设备名称" prop="serial.deviceName">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.deviceName"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.type === 'Serial'" label="波特率" prop="serial.baudRate">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.baudRate"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.type === 'Serial'" label="数据位" prop="serial.dataBit">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.dataBit"></el-input>
        </el-form-item>
        <el-form-item :disabled="this.title === '查看'" v-if="ruleForm.type === 'Serial'" label="停止位"
                      prop="serial.stopBit">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.stopBit"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.type === 'Serial'" label="检验" prop="serial.check">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.check" placeholder="N:无,E:偶校验:O:奇校验"></el-input>
        </el-form-item>
        <el-form-item label="超时时间" prop="timeout">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.timeout">
            <template slot="append">毫秒</template>
          </el-input>
        </el-form-item>
        <el-form-item label="通讯间隔" prop="interval">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.interval">
            <template slot="append">毫秒</template>
          </el-input>
        </el-form-item>
      </el-form>
      <span v-if="this.title !== '查看'" slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="handleOk">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {addCollector, updateCollector, deleteCollector, findCollectors} from '@/api/collector'

export default {
  data() {
    return {
      data: [],
      dialogVisible: false,
      title: '',
      ruleForm: {
        name: null,
        type: null,
        serial: {
          name: null,
          deviceName: null,
          baudRate: null,
          dataBit: null,
          stopBit: null,
          check: null,
        },
        tcpClient: {
          name: null,
          ip: null,
          port: null,
        },
        timeout: null,
        interval: null,
      },
      rules: {
        name: [{required: true, message: '请输入接口名称', trigger: 'blur'},],
        type: [{required: true, message: '选择接口类型', trigger: 'blur'},],
        tcpClient: {
          ip: [{
            required: true,
            message: '请输入Ip地址',
            trigger: 'blur'
          }, {
            // pattern: /^([0-9].|[1-9][0-9].|1[0-9]{1,2}.|2[0-4][0-9]|25[0-5].){3}([0-9]|[1-9][0-9]|1[0-9]{1,2}|2[0-4][0-9]|25[0-5])$/,
            pattern: /^((\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])$/,
            message: 'Ip地址格式错误'
          }],
          port: [{required: true, message: '请输入端口号', trigger: 'blur'}, {pattern: /^\d{1,4}$/, message: '端口号格式错误'}],
          name: [{required: true, message: '请输入名称', trigger: 'blur'},],
        },
        serial: {
          name: [{required: true, message: '请输入串口名称', trigger: 'blur'},],
          deviceName: [{required: true, message: '请输入设备名称', trigger: 'blur'},],
          baudRate: [{required: true, message: '请输入波特率', trigger: 'blur'},],
          dataBit: [{required: true, message: '请输入数据位', trigger: 'blur'},],
          stopBit: [{required: true, message: '请输入停止位', trigger: 'blur'},],
          check: [{required: true, message: '请输入检验', trigger: 'blur'},],
        },

        timeout: [{required: true, message: '请输入超时时间', trigger: 'blur'}, {pattern: /^[0-9]*$/, message: '超时时间必须为数字值'}],
        interval: [{required: true, message: '请输入通讯间隔', trigger: 'blur'}, {pattern: /^[0-9]*$/, message: '通讯间隔必须为数字值'}],
      }
    }
  },
  methods: {
    handleClick(row) {
      this.title = '查看'
      this.ruleForm = JSON.parse(JSON.stringify(row))
      this.dialogVisible = true
    },
    // 编辑
    edit(row) {
      this.title = '编辑'
      this.ruleForm = JSON.parse(JSON.stringify(row))
      this.dialogVisible = true

    },
    // 新增
    add() {
      this.ruleForm = {
        name: null,
        type: null,
        serial: {
          name: null,
          deviceName: '/dev/ttymxc3',
          baudRate: 9600,
          dataBit: 8,
          stopBit: "1",
          check: 'N',
        },
        tcpClient: {
          name: null,
          ip: "192.168.100.21",
          port: 7000,
        },
        timeout: 3000,
        interval: 200,
      }
      this.title = '新增'
      this.dialogVisible = true
    },
    // 删除
    cancel(row) {
      deleteCollector(row.name).then(
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
      findCollectors().then(
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
    //保存（新增/编辑）
    handleOk() {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          this.ruleForm.timeout = Number(this.ruleForm.timeout)
          this.ruleForm.interval = Number(this.ruleForm.interval)
          if (this.ruleForm.type === 'Serial') {
            this.ruleForm.serial.baudRate = Number(this.ruleForm.serial.baudRate)
            this.ruleForm.serial.dataBit = Number(this.ruleForm.serial.dataBit)
            // this.ruleForm.serial.stopBit = Number(this.ruleForm.serial.stopBit)
          } else {
            this.ruleForm.tcpClient.port = Number(this.ruleForm.tcpClient.port)
          }
          if (this.title === '新增') {
            addCollector(this.ruleForm).then(
                res => {
                  this.retrieve()
                  this.$message.info(res.message)
                  this.$refs['ruleForm'].resetFields();
                  this.dialogVisible = false
                }
            ).catch(
                err => {
                  this.$message.error(err.data.message)
                  this.$refs['ruleForm'].resetFields();
                  this.dialogVisible = false
                }
            )

          } else if (this.title === '编辑') {
            updateCollector(this.ruleForm).then(
                res => {
                  this.retrieve()
                  this.$message.info(res.message)
                  this.$refs['ruleForm'].resetFields();
                  this.dialogVisible = false
                }
            ).catch(
                err => {
                  this.$message.error('编辑失败')
                  this.$refs['ruleForm'].resetFields();
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
  }
}
</script>

<style lang="scss" scoped>
.rightMargin {
  margin-right: 10px;
}
</style>
