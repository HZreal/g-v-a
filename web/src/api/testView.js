import service from '@/utils/request'


export const testView = (data) => {
  return service({
    url: '/testGroup/test111',
    method: 'post',
    data
  })
}
