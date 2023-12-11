<template>
  <div class="container">
    <div class="box">
      <div>
        <div style="display: flex;justify-content: space-between">
          <div>
            <el-button style="margin-bottom: 20px" @click="add()" type="primary" icon="el-icon-plus" size="small">新增
            </el-button>
          </div>
          <!--          <div>-->
          <!--            <el-button style="margin-bottom: 20px" @click="importType()" type="primary" size="small">导入</el-button>-->
          <!--            <el-button style="margin-bottom: 20px" @click="exportType()" type="primary" size="small">导出</el-button>-->
          <!--          </div>-->
        </div>
        <el-table
          ref="singleTable"
          :data="data"
          max-height="700"
          border
          highlight-current-row
          @current-change="handleCurrentChange"
          style="width: 100%">
          <el-table-column prop="name" label="类型名称"></el-table-column>
          <el-table-column prop="driver" label="关联插件"></el-table-column>
          <el-table-column label="操作" width="260">
            <template slot-scope="scope">
              <div style="display: flex">
                <el-tag class="rightMargin" @click="handleClick(scope.row)" style="cursor: pointer" type="warning"
                        size="medium">查看
                </el-tag>
                <el-tag style="cursor: pointer" size="medium" class="rightMargin" type="success"
                        @click="edit(scope.row)"> 编辑
                </el-tag>
                <el-popconfirm title="确定删除吗？" @onConfirm="cancel(scope.row)">
                  <el-tag style="cursor: pointer" size="medium" class="rightMargin" slot="reference" type="danger">删除
                  </el-tag>
                </el-popconfirm>
                <el-upload
                  class="upload-demo"
                  :action="`api/device-type/upload/${scope.row.name}`"
                  :headers="{token: token}"
                  :show-file-list="false"
                  name="FileName"
                  :on-success="handleSuccess"
                  accept=".zip"
                >
                  <el-tag style="cursor: pointer" size="medium">上传驱动</el-tag>
                </el-upload>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div>
        <el-button style="margin-bottom: 20px" @click="add2()" type="primary" icon="el-icon-plus" size="small">新增
        </el-button>
        <el-button style="margin-bottom: 20px;" @click="exportProperties()" icon="el-icon-download" size="small">导出
        </el-button>
        <el-table :data="data2" max-height="700" border style="width: 100%">
          <el-table-column prop="name" label="属性名称" width="160"></el-table-column>
          <el-table-column prop="description" label="描述"></el-table-column>
          <el-table-column prop="type" label="数据类型"></el-table-column>
          <el-table-column prop="length" label="数据长度"></el-table-column>
          <el-table-column prop="decimal" label="小数位数"></el-table-column>
          <el-table-column prop="unit" label="计量单位">
          </el-table-column>
          <el-table-column prop="value" label="最新数值"></el-table-column>
          <el-table-column prop="reported" label="是否上报">
            <template slot-scope="scope">
              <el-checkbox disabled v-model="scope.row.reported"></el-checkbox>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="200">
            <template slot-scope="scope">
              <el-tag class="rightMargin" @click="handleClick2(scope.row)" style="cursor: pointer" type="warning"
                      size="medium">查看
              </el-tag>
              <el-tag style="cursor: pointer" size="medium" class="rightMargin" type="success"
                      @click="edit2(scope.row,scope.$index)">编辑
              </el-tag>
              <el-popconfirm
                title="确定删除吗？"
                @onConfirm="cancel2(scope.$index + 1)"
              >
                <el-tag style="cursor: pointer" size="medium" class="rightMargin" slot="reference" type="danger">删除
                </el-tag>
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
    <el-dialog
      :title="title"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="80px">
        <el-form-item label="类型名称" prop="name">
          <el-input :disabled="title === '编辑' || this.title === '详情'" v-model="ruleForm.name"></el-input>
        </el-form-item>
      </el-form>
      <span v-if="this.title !== '详情'" slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="handleOk">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog :title="title" :visible.sync="dialogVisible2" width="30%" :before-close="handleClose2">
      <el-form ref="ruleForm2" :model="ruleForm2" :rules="rules2" label-width="80px">
        <el-form-item label="属性名称" prop="name">
          <!--          <el-input :disabled="this.title === '详情'" v-model="ruleForm2.name"></el-input>-->
          <el-select filterable allow-create v-model="ruleForm2.name" placeholder="请选择" @change="form2NameChange"
                     style="width: 260px">
            <el-option v-for="(item, index) in ruleFormNameOptions" :key="index" :label="item.value"
                       :value="item.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input :disabled="this.title === '详情'" v-model="ruleForm2.description"></el-input>
        </el-form-item>
        <el-form-item label="数据类型" prop="type">
          <el-input :disabled="this.title === '详情'" v-model="ruleForm2.type"
                    placeholder="double,int,uint,string"></el-input>
        </el-form-item>
        <el-form-item label="数据长度" prop="length">
          <el-input :disabled="this.title === '详情'" v-model="ruleForm2.length"></el-input>
        </el-form-item>
        <el-form-item label="小数位数" prop="decimal">
          <el-input :disabled="this.title === '详情'" v-model="ruleForm2.decimal"></el-input>
        </el-form-item>
        <el-form-item label="计量单位" prop="unit">
          <el-input :disabled="this.title === '详情'" v-model="ruleForm2.unit"></el-input>
        </el-form-item>
        <el-form-item label="是否上报" prop="reported">
          <el-checkbox :disabled="this.title === '详情'" v-model="ruleForm2.reported"></el-checkbox>
        </el-form-item>
      </el-form>
      <span v-if="this.title !== '详情'" slot="footer" class="dialog-footer">
        <el-button @click="handleClose2">取 消</el-button>
        <el-button type="primary" @click="handleOk2">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {
  addDeviceProperty,
  addDeviceType,
  deleteDeviceProperty,
  deleteDeviceType,
  findDevicePropertys,
  findDeviceTypes,
  updateDeviceProperty,
  updateDeviceType,
} from "@/api/device-type"

