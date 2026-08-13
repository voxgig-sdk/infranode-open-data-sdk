# InfranodeOpenData TypeScript SDK Reference

Complete API reference for the InfranodeOpenData TypeScript SDK.


## InfranodeOpenDataSDK

### Constructor

```ts
new InfranodeOpenDataSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `InfranodeOpenDataSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = InfranodeOpenDataSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `InfranodeOpenDataSDK` instance in test mode.


### Instance Methods

#### `City(data?: object)`

Create a new `City` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CityEntity` instance.

#### `Compare(data?: object)`

Create a new `Compare` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CompareEntity` instance.

#### `Health(data?: object)`

Create a new `Health` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `HealthEntity` instance.

#### `Live(data?: object)`

Create a new `Live` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LiveEntity` instance.

#### `Meta(data?: object)`

Create a new `Meta` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MetaEntity` instance.

#### `Station(data?: object)`

Create a new `Station` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `StationEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `InfranodeOpenDataSDK.test()`.

**Returns:** `InfranodeOpenDataSDK` instance in test mode.


---

## CityEntity

```ts
const city = client.City()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `any` | Yes |  |
| `meta` | `Record<string, any>` | Yes |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `accident` | `/api/v1/cities/{slug}/accidents` | `client.City().load({ $action: 'accident', ... })` |
| `air` | `/api/v1/cities/{slug}/air` | `client.City().load({ $action: 'air', ... })` |
| `air_uba` | `/api/v1/cities/{slug}/air-uba` | `client.City().load({ $action: 'air_uba', ... })` |
| `base` | `/api/v1/cities/{slug}/base` | `client.City().load({ $action: 'base', ... })` |
| `bathing_water` | `/api/v1/cities/{slug}/bathing-water` | `client.City().load({ $action: 'bathing_water', ... })` |
| `bike_count` | `/api/v1/cities/{slug}/bike-counts` | `client.City().load({ $action: 'bike_count', ... })` |
| `business_registration` | `/api/v1/cities/{slug}/business-registrations` | `client.City().load({ $action: 'business_registration', ... })` |
| `charging` | `/api/v1/cities/{slug}/charging` | `client.City().load({ $action: 'charging', ... })` |
| `charging_status` | `/api/v1/cities/{slug}/charging-status` | `client.City().load({ $action: 'charging_status', ... })` |
| `civil_protection_warning` | `/api/v1/cities/{slug}/civil-protection-warnings` | `client.City().load({ $action: 'civil_protection_warning', ... })` |
| `construction` | `/api/v1/cities/{slug}/construction` | `client.City().load({ $action: 'construction', ... })` |
| `council_paper` | `/api/v1/cities/{slug}/council-papers` | `client.City().load({ $action: 'council_paper', ... })` |
| `crime_stat` | `/api/v1/cities/{slug}/crime-stats` | `client.City().load({ $action: 'crime_stat', ... })` |
| `demographic` | `/api/v1/cities/{slug}/demographics` | `client.City().load({ $action: 'demographic', ... })` |
| `district_heating` | `/api/v1/cities/{slug}/district-heating` | `client.City().load({ $action: 'district_heating', ... })` |
| `drinking_water` | `/api/v1/cities/{slug}/drinking-water` | `client.City().load({ $action: 'drinking_water', ... })` |
| `education` | `/api/v1/cities/{slug}/education` | `client.City().load({ $action: 'education', ... })` |
| `election` | `/api/v1/cities/{slug}/election` | `client.City().load({ $action: 'election', ... })` |
| `energy` | `/api/v1/cities/{slug}/energy` | `client.City().load({ $action: 'energy', ... })` |
| `event` | `/api/v1/cities/{slug}/events` | `client.City().load({ $action: 'event', ... })` |
| `fire_danger` | `/api/v1/cities/{slug}/fire-danger` | `client.City().load({ $action: 'fire_danger', ... })` |
| `flood` | `/api/v1/cities/{slug}/flood` | `client.City().load({ $action: 'flood', ... })` |
| `fuel_price` | `/api/v1/cities/{slug}/fuel-prices` | `client.City().load({ $action: 'fuel_price', ... })` |
| `geo` | `/api/v1/cities/{slug}/geo` | `client.City().load({ $action: 'geo', ... })` |
| `government_office` | `/api/v1/cities/{slug}/government-offices` | `client.City().load({ $action: 'government_office', ... })` |
| `health` | `/api/v1/cities/{slug}/health` | `client.City().load({ $action: 'health', ... })` |
| `heritage` | `/api/v1/cities/{slug}/heritage` | `client.City().load({ $action: 'heritage', ... })` |
| `holiday` | `/api/v1/cities/{slug}/holidays` | `client.City().load({ $action: 'holiday', ... })` |
| `hospitals_atla` | `/api/v1/cities/{slug}/hospitals-atlas` | `client.City().load({ $action: 'hospitals_atla', ... })` |
| `icu_live` | `/api/v1/cities/{slug}/icu-live` | `client.City().load({ $action: 'icu_live', ... })` |
| `indicator` | `/api/v1/cities/{slug}/indicators` | `client.City().load({ $action: 'indicator', ... })` |
| `insolvency` | `/api/v1/cities/{slug}/insolvencies` | `client.City().load({ $action: 'insolvency', ... })` |
| `land_value` | `/api/v1/cities/{slug}/land-values` | `client.City().load({ $action: 'land_value', ... })` |
| `market` | `/api/v1/cities/{slug}/markets` | `client.City().load({ $action: 'market', ... })` |
| `office_wait_time` | `/api/v1/cities/{slug}/office-wait-times` | `client.City().load({ $action: 'office_wait_time', ... })` |
| `overview` | `/api/v1/cities/{slug}/overview` | `client.City().load({ $action: 'overview', ... })` |
| `parcel_locker` | `/api/v1/cities/{slug}/parcel-lockers` | `client.City().load({ $action: 'parcel_locker', ... })` |
| `parking` | `/api/v1/cities/{slug}/parking` | `client.City().load({ $action: 'parking', ... })` |
| `playground` | `/api/v1/cities/{slug}/playgrounds` | `client.City().load({ $action: 'playground', ... })` |
| `poi` | `/api/v1/cities/{slug}/pois` | `client.City().load({ $action: 'poi', ... })` |
| `pollen_uv` | `/api/v1/cities/{slug}/pollen-uv` | `client.City().load({ $action: 'pollen_uv', ... })` |
| `population_density` | `/api/v1/cities/{slug}/population-density` | `client.City().load({ $action: 'population_density', ... })` |
| `post_box` | `/api/v1/cities/{slug}/post-boxes` | `client.City().load({ $action: 'post_box', ... })` |
| `post_office` | `/api/v1/cities/{slug}/post-offices` | `client.City().load({ $action: 'post_office', ... })` |
| `power_load` | `/api/v1/cities/{slug}/power-load` | `client.City().load({ $action: 'power_load', ... })` |
| `power_price` | `/api/v1/cities/{slug}/power-price` | `client.City().load({ $action: 'power_price', ... })` |
| `public_tender` | `/api/v1/cities/{slug}/public-tenders` | `client.City().load({ $action: 'public_tender', ... })` |
| `public_toilet` | `/api/v1/cities/{slug}/public-toilets` | `client.City().load({ $action: 'public_toilet', ... })` |
| `public_wifi` | `/api/v1/cities/{slug}/public-wifi` | `client.City().load({ $action: 'public_wifi', ... })` |
| `recycling_centre` | `/api/v1/cities/{slug}/recycling-centres` | `client.City().load({ $action: 'recycling_centre', ... })` |
| `road_event` | `/api/v1/cities/{slug}/road-events` | `client.City().load({ $action: 'road_event', ... })` |
| `sharing` | `/api/v1/cities/{slug}/sharing` | `client.City().load({ $action: 'sharing', ... })` |
| `solar` | `/api/v1/cities/{slug}/solar` | `client.City().load({ $action: 'solar', ... })` |
| `solar_roof` | `/api/v1/cities/{slug}/solar-roofs` | `client.City().load({ $action: 'solar_roof', ... })` |
| `station` | `/api/v1/cities/{slug}/stations` | `client.City().load({ $action: 'station', ... })` |
| `station_arrival` | `/api/v1/cities/{slug}/station-arrivals` | `client.City().load({ $action: 'station_arrival', ... })` |
| `station_departure` | `/api/v1/cities/{slug}/station-departures` | `client.City().load({ $action: 'station_departure', ... })` |
| `station_facility` | `/api/v1/cities/{slug}/station-facilities` | `client.City().load({ $action: 'station_facility', ... })` |
| `tax_rate` | `/api/v1/cities/{slug}/tax-rates` | `client.City().load({ $action: 'tax_rate', ... })` |
| `tourism` | `/api/v1/cities/{slug}/tourism` | `client.City().load({ $action: 'tourism', ... })` |
| `traffic` | `/api/v1/cities/{slug}/traffic` | `client.City().load({ $action: 'traffic', ... })` |
| `transit` | `/api/v1/cities/{slug}/transit` | `client.City().load({ $action: 'transit', ... })` |
| `tree_cadastre` | `/api/v1/cities/{slug}/tree-cadastre` | `client.City().load({ $action: 'tree_cadastre', ... })` |
| `unemployment` | `/api/v1/cities/{slug}/unemployment` | `client.City().load({ $action: 'unemployment', ... })` |
| `vehicle_registration` | `/api/v1/cities/{slug}/vehicle-registrations` | `client.City().load({ $action: 'vehicle_registration', ... })` |
| `water_level` | `/api/v1/cities/{slug}/water-level` | `client.City().load({ $action: 'water_level', ... })` |
| `weather` | `/api/v1/cities/{slug}/weather` | `client.City().load({ $action: 'weather', ... })` |
| `weather_warning` | `/api/v1/cities/{slug}/weather-warnings` | `client.City().load({ $action: 'weather_warning', ... })` |
| `webcam` | `/api/v1/cities/{slug}/webcams` | `client.City().load({ $action: 'webcam', ... })` |

