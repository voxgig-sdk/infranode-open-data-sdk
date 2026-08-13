# InfranodeOpenData SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

InfranodeOpenDataUtility.registrar = ->(u) {
  u.clean = InfranodeOpenDataUtilities::Clean
  u.done = InfranodeOpenDataUtilities::Done
  u.make_error = InfranodeOpenDataUtilities::MakeError
  u.feature_add = InfranodeOpenDataUtilities::FeatureAdd
  u.feature_hook = InfranodeOpenDataUtilities::FeatureHook
  u.feature_init = InfranodeOpenDataUtilities::FeatureInit
  u.fetcher = InfranodeOpenDataUtilities::Fetcher
  u.make_fetch_def = InfranodeOpenDataUtilities::MakeFetchDef
  u.make_context = InfranodeOpenDataUtilities::MakeContext
  u.make_options = InfranodeOpenDataUtilities::MakeOptions
  u.make_request = InfranodeOpenDataUtilities::MakeRequest
  u.make_response = InfranodeOpenDataUtilities::MakeResponse
  u.make_result = InfranodeOpenDataUtilities::MakeResult
  u.make_point = InfranodeOpenDataUtilities::MakePoint
  u.make_spec = InfranodeOpenDataUtilities::MakeSpec
  u.make_url = InfranodeOpenDataUtilities::MakeUrl
  u.param = InfranodeOpenDataUtilities::Param
  u.prepare_auth = InfranodeOpenDataUtilities::PrepareAuth
  u.prepare_body = InfranodeOpenDataUtilities::PrepareBody
  u.prepare_headers = InfranodeOpenDataUtilities::PrepareHeaders
  u.prepare_method = InfranodeOpenDataUtilities::PrepareMethod
  u.prepare_params = InfranodeOpenDataUtilities::PrepareParams
  u.prepare_path = InfranodeOpenDataUtilities::PreparePath
  u.prepare_query = InfranodeOpenDataUtilities::PrepareQuery
  u.graphql_body = InfranodeOpenDataUtilities::GraphqlBody
  u.graphql_errors = InfranodeOpenDataUtilities::GraphqlErrors
  u.result_basic = InfranodeOpenDataUtilities::ResultBasic
  u.result_body = InfranodeOpenDataUtilities::ResultBody
  u.result_headers = InfranodeOpenDataUtilities::ResultHeaders
  u.transform_request = InfranodeOpenDataUtilities::TransformRequest
  u.transform_response = InfranodeOpenDataUtilities::TransformResponse
}
