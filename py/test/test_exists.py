# InfranodeOpenData SDK exists test

import pytest
from infranodeopendata_sdk import InfranodeOpenDataSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = InfranodeOpenDataSDK.test(None, None)
        assert testsdk is not None
