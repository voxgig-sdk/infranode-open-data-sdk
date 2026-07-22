package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewCityEntityFunc func(client *InfranodeOpenDataSDK, entopts map[string]any) InfranodeOpenDataEntity

var NewCompareEntityFunc func(client *InfranodeOpenDataSDK, entopts map[string]any) InfranodeOpenDataEntity

var NewHealthEntityFunc func(client *InfranodeOpenDataSDK, entopts map[string]any) InfranodeOpenDataEntity

var NewLiveEntityFunc func(client *InfranodeOpenDataSDK, entopts map[string]any) InfranodeOpenDataEntity

var NewMetaEntityFunc func(client *InfranodeOpenDataSDK, entopts map[string]any) InfranodeOpenDataEntity

var NewStationEntityFunc func(client *InfranodeOpenDataSDK, entopts map[string]any) InfranodeOpenDataEntity

