-- InfranodeOpenData SDK exists test

local sdk = require("infranode-open-data_sdk")

describe("InfranodeOpenDataSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
