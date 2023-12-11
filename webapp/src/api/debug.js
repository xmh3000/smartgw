import request from '@/utils/request'

const api = '/api/debug/test'
const api2 = '/api/debug'

export function test(data) {
    return request({url: api, method: 'post', data})
}
export function systemStatus() {
  return request({url: api2 + '/system-status', method: 'post' })
}
export function systemReboot(data) {
  return request({url: api2 + '/system-reboot', method: 'post'})
}
export function systemNtp(data) {
  return request({url: api2 + '/system-ntp', method: 'post'})
}
