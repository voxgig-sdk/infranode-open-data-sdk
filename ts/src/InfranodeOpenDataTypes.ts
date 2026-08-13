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

  // Selects a custom action instead of the plain load:
  //   'accident' | 'air' | 'air_uba' | 'base' | 'bathing_water' | 'bike_count' | 'business_registration' | 'charging' | 'charging_status' | 'civil_protection_warning' | 'construction' | 'council_paper' | 'crime_stat' | 'demographic' | 'district_heating' | 'drinking_water' | 'education' | 'election' | 'energy' | 'event' | 'fire_danger' | 'flood' | 'fuel_price' | 'geo' | 'government_office' | 'health' | 'heritage' | 'holiday' | 'hospitals_atla' | 'icu_live' | 'indicator' | 'insolvency' | 'land_value' | 'market' | 'office_wait_time' | 'overview' | 'parcel_locker' | 'parking' | 'playground' | 'poi' | 'pollen_uv' | 'population_density' | 'post_box' | 'post_office' | 'power_load' | 'power_price' | 'public_tender' | 'public_toilet' | 'public_wifi' | 'recycling_centre' | 'road_event' | 'sharing' | 'solar' | 'solar_roof' | 'station' | 'station_arrival' | 'station_departure' | 'station_facility' | 'tax_rate' | 'tourism' | 'traffic' | 'transit' | 'tree_cadastre' | 'unemployment' | 'vehicle_registration' | 'water_level' | 'weather' | 'weather_warning' | 'webcam'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
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
  redis: boolean
  status: string
  version: string
}

export interface HealthLoadMatch {
  redis?: boolean
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

  // Selects a custom action instead of the plain load:
  //   'air' | 'air_uba' | 'baustellen' | 'departure' | 'ereignisse' | 'flood' | 'traffic' | 'traffic_flow' | 'transit_departure' | 'water_level' | 'webcam'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
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

  // Selects a custom action instead of the plain load:
  //   'arrival' | 'departure'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

