

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { InfranodeOpenDataSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('CompareEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when INFRANODE_OPEN_DATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('INFRANODE_OPEN_DATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = InfranodeOpenDataSDK.test()
    const ent = testsdk.Compare()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.INFRANODE_OPEN_DATA_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'compare.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"city","req":true,"type":"`$STRING`","index$":0},{"active":true,"name":"data","req":false,"type":"`$OBJECT`","index$":1},{"active":true,"name":"source_status","req":true,"type":"`$STRING`","index$":2}],"name":"compare","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"header":[{"active":true,"kind":"header","name":"if_none_match","orig":"if_none_match","reqd":false,"type":"`$STRING`"}],"query":[{"active":true,"example":"berlin,koeln,hamburg","kind":"query","name":"city","orig":"city","reqd":true,"type":"`$STRING`","index$":0},{"active":true,"example":50,"kind":"query","name":"limit","orig":"limit","reqd":false,"type":"`$INTEGER`","index$":1},{"active":true,"example":0,"kind":"query","name":"offset","orig":"offset","reqd":false,"type":"`$INTEGER`","index$":2},{"active":true,"example":"asc","kind":"query","name":"order","orig":"order","reqd":false,"type":"`$STRING`","index$":3},{"active":true,"example":1,"kind":"query","name":"page","orig":"page","reqd":false,"type":"`$INTEGER`","index$":4},{"active":true,"kind":"query","name":"resource","orig":"resource","reqd":true,"type":"`$STRING`","index$":5},{"active":true,"kind":"query","name":"sort","orig":"sort","reqd":false,"type":"`$STRING`","index$":6}]},"contract":{"id":"GET /api/v1/compare","json":"{\"operationId\":\"compareCities\",\"parameters\":[{\"description\":\"Kommaseparierte Stadt-Slugs (auf MAX_CITIES begrenzt).\",\"example\":\"berlin,koeln,hamburg\",\"in\":\"query\",\"name\":\"cities\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Zu vergleichende Ressource; unbekannt -> 400 invalid_request. weather/air = Live-Adapter (DWD/UBA); indicators, demographics, unemployment, tourism, charging-status und weather-warnings delegieren an die jeweiligen Stadt-Endpunkte (identisches Verhalten inkl. Attribution). charging-status liefert im Compare nur die Aggregate (points_omitted statt Einzelpunkt-Liste).\",\"in\":\"query\",\"name\":\"resource\",\"required\":true,\"schema\":{\"enum\":[\"weather\",\"air\",\"indicators\",\"demographics\",\"unemployment\",\"tourism\",\"charging-status\",\"weather-warnings\"],\"type\":\"string\"}},{\"in\":\"query\",\"name\":\"page\",\"schema\":{\"default\":1,\"minimum\":1,\"type\":\"integer\"}},{\"in\":\"query\",\"name\":\"limit\",\"schema\":{\"default\":50,\"minimum\":1,\"type\":\"integer\"}},{\"in\":\"query\",\"name\":\"offset\",\"schema\":{\"default\":0,\"minimum\":0,\"type\":\"integer\"}},{\"in\":\"query\",\"name\":\"sort\",\"schema\":{\"enum\":[\"city\",\"source_status\"],\"type\":\"string\"}},{\"in\":\"query\",\"name\":\"order\",\"schema\":{\"default\":\"asc\",\"enum\":[\"asc\",\"desc\"],\"type\":\"string\"}},{\"description\":\"Conditional GET; If-None-Match == ETag -> 304 Not Modified ohne Body.\",\"in\":\"header\",\"name\":\"If-None-Match\",\"required\":false,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"data\":{\"items\":{\"properties\":{\"city\":{\"type\":\"string\"},\"data\":{\"nullable\":true,\"type\":\"object\"},\"source_status\":{\"enum\":[\"ok\",\"disabled\",\"no_data\",\"error\",\"not_found\"],\"type\":\"string\"}},\"required\":[\"city\",\"source_status\"],\"type\":\"object\"},\"type\":\"array\"},\"meta\":{\"properties\":{\"cache_status\":{\"description\":\"Cache-Herkunft (hit/miss/stale), sofern die Route cacht.\",\"nullable\":true,\"type\":\"string\"},\"correlation_id\":{\"nullable\":true,\"type\":\"string\"},\"covered_cities\":{\"description\":\"Nur bei source_status=\\\"not_covered\\\": die Stadt-Slugs, die dieser teilabgedeckte Endpunkt tatsächlich bedient (z. B. flood, webcams, traffic, road-events).\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pagination\":{\"description\":\"Bei paginierbaren Datenart-Listen (charging, energy, events, transit, OSM-Feature-Endpunkte): der EHRLICH ausgewiesene Ausschnitt (keine stille Kappung). KANAL-ABHÄNGIGER Default für gleiche URL: direktes REST liefert die volle Liste (limit=null, returned==total, truncated=false); GPT-Actions (erkannt am OpenAI-Header) bekommen ein serverseitig gebundenes Default-Limit; MCP-Aufrufe sind gebunden, weil der MCP-Client selbst ein Default-Limit setzt. Kanalunabhängige Overrides: limit=all (oder ?all=1) erzwingt die volle Liste, limit + offset blättern gezielt. meta.pagination ist bei ALLEN Kanälen gesetzt.\",\"properties\":{\"limit\":{\"description\":\"Angewandtes Seiten-Limit; null bei Vollausgabe (REST-Default oder limit=all).\",\"nullable\":true,\"type\":\"integer\"},\"offset\":{\"description\":\"Start-Offset der Seite.\",\"type\":\"integer\"},\"returned\":{\"description\":\"Zahl der auf dieser Seite ausgelieferten Einträge.\",\"type\":\"integer\"},\"total\":{\"description\":\"Gesamtzahl der Einträge (voller Bestand).\",\"type\":\"integer\"},\"truncated\":{\"description\":\"true, wenn hinter dieser Seite noch Einträge liegen (offset + returned < total); bei Vollausgabe false.\",\"type\":\"boolean\"}},\"type\":\"object\"},\"source_status\":{\"description\":\"Ehrlicher Quellen-Status der Antwort. \\\"ok\\\" (Daten vorhanden), \\\"no_data\\\" (Quelle erreichbar/abgedeckt, aber gerade keine Daten), \\\"disabled\\\" (Quelle per Toggle aus), \\\"not_ingested\\\" (kein Snapshot), \\\"not_covered\\\" (Stadt ist für diesen teilabgedeckten Endpunkt strukturell nicht abgedeckt; data=null, siehe covered_cities). Klar unterscheidbar vom 404 (Stadt unbekannt).\",\"enum\":[\"ok\",\"no_data\",\"disabled\",\"not_ingested\",\"not_covered\"],\"type\":\"string\"}},\"type\":\"object\"}},\"required\":[\"data\",\"meta\"],\"type\":\"object\"}}},\"description\":\"Vergleichsliste mit per-Stadt source_status\",\"headers\":{\"Cache-Control\":{\"description\":\"Cache-Control je Ressource (\\\"public, max-age=<ttl>, stale-while-revalidate=<ttl>, stale-if-error=<ttl>\\\"). TTL aus der serverseitigen CACHE_TTL-Map (z.B. wikidata 86400 s, dwd 1800 s, uba 600 s, default 300 s). stale-while-revalidate/stale-if-error erlauben einem Shared Cache (Cloudflare), bei Ablauf bzw. Origin-Fehler kurz die letzte gute Antwort weiterzuliefern. Echtzeit-Endpunkte (/api/v1/live/*) liefern stattdessen \\\"no-store\\\".\",\"schema\":{\"type\":\"string\"}},\"ETag\":{\"description\":\"Stabiler Entity-Tag des Response-Bodys (sha256-basiert). Nur auf erfolgreichen GET-Reads (200) gesetzt, nie auf Fehler-Envelopes/503.\",\"schema\":{\"type\":\"string\"}}}},\"304\":{\"description\":\"Not Modified. If-None-Match stimmte mit dem aktuellen ETag überein; es wird kein Body geliefert (ETag + Cache-Control bleiben erhalten).\",\"headers\":{\"Cache-Control\":{\"description\":\"Cache-Control je Ressource (\\\"public, max-age=<ttl>, stale-while-revalidate=<ttl>, stale-if-error=<ttl>\\\"). TTL aus der serverseitigen CACHE_TTL-Map (z.B. wikidata 86400 s, dwd 1800 s, uba 600 s, default 300 s). stale-while-revalidate/stale-if-error erlauben einem Shared Cache (Cloudflare), bei Ablauf bzw. Origin-Fehler kurz die letzte gute Antwort weiterzuliefern. Echtzeit-Endpunkte (/api/v1/live/*) liefern stattdessen \\\"no-store\\\".\",\"schema\":{\"type\":\"string\"}},\"ETag\":{\"description\":\"Stabiler Entity-Tag des Response-Bodys (sha256-basiert). Nur auf erfolgreichen GET-Reads (200) gesetzt, nie auf Fehler-Envelopes/503.\",\"schema\":{\"type\":\"string\"}}}},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"error\":{\"properties\":{\"code\":{\"example\":\"upstream_unavailable\",\"type\":\"string\"},\"hint\":{\"nullable\":true,\"type\":\"string\"},\"message\":{\"example\":\"Simulated upstream failure\",\"type\":\"string\"}},\"required\":[\"code\",\"message\"],\"type\":\"object\"},\"meta\":{\"properties\":{\"correlation_id\":{\"nullable\":true,\"type\":\"string\"},\"generated_at\":{\"format\":\"date-time\",\"type\":\"string\"}},\"type\":\"object\"}},\"required\":[\"error\",\"meta\"],\"type\":\"object\"}}},\"description\":\"Unbekannte resource oder leeres cities (invalid_request)\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/compare","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"compare"}],"select":{"exist":["city","if_none_match","limit","offset","order","page","resource","sort"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"compare","name__orig":"compare","Name":"Compare","name_":"compare","name-":"compare","NAME":"COMPARE","index$":1}, {"active":true,"entity":"compare","key$":"BasicCompareFlow","kind":"basic","name":"BasicCompareFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"compare_ref01"}}],"index$":0}]}, 'Compare')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let compare_ref01_data = Object.values(setup.data.existing.compare)[0] as any

    // LIST
    const compare_ref01_ent = client.Compare()
    const compare_ref01_match: any = {}

    const compare_ref01_list = (await compare_ref01_ent.list(compare_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/compare/CompareTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = InfranodeOpenDataSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['compare01','compare02','compare03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'INFRANODE_OPEN_DATA_TEST_COMPARE_ENTID': idmap,
    'INFRANODE_OPEN_DATA_TEST_LIVE': 'FALSE',
    'INFRANODE_OPEN_DATA_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['INFRANODE_OPEN_DATA_TEST_COMPARE_ENTID']

  const live = 'TRUE' === env.INFRANODE_OPEN_DATA_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['INFRANODE_OPEN_DATA_TEST_COMPARE_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new InfranodeOpenDataSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.INFRANODE_OPEN_DATA_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
