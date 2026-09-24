
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { InfranodeOpenDataSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = InfranodeOpenDataSDK.test()
    equal(testsdk instanceof InfranodeOpenDataSDK, true,
      'InfranodeOpenDataSDK.test() must return a client synchronously')
  })

})
