import request from '@/utils/request'

const api = '/api/ethernet'

export function addnetWork(data) {
  return request({url: api, method: 'post', data})
}

export function updatenetWork(data) {
  return request({url: api, method: 'put', data})
}

export function deletenetWork(name) {
  return request({url: `${api}/${name}`, method: 'delete'})
}

export function findnetWork(name) {
  return request({url: `${api}/${name}`, method: 'get'})
}

export function findnetWorks(data) {
  return request({url: `${api}s`, method: 'get'})
}
