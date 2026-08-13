# InfranodeOpenData TypeScript SDK



The TypeScript SDK for the InfranodeOpenData API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.City()` — each with a small set of operations (`list`, `load`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/infranode-open-data-sdk/releases](https://github.com/voxgig-sdk/infranode-open-data-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { InfranodeOpenDataSDK } from '@voxgig-sdk/infranode-open-data'

const client = new InfranodeOpenDataSDK()
```

### 2. List city records

`list()` resolves to an array of City ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const citys = await client.City().list()

for (const city of citys) {
  console.log(city)
}
```

### 3. Load a station

Station is nested under eva, so provide the `eva`.
`load()` returns the entity directly and throws on failure:

```ts
try {
  const station = await client.Station().load({
    eva: 'example_eva',
  })
  console.log(station)
} catch (err) {
  console.error('load failed:', err)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const citys = await client.City().list()
  console.log(citys)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = InfranodeOpenDataSDK.test()

const city = await client.City().list()
// city is the entity, populated with mock response data
// — call city.data() for the record itself
console.log(city)
```

You can also use the instance method:

```ts
const client = new InfranodeOpenDataSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.City()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new InfranodeOpenDataSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
INFRANODE_OPEN_DATA_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### InfranodeOpenDataSDK

#### Constructor

```ts
new InfranodeOpenDataSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `City(data?)` | `CityEntity` | Create a City entity instance. |
| `Compare(data?)` | `CompareEntity` | Create a Compare entity instance. |
| `Health(data?)` | `HealthEntity` | Create a Health entity instance. |
| `Live(data?)` | `LiveEntity` | Create a Live entity instance. |
| `Meta(data?)` | `MetaEntity` | Create a Meta entity instance. |
| `Station(data?)` | `StationEntity` | Create a Station entity instance. |
| `tester(testopts?, sdkopts?)` | `InfranodeOpenDataSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `InfranodeOpenDataSDK.test(testopts?, sdkopts?)` | `InfranodeOpenDataSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): InfranodeOpenDataSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` resolves to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### City

| Field | Description |
| --- | --- |
| `data` |  |
| `meta` |  |

Operations: list, load.

API path: `/api/v1/cities`

#### Compare

| Field | Description |
| --- | --- |
| `city` |  |
| `data` |  |
| `source_status` |  |

Operations: list.

API path: `/api/v1/compare`

#### Health

| Field | Description |
| --- | --- |
| `redis` |  |
| `status` |  |
| `version` |  |

Operations: load.

API path: `/api/v1/health`

#### Live

| Field | Description |
| --- | --- |
| `data` |  |
| `meta` |  |

Operations: load.

API path: `/api/v1/live/{city}/transit/routes/{route_id}/status`

#### Meta

| Field | Description |
| --- | --- |
| `breaker_state` |  |
| `enabled` |  |
| `source` |  |

Operations: list, load.

API path: `/api/v1/sources`

#### Station

| Field | Description |
| --- | --- |
| `data` |  |
| `meta` |  |

Operations: load.

API path: `/api/v1/stations/{eva}/arrivals`



## Entities


### City

Create an instance: `const city = client.City()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `any` |  |
| `meta` | `Record<string, any>` |  |

#### Example: Load

```ts
const city = await client.City().load({ id: 'city_id' })
```

#### Example: List

```ts
const citys = await client.City().list()
```


### Compare

Create an instance: `const compare = client.Compare()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `city` | `string` |  |
| `data` | `Record<string, any>` |  |
| `source_status` | `string` |  |

#### Example: List

```ts
const compares = await client.Compare().list()
```


### Health

Create an instance: `const health = client.Health()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `redis` | `boolean` |  |
| `status` | `string` |  |
| `version` | `string` |  |

#### Example: Load

```ts
const health = await client.Health().load()
```


### Live

Create an instance: `const live = client.Live()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `any` |  |
| `meta` | `Record<string, any>` |  |

#### Example: Load

```ts
const live = await client.Live().load()
```


### Meta

Create an instance: `const meta = client.Meta()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `breaker_state` | `string` |  |
| `enabled` | `boolean` |  |
| `source` | `string` |  |

#### Example: Load

```ts
const meta = await client.Meta().load()
```

#### Example: List

```ts
const metas = await client.Meta().list()
```


### Station

Create an instance: `const station = client.Station()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `any` |  |
| `meta` | `Record<string, any>` |  |

#### Example: Load

```ts
const station = await client.Station().load({ eva: 'eva' })
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
infranode-open-data/
├── src/
│   ├── InfranodeOpenDataSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { InfranodeOpenDataSDK } from '@voxgig-sdk/infranode-open-data'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const city = client.City()
await city.list()

// city.data() now returns the city data from the last `list`
// city.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
