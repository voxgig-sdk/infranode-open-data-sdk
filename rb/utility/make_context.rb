# InfranodeOpenData SDK utility: make_context
require_relative '../core/context'
module InfranodeOpenDataUtilities
  MakeContext = ->(ctxmap, basectx) {
    InfranodeOpenDataContext.new(ctxmap, basectx)
  }
end
