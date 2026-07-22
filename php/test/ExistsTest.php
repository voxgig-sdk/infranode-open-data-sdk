<?php
declare(strict_types=1);

// InfranodeOpenData SDK exists test

require_once __DIR__ . '/../infranodeopendata_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = InfranodeOpenDataSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
