# InfraNode API

Open-Source-Proxy-REST-API für normalisierte offene Daten deutscher Großstädte. Bekannt sind 84 Städte über 100.000 Einwohner, davon 28 Kern-Städte mit voller Quellen-Abdeckung; weitere Städte werden über AGS- und geobasierte Quellen bedient. Jede Antwort folgt demselben kanonischen Envelope (data/meta/attribution) mit Lizenz-Tag und ehrlichem source_status (ok/no_data/disabled/not_covered). Teilabgedeckte Endpunkte (flood, webcams, traffic, road-events) liefern für eine nicht-abgedeckte Stadt 200 source_status=&quot;not_covered&quot; + meta.covered_cities statt eines leeren &quot;ok&quot;; die Abdeckung listet die Seite /abdeckung. Pfade tragen den vollen /api/v1-Prefix, damit die ausgelieferte Spec driftfrei zu den live registrierten Routen bleibt. Paginierung ist kanal-abhängig: direktes REST liefert Datenart-Listen voll, GPT-Actions und MCP sind serverseitig bzw. client-seitig auf ein Default-Limit gebunden; limit=all (bzw. ?all=1) erzwingt kanalunabhängig die Vollausgabe. Details siehe meta.pagination (Meta-Schema). Der Stadt-Slug im Pfad wird tolerant aufgelöst: der deutsche Name mit oder ohne Umlaute, beliebige Groß/Kleinschreibung, gängige englische Exonyme und Kurzformen führen zum kanonischen Slug (München/münchen/munich/munchen -&gt; muenchen, cologne -&gt; koeln, frankfurt -&gt; frankfurt-am-main). Ein unbekannter Name liefert 404 mit einem &quot;Meintest du ...?&quot;-Hinweis auf den nächstliegenden Slug; die kanonischen Slugs listet GET /api/v1/cities.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 6 entities and 104 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### City

Results: Stadt-Liste; Auftragsvergabe-Envelope (aggregierte notices-Liste + count, oder null bei no_data/disabled); Ratsinformations-Envelope (aggregierte papers-Liste + count, oder null bei no_data/disabled); Transit-Envelope. data ist eine Liste von CanonicalRecord (TransitStopPayload) oder null/leer, je nach source_status (&quot;ok&quot;|&quot;disabled&quot;|&quot;not_ingested&quot;).; Katalog-Envelope (CanonicalRecord mit StationCatalogPayload) oder null.; Verkehrs-Envelope (CanonicalRecord oder null); POI-Envelope (CanonicalRecord oder null); Stadt-Eintrag; Unfall-Envelope (CanonicalRecord mit AccidentPayload oder null); Luftqualitäts-Envelope (CanonicalRecord oder null); Stammdaten-Envelope. data ist der CanonicalRecord oder null, wenn die Quelle deaktiviert ist (meta.source_status=&quot;disabled&quot;).; Badegewässer-Envelope (CanonicalRecord oder null); Bike-Counts-Envelope (CanonicalRecord mit CountStationPayload) oder null.; Gewerbeanzeigen-Envelope (CanonicalRecord mit BusinessRegistrationsPayload) oder null.; Ladesäulen-Envelope (CanonicalRecord oder null); Ladebelegungs-Envelope (CanonicalRecord oder null); Warnungen-Envelope (CanonicalRecord mit CivilProtectionWarningPayload: ars, coverage_granularity, count, warnings) oder null bei source_status disabled. source_status ok/disabled.; Regionalstatistik-Envelope (CanonicalRecord oder null); Kriminalstatistik-Envelope (CanonicalRecord mit CrimeStatsPayload oder null); Demografie-Envelope (CanonicalRecord oder null); District-Heating-Envelope (CanonicalRecord mit DistrictHeatingPayload) oder null.; OSM-Feature-Envelope (CanonicalRecord oder null); Wahl-Envelope (CanonicalRecord oder null); Energie-Envelope. data ist eine Liste von CanonicalRecord (EnergyAssetPayload) oder null/leer, je nach source_status.; Event-Envelope (CanonicalRecord oder null); Waldbrand-Envelope (CanonicalRecord oder null); Hochwasser-Envelope (CanonicalRecord oder null); Spritpreis-Envelope (CanonicalRecord mit FuelPricePayload) oder null.; Geo-Envelope (CanonicalRecord oder null); Krankenhaus-Envelope (CanonicalRecord oder null); Denkmal-Envelope (CanonicalRecord oder null); Feiertags-Envelope (CanonicalRecord oder null); Krankenhausatlas-Envelope (CanonicalRecord oder null); ICU-Envelope (CanonicalRecord oder null); Indikatoren-Envelope (CanonicalRecord mit IndicatorsPayload) oder null.; Insolvenzen-Envelope (CanonicalRecord mit InsolvenciesPayload) oder null.; Bodenrichtwerte-Envelope (CanonicalRecord mit LandValuesPayload) oder null.; Office-Wait-Times-Envelope (CanonicalRecord mit OfficeWaitTimesPayload) oder null.; Überblick-Envelope (Basis + Katalog + Highlights + Summary); Parking-Envelope (CanonicalRecord mit ParkingPayload) oder null.; Pollen-/UV-Envelope (CanonicalRecord oder null); Einwohnerdichte-Envelope (CanonicalRecord oder null); Power-Envelope (CanonicalRecord mit PowerPayload) oder null.; Road-Event-Envelope (CanonicalRecord oder null); Sharing-Envelope (CanonicalRecord mit SharingPayload) oder null.; Solar-Envelope (CanonicalRecord mit SolarPayload) oder null.; Solar-Roofs-Envelope (CanonicalRecord mit SolarRoofsPayload) oder null.; Ankunfts-Envelope (CanonicalRecord mit StationArrivalsPayload) oder null.; Abfahrts-Envelope (CanonicalRecord mit StationDeparturesPayload) oder null.; Aufzug-/Rolltreppen-Envelope (CanonicalRecord oder null); Hebesatz-Envelope (CanonicalRecord mit TaxRatesPayload) oder null.; Baumkataster-Envelope (CanonicalRecord oder null); Fahrzeug-Envelope. data ist ein CanonicalRecord (VehicleRegistrationPayload) oder null, je nach source_status.; Pegelstand-Envelope (CanonicalRecord oder null); Wetter-Envelope (CanonicalRecord oder null); Warnungen-Envelope (CanonicalRecord mit WeatherWarningPayload) oder null.; Webcam-Envelope (CanonicalRecord oder null).

