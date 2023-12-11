<template>
  <div class="container">
    <el-button disabled style="margin-bottom: 20px" @click="add()" type="primary" icon="el-icon-plus" size="small">新增
    </el-button>
    <el-table
        :data="data"
        max-height="700"
        border
        style="width: 100%">
      <el-table-column prop="name" label="服务名称" width="180"></el-table-column>
      <el-table-column prop="reportName" label="平台名称"></el-table-column>
      <el-table-column prop="ip" label="Ip地址"></el-table-column>
      <el-table-column prop="port" label="端口号"></el-table-column>
      <el-table-column prop="clientID" label="网关编号"></el-table-column>
      <el-table-column prop="username" label="用户名"></el-table-column>
      <el-table-column prop="password" label="密码"></el-table-column>
      <el-table-column prop="cron" label="定时策略"></el-table-column>
      <el-table-column prop="status" label="状态">
        <template slot-scope="scope">
          <span style="display: flex;align-items: center">
              <i v-if="scope.row.status === 1" class="el-icon-video-play"></i>
             <i v-else class="el-icon-video-pause"></i>
            {{ scope.row.status === 0 ? '已停止' : '运行中' }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template slot-scope="scope">
          <el-tag style="cursor: pointer" size="medium" class="rightMargin" type="success" @click="edit(scope.row)">编辑
          </el-tag>
          <el-popconfirm title="确定删除吗？" @onConfirm="cancel(scope.row)" :disabled="true">
            <el-tag style="cursor: pointer" size="medium" class="rightMargin" slot="reference" type="danger">删除</el-tag>
          </el-popconfirm>
          <el-popconfirm :title="`确定${scope.row.status === 0?'启动':'停止'}吗？`" @onConfirm="handleClick(scope.row)">
            <el-tag style="cursor: pointer" slot="reference" size="medium">{{ scope.row.status === 0 ? '启动' : '停止' }}
            </el-tag>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog :title="title" :visible.sync="dialogVisible" width="30%" :before-close="handleClose">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="80px">
        <el-form-item label="服务名称" prop="name">
          <el-input :disabled="title === '编辑'" v-model="ruleForm.name"></el-input>
        </el-form-item>
        <el-form-item label="平台名称" prop="reportName">
          <el-input disabled v-model="ruleForm.reportName"></el-input>
        </el-form-item>
        <el-form-item label="Ip地址" prop="ip">
          <el-input v-model="ruleForm.ip"></el-input>
        </el-form-item>
        <el-form-item label="端口号" prop="port">
          <el-input v-model="ruleForm.port"></el-input>
        </el-form-item>
        <el-form-item label="网关编号" prop="clientID">
          <el-input v-model="ruleForm.clientID"></el-input>
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input disabled v-model="ruleForm.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input disabled v-model="ruleForm.password"></el-input>
        </el-form-item>
        <el-form-item label="定时策略" prop="cron">
          <el-input v-model="ruleForm.cron"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="handleOk">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {
  addReportTask,
  updateReportTask,
  deleteReportTask,
  findReportTask,
  findReportTasks,
  startReportTask,
  stopReportTask
} from "@/api/report-task";

export default {

  data() {
    return {
      data: [],
      dialogVisible: false,
      title: '',
      ruleForm: {
        name: null,
        reportName: null,
        ip: null,
        port: null,
        clientID: null,
        username: null,
        password: null,
        cron: null,
        status: null,
      },
      rules: {
        name: [{required: true, message: '请输入服务名称', trigger: 'blur'},],
        reportName: [{required: true, message: '请输入平台名称', trigger: 'blur'},],
        ip: [{required: true, message: '请输入Ip地址', trigger: 'blur'}, {
          //pattern: /^([0-9].|[1-9][0-9].|1[0-9]{1,2}.|2[0-4][0-9]|25[0-5].){3}([0-9]|[1-9][0-9]|1[0-9]{1,2}|2[0-4][0-9]|25[0-5])$/,
          pattern: /^((\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])$/,
          message: 'Ip地址格式错误'
        }],
        port: [{required: true, message: '请输入端口号', trigger: 'blur'}, {pattern: /^\d{1,4}$/, message: '端口号格式错误'}],
        clientID: [{required: true, message: '请输入网关编号', trigger: 'blur'},],
        username: [{required: true, message: '请输入用户名', trigger: 'blur'},],
        password: [{required: true, message: '请输入密码', trigger: 'blur'},],
        cron: [{required: true, message: '请输入定时策略', trigger: 'blur'},],
      }
    }
  },
  methods: {
    // 启动/停止
    handleClick(row) {
      if (row.status === 0) {
        startReportTask(row.name).then(
            res => {
              this.$message.info(res.message)
              this.retrieve()
            }
        )
      } else {
        stopReportTask(row.name).then(
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
        name: '数据上报',
        reportName: '物联平台',
        ip: null,
        port: 1883,
        username: 'admin',
        password: 'admin',
        cron: '@every 1m30s',
        status: 0,
        clientID: null,
      },
          this.title = '新增'
      this.dialogVisible = true
    },
    // 删除
    cancel(row) {
      deleteReportTask(row.name).then(
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
      findReportTasks().then(
          res => {
            this.data = res.data
            // this.$message.info(res.message)
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
          this.ruleForm.port = Number(this.ruleForm.port)
          this.ruleForm.status = Number(this.ruleForm.status)
          if (this.title === '新增') {
            addReportTask(this.ruleForm).then(
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
            updateReportTask(this.ruleForm).then(
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
</style>
