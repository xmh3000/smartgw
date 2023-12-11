<template>
  <div ref="echartDivRef" class="dashboardLine" />
</template>

<script>
import * as echarts from 'echarts'

export default {
  name: 'DashboardLine',
  data() {
    return {
      myChart: echarts.ECharts
    }
  },
  mounted() {
    this.myChart = echarts.init(this.$refs['echartDivRef'])
    this.updateTrend()
    window.addEventListener('resize', () => {
      this.myChart.resize()
    })
  },
  methods: {
    updateTrend() {
      this.myChart.setOption({
        legend: {},
        tooltip: {
          trigger: 'axis'
        }, xAxis: {
          type: 'category',
          data: []
        }, yAxis: {
          type: 'value'
        }, series: [{
          data: [],
          name: [],
          type: 'line',
          smooth: true
        }]
      })
    }
  }
}
</script>

<style scoped>
.dashboardLine {
  width: 100%;
  height: 100%;
}
</style>
