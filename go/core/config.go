package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "InfranodeOpenData",
			"slug": "infranode-open-data",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
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
						"name": "data",
						"title": "Data",
						"type": "`$ANY`",
						"req": true,
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "meta",
						"title": "Meta",
						"type": "`$OBJECT`",
						"req": true,
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "city",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/public-tenders",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "public-tenders",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"public-tenders",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 50,
										},
										map[string]any{
											"name": "match",
											"orig": "match",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 0,
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "since",
											"orig": "since",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
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
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/council-papers",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "council-papers",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"council-papers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 50,
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 0,
										},
										map[string]any{
											"name": "paper_type",
											"orig": "paper_type",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "since",
											"orig": "since",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
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
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/transit",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "transit",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"transit",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "near",
											"orig": "near",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "radius_m",
											"orig": "radius_m",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1000,
										},
									},
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
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/tenders",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "tenders",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"tenders",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 50,
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 0,
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "since",
											"orig": "since",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
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
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/stations",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "stations",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"stations",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "station",
									"exist": []any{
										"limit",
										"q",
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/traffic",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "traffic",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"traffic",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "full",
											"orig": "full",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "include",
											"orig": "include",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "traffic",
									"exist": []any{
										"full",
										"include",
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/pois",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "pois",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"pois",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "poi",
									"exist": []any{
										"slug",
										"type",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/accidents",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "accidents",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"accidents",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "accident",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/air",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "air",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"air",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "air",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/air-uba",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "air-uba",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"air-uba",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "air_uba",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/base",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "base",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"base",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "base",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/bathing-water",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "bathing-water",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"bathing-water",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "bathing_water",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/bike-counts",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "bike-counts",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"bike-counts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "bike_count",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/business-registrations",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "business-registrations",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"business-registrations",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "business_registration",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/charging",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "charging",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"charging",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "charging",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/charging-status",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "charging-status",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"charging-status",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "charging_status",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/civil-protection-warnings",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "civil-protection-warnings",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"civil-protection-warnings",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "civil_protection_warning",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/construction",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "construction",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"construction",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "construction",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/crime-stats",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "crime-stats",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"crime-stats",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "crime_stat",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/demographics",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "demographics",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"demographics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "demographic",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/district-heating",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "district-heating",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"district-heating",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "district_heating",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/drinking-water",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "drinking-water",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"drinking-water",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "drinking_water",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/education",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "education",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"education",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "education",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/election",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "election",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"election",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "election",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/energy",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "energy",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"energy",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "energy",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/events",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "events",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"events",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "event",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/fire-danger",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "fire-danger",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"fire-danger",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "fire_danger",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/flood",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "flood",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"flood",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "flood",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/fuel-prices",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "fuel-prices",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"fuel-prices",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "fuel_price",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/geo",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "geo",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"geo",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "geo",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/government-offices",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "government-offices",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"government-offices",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "government_office",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/health",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "health",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"health",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "health",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/heritage",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "heritage",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"heritage",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "heritage",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/holidays",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "holidays",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"holidays",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "holiday",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/hospitals-atlas",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "hospitals-atlas",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"hospitals-atlas",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "hospitals_atla",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/icu-live",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "icu-live",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"icu-live",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "icu_live",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/indicators",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "indicators",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"indicators",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "indicator",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/insolvencies",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "insolvencies",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"insolvencies",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "insolvency",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/land-values",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "land-values",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"land-values",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "land_value",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/markets",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "markets",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"markets",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "market",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/office-wait-times",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "office-wait-times",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"office-wait-times",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "office_wait_time",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/overview",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "overview",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"overview",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "overview",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/parcel-lockers",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "parcel-lockers",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"parcel-lockers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "parcel_locker",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/parking",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "parking",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"parking",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "parking",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/playgrounds",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "playgrounds",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"playgrounds",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "playground",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/pollen-uv",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "pollen-uv",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"pollen-uv",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "pollen_uv",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/population-density",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "population-density",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"population-density",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "population_density",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/post-boxes",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "post-boxes",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"post-boxes",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "post_box",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/post-offices",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "post-offices",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"post-offices",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "post_office",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/power-load",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "power-load",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"power-load",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "power_load",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/power-price",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "power-price",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"power-price",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "power_price",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/public-toilets",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "public-toilets",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"public-toilets",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "public_toilet",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/public-wifi",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "public-wifi",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"public-wifi",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "public_wifi",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/recycling-centres",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "recycling-centres",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"recycling-centres",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "recycling_centre",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/road-events",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "road-events",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"road-events",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "road_event",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/sharing",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "sharing",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"sharing",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "sharing",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/solar",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "solar",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"solar",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "solar",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/solar-roofs",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "solar-roofs",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"solar-roofs",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "solar_roof",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/station-arrivals",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "station-arrivals",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"station-arrivals",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "station_arrival",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/station-departures",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "station-departures",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"station-departures",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "station_departure",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/station-facilities",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "station-facilities",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"station-facilities",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "station_facility",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/tax-rates",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "tax-rates",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"tax-rates",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "tax_rate",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/tourism",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "tourism",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"tourism",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "tourism",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/tree-cadastre",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "tree-cadastre",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"tree-cadastre",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "tree_cadastre",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/unemployment",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "unemployment",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"unemployment",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "unemployment",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/vehicle-registrations",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "vehicle-registrations",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"vehicle-registrations",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "vehicle_registration",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/water-level",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "water-level",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"water-level",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "water_level",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/weather",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "weather",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"weather",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "weather",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/weather-warnings",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "weather-warnings",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"weather-warnings",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "weather_warning",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/cities/{slug}/webcams",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "cities",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "webcams",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"cities",
									"{slug}",
									"webcams",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "webcam",
									"exist": []any{
										"slug",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"compare": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "city",
						"title": "City",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "source_status",
						"title": "Source Status",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "compare",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/compare",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "compare",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"compare",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "if_none_match",
											"orig": "if_none_match",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "city",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
											"example": "berlin,koeln,hamburg",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 50,
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 0,
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
											"example": "asc",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "resource",
											"orig": "resource",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
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
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"health": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "redis",
						"title": "Redis",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "true wenn Redis erreichbar (Ping erfolgreich)",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "version",
						"title": "Version",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "health",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/health",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "health",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"health",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"live": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$ANY`",
						"req": true,
					},
					map[string]any{
						"name": "meta",
						"title": "Meta",
						"type": "`$OBJECT`",
						"req": true,
					},
				},
				"name": "live",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{city}/transit/routes/{route_id}/status",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "live_id",
									},
									map[string]any{
										"lit": "transit",
									},
									map[string]any{
										"lit": "routes",
									},
									map[string]any{
										"var": "route_id",
									},
									map[string]any{
										"lit": "status",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "live_id",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "route_id",
											"orig": "route_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"live_id",
										"route_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{city}/transit/departures",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "live_id",
									},
									map[string]any{
										"lit": "transit",
									},
									map[string]any{
										"lit": "departures",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "live_id",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "stop_id",
											"orig": "stop_id",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "transit_departure",
									"exist": []any{
										"live_id",
										"stop_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{city}/transit/trips/{trip_id}",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "live_id",
									},
									map[string]any{
										"lit": "transit",
									},
									map[string]any{
										"lit": "trips",
									},
									map[string]any{
										"var": "trip_id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "live_id",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "trip_id",
											"orig": "trip_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"live_id",
										"trip_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{slug}/departures",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "departures",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"departures",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "station",
											"orig": "station",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "departure",
									"exist": []any{
										"slug",
										"station",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{slug}/air",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "air",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"air",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "air",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{slug}/air-uba",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "air-uba",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"air-uba",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "air_uba",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{city}/baustellen",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "city",
									},
									map[string]any{
										"lit": "baustellen",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{city}",
									"baustellen",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "city",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "baustellen",
									"exist": []any{
										"city",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{city}/ereignisse",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "city",
									},
									map[string]any{
										"lit": "ereignisse",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{city}",
									"ereignisse",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "city",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "ereignisse",
									"exist": []any{
										"city",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{slug}/flood",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "flood",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"flood",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "flood",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{slug}/traffic",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "traffic",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"traffic",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "traffic",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{city}/traffic-flow",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "city",
									},
									map[string]any{
										"lit": "traffic-flow",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{city}",
									"traffic-flow",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "city",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "traffic_flow",
									"exist": []any{
										"city",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{slug}/water-level",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "water-level",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"water-level",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "water_level",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/{slug}/webcams",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"var": "slug",
									},
									map[string]any{
										"lit": "webcams",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"{slug}",
									"webcams",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "webcam",
									"exist": []any{
										"slug",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/frankfurt-am-main/departures",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "frankfurt-am-main",
									},
									map[string]any{
										"lit": "departures",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"frankfurt-am-main",
									"departures",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "station",
											"orig": "station",
											"type": "`$STRING`",
											"kind": "query",
											"example": "Frankfurt (Main) Hauptbahnhof",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"station",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/hamburg/departures",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "hamburg",
									},
									map[string]any{
										"lit": "departures",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"hamburg",
									"departures",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "station",
											"orig": "station",
											"type": "`$STRING`",
											"kind": "query",
											"example": "Hamburg Hauptbahnhof",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"station",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/nuernberg/departures",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "nuernberg",
									},
									map[string]any{
										"lit": "departures",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"nuernberg",
									"departures",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "stop_id",
											"orig": "stop_id",
											"type": "`$STRING`",
											"kind": "query",
											"example": "510",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"stop_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/berlin/verkehrsmeldungen",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "berlin",
									},
									map[string]any{
										"lit": "verkehrsmeldungen",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"berlin",
									"verkehrsmeldungen",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/dortmund/parking",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "dortmund",
									},
									map[string]any{
										"lit": "parking",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"dortmund",
									"parking",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/eround/charging",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "eround",
									},
									map[string]any{
										"lit": "charging",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"eround",
									"charging",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/frankfurt-am-main/parking",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "frankfurt-am-main",
									},
									map[string]any{
										"lit": "parking",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"frankfurt-am-main",
									"parking",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/hamburg/verkehrslage",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "hamburg",
									},
									map[string]any{
										"lit": "verkehrslage",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"hamburg",
									"verkehrslage",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/hannover/verkehrsmeldungen",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "hannover",
									},
									map[string]any{
										"lit": "verkehrsmeldungen",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"hannover",
									"verkehrsmeldungen",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/kiel/zaehlstellen",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "kiel",
									},
									map[string]any{
										"lit": "zaehlstellen",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"kiel",
									"zaehlstellen",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/koeln/umweltzone",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "koeln",
									},
									map[string]any{
										"lit": "umweltzone",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"koeln",
									"umweltzone",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/magdeburg/parking",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "magdeburg",
									},
									map[string]any{
										"lit": "parking",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"magdeburg",
									"parking",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/live/wuppertal/parking",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "live",
									},
									map[string]any{
										"lit": "wuppertal",
									},
									map[string]any{
										"lit": "parking",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"live",
									"wuppertal",
									"parking",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"meta": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "breaker_state",
						"title": "Breaker State",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "enabled",
						"title": "Enabled",
						"type": "`$BOOLEAN`",
						"req": true,
					},
					map[string]any{
						"name": "source",
						"title": "Source",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "meta",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/sources",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "sources",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"sources",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meta`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "if_none_match",
											"orig": "if_none_match",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 50,
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 0,
										},
										map[string]any{
											"name": "order",
											"orig": "order",
											"type": "`$STRING`",
											"kind": "query",
											"example": "asc",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
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
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/openapi.yaml",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "openapi.yaml",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"openapi.yaml",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"station": map[string]any{
				"fields": []any{},
				"name": "station",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/stations/{eva}/arrivals",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "stations",
									},
									map[string]any{
										"var": "eva",
									},
									map[string]any{
										"lit": "arrivals",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"stations",
									"{eva}",
									"arrivals",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "eva",
											"orig": "eva",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "arrival",
									"exist": []any{
										"eva",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/stations/{eva}/departures",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "stations",
									},
									map[string]any{
										"var": "eva",
									},
									map[string]any{
										"lit": "departures",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"stations",
									"{eva}",
									"departures",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "eva",
											"orig": "eva",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "departure",
									"exist": []any{
										"eva",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
