<template>
  <div class="container">
    <div class="box">
      <div class="box1">
        <el-input
          type="textarea"
          :autosize="{ minRows: 16}"
          placeholder="接收区"
          v-model="textarea1"
          class="textarea"
        >
        </el-input>
        <div class="box2">
          <el-select v-model="value" placeholder="请选择" class="marginBottom">
            <el-option
              v-for="item in data"
              :key="item.name"
              :label="item.name"
              :value="item.name">
            </el-option>
          </el-select>
          <el-button class="but1" @click="clearGet">清空接收区</el-button>
        </div>
      </div>
      <div class="box1">
        <el-input
          type="textarea"
          :autosize="{ minRows: 6}"
          placeholder="发送区"
          v-model="textarea2"
          class="textarea">
        </el-input>
        <div class="box2">
          <el-button class="marginBottom but2"  @click="sendDirectData">发送</el-button>
          <el-button class="but1" @click="clearSend">清空发送区</el-button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>

import {findCollectors} from "@/api/collector";
import {test} from "@/api/debug"

export default {
  data(){
    return {
      data:[],
      textarea1:'',
      textarea2:'',
      value:''
    }

  },
  methods:{
    clearGet(){
      this.textarea1=''
    },
    sendDirectData(){
      let data = {
        collectorName:this.value,
        data:this.textarea2
      }
      test(data).then(
        res=>{
          if(res.code === 0){
            this.textarea1 = res.data
          }
          this.$message.info(res.message)
        }
      )
    },
    clearSend(){
      this.textarea2=''
    }
  },
  mounted(){
    findCollectors().then(
      res => {
        this.data = res.data
        if(this.data.length){
          this.value = res.data[0].name
        }
        // this.$message.info(res.message)
      }
    )
  }
}
</script>

<style scoped>
.box{
  width: 65%;
  margin: 20px auto;

}
.box1{
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}
.textarea {
  width: 70%;
}
::v-deep .textarea .el-textarea__inner{
  border: 1px solid #0990d2;
  border-radius: 10px;
}
.box2{
  width: 30%;
  padding: 0 10px;
  display: flex;
  flex-direction: column;
}
.marginBottom {
  margin-bottom: 20px;
}
.but1 {
  background-color: #ffb800;
  border-radius: 50px;
  color: #fff;
}
.but2 {
  background-color: #1e9fff;
  border-radius: 50px;
  color: #fff;
}
</style>
