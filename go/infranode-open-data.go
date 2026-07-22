package voxgiginfranodeopendatasdk

import (
	"github.com/voxgig-sdk/infranode-open-data-sdk/go/core"
	"github.com/voxgig-sdk/infranode-open-data-sdk/go/entity"
	"github.com/voxgig-sdk/infranode-open-data-sdk/go/feature"
	_ "github.com/voxgig-sdk/infranode-open-data-sdk/go/utility"
)

// Type aliases preserve external API.
type InfranodeOpenDataSDK = core.InfranodeOpenDataSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type InfranodeOpenDataEntity = core.InfranodeOpenDataEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type InfranodeOpenDataError = core.InfranodeOpenDataError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewCityEntityFunc = func(client *core.InfranodeOpenDataSDK, entopts map[string]any) core.InfranodeOpenDataEntity {
		return entity.NewCityEntity(client, entopts)
	}
	core.NewCompareEntityFunc = func(client *core.InfranodeOpenDataSDK, entopts map[string]any) core.InfranodeOpenDataEntity {
		return entity.NewCompareEntity(client, entopts)
	}
	core.NewHealthEntityFunc = func(client *core.InfranodeOpenDataSDK, entopts map[string]any) core.InfranodeOpenDataEntity {
		return entity.NewHealthEntity(client, entopts)
	}
	core.NewLiveEntityFunc = func(client *core.InfranodeOpenDataSDK, entopts map[string]any) core.InfranodeOpenDataEntity {
		return entity.NewLiveEntity(client, entopts)
	}
	core.NewMetaEntityFunc = func(client *core.InfranodeOpenDataSDK, entopts map[string]any) core.InfranodeOpenDataEntity {
		return entity.NewMetaEntity(client, entopts)
	}
	core.NewStationEntityFunc = func(client *core.InfranodeOpenDataSDK, entopts map[string]any) core.InfranodeOpenDataEntity {
		return entity.NewStationEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewInfranodeOpenDataSDK = core.NewInfranodeOpenDataSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewInfranodeOpenDataSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *InfranodeOpenDataSDK  { return NewInfranodeOpenDataSDK(nil) }
func Test() *InfranodeOpenDataSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
