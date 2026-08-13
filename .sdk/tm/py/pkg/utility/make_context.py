# InfranodeOpenData SDK utility: make_context

from projectname_sdk.core.context import InfranodeOpenDataContext


def make_context_util(ctxmap, basectx):
    return InfranodeOpenDataContext(ctxmap, basectx)
