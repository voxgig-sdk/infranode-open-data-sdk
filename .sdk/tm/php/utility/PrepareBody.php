<?php
declare(strict_types=1);

// InfranodeOpenData SDK utility: prepare_body

class InfranodeOpenDataPrepareBody
{
    public static function call(InfranodeOpenDataContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
