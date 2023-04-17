import request from '@/utils/request'

export function login(data) {
  return request({
    url: 'login',
    method: 'post',
    data
  })
}

export function find(username) {
  return request({
    url: 'api/user/' + username,
    method: 'get'
  })
}
