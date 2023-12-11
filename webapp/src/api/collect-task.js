import request from '@/utils/request'

const api = '/api/collect-task'

export function addCollectTask(data) {
  return request({url: api, method: 'post', data})
}

export function updateCollectTask(data) {
  return request({url: api, method: 'put', data})
}

export function deleteCollectTask(name) {
  return request({url: `${api}/${name}`, method: 'delete'})
}

export function findCollectTask(name) {
  return request({url: `${api}/${name}`, method: 'get'})
}

export function findCollectTasks(data) {
  return request({url: `${api}s`, method: 'get'})
}

// 启动任务
export function startCollectTask(name) {
  return request({url: `${api}/start/${name}`, method: 'get'})
}

// 停止任务
export function stopCollectTask(name) {
  return request({url: `${api}/stop/${name}`, method: 'get'})
}
