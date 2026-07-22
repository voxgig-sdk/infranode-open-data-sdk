# InfranodeOpenData SDK utility: make_context

from core.context import InfranodeOpenDataContext


def make_context_util(ctxmap, basectx):
    return InfranodeOpenDataContext(ctxmap, basectx)
