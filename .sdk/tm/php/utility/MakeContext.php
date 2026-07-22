<?php
declare(strict_types=1);

// InfranodeOpenData SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class InfranodeOpenDataMakeContext
{
    public static function call(array $ctxmap, ?InfranodeOpenDataContext $basectx): InfranodeOpenDataContext
    {
        return new InfranodeOpenDataContext($ctxmap, $basectx);
    }
}
