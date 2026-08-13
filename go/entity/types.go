// Typed models for the InfranodeOpenData SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/infranode-open-data-sdk/go/core"
)

// City is the typed data model for the city entity.
type City struct {
	Data any `json:"data"`
	Meta map[string]any `json:"meta"`
}

// CityLoadMatch is the typed request payload for City.LoadTyped.
type CityLoadMatch struct {
	Id *string `json:"id,omitempty"`
}

// CityListMatch is the typed request payload for City.ListTyped.
type CityListMatch struct {
	Data *any `json:"data,omitempty"`
	Meta *map[string]any `json:"meta,omitempty"`
}

// Compare is the typed data model for the compare entity.
type Compare struct {
	City string `json:"city"`
	Data *map[string]any `json:"data,omitempty"`
	SourceStatus string `json:"source_status"`
}

// CompareListMatch is the typed request payload for Compare.ListTyped.
type CompareListMatch struct {
	City *string `json:"city,omitempty"`
	Data *map[string]any `json:"data,omitempty"`
	SourceStatus *string `json:"source_status,omitempty"`
}

// Health is the typed data model for the health entity.
type Health struct {
	Redis bool `json:"redis"`
	Status string `json:"status"`
	Version string `json:"version"`
}

// HealthLoadMatch is the typed request payload for Health.LoadTyped.
type HealthLoadMatch struct {
	Redis *bool `json:"redis,omitempty"`
	Status *string `json:"status,omitempty"`
	Version *string `json:"version,omitempty"`
}

// Live is the typed data model for the live entity.
type Live struct {
	Data any `json:"data"`
	Meta map[string]any `json:"meta"`
}

// LiveLoadMatch is the typed request payload for Live.LoadTyped.
type LiveLoadMatch struct {
	LiveId *string `json:"live_id,omitempty"`
	RouteId *string `json:"route_id,omitempty"`
	TripId *string `json:"trip_id,omitempty"`
}

// Meta is the typed data model for the meta entity.
type Meta struct {
	BreakerState string `json:"breaker_state"`
	Enabled bool `json:"enabled"`
	Source string `json:"source"`
}

// MetaLoadMatch is the typed request payload for Meta.LoadTyped.
type MetaLoadMatch struct {
	BreakerState *string `json:"breaker_state,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Source *string `json:"source,omitempty"`
}

// MetaListMatch is the typed request payload for Meta.ListTyped.
type MetaListMatch struct {
	BreakerState *string `json:"breaker_state,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Source *string `json:"source,omitempty"`
}

// Station is the typed data model for the station entity.
type Station struct {
	Data any `json:"data"`
	Meta map[string]any `json:"meta"`
}

// StationLoadMatch is the typed request payload for Station.LoadTyped.
type StationLoadMatch struct {
	Eva string `json:"eva"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
