import request from '@/utils/request'

export function login(data) {
  return request({
    url: 'login',
    method: 'post',
    data
  })
}

export function findUser() {
  return request({
    url: 'api/user/admin',
    method: 'get',
  })
}
