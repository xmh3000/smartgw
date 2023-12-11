import request from '@/utils/request'

const api = '/api/device'

export function addDevice(data) {
  return request({url: api, method: 'post', data})
}

export function updateDevice(data) {
  return request({url: api, method: 'put', data})
}

export function deleteDevice(name) {
  return request({url: `${api}/${name}`, method: 'delete'})
}

export function findDevice(name) {
  return request({url: `${api}/${name}`, method: 'get'})
}

export function findDevices() {
  return request({url: `${api}s`, method: 'get'})
}

// 导出
export function exportDevices() {
  return request({url: `${api}s/export`, method: 'get'})
}

