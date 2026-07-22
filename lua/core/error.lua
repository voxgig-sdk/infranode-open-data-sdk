-- InfranodeOpenData SDK error

local InfranodeOpenDataError = {}
InfranodeOpenDataError.__index = InfranodeOpenDataError


function InfranodeOpenDataError.new(code, msg, ctx)
  local self = setmetatable({}, InfranodeOpenDataError)
  self.is_sdk_error = true
  self.sdk = "InfranodeOpenData"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function InfranodeOpenDataError:error()
  return self.msg
end


function InfranodeOpenDataError:__tostring()
  return self.msg
end


return InfranodeOpenDataError
