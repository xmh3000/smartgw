import request from '@/utils/request'

const api = '/api/device-type'
const api2 = '/api/device-property'

export function addDeviceType(data) {
  return request({url: api, method: 'post', data})
}

export function updateDeviceType(data) {
  return request({url: api, method: 'put', data})
}

export function deleteDeviceType(name) {
  return request({url: `${api}/${name}`, method: 'delete'})
}

export function findDeviceType(name) {
  return request({url: `${api}/${name}`, method: 'get',})
}

export function findDeviceTypes() {
  return request({url: `${api}s`, method: 'get'})
}

export function addDeviceProperty(data, name) {
  return request({url: `${api2}/${name}`, method: 'post', data})
}

export function updateDeviceProperty(data, name, propertyid) {
  return request({url: `${api2}/${name}/${propertyid}`, method: 'put', data})
}

export function deleteDeviceProperty(name, propertyid) {
  return request({url: `${api2}/${name}/${propertyid}`, method: 'delete'})
}

export function findDeviceProperty(name, propertyid) {
  return request({url: `${api2}/${name}/${propertyid}`, method: 'get'})
}

export function findDevicePropertys(name) {
  return request({url: `${api2}/${name}`, method: 'get'})
}
