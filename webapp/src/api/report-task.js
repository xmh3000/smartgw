import request from '@/utils/request'

const api = '/api/report-task'

export function addReportTask(data) {
  return request({url: api, method: 'post', data})
}

export function updateReportTask(data) {
  return request({url: api, method: 'put', data})
}

export function deleteReportTask(name) {
  return request({url: `${api}/${name}`, method: 'delete'})
}

export function findReportTask(name) {
  return request({url: `${api}/${name}`, method: 'get'})
}

export function findReportTasks(data) {
  return request({url: `${api}s`, method: 'get'})
}

// 启动任务
export function startReportTask(name) {
  return request({url: `${api}/start/${name}`, method: 'get'})
}

// 停止任务
export function stopReportTask(name) {
  return request({url: `${api}/stop/${name}`, method: 'get'})
}