An action returns that action's OWN response, which is not necessarily a
City record — check the API definition for its shape.

```ts
const result = await client.City().load({
  $action: 'accident',
  /* ...the action's own arguments */
})
```

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.City().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.City().load({ id: 'city_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CityEntity` instance with the same client and
options.

#### `client()`

Return the parent `InfranodeOpenDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CompareEntity

```ts
const compare = client.Compare()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `city` | `string` | Yes |  |
| `data` | `Record<string, any>` | No |  |
| `source_status` | `string` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Compare().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CompareEntity` instance with the same client and
options.

#### `client()`

Return the parent `InfranodeOpenDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## HealthEntity

```ts
const health = client.Health()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `redis` | `boolean` | Yes |  |
| `status` | `string` | Yes |  |
| `version` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Health().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `HealthEntity` instance with the same client and
options.

#### `client()`

Return the parent `InfranodeOpenDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LiveEntity

```ts
const live = client.Live()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `any` | Yes |  |
| `meta` | `Record<string, any>` | Yes |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `air` | `/api/v1/live/{slug}/air` | `client.Live().load({ $action: 'air', ... })` |
| `air_uba` | `/api/v1/live/{slug}/air-uba` | `client.Live().load({ $action: 'air_uba', ... })` |
| `baustellen` | `/api/v1/live/{city}/baustellen` | `client.Live().load({ $action: 'baustellen', ... })` |
| `departure` | `/api/v1/live/{slug}/departures` | `client.Live().load({ $action: 'departure', ... })` |
| `ereignisse` | `/api/v1/live/{city}/ereignisse` | `client.Live().load({ $action: 'ereignisse', ... })` |
| `flood` | `/api/v1/live/{slug}/flood` | `client.Live().load({ $action: 'flood', ... })` |
| `traffic` | `/api/v1/live/{slug}/traffic` | `client.Live().load({ $action: 'traffic', ... })` |
| `traffic_flow` | `/api/v1/live/{city}/traffic-flow` | `client.Live().load({ $action: 'traffic_flow', ... })` |
| `transit_departure` | `/api/v1/live/{city}/transit/departures` | `client.Live().load({ $action: 'transit_departure', ... })` |
| `water_level` | `/api/v1/live/{slug}/water-level` | `client.Live().load({ $action: 'water_level', ... })` |
| `webcam` | `/api/v1/live/{slug}/webcams` | `client.Live().load({ $action: 'webcam', ... })` |

An action returns that action's OWN response, which is not necessarily a
Live record — check the API definition for its shape.

```ts
const result = await client.Live().load({
  $action: 'air',
  /* ...the action's own arguments */
})
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Live().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LiveEntity` instance with the same client and
options.

#### `client()`

Return the parent `InfranodeOpenDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MetaEntity

```ts
const meta = client.Meta()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `breaker_state` | `string` | Yes |  |
| `enabled` | `boolean` | Yes |  |
| `source` | `string` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Meta().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Meta().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MetaEntity` instance with the same client and
options.

#### `client()`

Return the parent `InfranodeOpenDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## StationEntity

```ts
const station = client.Station()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `any` | Yes |  |
| `meta` | `Record<string, any>` | Yes |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `arrival` | `/api/v1/stations/{eva}/arrivals` | `client.Station().load({ $action: 'arrival', ... })` |
| `departure` | `/api/v1/stations/{eva}/departures` | `client.Station().load({ $action: 'departure', ... })` |

An action returns that action's OWN response, which is not necessarily a
Station record — check the API definition for its shape.

```ts
const result = await client.Station().load({
  $action: 'arrival',
  /* ...the action's own arguments */
})
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Station().load({ eva: 'eva' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `StationEntity` instance with the same client and
options.

#### `client()`

Return the parent `InfranodeOpenDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new InfranodeOpenDataSDK({
  feature: {
    test: { active: true },
  }
})
```

