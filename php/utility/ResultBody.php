<?php
declare(strict_types=1);

// InfranodeOpenData SDK utility: result_body

class InfranodeOpenDataResultBody
{
    public static function call(InfranodeOpenDataContext $ctx): ?InfranodeOpenDataResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
