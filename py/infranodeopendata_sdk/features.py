# InfranodeOpenData SDK feature factory

from infranodeopendata_sdk.feature.base_feature import InfranodeOpenDataBaseFeature
from infranodeopendata_sdk.feature.test_feature import InfranodeOpenDataTestFeature


def _make_feature(name):
    features = {
        "base": lambda: InfranodeOpenDataBaseFeature(),
        "test": lambda: InfranodeOpenDataTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
