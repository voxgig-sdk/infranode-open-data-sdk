# InfranodeOpenData SDK feature factory

from feature.base_feature import InfranodeOpenDataBaseFeature
from feature.test_feature import InfranodeOpenDataTestFeature


def _make_feature(name):
    features = {
        "base": lambda: InfranodeOpenDataBaseFeature(),
        "test": lambda: InfranodeOpenDataTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
