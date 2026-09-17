-- Typed models for the InfranodeOpenData SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class City
---@field data any
---@field id? string
---@field meta table

---@class CityLoadMatch
---@field id string

---@class CityListMatch
---@field data? any
---@field id? string
---@field meta? table

---@class Compare
---@field city string
---@field data? table
---@field source_status string

---@class CompareListMatch
---@field city string
---@field limit? number
---@field offset? number
---@field order? string
---@field page? number
---@field resource string
---@field sort? string

---@class Health
---@field redis boolean
---@field status string
---@field version string

---@class HealthLoadMatch
---@field redis? boolean
---@field status? string
---@field version? string

---@class Live
---@field data any
---@field meta table

---@class LiveLoadMatch
---@field live_id string
---@field trip_id string

---@class Meta
---@field breaker_state string
---@field enabled boolean
---@field source string

---@class MetaLoadMatch
---@field breaker_state? string
---@field enabled? boolean
---@field source? string

---@class MetaListMatch
---@field limit? number
---@field offset? number
---@field order? string
---@field page? number
---@field sort? string

---@class Station

---@class StationLoadMatch
---@field eva string

local M = {}

return M
