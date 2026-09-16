"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('MetaEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when INFRANODE_OPEN_DATA_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('INFRANODE_OPEN_DATA_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.InfranodeOpenDataSDK.test();
        const ent = testsdk.Meta();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.INFRANODE_OPEN_DATA_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'meta.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "breaker_state", "req": true, "type": "`$STRING`", "index$": 0 }, { "active": true, "name": "enabled", "req": true, "type": "`$BOOLEAN`", "index$": 1 }, { "active": true, "name": "source", "req": true, "type": "`$STRING`", "index$": 2 }], "name": "meta", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": { "header": [{ "active": true, "kind": "header", "name": "if_none_match", "orig": "if_none_match", "reqd": false, "type": "`$STRING`" }], "query": [{ "active": true, "example": 50, "kind": "query", "name": "limit", "orig": "limit", "reqd": false, "type": "`$INTEGER`", "index$": 0 }, { "active": true, "example": 0, "kind": "query", "name": "offset", "orig": "offset", "reqd": false, "type": "`$INTEGER`", "index$": 1 }, { "active": true, "example": "asc", "kind": "query", "name": "order", "orig": "order", "reqd": false, "type": "`$STRING`", "index$": 2 }, { "active": true, "example": 1, "kind": "query", "name": "page", "orig": "page", "reqd": false, "type": "`$INTEGER`", "index$": 3 }, { "active": true, "kind": "query", "name": "sort", "orig": "sort", "reqd": false, "type": "`$STRING`", "index$": 4 }] }, "contract": { "id": "GET /api/v1/sources", "json": "{\"operationId\":\"getSources\",\"parameters\":[{\"description\":\"Conditional GET. Stimmt der Wert mit dem aktuellen ETag überein, antwortet der Server mit 304 Not Modified ohne Body.\",\"in\":\"header\",\"name\":\"If-None-Match\",\"required\":false,\"schema\":{\"type\":\"string\"}},{\"in\":\"query\",\"name\":\"page\",\"schema\":{\"default\":1,\"minimum\":1,\"type\":\"integer\"}},{\"description\":\"Seitengröße. Wird auf MAX_LIMIT (200) gedeckelt: zu große Werte liefern eine 200er-Seite mit 200 Einträgen statt eines Fehlers.\",\"in\":\"query\",\"name\":\"limit\",\"schema\":{\"default\":50,\"minimum\":1,\"type\":\"integer\"}},{\"description\":\"Ein zu großer Offset liefert eine leere Seite (200), nie 500.\",\"in\":\"query\",\"name\":\"offset\",\"schema\":{\"default\":0,\"minimum\":0,\"type\":\"integer\"}},{\"description\":\"Nur die gelisteten Felder sind erlaubt; ein unbekanntes Feld wird mit 400 (invalid_request) abgewiesen, bevor es ausgewertet wird.\",\"in\":\"query\",\"name\":\"sort\",\"schema\":{\"enum\":[\"source\",\"enabled\",\"license\"],\"type\":\"string\"}},{\"in\":\"query\",\"name\":\"order\",\"schema\":{\"default\":\"asc\",\"enum\":[\"asc\",\"desc\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"data\":{\"items\":{\"properties\":{\"breaker_state\":{\"enum\":[\"CLOSED\",\"OPEN\",\"HALF_OPEN\"],\"type\":\"string\"},\"enabled\":{\"type\":\"boolean\"},\"source\":{\"example\":\"wikidata\",\"type\":\"string\"}},\"required\":[\"source\",\"enabled\",\"breaker_state\"],\"type\":\"object\"},\"type\":\"array\"},\"meta\":{\"properties\":{\"cache_status\":{\"description\":\"Cache-Herkunft (hit/miss/stale), sofern die Route cacht.\",\"nullable\":true,\"type\":\"string\"},\"correlation_id\":{\"nullable\":true,\"type\":\"string\"},\"covered_cities\":{\"description\":\"Nur bei source_status=\\\"not_covered\\\": die Stadt-Slugs, die dieser teilabgedeckte Endpunkt tatsächlich bedient (z. B. flood, webcams, traffic, road-events).\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pagination\":{\"description\":\"Bei paginierbaren Datenart-Listen (charging, energy, events, transit, OSM-Feature-Endpunkte): der EHRLICH ausgewiesene Ausschnitt (keine stille Kappung). KANAL-ABHÄNGIGER Default für gleiche URL: direktes REST liefert die volle Liste (limit=null, returned==total, truncated=false); GPT-Actions (erkannt am OpenAI-Header) bekommen ein serverseitig gebundenes Default-Limit; MCP-Aufrufe sind gebunden, weil der MCP-Client selbst ein Default-Limit setzt. Kanalunabhängige Overrides: limit=all (oder ?all=1) erzwingt die volle Liste, limit + offset blättern gezielt. meta.pagination ist bei ALLEN Kanälen gesetzt.\",\"properties\":{\"limit\":{\"description\":\"Angewandtes Seiten-Limit; null bei Vollausgabe (REST-Default oder limit=all).\",\"nullable\":true,\"type\":\"integer\"},\"offset\":{\"description\":\"Start-Offset der Seite.\",\"type\":\"integer\"},\"returned\":{\"description\":\"Zahl der auf dieser Seite ausgelieferten Einträge.\",\"type\":\"integer\"},\"total\":{\"description\":\"Gesamtzahl der Einträge (voller Bestand).\",\"type\":\"integer\"},\"truncated\":{\"description\":\"true, wenn hinter dieser Seite noch Einträge liegen (offset + returned < total); bei Vollausgabe false.\",\"type\":\"boolean\"}},\"type\":\"object\"},\"source_status\":{\"description\":\"Ehrlicher Quellen-Status der Antwort. \\\"ok\\\" (Daten vorhanden), \\\"no_data\\\" (Quelle erreichbar/abgedeckt, aber gerade keine Daten), \\\"disabled\\\" (Quelle per Toggle aus), \\\"not_ingested\\\" (kein Snapshot), \\\"not_covered\\\" (Stadt ist für diesen teilabgedeckten Endpunkt strukturell nicht abgedeckt; data=null, siehe covered_cities). Klar unterscheidbar vom 404 (Stadt unbekannt).\",\"enum\":[\"ok\",\"no_data\",\"disabled\",\"not_ingested\",\"not_covered\"],\"type\":\"string\"}},\"type\":\"object\"}},\"required\":[\"data\",\"meta\"],\"type\":\"object\"}}},\"description\":\"Liste der Quellen-Status\",\"headers\":{\"Cache-Control\":{\"description\":\"Cache-Control je Ressource (\\\"public, max-age=<ttl>, stale-while-revalidate=<ttl>, stale-if-error=<ttl>\\\"). TTL aus der serverseitigen CACHE_TTL-Map (z.B. wikidata 86400 s, dwd 1800 s, uba 600 s, default 300 s). stale-while-revalidate/stale-if-error erlauben einem Shared Cache (Cloudflare), bei Ablauf bzw. Origin-Fehler kurz die letzte gute Antwort weiterzuliefern. Echtzeit-Endpunkte (/api/v1/live/*) liefern stattdessen \\\"no-store\\\".\",\"schema\":{\"type\":\"string\"}},\"ETag\":{\"description\":\"Stabiler Entity-Tag des Response-Bodys (sha256-basiert). Nur auf erfolgreichen GET-Reads (200) gesetzt, nie auf Fehler-Envelopes/503.\",\"schema\":{\"type\":\"string\"}}}},\"304\":{\"description\":\"Not Modified. If-None-Match stimmte mit dem aktuellen ETag überein; es wird kein Body geliefert (ETag + Cache-Control bleiben erhalten).\",\"headers\":{\"Cache-Control\":{\"description\":\"Cache-Control je Ressource (\\\"public, max-age=<ttl>, stale-while-revalidate=<ttl>, stale-if-error=<ttl>\\\"). TTL aus der serverseitigen CACHE_TTL-Map (z.B. wikidata 86400 s, dwd 1800 s, uba 600 s, default 300 s). stale-while-revalidate/stale-if-error erlauben einem Shared Cache (Cloudflare), bei Ablauf bzw. Origin-Fehler kurz die letzte gute Antwort weiterzuliefern. Echtzeit-Endpunkte (/api/v1/live/*) liefern stattdessen \\\"no-store\\\".\",\"schema\":{\"type\":\"string\"}},\"ETag\":{\"description\":\"Stabiler Entity-Tag des Response-Bodys (sha256-basiert). Nur auf erfolgreichen GET-Reads (200) gesetzt, nie auf Fehler-Envelopes/503.\",\"schema\":{\"type\":\"string\"}}}}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/api/v1/sources", "segments": [{ "lit": "api" }, { "lit": "v1" }, { "lit": "sources" }], "select": { "exist": ["if_none_match", "limit", "offset", "order", "page", "sort"] }, "transform": { "req": "`reqdata`", "res": "`body.meta`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": {}, "contract": { "id": "GET /api/v1/openapi.yaml", "json": "{\"operationId\":\"getOpenapiYaml\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/yaml\":{\"schema\":{\"type\":\"string\"}}},\"description\":\"Die Spec als YAML\"}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/api/v1/openapi.yaml", "segments": [{ "lit": "api" }, { "lit": "v1" }, { "lit": "openapi.yaml" }], "select": {}, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "meta", "name__orig": "meta", "Name": "Meta", "name_": "meta", "name-": "meta", "NAME": "META", "index$": 4 }, { "active": true, "entity": "meta", "key$": "BasicMetaFlow", "kind": "basic", "name": "BasicMetaFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "meta_ref01" } }], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "meta_ref01", "srcdatavar": "meta_ref01_data", "suffix": "_dt0" }, "match": {}, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-meta_ref01" } }], "index$": 1 }] }, 'Meta');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let meta_ref01_data = Object.values(setup.data.existing.meta)[0];
        // LIST
        const meta_ref01_ent = client.Meta();
        const meta_ref01_match = {};
        const meta_ref01_list = (await meta_ref01_ent.list(meta_ref01_match)).map((e) => e.data());
        // LOAD
        const meta_ref01_match_dt0 = {};
        const meta_ref01_data_dt0 = (await meta_ref01_ent.load(meta_ref01_match_dt0)).data();
        (0, node_assert_1.default)(null != meta_ref01_data_dt0);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/meta/MetaTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.InfranodeOpenDataSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['meta01', 'meta02', 'meta03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'INFRANODE_OPEN_DATA_TEST_META_ENTID': idmap,
        'INFRANODE_OPEN_DATA_TEST_LIVE': 'FALSE',
        'INFRANODE_OPEN_DATA_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['INFRANODE_OPEN_DATA_TEST_META_ENTID'];
    const live = 'TRUE' === env.INFRANODE_OPEN_DATA_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['INFRANODE_OPEN_DATA_TEST_META_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.InfranodeOpenDataSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {},
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
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
    };
    return setup;
}
//# sourceMappingURL=MetaEntity.test.js.map