import {getToken} from '@/utils/auth'
import {exportJson} from "@/utils/export.js"

export default {
  data() {
    return {
      token: getToken(),
      index: null,
      title: '',
      dialogVisible: false,
      dialogVisible2: false,
      off: true,
      data: [],
      data2: [],
      name: null,
      ruleForm: {
        name: null,
        driver: null,
        properties: {}
      },
      ruleForm2: {
        name: null,
        description: null,
        type: null,
        length: null,
        decimal: null,
        unit: null,
        value: null,
        reported: true,
      },
      rules: {
        name: [{required: true, message: '请输入类型名称', trigger: 'blur'},],
        driver: [{required: true, message: '请输入关联插件', trigger: 'blur'},],
      },
      rules2: {
        name: [{required: true, message: '请输入名称', trigger: 'blur'},],
        description: [{required: true, message: '请输入描述', trigger: 'blur'},],
        type: [{required: true, message: '请输入类型', trigger: 'blur'},],
        length: [{required: true, message: '请输入长度', trigger: 'blur'}, {pattern: /^[0-9]*$/, message: '长度必须为数字值'}],
        decimal: [{required: true, message: '请输入小数位', trigger: 'blur'}, {pattern: /^[0-9]*$/, message: '小数位必须为数字值'}],
        unit: [{required: true, message: '请输入计量单位', trigger: 'blur'},],
        value: [{required: true, message: '请输入数值', trigger: 'blur'}, {pattern: /^[0-9]*$/, message: '数值必须为数字值'}],
      },

      ruleFormNameOptions: [
        {label: '总电能', value: 'dev_consumption', unit: 'kwh'},
        {label: '剩余金额', value: 'dev_remain_amt', unit: '元'},
        {label: '剩余电量', value: 'dev_remain_elec', unit: 'kwh'},
        {label: '欠费金额', value: 'dev_owe_amt', unit: '元'},
        {label: '欠费电量', value: 'dev_owe_elec', unit: 'kwh'},
        {label: '瞬时流量', value: 'dev_instant_flow', unit: 'm³'},
        {label: '总流量', value: 'dev_flow', unit: 'm³'},
        {label: '温度值', value: 'dev_temp', unit: '℃'},
        {label: '湿度值', value: 'dev_rh', unit: '%'},
        {label: 'PM2.5', value: 'dev_pm25', unit: 'mg/m³'},
        {label: 'CO2', value: 'dev_co2', unit: 'ppm'},
        {label: 'CH2O', value: 'dev_ch2o', unit: 'mg/m³'},
        {label: 'VOC', value: 'dev_voc', unit: 'mg/m³'},
        {label: 'O2', value: 'dev_o2', unit: '%'},
        {label: 'CO', value: 'dev_co', unit: 'mg/m³'},
        {label: '上下行', value: 'dev_updown', unit: 'int类型'},
        {label: '运行状态', value: 'dev_status', unit: 'int类型'},
        {label: '检修状态', value: 'dev_fix', unit: 'int类型'},
        {label: '故障状态', value: 'dev_err', unit: 'int类型'},
        {label: '停泊状态', value: 'dev_stp', unit: 'int类型'},
      ]
    }
  },
  methods: {
    //查询数据
    retrieve() {
      findDeviceTypes().then(
        res => {
          this.data = res.data
          if (this.name && !this.off) {
            this.$refs.singleTable.setCurrentRow(this.data.filter(item => item.name === this.name)[0]);
          }
          if (this.data.length) {
            if (this.off) {
              this.$refs.singleTable.setCurrentRow(this.data[0]);
              this.name = this.data[0].name
              this.off = false
            }
            this.retrieve2()
          } else {
            this.data2 = []
          }
        }
      )
    },
    retrieve2() {
      findDevicePropertys(this.name).then(
        res => {
          this.data2 = res.data
        }
      )
    },
    //导入
    importType() {

    },
    // 导出
    exportType() {

    },
    //选中
    handleCurrentChange(val) {
      this.name = val.name
      this.retrieve2()
    },
    //详情
    handleClick(row) {
      this.ruleForm = JSON.parse(JSON.stringify(row))
      this.title = '详情'
      this.dialogVisible = true
    },
    handleClick2(row) {
      this.title = '详情'
      this.ruleForm2 = JSON.parse(JSON.stringify(row))
      this.dialogVisible2 = true

    },
    //新增
    add() {
      this.ruleForm = {
        name: null,
        driver: null,
        properties: []
      }
      this.title = '新增'
      this.dialogVisible = true
    },
    add2() {
      this.index = this.data2.length + 1
      this.ruleForm2 = {
        name: null,
        description: null,
        type: 'double',
        length: 13,
        decimal: 2,
        unit: null,
        value: 0,
        reported: true,
      },
        this.title = '新增'
      this.dialogVisible2 = true
    },
    exportProperties() {
      exportJson(JSON.stringify(this.data2, null, '\t'), 'properties')
    },
    //编辑
    edit(row) {
      this.title = '编辑'
      this.ruleForm = JSON.parse(JSON.stringify(row))
      this.dialogVisible = true

    },
    edit2(row, index) {
      this.index = index + 1
      this.title = '编辑'
      this.ruleForm2 = JSON.parse(JSON.stringify(row))
      this.dialogVisible2 = true

    },
    // 取消
    handleClose() {
      this.$refs['ruleForm'].resetFields();
      this.dialogVisible = false
    },
    handleClose2() {
      this.$refs['ruleForm2'].resetFields();
      this.dialogVisible2 = false
    },
    //保存 编辑/新增
    handleOk() {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          if (this.title === '新增') {
            addDeviceType(this.ruleForm).then(
              res => {
                this.retrieve()
                this.$message.info(res.message)
                this.$refs['ruleForm'].resetFields();
                this.dialogVisible = false
              }
            )
          } else if (this.title === '编辑') {
            updateDeviceType(this.ruleForm).then(
              res => {
                this.retrieve()
                this.$message.info(res.message)
                this.$refs['ruleForm'].resetFields();
                this.dialogVisible = false
              }
            )
          }
        }
      })
    },
    handleOk2() {
      this.$refs['ruleForm2'].validate((valid) => {
        if (valid) {
          this.ruleForm2.length = Number(this.ruleForm2.length)
          this.ruleForm2.decimal = Number(this.ruleForm2.decimal)
          this.ruleForm2.value = Number(this.ruleForm2.value)
          if (this.title === '新增') {
            addDeviceProperty(this.ruleForm2, this.name).then(
              res => {
                this.retrieve2()
                this.$message.info(res.message)
                this.$refs['ruleForm2'].resetFields();
                this.dialogVisible2 = false
              }
            )
          } else if (this.title === '编辑') {
            updateDeviceProperty(this.ruleForm2, this.name, this.index).then(
              res => {
                this.retrieve2()
                this.$message.info(res.message)
                this.$refs['ruleForm2'].resetFields();
                this.dialogVisible2 = false
              }
            )
          }
        }
      })
    },
    //删除
    cancel(row) {
      deleteDeviceType(row.name).then(
        res => {
          if (res.code === 0 && row.name === this.name) {
            this.off = true
          }
          this.$message.info(res.message)
          this.retrieve()
        }
      )
    },
    cancel2(index) {
      deleteDeviceProperty(this.name, index).then(
        res => {
          this.$message.info(res.message)
          this.retrieve2()
        }
      )
    },
    handleSuccess(res) {
      this.$message.info(res.message)
      if (res.code === 0) {
        this.retrieve()
      }
    },

    form2NameChange() {
      this.ruleForm2.description = this.setForm2Attr(this.ruleForm2.name).description;
      this.ruleForm2.unit = this.setForm2Attr(this.ruleForm2.name).unit;
    },

    setForm2Attr(value) {
      let findEle = this.ruleFormNameOptions.filter(item => {
        return item.value == value;
      })
      return {
        description: findEle[0].label,
        unit: findEle[0].unit,
      };
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
  justify-content: space-between;
}

.box > div:nth-child(1) {
  width: 35%;
}

.box > div:nth-child(2) {
  width: calc(65% - 20px);
}

.rightMargin {
  margin-right: 10px;
}
</style>
