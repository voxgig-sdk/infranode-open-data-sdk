# InfranodeOpenData SDK feature factory

from infranodeopendata_sdk.feature.base_feature import InfranodeOpenDataBaseFeature
from infranodeopendata_sdk.feature.ratelimit_feature import InfranodeOpenDataRatelimitFeature
from infranodeopendata_sdk.feature.retry_feature import InfranodeOpenDataRetryFeature
from infranodeopendata_sdk.feature.test_feature import InfranodeOpenDataTestFeature
from infranodeopendata_sdk.feature.timeout_feature import InfranodeOpenDataTimeoutFeature


_FEATURES = {
    "base": lambda: InfranodeOpenDataBaseFeature(),
    "ratelimit": lambda: InfranodeOpenDataRatelimitFeature(),
    "retry": lambda: InfranodeOpenDataRetryFeature(),
    "test": lambda: InfranodeOpenDataTestFeature(),
    "timeout": lambda: InfranodeOpenDataTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
