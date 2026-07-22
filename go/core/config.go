package core

func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "InfranodeOpenData",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://infranode.dev",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"city": map[string]any{},
				"compare": map[string]any{},
				"health": map[string]any{},
				"live": map[string]any{},
				"meta": map[string]any{},
				"station": map[string]any{},
			},
		},
		"entity": map[string]any{
			"city": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "data",
						"req": true,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "meta",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 1,
					},
				},
				"name": "city",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/cities",
								"parts": []any{
									"api",
									"v1",
									"cities",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": 50,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "match",
											"orig": "match",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": 0,
											"kind": "query",
											"name": "offset",
											"orig": "offset",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "since",
											"orig": "since",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "status",
											"orig": "status",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/public-tenders",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"public-tenders",
								},
								"select": map[string]any{
									"$action": "public_tender",
									"exist": []any{
										"limit",
										"match",
										"offset",
										"q",
										"since",
										"slug",
										"status",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": 50,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 0,
											"kind": "query",
											"name": "offset",
											"orig": "offset",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "paper_type",
											"orig": "paper_type",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "since",
											"orig": "since",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/council-papers",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"council-papers",
								},
								"select": map[string]any{
									"$action": "council_paper",
									"exist": []any{
										"limit",
										"offset",
										"paper_type",
										"q",
										"since",
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "near",
											"orig": "near",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": 1000,
											"kind": "query",
											"name": "radius_m",
											"orig": "radius_m",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/transit",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"transit",
								},
								"select": map[string]any{
									"$action": "transit",
									"exist": []any{
										"limit",
										"near",
										"page",
										"q",
										"radius_m",
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 2,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": 50,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 0,
											"kind": "query",
											"name": "offset",
											"orig": "offset",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "since",
											"orig": "since",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "status",
											"orig": "status",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/tenders",
								"parts": []any{
									"api",
									"v1",
									"tenders",
								},
								"select": map[string]any{
									"exist": []any{
										"limit",
										"offset",
										"q",
										"since",
										"status",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 3,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/stations",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"stations",
								},
								"select": map[string]any{
									"$action": "station",
									"exist": []any{
										"limit",
										"q",
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 4,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "full",
											"orig": "full",
											"reqd": false,
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "include",
											"orig": "include",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/traffic",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"traffic",
								},
								"select": map[string]any{
									"$action": "traffic",
									"exist": []any{
										"full",
										"include",
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 5,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "type",
											"orig": "type",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/pois",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"pois",
								},
								"select": map[string]any{
									"$action": "poi",
									"exist": []any{
										"slug",
										"type",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 6,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"slug": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 7,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/accidents",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"accidents",
								},
								"select": map[string]any{
									"$action": "accident",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 8,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/air",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"air",
								},
								"select": map[string]any{
									"$action": "air",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 9,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/air-uba",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"air-uba",
								},
								"select": map[string]any{
									"$action": "air_uba",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 10,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/base",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"base",
								},
								"select": map[string]any{
									"$action": "base",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 11,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/bathing-water",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"bathing-water",
								},
								"select": map[string]any{
									"$action": "bathing_water",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 12,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/bike-counts",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"bike-counts",
								},
								"select": map[string]any{
									"$action": "bike_count",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 13,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/business-registrations",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"business-registrations",
								},
								"select": map[string]any{
									"$action": "business_registration",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 14,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/charging",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"charging",
								},
								"select": map[string]any{
									"$action": "charging",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 15,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/charging-status",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"charging-status",
								},
								"select": map[string]any{
									"$action": "charging_status",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 16,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/civil-protection-warnings",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"civil-protection-warnings",
								},
								"select": map[string]any{
									"$action": "civil_protection_warning",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 17,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/construction",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"construction",
								},
								"select": map[string]any{
									"$action": "construction",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 18,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/crime-stats",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"crime-stats",
								},
								"select": map[string]any{
									"$action": "crime_stat",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 19,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/demographics",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"demographics",
								},
								"select": map[string]any{
									"$action": "demographic",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 20,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/district-heating",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"district-heating",
								},
								"select": map[string]any{
									"$action": "district_heating",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 21,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/drinking-water",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"drinking-water",
								},
								"select": map[string]any{
									"$action": "drinking_water",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 22,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/education",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"education",
								},
								"select": map[string]any{
									"$action": "education",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 23,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/election",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"election",
								},
								"select": map[string]any{
									"$action": "election",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 24,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/energy",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"energy",
								},
								"select": map[string]any{
									"$action": "energy",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 25,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/events",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"events",
								},
								"select": map[string]any{
									"$action": "event",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 26,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/fire-danger",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"fire-danger",
								},
								"select": map[string]any{
									"$action": "fire_danger",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 27,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/flood",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"flood",
								},
								"select": map[string]any{
									"$action": "flood",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 28,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/fuel-prices",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"fuel-prices",
								},
								"select": map[string]any{
									"$action": "fuel_price",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 29,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/geo",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"geo",
								},
								"select": map[string]any{
									"$action": "geo",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 30,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/government-offices",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"government-offices",
								},
								"select": map[string]any{
									"$action": "government_office",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 31,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/health",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"health",
								},
								"select": map[string]any{
									"$action": "health",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 32,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/heritage",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"heritage",
								},
								"select": map[string]any{
									"$action": "heritage",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 33,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/holidays",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"holidays",
								},
								"select": map[string]any{
									"$action": "holiday",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 34,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/hospitals-atlas",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"hospitals-atlas",
								},
								"select": map[string]any{
									"$action": "hospitals_atla",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 35,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/icu-live",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"icu-live",
								},
								"select": map[string]any{
									"$action": "icu_live",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 36,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/indicators",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"indicators",
								},
								"select": map[string]any{
									"$action": "indicator",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 37,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/insolvencies",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"insolvencies",
								},
								"select": map[string]any{
									"$action": "insolvency",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 38,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/land-values",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"land-values",
								},
								"select": map[string]any{
									"$action": "land_value",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 39,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/markets",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"markets",
								},
								"select": map[string]any{
									"$action": "market",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 40,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/office-wait-times",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"office-wait-times",
								},
								"select": map[string]any{
									"$action": "office_wait_time",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 41,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/overview",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"overview",
								},
								"select": map[string]any{
									"$action": "overview",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 42,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/parcel-lockers",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"parcel-lockers",
								},
								"select": map[string]any{
									"$action": "parcel_locker",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 43,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/parking",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"parking",
								},
								"select": map[string]any{
									"$action": "parking",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 44,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/playgrounds",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"playgrounds",
								},
								"select": map[string]any{
									"$action": "playground",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 45,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/pollen-uv",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"pollen-uv",
								},
								"select": map[string]any{
									"$action": "pollen_uv",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 46,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/population-density",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"population-density",
								},
								"select": map[string]any{
									"$action": "population_density",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 47,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/post-boxes",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"post-boxes",
								},
								"select": map[string]any{
									"$action": "post_box",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 48,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/post-offices",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"post-offices",
								},
								"select": map[string]any{
									"$action": "post_office",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 49,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/power-load",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"power-load",
								},
								"select": map[string]any{
									"$action": "power_load",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 50,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/power-price",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"power-price",
								},
								"select": map[string]any{
									"$action": "power_price",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 51,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/public-toilets",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"public-toilets",
								},
								"select": map[string]any{
									"$action": "public_toilet",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 52,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/public-wifi",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"public-wifi",
								},
								"select": map[string]any{
									"$action": "public_wifi",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 53,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/recycling-centres",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"recycling-centres",
								},
								"select": map[string]any{
									"$action": "recycling_centre",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 54,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/road-events",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"road-events",
								},
								"select": map[string]any{
									"$action": "road_event",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 55,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/sharing",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"sharing",
								},
								"select": map[string]any{
									"$action": "sharing",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 56,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/solar",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"solar",
								},
								"select": map[string]any{
									"$action": "solar",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 57,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/solar-roofs",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"solar-roofs",
								},
								"select": map[string]any{
									"$action": "solar_roof",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 58,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/station-arrivals",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"station-arrivals",
								},
								"select": map[string]any{
									"$action": "station_arrival",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 59,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/station-departures",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"station-departures",
								},
								"select": map[string]any{
									"$action": "station_departure",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 60,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/station-facilities",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"station-facilities",
								},
								"select": map[string]any{
									"$action": "station_facility",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 61,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/tax-rates",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"tax-rates",
								},
								"select": map[string]any{
									"$action": "tax_rate",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 62,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/tourism",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"tourism",
								},
								"select": map[string]any{
									"$action": "tourism",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 63,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/tree-cadastre",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"tree-cadastre",
								},
								"select": map[string]any{
									"$action": "tree_cadastre",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 64,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/unemployment",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"unemployment",
								},
								"select": map[string]any{
									"$action": "unemployment",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 65,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/vehicle-registrations",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"vehicle-registrations",
								},
								"select": map[string]any{
									"$action": "vehicle_registration",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 66,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/water-level",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"water-level",
								},
								"select": map[string]any{
									"$action": "water_level",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 67,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/weather",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"weather",
								},
								"select": map[string]any{
									"$action": "weather",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 68,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/weather-warnings",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"weather-warnings",
								},
								"select": map[string]any{
									"$action": "weather_warning",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 69,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/webcams",
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"webcams",
								},
								"select": map[string]any{
									"$action": "webcam",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 70,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"city",
						},
					},
				},
			},
			"compare": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "city",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "data",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "source_status",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
				},
				"name": "compare",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "if_none_match",
											"orig": "if_none_match",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": "berlin,koeln,hamburg",
											"kind": "query",
											"name": "city",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": 50,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 0,
											"kind": "query",
											"name": "offset",
											"orig": "offset",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": "asc",
											"kind": "query",
											"name": "order",
											"orig": "order",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "resource",
											"orig": "resource",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/compare",
								"parts": []any{
									"api",
									"v1",
									"compare",
								},
								"select": map[string]any{
									"exist": []any{
										"city",
										"if_none_match",
										"limit",
										"offset",
										"order",
										"page",
										"resource",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"health": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "redi",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": true,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "version",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
				},
				"name": "health",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/health",
								"parts": []any{
									"api",
									"v1",
									"health",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"live": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "data",
						"req": true,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "meta",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 1,
					},
				},
				"name": "live",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "live_id",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "route_id",
											"orig": "route_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{city}/transit/routes/{route_id}/status",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{live_id}",
									"transit",
									"routes",
									"{route_id}",
									"status",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"city": "live_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"live_id",
										"route_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "live_id",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "stop_id",
											"orig": "stop_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{city}/transit/departures",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{live_id}",
									"transit",
									"departures",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"city": "live_id",
									},
								},
								"select": map[string]any{
									"$action": "transit_departure",
									"exist": []any{
										"live_id",
										"stop_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "live_id",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "trip_id",
											"orig": "trip_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{city}/transit/trips/{trip_id}",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{live_id}",
									"transit",
									"trips",
									"{trip_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"city": "live_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"live_id",
										"trip_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 2,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "station",
											"orig": "station",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{slug}/departures",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"departures",
								},
								"select": map[string]any{
									"$action": "departure",
									"exist": []any{
										"slug",
										"station",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 3,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{slug}/air",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"air",
								},
								"select": map[string]any{
									"$action": "air",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 4,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{slug}/air-uba",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"air-uba",
								},
								"select": map[string]any{
									"$action": "air_uba",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 5,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "city",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{city}/baustellen",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{city}",
									"baustellen",
								},
								"select": map[string]any{
									"$action": "baustellen",
									"exist": []any{
										"city",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 6,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "city",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{city}/ereignisse",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{city}",
									"ereignisse",
								},
								"select": map[string]any{
									"$action": "ereignisse",
									"exist": []any{
										"city",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 7,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{slug}/flood",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"flood",
								},
								"select": map[string]any{
									"$action": "flood",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 8,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{slug}/traffic",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"traffic",
								},
								"select": map[string]any{
									"$action": "traffic",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 9,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "city",
											"orig": "city",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{city}/traffic-flow",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{city}",
									"traffic-flow",
								},
								"select": map[string]any{
									"$action": "traffic_flow",
									"exist": []any{
										"city",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 10,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{slug}/water-level",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"water-level",
								},
								"select": map[string]any{
									"$action": "water_level",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 11,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "slug",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/{slug}/webcams",
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"webcams",
								},
								"select": map[string]any{
									"$action": "webcam",
									"exist": []any{
										"slug",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 12,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "Frankfurt (Main) Hauptbahnhof",
											"kind": "query",
											"name": "station",
											"orig": "station",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/frankfurt-am-main/departures",
								"parts": []any{
									"api",
									"v1",
									"live",
									"frankfurt-am-main",
									"departures",
								},
								"select": map[string]any{
									"exist": []any{
										"station",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 13,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "Hamburg Hauptbahnhof",
											"kind": "query",
											"name": "station",
											"orig": "station",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/hamburg/departures",
								"parts": []any{
									"api",
									"v1",
									"live",
									"hamburg",
									"departures",
								},
								"select": map[string]any{
									"exist": []any{
										"station",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 14,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"active": true,
											"example": "510",
											"kind": "query",
											"name": "stop_id",
											"orig": "stop_id",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/live/nuernberg/departures",
								"parts": []any{
									"api",
									"v1",
									"live",
									"nuernberg",
									"departures",
								},
								"select": map[string]any{
									"exist": []any{
										"stop_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 15,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/berlin/verkehrsmeldungen",
								"parts": []any{
									"api",
									"v1",
									"live",
									"berlin",
									"verkehrsmeldungen",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 16,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/dortmund/parking",
								"parts": []any{
									"api",
									"v1",
									"live",
									"dortmund",
									"parking",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 17,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/eround/charging",
								"parts": []any{
									"api",
									"v1",
									"live",
									"eround",
									"charging",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 18,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/frankfurt-am-main/parking",
								"parts": []any{
									"api",
									"v1",
									"live",
									"frankfurt-am-main",
									"parking",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 19,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/hamburg/verkehrslage",
								"parts": []any{
									"api",
									"v1",
									"live",
									"hamburg",
									"verkehrslage",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 20,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/hannover/verkehrsmeldungen",
								"parts": []any{
									"api",
									"v1",
									"live",
									"hannover",
									"verkehrsmeldungen",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 21,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/kiel/zaehlstellen",
								"parts": []any{
									"api",
									"v1",
									"live",
									"kiel",
									"zaehlstellen",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 22,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/koeln/umweltzone",
								"parts": []any{
									"api",
									"v1",
									"live",
									"koeln",
									"umweltzone",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 23,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/magdeburg/parking",
								"parts": []any{
									"api",
									"v1",
									"live",
									"magdeburg",
									"parking",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 24,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/live/wuppertal/parking",
								"parts": []any{
									"api",
									"v1",
									"live",
									"wuppertal",
									"parking",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 25,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"live",
						},
						[]any{
							"live",
							"route",
						},
						[]any{
							"live",
							"trip",
						},
					},
				},
			},
			"meta": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "breaker_state",
						"req": true,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "enabled",
						"req": true,
						"type": "`$BOOLEAN`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "source",
						"req": true,
						"type": "`$STRING`",
						"index$": 2,
					},
				},
				"name": "meta",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"active": true,
											"kind": "header",
											"name": "if_none_match",
											"orig": "if_none_match",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": 50,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 0,
											"kind": "query",
											"name": "offset",
											"orig": "offset",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": "asc",
											"kind": "query",
											"name": "order",
											"orig": "order",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/sources",
								"parts": []any{
									"api",
									"v1",
									"sources",
								},
								"select": map[string]any{
									"exist": []any{
										"if_none_match",
										"limit",
										"offset",
										"order",
										"page",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meta`",
								},
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
								"method": "GET",
								"orig": "/api/v1/openapi.yaml",
								"parts": []any{
									"api",
									"v1",
									"openapi.yaml",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"station": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "data",
						"req": true,
						"type": "`$ANY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "meta",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 1,
					},
				},
				"name": "station",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "eva",
											"orig": "eva",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/stations/{eva}/arrivals",
								"parts": []any{
									"api",
									"v1",
									"stations",
									"{eva}",
									"arrivals",
								},
								"select": map[string]any{
									"$action": "arrival",
									"exist": []any{
										"eva",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "eva",
											"orig": "eva",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
								"method": "GET",
								"orig": "/api/v1/stations/{eva}/departures",
								"parts": []any{
									"api",
									"v1",
									"stations",
									"{eva}",
									"departures",
								},
								"select": map[string]any{
									"$action": "departure",
									"exist": []any{
										"eva",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 1,
							},
						},
						"key$": "load",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"station",
						},
					},
				},
			},
		},
	}
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
