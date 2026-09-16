# InfranodeOpenData SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module InfranodeOpenDataFeatures
  def self.make_feature(name)
    case name
    when "base"
      InfranodeOpenDataBaseFeature.new
    when "ratelimit"
      InfranodeOpenDataRatelimitFeature.new
    when "retry"
      InfranodeOpenDataRetryFeature.new
    when "test"
      InfranodeOpenDataTestFeature.new
    when "timeout"
      InfranodeOpenDataTimeoutFeature.new
    else
      InfranodeOpenDataBaseFeature.new
    end
  end
end
