

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


loadEnvLocal(__dirname + '/../../../.env.local')


describe('MetaEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when INFRANODE_OPEN_DATA_TEST_LIVE=TRUE.
  afterEach(liveDelay('INFRANODE_OPEN_DATA_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = InfranodeOpenDataSDK.test()
    const ent = testsdk.Meta()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.INFRANODE_OPEN_DATA_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'meta.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"breaker_state":{"a":true,"h":"Breaker State","n":"breaker_state","r":true,"t":"`$STRING`","key$":"breaker_state","index$":0},"enabled":{"a":true,"h":"Enabled","n":"enabled","r":true,"t":"`$BOOLEAN`","key$":"enabled","index$":1},"source":{"a":true,"h":"Source","n":"source","r":true,"t":"`$STRING`","key$":"source","index$":2}},"name":"meta","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /api/v1/sources","source":"openapi3","version":2},"g":{"header":[{"a":true,"k":"header","n":"if_none_match","or":"if_none_match","r":false,"t":"`$STRING`","index$":0}],"query":[{"a":true,"ex":50,"k":"query","n":"limit","or":"limit","r":false,"t":"`$INTEGER`","index$":0},{"a":true,"ex":0,"k":"query","n":"offset","or":"offset","r":false,"t":"`$INTEGER`","index$":1},{"a":true,"ex":"asc","k":"query","n":"order","or":"order","r":false,"t":"`$STRING`","index$":2},{"a":true,"ex":1,"k":"query","n":"page","or":"page","r":false,"t":"`$INTEGER`","index$":3},{"a":true,"k":"query","n":"sort","or":"sort","r":false,"t":"`$STRING`","index$":4}]},"k":"http","m":"GET","o":"/api/v1/sources","q":{"exist":["if_none_match","limit","offset","order","page","sort"]},"r":{},"s":[{"lit":"api"},{"lit":"v1"},{"lit":"sources"}],"t":{"req":"`reqdata`","res":"`body.meta`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /api/v1/openapi.yaml","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/api/v1/openapi.yaml","q":{},"r":{},"s":[{"lit":"api"},{"lit":"v1"},{"lit":"openapi.yaml"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"meta","name__orig":"meta","Name":"Meta","name_":"meta","name-":"meta","NAME":"META","index$":4}, {"active":true,"entity":"meta","key$":"BasicMetaFlow","kind":"basic","name":"BasicMetaFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"meta_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"meta_ref01","srcdatavar":"meta_ref01_data","suffix":"_dt0"},"m":{},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-meta_ref01"}}],"index$":1}]}, 'Meta', {"GET /api/v1/sources":{"protocol":"http","operationId":"getSources","responses":{"200":{"description":"Liste der Quellen-Status","headers":{"ETag":{"description":"Stabiler Entity-Tag des Response-Bodys (sha256-basiert). Nur auf erfolgreichen GET-Reads (200) gesetzt, nie auf Fehler-Envelopes/503.","schema":{"type":"string"},"x-ref":"#/components/headers/ETag"},"Cache-Control":{"description":"Cache-Control je Ressource (\"public, max-age=<ttl>, stale-while-revalidate=<ttl>, stale-if-error=<ttl>\"). TTL aus der serverseitigen CACHE_TTL-Map (z.B. wikidata 86400 s, dwd 1800 s, uba 600 s, default 300 s). stale-while-revalidate/stale-if-error erlauben einem Shared Cache (Cloudflare), bei Ablauf bzw. Origin-Fehler kurz die letzte gute Antwort weiterzuliefern. Echtzeit-Endpunkte (/api/v1/live/*) liefern stattdessen \"no-store\".","schema":{"type":"string"},"x-ref":"#/components/headers/CacheControl"}},"content":{"application/json":{"schema":{"type":"object","required":["data","meta"],"properties":{"data":{"items":{"properties":{"breaker_state":{"enum":["CLOSED","OPEN","HALF_OPEN"],"type":"string","key$":"breaker_state"},"enabled":{"type":"boolean","key$":"enabled"},"source":{"example":"wikidata","type":"string","key$":"source"}},"required":["source","enabled","breaker_state"],"type":"object","x-ref":"#/components/schemas/SourceStatus","index$":0},"key$":"data","type":"array"},"meta":{"key$":"meta","properties":{"cache_status":{"description":"Cache-Herkunft (hit/miss/stale), sofern die Route cacht.","nullable":true,"type":"string"},"correlation_id":{"nullable":true,"type":"string"},"covered_cities":{"description":"Nur bei source_status=\"not_covered\": die Stadt-Slugs, die dieser teilabgedeckte Endpunkt tatsächlich bedient (z. B. flood, webcams, traffic, road-events).","items":{"type":"string"},"type":"array"},"pagination":{"description":"Bei paginierbaren Datenart-Listen (charging, energy, events, transit, OSM-Feature-Endpunkte): der EHRLICH ausgewiesene Ausschnitt (keine stille Kappung). KANAL-ABHÄNGIGER Default für gleiche URL: direktes REST liefert die volle Liste (limit=null, returned==total, truncated=false); GPT-Actions (erkannt am OpenAI-Header) bekommen ein serverseitig gebundenes Default-Limit; MCP-Aufrufe sind gebunden, weil der MCP-Client selbst ein Default-Limit setzt. Kanalunabhängige Overrides: limit=all (oder ?all=1) erzwingt die volle Liste, limit + offset blättern gezielt. meta.pagination ist bei ALLEN Kanälen gesetzt.","properties":{"limit":{"description":"Angewandtes Seiten-Limit; null bei Vollausgabe (REST-Default oder limit=all).","nullable":true,"type":"integer"},"offset":{"description":"Start-Offset der Seite.","type":"integer"},"returned":{"description":"Zahl der auf dieser Seite ausgelieferten Einträge.","type":"integer"},"total":{"description":"Gesamtzahl der Einträge (voller Bestand).","type":"integer"},"truncated":{"description":"true, wenn hinter dieser Seite noch Einträge liegen (offset + returned < total); bei Vollausgabe false.","type":"boolean"}},"type":"object","x-description-en":"For paginable data-type lists (charging, energy, events, transit, OSM feature endpoints): the honestly reported slice (no silent capping). CHANNEL-DEPENDENT default for the same URL: direct REST returns the full list (limit=null, returned==total, truncated=false); GPT Actions (detected via the OpenAI header) get a server-side bound default limit; MCP calls are bound because the MCP client sets a default limit itself. Channel-independent overrides: limit=all (or ?all=1) forces the full list, limit + offset page explicitly. meta.pagination is set on ALL channels."},"source_status":{"description":"Ehrlicher Quellen-Status der Antwort. \"ok\" (Daten vorhanden), \"no_data\" (Quelle erreichbar/abgedeckt, aber gerade keine Daten), \"disabled\" (Quelle per Toggle aus), \"not_ingested\" (kein Snapshot), \"not_covered\" (Stadt ist für diesen teilabgedeckten Endpunkt strukturell nicht abgedeckt; data=null, siehe covered_cities). Klar unterscheidbar vom 404 (Stadt unbekannt).","enum":["ok","no_data","disabled","not_ingested","not_covered"],"type":"string"}},"type":"object","x-ref":"#/components/schemas/Meta"}}}}}},"304":{"description":"Not Modified. If-None-Match stimmte mit dem aktuellen ETag überein; es wird kein Body geliefert (ETag + Cache-Control bleiben erhalten).","headers":{"ETag":{"description":"Stabiler Entity-Tag des Response-Bodys (sha256-basiert). Nur auf erfolgreichen GET-Reads (200) gesetzt, nie auf Fehler-Envelopes/503.","schema":{"type":"string"},"x-ref":"#/components/headers/ETag"},"Cache-Control":{"description":"Cache-Control je Ressource (\"public, max-age=<ttl>, stale-while-revalidate=<ttl>, stale-if-error=<ttl>\"). TTL aus der serverseitigen CACHE_TTL-Map (z.B. wikidata 86400 s, dwd 1800 s, uba 600 s, default 300 s). stale-while-revalidate/stale-if-error erlauben einem Shared Cache (Cloudflare), bei Ablauf bzw. Origin-Fehler kurz die letzte gute Antwort weiterzuliefern. Echtzeit-Endpunkte (/api/v1/live/*) liefern stattdessen \"no-store\".","schema":{"type":"string"},"x-ref":"#/components/headers/CacheControl"}},"x-ref":"#/components/responses/NotModified"}},"parameters":[{"in":"header","name":"If-None-Match","required":false,"schema":{"type":"string"},"description":"Conditional GET. Stimmt der Wert mit dem aktuellen ETag überein, antwortet der Server mit 304 Not Modified ohne Body.","x-description-en":"Conditional GET. If the value matches the current ETag, the server responds with 304 Not Modified and no body.","index$":0},{"in":"query","name":"page","schema":{"type":"integer","minimum":1,"default":1},"index$":1},{"in":"query","name":"limit","schema":{"type":"integer","minimum":1,"default":50},"description":"Seitengröße. Wird auf MAX_LIMIT (200) gedeckelt: zu große Werte liefern eine 200er-Seite mit 200 Einträgen statt eines Fehlers.","x-description-en":"Page size. Capped at MAX_LIMIT (200): larger values return a 200 page of 200 items instead of an error.","index$":2},{"in":"query","name":"offset","schema":{"type":"integer","minimum":0,"default":0},"description":"Ein zu großer Offset liefert eine leere Seite (200), nie 500.","x-description-en":"An offset past the end returns an empty page (200), never 500.","index$":3},{"in":"query","name":"sort","schema":{"type":"string","enum":["source","enabled","license"]},"description":"Nur die gelisteten Felder sind erlaubt; ein unbekanntes Feld wird mit 400 (invalid_request) abgewiesen, bevor es ausgewertet wird.","x-description-en":"Only the listed fields are allowed; an unknown field is rejected with 400 (invalid_request) before it is evaluated.","index$":4},{"in":"query","name":"order","schema":{"type":"string","enum":["asc","desc"],"default":"asc"},"index$":5}],"securitySource":"unspecified"},"GET /api/v1/openapi.yaml":{"protocol":"http","operationId":"getOpenapiYaml","responses":{"200":{"description":"Die Spec als YAML","content":{"application/yaml":{"schema":{"type":"string"}}}}},"parameters":[],"securitySource":"unspecified"}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let meta_ref01_data = Object.values(setup.data.existing.meta)[0] as any

    // LIST
    const meta_ref01_ent = client.Meta()
    const meta_ref01_match: any = {}

    const meta_ref01_list = (await meta_ref01_ent.list(meta_ref01_match)).map((e: any) => e.data())


    // LOAD
    const meta_ref01_match_dt0: any = {}
    const meta_ref01_data_dt0 = (await meta_ref01_ent.load(meta_ref01_match_dt0)).data()
    assert(null != meta_ref01_data_dt0)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/meta/MetaTestData.json')

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
    ['meta01','meta02','meta03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'INFRANODE_OPEN_DATA_TEST_META_ENTID': idmap,
    'INFRANODE_OPEN_DATA_TEST_LIVE': 'FALSE',
    'INFRANODE_OPEN_DATA_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['INFRANODE_OPEN_DATA_TEST_META_ENTID']

  const live = 'TRUE' === env.INFRANODE_OPEN_DATA_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['INFRANODE_OPEN_DATA_TEST_META_ENTID']
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
  
