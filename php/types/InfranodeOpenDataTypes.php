<?php
declare(strict_types=1);

// Typed models for the InfranodeOpenData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** City entity data model. */
class City
{
    public mixed $data;
    public array $meta;
}

/** Request payload for City#load. */
class CityLoadMatch
{
    public ?string $id = null;
}

/** Request payload for City#list. */
class CityListMatch
{
    public mixed $data = null;
    public ?array $meta = null;
}

/** Compare entity data model. */
class Compare
{
    public string $city;
    public ?array $data = null;
    public string $source_status;
}

/** Request payload for Compare#list. */
class CompareListMatch
{
    public ?string $city = null;
    public ?array $data = null;
    public ?string $source_status = null;
}

/** Health entity data model. */
class Health
{
    public bool $redis;
    public string $status;
    public string $version;
}

/** Request payload for Health#load. */
class HealthLoadMatch
{
    public ?bool $redis = null;
    public ?string $status = null;
    public ?string $version = null;
}

/** Live entity data model. */
class Live
{
    public mixed $data;
    public array $meta;
}

/** Request payload for Live#load. */
class LiveLoadMatch
{
    public ?string $live_id = null;
    public ?string $route_id = null;
    public ?string $trip_id = null;
}

/** Meta entity data model. */
class Meta
{
    public string $breaker_state;
    public bool $enabled;
    public string $source;
}

/** Request payload for Meta#load. */
class MetaLoadMatch
{
    public ?string $breaker_state = null;
    public ?bool $enabled = null;
    public ?string $source = null;
}

/** Request payload for Meta#list. */
class MetaListMatch
{
    public ?string $breaker_state = null;
    public ?bool $enabled = null;
    public ?string $source = null;
}

/** Station entity data model. */
class Station
{
    public mixed $data;
    public array $meta;
}

/** Request payload for Station#load. */
class StationLoadMatch
{
    public string $eva;
}

