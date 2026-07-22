<?php
declare(strict_types=1);

// InfranodeOpenData SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class InfranodeOpenDataFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new InfranodeOpenDataBaseFeature();
            case "test":
                return new InfranodeOpenDataTestFeature();
            default:
                return new InfranodeOpenDataBaseFeature();
        }
    }
}
