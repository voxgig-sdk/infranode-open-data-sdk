# Typed models for the InfranodeOpenData SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class CityRequired(TypedDict):
    data: Any
    meta: dict


class City(CityRequired, total=False):
    id: str


class CityLoadMatch(TypedDict):
    id: str


class CityListMatch(TypedDict, total=False):
    data: Any
    id: str
    meta: dict


class CompareRequired(TypedDict):
    city: str
    source_status: str


class Compare(CompareRequired, total=False):
    data: dict


class CompareListMatchRequired(TypedDict):
    city: str
    resource: str


class CompareListMatch(CompareListMatchRequired, total=False):
    limit: int
    offset: int
    order: str
    page: int
    sort: str


class Health(TypedDict):
    redis: bool
    status: str
    version: str


class HealthLoadMatch(TypedDict, total=False):
    redis: bool
    status: str
    version: str


class Live(TypedDict):
    data: Any
    meta: dict


class LiveLoadMatch(TypedDict):
    live_id: str
    trip_id: str


class Meta(TypedDict):
    breaker_state: str
    enabled: bool
    source: str


class MetaLoadMatch(TypedDict, total=False):
    breaker_state: str
    enabled: bool
    source: str


class MetaListMatch(TypedDict, total=False):
    limit: int
    offset: int
    order: str
    page: int
    sort: str


class Station(TypedDict):
    data: Any
    meta: dict


class StationLoadMatch(TypedDict):
    eva: str
