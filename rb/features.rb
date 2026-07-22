# InfranodeOpenData SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module InfranodeOpenDataFeatures
  def self.make_feature(name)
    case name
    when "base"
      InfranodeOpenDataBaseFeature.new
    when "test"
      InfranodeOpenDataTestFeature.new
    else
      InfranodeOpenDataBaseFeature.new
    end
  end
end
