# frozen_string_literal: true

# Typed models for the InfranodeOpenData SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# City entity data model.
#
# @!attribute [rw] data
#   @return [Object]
#
# @!attribute [rw] meta
#   @return [Hash]
City = Struct.new(
  :data,
  :meta,
  keyword_init: true
)

# Request payload for City#load.
#
# @!attribute [rw] id
#   @return [String, nil]
CityLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for City#list.
#
# @!attribute [rw] data
#   @return [Object, nil]
#
# @!attribute [rw] meta
#   @return [Hash, nil]
CityListMatch = Struct.new(
  :data,
  :meta,
  keyword_init: true
)

# Compare entity data model.
#
# @!attribute [rw] city
#   @return [String]
#
# @!attribute [rw] data
#   @return [Hash, nil]
#
# @!attribute [rw] source_status
#   @return [String]
Compare = Struct.new(
  :city,
  :data,
  :source_status,
  keyword_init: true
)

# Request payload for Compare#list.
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] data
#   @return [Hash, nil]
#
# @!attribute [rw] source_status
#   @return [String, nil]
CompareListMatch = Struct.new(
  :city,
  :data,
  :source_status,
  keyword_init: true
)

# Health entity data model.
#
# @!attribute [rw] redi
#   @return [Boolean]
#
# @!attribute [rw] status
#   @return [String]
#
# @!attribute [rw] version
#   @return [String]
Health = Struct.new(
  :redi,
  :status,
  :version,
  keyword_init: true
)

# Request payload for Health#load.
#
# @!attribute [rw] redi
#   @return [Boolean, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] version
#   @return [String, nil]
HealthLoadMatch = Struct.new(
  :redi,
  :status,
  :version,
  keyword_init: true
)

# Live entity data model.
#
# @!attribute [rw] data
#   @return [Object]
#
# @!attribute [rw] meta
#   @return [Hash]
Live = Struct.new(
  :data,
  :meta,
  keyword_init: true
)

# Request payload for Live#load.
#
# @!attribute [rw] live_id
#   @return [String, nil]
#
# @!attribute [rw] route_id
#   @return [String, nil]
#
# @!attribute [rw] trip_id
#   @return [String, nil]
LiveLoadMatch = Struct.new(
  :live_id,
  :route_id,
  :trip_id,
  keyword_init: true
)

# Meta entity data model.
#
# @!attribute [rw] breaker_state
#   @return [String]
#
# @!attribute [rw] enabled
#   @return [Boolean]
#
# @!attribute [rw] source
#   @return [String]
Meta = Struct.new(
  :breaker_state,
  :enabled,
  :source,
  keyword_init: true
)

# Request payload for Meta#load.
#
# @!attribute [rw] breaker_state
#   @return [String, nil]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
MetaLoadMatch = Struct.new(
  :breaker_state,
  :enabled,
  :source,
  keyword_init: true
)

# Request payload for Meta#list.
#
# @!attribute [rw] breaker_state
#   @return [String, nil]
#
# @!attribute [rw] enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
MetaListMatch = Struct.new(
  :breaker_state,
  :enabled,
  :source,
  keyword_init: true
)

# Station entity data model.
#
# @!attribute [rw] data
#   @return [Object]
#
# @!attribute [rw] meta
#   @return [Hash]
Station = Struct.new(
  :data,
  :meta,
  keyword_init: true
)

# Request payload for Station#load.
#
# @!attribute [rw] eva
#   @return [String]
StationLoadMatch = Struct.new(
  :eva,
  keyword_init: true
)

