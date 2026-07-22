// Typed models for the InfranodeOpenData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface City {
  data: any
  meta: Record<string, any>
}

export interface CityLoadMatch {
  id?: string
}

export interface CityListMatch {
  data?: any
  meta?: Record<string, any>
}

export interface Compare {
  city: string
  data?: Record<string, any>
  source_status: string
}

export interface CompareListMatch {
  city?: string
  data?: Record<string, any>
  source_status?: string
}

export interface Health {
  redi: boolean
  status: string
  version: string
}

export interface HealthLoadMatch {
  redi?: boolean
  status?: string
  version?: string
}

export interface Live {
  data: any
  meta: Record<string, any>
}

export interface LiveLoadMatch {
  live_id?: string
  route_id?: string
  trip_id?: string
}

export interface Meta {
  breaker_state: string
  enabled: boolean
  source: string
}

export interface MetaLoadMatch {
  breaker_state?: string
  enabled?: boolean
  source?: string
}

export interface MetaListMatch {
  breaker_state?: string
  enabled?: boolean
  source?: string
}

export interface Station {
  data: any
  meta: Record<string, any>
}

export interface StationLoadMatch {
  eva: string
}

