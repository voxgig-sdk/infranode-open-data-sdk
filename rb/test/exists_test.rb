# InfranodeOpenData SDK exists test

require "minitest/autorun"
require_relative "../InfranodeOpenData_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = InfranodeOpenDataSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
