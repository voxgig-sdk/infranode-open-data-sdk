
import { Context } from './Context'


class InfranodeOpenDataError extends Error {

  isInfranodeOpenDataError = true

  sdk = 'InfranodeOpenData'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  InfranodeOpenDataError
}