SDK operations: `list`, `load`.

Key fields to recognise:

- `meta`: meta trägt zusätzlich source_status (&quot;ok&quot;|&quot;disabled&quot;) und auf dem ok-Pfad cache_status (HIT/MISS/STALE/STALE-ON-ERROR).

### Compare

Results: Vergleichsliste mit per-Stadt source_status.

SDK operations: `list`.

### Health

Results: App läuft.

SDK operations: `load`.

Key fields to recognise:

- `redis`: true wenn Redis erreichbar (Ping erfolgreich)

### Live

Results: Live-Linienstatus-Envelope (CanonicalRecord oder null); Live-Abfahrten-Envelope (CanonicalRecord oder null); Live-Fahrt-Envelope (CanonicalRecord oder null); Luft-Envelope (CanonicalRecord oder null); Live-Baustellen-Envelope (CanonicalRecord oder null); Live-Ereignisse-Envelope (CanonicalRecord oder null); Hochwasser-Envelope (CanonicalRecord oder null); Verkehrs-Envelope (CanonicalRecord oder null); Live-Verkehrslage-Envelope (CanonicalRecord oder null); Pegelstand-Envelope (CanonicalRecord oder null); Webcam-Envelope (CanonicalRecord oder null); Live-Verkehrsmeldungen-Envelope (CanonicalRecord oder null); Live-Parkbelegung-Envelope (CanonicalRecord oder null); Live-Ladesäulen-Belegung-Envelope (CanonicalRecord oder null); Live-Zähldaten-Envelope (CanonicalRecord oder null); Live-Umweltzone-Envelope (CanonicalRecord oder null).

SDK operations: `load`.

### Meta

Results: Liste der Quellen-Status; Die Spec als YAML.

SDK operations: `list`, `load`.

### Station

Results: Ankunfts-Envelope (CanonicalRecord mit StationArrivalsPayload) oder null.; Abfahrts-Envelope (CanonicalRecord mit StationDeparturesPayload) oder null.

