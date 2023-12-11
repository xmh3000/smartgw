<template>
  <div class="container">
    <el-button disabled style="margin-bottom: 20px" @click="add()" type="primary" icon="el-icon-plus" size="small">新增</el-button>
    <el-table
      :data="data"
      max-height="700"
      border
      style="width: 100%">
      <el-table-column prop="name" label="网口名称" ></el-table-column>
      <el-table-column prop="index" label="索引" ></el-table-column>
      <el-table-column prop="MTU" label="MTU" ></el-table-column>
      <el-table-column prop="MAC" label="MAC地址" ></el-table-column>
      <el-table-column prop="flags" label="网卡标志" ></el-table-column>
      <el-table-column prop="IP" label="网络地址" ></el-table-column>
      <el-table-column prop="netmask" label="子网掩码" ></el-table-column>
      <el-table-column prop="gateway" label="默认网关" ></el-table-column>
      <el-table-column prop="configEnabled" label="配置">
        <template slot-scope="scope">
          {{ scope.row.configEnabled ? '是' : '否' }}
        </template>
      </el-table-column>
      <el-table-column prop="DHCPEnabled" label="模式">
        <template slot-scope="scope">
          {{ scope.row.DHCPEnabled ? '手动' : '自动' }}
        </template>
      </el-table-column>
      <el-table-column prop="configIP" label="网络地址"></el-table-column>
      <el-table-column prop="configNetmask" label="子网掩码"></el-table-column>
      <el-table-column prop="configGateway" label="默认网关"></el-table-column>
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
        <el-form-item label="网口名称" prop="name">
          <el-input :disabled="title === '编辑' || this.title === '查看'" v-model="ruleForm.name"></el-input>
        </el-form-item>
        <el-form-item label="索引" prop="index">
          <el-input :disabled="title === '编辑' || this.title === '查看'" v-model="ruleForm.index"></el-input>
        </el-form-item>
        <el-form-item label="MTU" prop="MTU">
          <el-input :disabled="title === '编辑' || this.title === '查看'" v-model="ruleForm.MTU"></el-input>
        </el-form-item>
        <el-form-item label="MAC地址" prop="MAC">
          <el-input :disabled="title === '编辑' || this.title === '查看'" v-model="ruleForm.MAC"></el-input>
        </el-form-item>
        <el-form-item label="网卡标志" prop="flags">
          <el-input :disabled="title === '编辑' || this.title === '查看'" v-model="ruleForm.flags"></el-input>
        </el-form-item>
        <el-form-item label="配置" prop="configEnabled">
          <el-switch :disabled="this.title === '查看'" v-model="ruleForm.configEnabled"></el-switch>
        </el-form-item>
        <el-form-item v-if="ruleForm.configEnabled" label="手动模式" prop="DHCPEnabled">
          <el-switch :disabled="this.title === '查看'" v-model="ruleForm.DHCPEnabled"></el-switch>
        </el-form-item>
        <el-form-item v-if="ruleForm.DHCPEnabled && ruleForm.configEnabled" label="网络地址" prop="ConfigIP">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.ConfigIP"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.DHCPEnabled && ruleForm.configEnabled" label="子网掩码" prop="ConfigNetmask">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.ConfigNetmask"></el-input>
        </el-form-item>
        <el-form-item v-if="ruleForm.DHCPEnabled && ruleForm.configEnabled" label="默认网关" prop="ConfigGateway">
          <el-input :disabled="this.title === '查看'" v-model="ruleForm.ConfigGateway"></el-input>
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
import {addnetWork, updatenetWork, deletenetWork, findnetWorks} from '@/api/net-work'
export default {
  data(){
    return {
      data: [],
      dialogVisible: false,
      title: '',
      ruleForm: {
        name: null,
        index:null,
        MTU:null,
        MAC:null,
        flags:null,
        IP:null,
        netmask:null,
        gateway:null,
        configEnabled: null,
        DHCPEnabled: null,
        ConfigIP: null,
        ConfigNetmask: null,
        ConfigGateway: null,
      },
      rules: {
        name: [{required: true, message: '请输入网口名称', trigger: 'blur'},],
        index: [{required: true, message: '请输入索引', trigger: 'blur'},],
        MTU: [{required: true, message: '请输入MTU', trigger: 'blur'},],
        MAC: [{required: true, message: '请输入MAC地址', trigger: 'blur'},],
        flags: [{required: true, message: '请输入网卡标志', trigger: 'blur'},],
        IP: [{required: true, message: '请输入网络地址', trigger: 'blur'},],
        netmask: [{required: true, message: '请输入子网掩码', trigger: 'blur'},],
        ConfigIP: [{required: true, message: '请输入网络地址', trigger: 'blur'},],
        ConfigNetmask: [{required: true, message: '请输入子网掩码', trigger: 'blur'},],
        ConfigGateway: [{required: true, message: '请输入默认网关', trigger: 'blur'},],
      }
    }
  },
  methods:{
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
        ruleForm: {
          name: null,
          index:null,
          MTU:null,
          MAC:null,
          flags:null,
          IP:null,
          netmask:null,
          gateway:null,
          configEnabled: false,
          DHCPEnabled: false,
          ConfigIP: null,
          ConfigNetmask: null,
          ConfigGateway: null,
        },
      }
      this.title = '新增'
      this.dialogVisible = true
    },
    // 删除
    cancel(row) {
      deletenetWork(row.name).then(
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
      findnetWorks().then(
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
          if (this.title === '新增') {
            addnetWork(this.ruleForm).then(
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
            updatenetWork(this.ruleForm).then(
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

<style scoped>
.rightMargin {
  margin-right: 10px;
}
</style>
