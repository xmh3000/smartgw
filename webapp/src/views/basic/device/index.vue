<template>
  <div class="container">
    <el-form :model="searchFrom" :inline="true" label-width="100px">
      <el-form-item label="设备类型：">
        <el-select v-model="searchFrom.deviceType" clearable placeholder="请选择设备类型" style="width: 150px">
          <el-option v-for="(item,index) in dataDeviceType" :key="index" :value="item.name">{{ item.name }}</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="采集接口：">
        <el-select v-model="searchFrom.collectorName" clearable placeholder="请选择采集接口" style="width: 150px">
          <el-option v-for="(item,index) in dataCollector" :key="index" :value="item.name">{{ item.name }}</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="独立开关：">
        <el-select v-model="searchFrom.alone" clearable placeholder="请选择" style="width: 150px">
          <el-option label="是" :value="1"></el-option>
          <el-option label="否" :value="0"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="在线状态：">
        <el-select v-model="searchFrom.online" clearable placeholder="请选择在线状态" style="width: 150px">
          <el-option label="在线" :value="1"></el-option>
          <el-option label="离线" :value="0"></el-option>
        </el-select>
      </el-form-item>
      <el-button type="primary" @click="filterTable">搜索</el-button>
      <el-button @click="resetSearch">清空条件</el-button>
    </el-form>
    <el-button style="margin-bottom: 20px;margin-right: 10px" @click="add()" type="primary" icon="el-icon-plus" size="small">新增
    </el-button>
    <el-upload style="display: inline-block; margin-right: 10px" :action="`/api/devices/import`" :headers="{token: token}" :show-file-list="false"
               name="FileName" :on-success="handleSuccess" accept=".json">
      <el-button style="margin-bottom: 20px;" icon="el-icon-upload" size="small">导入</el-button>
    </el-upload>
    <el-button style="margin-bottom: 20px;" @click="exportDevice()" icon="el-icon-download" size="small">导出</el-button>
    <el-table
      :data="data"
      max-height="700"
      border
      style="width: 100%">
      <el-table-column prop="name" label="设备名称">
      </el-table-column>
      <el-table-column prop="type.name" label="设备类型">
        <template slot-scope="scope">
          <span class="sp1" @click="lookRow(1,scope.row.type)">{{ scope.row.type.name }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="address" label="通讯地址"></el-table-column>
      <el-table-column prop="collector.name" label="采集接口">
        <template slot-scope="scope">
          <span class="sp1" @click="lookRow(2,scope.row.collector)">{{ scope.row.collector.name }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="alone" label="独立开关">
        <template slot-scope="scope">
          {{ scope.row.alone ? '是' : '否' }}
        </template>
      </el-table-column>
      <el-table-column width="500" label="通讯参数">
        <template slot-scope="scope">
          <span class="sp2"
                :title="`物理接口：${scope.row.serial.deviceName} 波特率：${scope.row.serial.baudRate} 数据位：${scope.row.serial.dataBit} 停止位：${scope.row.serial.stopBit} 校验位：${scope.row.serial.check}`"
                v-if="scope.row.alone">
            {{
              `物理接口：${scope.row.serial.deviceName} 波特率：${scope.row.serial.baudRate} 数据位：${scope.row.serial.dataBit} 停止位：${scope.row.serial.stopBit} 校验位：${scope.row.serial.check}`
            }}
          </span>
          <span
            :title="scope.row.collector.type === 'Serial' ? `物理接口：${scope.row.collector.serial.deviceName} 波特率：${scope.row.collector.serial.baudRate} 数据位：${scope.row.collector.serial.dataBit} 停止位：${scope.row.collector.serial.stopBit} 校验位：${scope.row.collector.serial.check}` : `目标IP：${scope.row.collector.tcpClient.ip} 端口：${scope.row.collector.tcpClient.port}`"
            class="sp2" v-else>
            {{
              scope.row.collector.type === 'Serial' ? `物理接口：${scope.row.collector.serial.deviceName} 波特率：${scope.row.collector.serial.baudRate} 数据位：${scope.row.collector.serial.dataBit} 停止位：${scope.row.collector.serial.stopBit} 校验位：${scope.row.collector.serial.check}` : `目标IP：${scope.row.collector.tcpClient.ip} 端口：${scope.row.collector.tcpClient.port}`
            }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="状态" width="70">
        <template slot-scope="scope">
          <span style="color: green" v-if="scope.row.online">在线</span>
          <span style="color: red" v-else>离线</span>
        </template>
      </el-table-column>
      <el-table-column width="180" prop="collectTime" label="最后采集时间">
        <template slot-scope="scope">
          {{ moment(scope.row.collectTime).format('yyyy-MM-DD hh:mm:ss SSS') }}
        </template>
      </el-table-column>
      <el-table-column prop="collectTotal" label="采集次数"></el-table-column>
      <el-table-column prop="collectSuccess" label="成功次数"></el-table-column>
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
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="120px">
        <el-form-item label="设备名称" prop="name">
          <el-input :disabled="title === '编辑' || this.title === '查看'" v-model="ruleForm.name"></el-input>
        </el-form-item>
        <el-form-item label="设备类型" prop="type.name">
          <el-select :disabled="this.title === '查看'" v-model="ruleForm.type.name" placeholder="">
            <el-option v-for="item in dataDeviceType" :key="item.name" :label="item.name"
                       :value="item.name"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="通讯地址" prop="address">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.address"></el-input>
        </el-form-item>
        <el-form-item label="采集接口" prop="collector.name">
          <el-select :disabled="this.title === '查看'" v-model="ruleForm.collector.name" placeholder="">
            <el-option v-for="item in dataCollector" :key="item.name" :label="item.name" :value="item.name"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="独立开关" prop="alone">
          <el-switch
            :disabled="this.title === '查看'"
            v-model="ruleForm.alone">
          </el-switch>
        </el-form-item>
        <el-form-item v-if="ruleForm.alone" label="串口名称" prop="serial.name">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.name"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.alone" label="设备名称" prop="serial.deviceName">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.deviceName"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.alone" label="波特率" prop="serial.baudRate">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.baudRate"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.alone" label="数据位" prop="serial.dataBit">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.dataBit"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.alone" :disabled="this.title === '查看'" label="停止位"
                      prop="serial.stopBit">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.stopBit"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.alone" label="检验" prop="serial.check">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.serial.check"></el-input>
        </el-form-item>
      </el-form>
      <span v-if="this.title !== '查看'" slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="handleOk">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog title="设备类型" :visible.sync="dialogVisible2" width="40%" :before-close="handleClose2">
      <el-form ref="form1" :model="row1" label-width="120px">
        <el-form-item label="类型名称">
          <el-input disabled v-model="row1.name"></el-input>
        </el-form-item>
        <el-form-item label="驱动程序目录">
          <el-input disabled v-model="row1.driver"></el-input>
        </el-form-item>
        <div>
          <el-table :data="row1.properties" max-height="300" border style="width: 100%">
            <el-table-column prop="name" label="属性名称" width="160"></el-table-column>
            <el-table-column prop="description" label="描述"></el-table-column>
            <el-table-column prop="type" label="数据类型"></el-table-column>
            <el-table-column prop="length" label="数据长度"></el-table-column>
            <el-table-column prop="decimal" label="小数位数"></el-table-column>
            <el-table-column prop="unit" label="计量单位"></el-table-column>
            <el-table-column prop="value" label="最新数值"></el-table-column>
            <el-table-column prop="reported" label="是否上报">
              <template slot-scope="scope">
                <el-checkbox disabled v-model="scope.row.reported"></el-checkbox>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-form>
    </el-dialog>
    <el-dialog title="采集接口" :visible.sync="dialogVisible3" width="30%" :before-close="handleClose2">
      <el-form ref="form2" :model="row2" label-width="80px">
        <el-form-item label="接口名称">
          <el-input disabled v-model="row2.name"></el-input>
        </el-form-item>
        <el-form-item label="接口类型">
          <el-input disabled v-model="row2.type==='Serial'?'串口':'Tcp客户端'"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type==='Serial'" label="串口名称">
          <el-input disabled v-model="row2.serial.name"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type==='Serial'" label="设备名称">
          <el-input disabled v-model="row2.serial.deviceName"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type==='Serial'" label="波特率">
          <el-input disabled v-model="row2.serial.baudRate"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type==='Serial'" label="数据位">
          <el-input disabled v-model="row2.serial.dataBit"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type==='Serial'" label="停止位">
          <el-input disabled v-model="row2.serial.stopBit"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type==='Serial'" label="检验">
          <el-input disabled v-model="row2.serial.check"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type === 'TcpClient'" label="名称">
          <el-input disabled v-model="row2.tcpClient.name"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type === 'TcpClient'" label="Ip地址">
          <el-input disabled v-model="row2.tcpClient.ip"></el-input>
        </el-form-item>
        <el-form-item v-if="row2.type ==='TcpClient'" label="端口号">
          <el-input disabled v-model="row2.tcpClient.port"></el-input>
        </el-form-item>
        <el-form-item label="超时时间">
          <el-input disabled v-model="row2.timeout"></el-input>
        </el-form-item>
        <el-form-item label="通讯间隔">
          <el-input disabled v-model="row2.timeout"></el-input>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import {addDevice, deleteDevice, exportDevices, findDevices, updateDevice} from '@/api/device'
import {findDeviceTypes} from "@/api/device-type"
import {findCollectors} from '@/api/collector'
import moment from 'moment'
import {exportJson} from "@/utils/export.js"
import {getToken} from '@/utils/auth'

export default {
  data() {
    return {
      token: getToken(),
      moment: moment,
      data: [],
      dataDeviceType: [],
      dataCollector: [],
      dialogVisible: false,
      dialogVisible2: false,
      dialogVisible3: false,
      title: '',
      row1: {},
      row2: {},
      ruleForm: {
        name: null,
        type: {
          name: null,
        },
        alone: false,
        collector: {
          name: null,
        },
        address: null,
        serial: {
          name: null,
          deviceName: null,
          baudRate: null,
          dataBit: null,
          stopBit: null,
          check: null,
        },
        collectTime: null,
        collectTotal: 0,
        collectSuccess: 0,
        reportTime: null,
        reportTotal: 0,
        reportSuccess: 0,
      },
      rules: {
        name: [{required: true, message: '请输入接口名称', trigger: 'blur'},],
        type: {name: [{required: true, message: '选择设备类型', trigger: 'blur'},]},
        address: [{required: true, message: '请输入通讯地址', trigger: 'blur'},],
        collector: {name: [{required: true, message: '选择采集接口', trigger: 'blur'},]},
        serial: {
          name: [{required: true, message: '请输入串口名称', trigger: 'blur'},],
          deviceName: [{required: true, message: '请输入设备名称', trigger: 'blur'},],
          baudRate: [{required: true, message: '请输入波特率', trigger: 'blur'}, , {
            pattern: /^[0-9]*$/,
            message: '波特率必须为数字值'
          }],
          dataBit: [{required: true, message: '请输入数据位', trigger: 'blur'}, , {
            pattern: /^[0-9]*$/,
            message: '数据位必须为数字值'
          }],
          stopBit: [{required: true, message: '请输入停止位', trigger: 'blur'}, , {
            pattern: /^[0-9]*$/,
            message: '停止位必须为数字值'
          }],
          check: [{required: true, message: '请输入检验', trigger: 'blur'},],
        },
      },
      searchFrom: {
        deviceType: '',
        collectorName: '',
        alone: '',
        online: ''
      },
      originTableData: []
    }
  },
  methods: {
    handleSuccess(res) {
      this.$message.info(res.message)
      this.retrieve()
    },
    lookRow(index, row) {
      if (index === 1) {
        this.row1 = JSON.parse(JSON.stringify(row))
        this.dialogVisible2 = true

      } else {
        this.row2 = JSON.parse(JSON.stringify(row))
        this.dialogVisible3 = true
      }

    },
    handleClick(row) {
      this.title = '查看'
      this.ruleForm = {...row}
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
        type: {
          name: null,
        },
        alone: false,
        collector: {
          name: null,
        },
        address: null,
        serial: {
          name: null,
          deviceName: '/dev/ttymxc3',
          baudRate: 9600,
          dataBit: 8,
          stopBit: "1",
          check: 'N',
        },
        online: false,
        collectTime: null,
        collectTotal: 0,
        collectSuccess: 0,
        reportTime: null,
        reportTotal: 0,
        reportSuccess: 0,
      }
      this.title = '新增'
      this.dialogVisible = true
    },
    // 删除
    cancel(row) {
      deleteDevice(row.name).then(
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
      findDevices().then(
        res => {
          res.data.forEach((item, index) => {
            res.data[index]['deviceType'] = item.type.name;
            res.data[index]['collectorName'] = item.collector.name;
          })
          this.data = res.data;
          this.originTableData = JSON.parse(JSON.stringify(res.data));
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
      this.dialogVisible3 = false
    },
    //保存（新增/编辑）
    handleOk() {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          this.ruleForm.collector = this.dataCollector.filter(item => item.name === this.ruleForm.collector.name)[0]
          this.ruleForm.type = this.dataDeviceType.filter(item => item.name === this.ruleForm.type.name)[0]
          if (this.ruleForm.alone) {
            this.ruleForm.serial.baudRate = Number(this.ruleForm.serial.baudRate)
            this.ruleForm.serial.dataBit = Number(this.ruleForm.serial.dataBit)
            // this.ruleForm.serial.stopBit = Number(this.ruleForm.serial.stopBit)
          }
          if (this.title === '新增') {
            addDevice(this.ruleForm).then(
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
            updateDevice(this.ruleForm).then(
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
    },

    filterTable() {
      let result = this.originTableData;
      Object.keys(this.searchFrom).forEach(key => {
        if (this.searchFrom[key] != "") {
          result = this.filterFun(result, key, this.searchFrom[key]);
        }
      })
      this.data = result;
    },

    filterFun(arr, key, value) {
      let result = [];
      result = arr.filter(item => {
        return item[key] == value
      })
      return result;
    },

    resetSearch() {
      this.searchFrom.alone = "";
      this.searchFrom.collectorName = "";
      this.searchFrom.deviceType = "";
      this.searchFrom.online = "";
      this.data = this.originTableData;
    },
    exportDevice() {
      exportDevices().then(
        res => {
          const data = res.data
          if (res.code == 0) {
            exportJson(JSON.stringify(data, null, '\t'), 'devices')
            this.retrieve()
          }
          this.$message.success(res.message)
        }
      ).catch(
        err => {
          this.$message.error('导出数据失败！')
        }
      )
    },
  },

  created() {
    findDeviceTypes().then(
      res => {
        this.dataDeviceType = res.data
      }
    )
    findCollectors().then(
      res => {
        this.dataCollector = res.data
      }
    )
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

.sp1 {
  color: #20a0ff;
  cursor: pointer;
}

.sp2 {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
