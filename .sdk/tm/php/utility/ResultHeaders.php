<?php
declare(strict_types=1);

// InfranodeOpenData SDK utility: result_headers

class InfranodeOpenDataResultHeaders
{
    public static function call(InfranodeOpenDataContext $ctx): ?InfranodeOpenDataResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