SDK operations: `load`.

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| City | `list` | `GET /api/v1/cities` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/public-tenders` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/council-papers` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/transit` | See reference |
| City | `load` | `GET /api/v1/tenders` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/stations` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/traffic` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/pois` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/accidents` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/air` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/air-uba` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/base` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/bathing-water` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/bike-counts` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/business-registrations` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/charging` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/charging-status` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/civil-protection-warnings` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/construction` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/crime-stats` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/demographics` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/district-heating` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/drinking-water` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/education` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/election` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/energy` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/events` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/fire-danger` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/flood` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/fuel-prices` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/geo` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/government-offices` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/health` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/heritage` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/holidays` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/hospitals-atlas` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/icu-live` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/indicators` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/insolvencies` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/land-values` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/markets` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/office-wait-times` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/overview` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/parcel-lockers` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/parking` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/playgrounds` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/pollen-uv` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/population-density` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/post-boxes` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/post-offices` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/power-load` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/power-price` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/public-toilets` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/public-wifi` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/recycling-centres` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/road-events` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/sharing` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/solar` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/solar-roofs` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/station-arrivals` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/station-departures` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/station-facilities` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/tax-rates` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/tourism` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/tree-cadastre` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/unemployment` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/vehicle-registrations` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/water-level` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/weather` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/weather-warnings` | See reference |
| City | `load` | `GET /api/v1/cities/{slug}/webcams` | See reference |
| Compare | `list` | `GET /api/v1/compare` | See reference |
| Health | `load` | `GET /api/v1/health` | See reference |
| Live | `load` | `GET /api/v1/live/{city}/transit/routes/{route_id}/status` | See reference |
| Live | `load` | `GET /api/v1/live/{city}/transit/departures` | See reference |
| Live | `load` | `GET /api/v1/live/{city}/transit/trips/{trip_id}` | See reference |
| Live | `load` | `GET /api/v1/live/{slug}/departures` | See reference |
| Live | `load` | `GET /api/v1/live/{slug}/air` | See reference |
| Live | `load` | `GET /api/v1/live/{slug}/air-uba` | See reference |
| Live | `load` | `GET /api/v1/live/{city}/baustellen` | See reference |
| Live | `load` | `GET /api/v1/live/{city}/ereignisse` | See reference |
| Live | `load` | `GET /api/v1/live/{slug}/flood` | See reference |
| Live | `load` | `GET /api/v1/live/{slug}/traffic` | See reference |
| Live | `load` | `GET /api/v1/live/{city}/traffic-flow` | See reference |
| Live | `load` | `GET /api/v1/live/{slug}/water-level` | See reference |
| Live | `load` | `GET /api/v1/live/{slug}/webcams` | See reference |
| Live | `load` | `GET /api/v1/live/frankfurt-am-main/departures` | See reference |
| Live | `load` | `GET /api/v1/live/hamburg/departures` | See reference |
| Live | `load` | `GET /api/v1/live/nuernberg/departures` | See reference |
| Live | `load` | `GET /api/v1/live/berlin/verkehrsmeldungen` | See reference |
| Live | `load` | `GET /api/v1/live/dortmund/parking` | See reference |
| Live | `load` | `GET /api/v1/live/eround/charging` | See reference |
| Live | `load` | `GET /api/v1/live/frankfurt-am-main/parking` | See reference |
| Live | `load` | `GET /api/v1/live/hamburg/verkehrslage` | See reference |
| Live | `load` | `GET /api/v1/live/hannover/verkehrsmeldungen` | See reference |
| Live | `load` | `GET /api/v1/live/kiel/zaehlstellen` | See reference |
| Live | `load` | `GET /api/v1/live/koeln/umweltzone` | See reference |
| Live | `load` | `GET /api/v1/live/magdeburg/parking` | See reference |
| Live | `load` | `GET /api/v1/live/wuppertal/parking` | See reference |
| Meta | `list` | `GET /api/v1/sources` | See reference |
| Meta | `load` | `GET /api/v1/openapi.yaml` | See reference |
| Station | `load` | `GET /api/v1/stations/{eva}/arrivals` | See reference |
| Station | `load` | `GET /api/v1/stations/{eva}/departures` | See reference |

## Connect to the API

- Öffentliche gehostete API, keylos (Standard): `https://infranode.dev`
- Lokaler Caddy-Ingress (Docker Compose): `http://localhost`

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| Golang | `go/` | Build from source |
| Lua | `lua/` | Build from source |
| PHP | `php/` | Build from source |
| Python | `py/` | Build from source |
| Ruby | `rb/` | Build from source |
| TypeScript | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### Go CLI

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### Go MCP server

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `infranode-open-data_list`: List records for an entity. Supported entities: `city`, `compare`, `meta`.
- `infranode-open-data_load`: Load one record for an entity. Supported entities: `city`, `health`, `live`, `meta`, `station`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- `ratelimit`: Client-side rate limiting via a token bucket
- `retry`: Automatic retry of transient failures with exponential backoff
- `test`: In-memory mock transport for testing without a live server
- `timeout`: Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the first-call guide for the setup sequence.
- Read the authentication guide before using protected routes.
- Use the API reference for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

