<?php
declare(strict_types=1);

// InfranodeOpenData SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

InfranodeOpenDataUtility::setRegistrar(function (InfranodeOpenDataUtility $u): void {
    $u->clean = [InfranodeOpenDataClean::class, 'call'];
    $u->done = [InfranodeOpenDataDone::class, 'call'];
    $u->make_error = [InfranodeOpenDataMakeError::class, 'call'];
    $u->feature_add = [InfranodeOpenDataFeatureAdd::class, 'call'];
    $u->feature_hook = [InfranodeOpenDataFeatureHook::class, 'call'];
    $u->feature_init = [InfranodeOpenDataFeatureInit::class, 'call'];
    $u->fetcher = [InfranodeOpenDataFetcher::class, 'call'];
    $u->make_fetch_def = [InfranodeOpenDataMakeFetchDef::class, 'call'];
    $u->make_context = [InfranodeOpenDataMakeContext::class, 'call'];
    $u->make_options = [InfranodeOpenDataMakeOptions::class, 'call'];
    $u->make_request = [InfranodeOpenDataMakeRequest::class, 'call'];
    $u->make_response = [InfranodeOpenDataMakeResponse::class, 'call'];
    $u->make_result = [InfranodeOpenDataMakeResult::class, 'call'];
    $u->make_point = [InfranodeOpenDataMakePoint::class, 'call'];
    $u->make_spec = [InfranodeOpenDataMakeSpec::class, 'call'];
    $u->make_url = [InfranodeOpenDataMakeUrl::class, 'call'];
    $u->param = [InfranodeOpenDataParam::class, 'call'];
    $u->prepare_auth = [InfranodeOpenDataPrepareAuth::class, 'call'];
    $u->prepare_body = [InfranodeOpenDataPrepareBody::class, 'call'];
    $u->prepare_headers = [InfranodeOpenDataPrepareHeaders::class, 'call'];
    $u->prepare_method = [InfranodeOpenDataPrepareMethod::class, 'call'];
    $u->prepare_params = [InfranodeOpenDataPrepareParams::class, 'call'];
    $u->prepare_path = [InfranodeOpenDataPreparePath::class, 'call'];
    $u->prepare_query = [InfranodeOpenDataPrepareQuery::class, 'call'];
    $u->result_basic = [InfranodeOpenDataResultBasic::class, 'call'];
    $u->result_body = [InfranodeOpenDataResultBody::class, 'call'];
    $u->result_headers = [InfranodeOpenDataResultHeaders::class, 'call'];
    $u->transform_request = [InfranodeOpenDataTransformRequest::class, 'call'];
    $u->transform_response = [InfranodeOpenDataTransformResponse::class, 'call'];
});
