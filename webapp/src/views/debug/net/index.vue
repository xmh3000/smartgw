<template>
  <div class="container">
    <el-card class="box-card">
      <div slot="header">
        系统升级
      </div>
      <el-form label-position="left" label-width="120px">
        <el-form-item label="请选择系统文件:">
          <el-upload
              class="upload-demo" :action="`/api/debug/upgrade`" :headers="{token: token}" :show-file-list="false"
              name="FileName" :on-success="handleSuccess" accept=".zip">
            <el-button type="primary">上传系统</el-button>
          </el-upload>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import {getToken} from '@/utils/auth'

export default {
  data() {
    return {
      token: getToken()
    }
  },
  methods: {
    handleSuccess(res) {
      this.$message.info(res.message)
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
