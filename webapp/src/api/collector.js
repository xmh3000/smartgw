import request from '@/utils/request'

const api = '/api/collector'

export function addCollector(data) {
  return request({url: api, method: 'post', data})
}

export function updateCollector(data) {
  return request({url: api, method: 'put', data})
}

export function deleteCollector(name) {
  return request({url: `${api}/${name}`, method: 'delete'})
}

export function findCollector(name) {
  return request({url: `${api}/${name}`, method: 'get'})
}

export function findCollectors(data) {
  return request({url: `${api}s`, method: 'get'})
}
