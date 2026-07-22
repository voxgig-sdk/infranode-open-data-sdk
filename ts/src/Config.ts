
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }


  main = {
    name: 'ProjectName',
  }


  feature = {
     test:     {
      "options": {
        "active": false
      }
    },

  }


  options = {
    base: 'https://infranode.dev',

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      city: {
      },

      compare: {
      },

      health: {
      },

      live: {
      },

      meta: {
      },

      station: {
      },

    }
  }


  entity = {
    "city": {
      "fields": [
        {
          "active": true,
          "name": "data",
          "req": true,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "meta",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 1
        }
      ],
      "name": "city",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/cities",
              "parts": [
                "api",
                "v1",
                "cities"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": 50,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "match",
                    "orig": "match",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 0,
                    "kind": "query",
                    "name": "offset",
                    "orig": "offset",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "since",
                    "orig": "since",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/public-tenders",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "public-tenders"
              ],
              "select": {
                "$action": "public_tender",
                "exist": [
                  "limit",
                  "match",
                  "offset",
                  "q",
                  "since",
                  "slug",
                  "status"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": 50,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 0,
                    "kind": "query",
                    "name": "offset",
                    "orig": "offset",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "paper_type",
                    "orig": "paper_type",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "since",
                    "orig": "since",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/council-papers",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "council-papers"
              ],
              "select": {
                "$action": "council_paper",
                "exist": [
                  "limit",
                  "offset",
                  "paper_type",
                  "q",
                  "since",
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "near",
                    "orig": "near",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 1000,
                    "kind": "query",
                    "name": "radius_m",
                    "orig": "radius_m",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/transit",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "transit"
              ],
              "select": {
                "$action": "transit",
                "exist": [
                  "limit",
                  "near",
                  "page",
                  "q",
                  "radius_m",
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": 50,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 0,
                    "kind": "query",
                    "name": "offset",
                    "orig": "offset",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "since",
                    "orig": "since",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/tenders",
              "parts": [
                "api",
                "v1",
                "tenders"
              ],
              "select": {
                "exist": [
                  "limit",
                  "offset",
                  "q",
                  "since",
                  "status"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/stations",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "stations"
              ],
              "select": {
                "$action": "station",
                "exist": [
                  "limit",
                  "q",
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 4
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "full",
                    "orig": "full",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "include",
                    "orig": "include",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/traffic",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "traffic"
              ],
              "select": {
                "$action": "traffic",
                "exist": [
                  "full",
                  "include",
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 5
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "type",
                    "orig": "type",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/pois",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "pois"
              ],
              "select": {
                "$action": "poi",
                "exist": [
                  "slug",
                  "type"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 6
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}",
              "parts": [
                "api",
                "v1",
                "cities",
                "{id}"
              ],
              "rename": {
                "param": {
                  "slug": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 7
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/accidents",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "accidents"
              ],
              "select": {
                "$action": "accident",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 8
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/air",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "air"
              ],
              "select": {
                "$action": "air",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 9
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/air-uba",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "air-uba"
              ],
              "select": {
                "$action": "air_uba",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 10
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/base",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "base"
              ],
              "select": {
                "$action": "base",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 11
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/bathing-water",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "bathing-water"
              ],
              "select": {
                "$action": "bathing_water",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 12
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/bike-counts",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "bike-counts"
              ],
              "select": {
                "$action": "bike_count",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 13
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/business-registrations",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "business-registrations"
              ],
              "select": {
                "$action": "business_registration",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 14
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/charging",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "charging"
              ],
              "select": {
                "$action": "charging",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 15
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/charging-status",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "charging-status"
              ],
              "select": {
                "$action": "charging_status",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 16
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/civil-protection-warnings",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "civil-protection-warnings"
              ],
              "select": {
                "$action": "civil_protection_warning",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 17
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/construction",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "construction"
              ],
              "select": {
                "$action": "construction",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 18
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/crime-stats",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "crime-stats"
              ],
              "select": {
                "$action": "crime_stat",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 19
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/demographics",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "demographics"
              ],
              "select": {
                "$action": "demographic",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 20
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/district-heating",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "district-heating"
              ],
              "select": {
                "$action": "district_heating",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 21
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/drinking-water",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "drinking-water"
              ],
              "select": {
                "$action": "drinking_water",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 22
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/education",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "education"
              ],
              "select": {
                "$action": "education",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 23
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/election",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "election"
              ],
              "select": {
                "$action": "election",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 24
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/energy",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "energy"
              ],
              "select": {
                "$action": "energy",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 25
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/events",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "events"
              ],
              "select": {
                "$action": "event",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 26
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/fire-danger",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "fire-danger"
              ],
              "select": {
                "$action": "fire_danger",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 27
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/flood",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "flood"
              ],
              "select": {
                "$action": "flood",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 28
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/fuel-prices",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "fuel-prices"
              ],
              "select": {
                "$action": "fuel_price",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 29
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/geo",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "geo"
              ],
              "select": {
                "$action": "geo",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 30
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/government-offices",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "government-offices"
              ],
              "select": {
                "$action": "government_office",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 31
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/health",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "health"
              ],
              "select": {
                "$action": "health",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 32
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/heritage",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "heritage"
              ],
              "select": {
                "$action": "heritage",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 33
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/holidays",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "holidays"
              ],
              "select": {
                "$action": "holiday",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 34
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/hospitals-atlas",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "hospitals-atlas"
              ],
              "select": {
                "$action": "hospitals_atla",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 35
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/icu-live",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "icu-live"
              ],
              "select": {
                "$action": "icu_live",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 36
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/indicators",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "indicators"
              ],
              "select": {
                "$action": "indicator",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 37
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/insolvencies",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "insolvencies"
              ],
              "select": {
                "$action": "insolvency",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 38
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/land-values",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "land-values"
              ],
              "select": {
                "$action": "land_value",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 39
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/markets",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "markets"
              ],
              "select": {
                "$action": "market",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 40
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/office-wait-times",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "office-wait-times"
              ],
              "select": {
                "$action": "office_wait_time",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 41
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/overview",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "overview"
              ],
              "select": {
                "$action": "overview",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 42
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/parcel-lockers",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "parcel-lockers"
              ],
              "select": {
                "$action": "parcel_locker",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 43
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/parking",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "parking"
              ],
              "select": {
                "$action": "parking",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 44
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/playgrounds",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "playgrounds"
              ],
              "select": {
                "$action": "playground",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 45
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/pollen-uv",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "pollen-uv"
              ],
              "select": {
                "$action": "pollen_uv",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 46
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/population-density",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "population-density"
              ],
              "select": {
                "$action": "population_density",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 47
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/post-boxes",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "post-boxes"
              ],
              "select": {
                "$action": "post_box",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 48
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/post-offices",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "post-offices"
              ],
              "select": {
                "$action": "post_office",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 49
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/power-load",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "power-load"
              ],
              "select": {
                "$action": "power_load",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 50
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/power-price",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "power-price"
              ],
              "select": {
                "$action": "power_price",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 51
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/public-toilets",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "public-toilets"
              ],
              "select": {
                "$action": "public_toilet",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 52
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/public-wifi",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "public-wifi"
              ],
              "select": {
                "$action": "public_wifi",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 53
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/recycling-centres",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "recycling-centres"
              ],
              "select": {
                "$action": "recycling_centre",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 54
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/road-events",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "road-events"
              ],
              "select": {
                "$action": "road_event",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 55
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/sharing",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "sharing"
              ],
              "select": {
                "$action": "sharing",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 56
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/solar",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "solar"
              ],
              "select": {
                "$action": "solar",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 57
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/solar-roofs",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "solar-roofs"
              ],
              "select": {
                "$action": "solar_roof",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 58
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/station-arrivals",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "station-arrivals"
              ],
              "select": {
                "$action": "station_arrival",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 59
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/station-departures",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "station-departures"
              ],
              "select": {
                "$action": "station_departure",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 60
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/station-facilities",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "station-facilities"
              ],
              "select": {
                "$action": "station_facility",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 61
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/tax-rates",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "tax-rates"
              ],
              "select": {
                "$action": "tax_rate",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 62
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/tourism",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "tourism"
              ],
              "select": {
                "$action": "tourism",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 63
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/tree-cadastre",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "tree-cadastre"
              ],
              "select": {
                "$action": "tree_cadastre",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 64
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/unemployment",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "unemployment"
              ],
              "select": {
                "$action": "unemployment",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 65
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/vehicle-registrations",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "vehicle-registrations"
              ],
              "select": {
                "$action": "vehicle_registration",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 66
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/water-level",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "water-level"
              ],
              "select": {
                "$action": "water_level",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 67
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/weather",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "weather"
              ],
              "select": {
                "$action": "weather",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 68
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/weather-warnings",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "weather-warnings"
              ],
              "select": {
                "$action": "weather_warning",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 69
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/cities/{slug}/webcams",
              "parts": [
                "api",
                "v1",
                "cities",
                "{slug}",
                "webcams"
              ],
              "select": {
                "$action": "webcam",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 70
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": [
          [
            "city"
          ]
        ]
      }
    },
    "compare": {
      "fields": [
        {
          "active": true,
          "name": "city",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "data",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 1
        },
        {
          "active": true,
          "name": "source_status",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        }
      ],
      "name": "compare",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "header": [
                  {
                    "active": true,
                    "kind": "header",
                    "name": "if_none_match",
                    "orig": "if_none_match",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": "berlin,koeln,hamburg",
                    "kind": "query",
                    "name": "city",
                    "orig": "city",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 50,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 0,
                    "kind": "query",
                    "name": "offset",
                    "orig": "offset",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": "asc",
                    "kind": "query",
                    "name": "order",
                    "orig": "order",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "resource",
                    "orig": "resource",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/compare",
              "parts": [
                "api",
                "v1",
                "compare"
              ],
              "select": {
                "exist": [
                  "city",
                  "if_none_match",
                  "limit",
                  "offset",
                  "order",
                  "page",
                  "resource",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "health": {
      "fields": [
        {
          "active": true,
          "name": "redi",
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 0
        },
        {
          "active": true,
          "name": "status",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "version",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        }
      ],
      "name": "health",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/health",
              "parts": [
                "api",
                "v1",
                "health"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "live": {
      "fields": [
        {
          "active": true,
          "name": "data",
          "req": true,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "meta",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 1
        }
      ],
      "name": "live",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "live_id",
                    "orig": "city",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "route_id",
                    "orig": "route_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{city}/transit/routes/{route_id}/status",
              "parts": [
                "api",
                "v1",
                "live",
                "{live_id}",
                "transit",
                "routes",
                "{route_id}",
                "status"
              ],
              "rename": {
                "param": {
                  "city": "live_id"
                }
              },
              "select": {
                "exist": [
                  "live_id",
                  "route_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "live_id",
                    "orig": "city",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "stop_id",
                    "orig": "stop_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{city}/transit/departures",
              "parts": [
                "api",
                "v1",
                "live",
                "{live_id}",
                "transit",
                "departures"
              ],
              "rename": {
                "param": {
                  "city": "live_id"
                }
              },
              "select": {
                "$action": "transit_departure",
                "exist": [
                  "live_id",
                  "stop_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "live_id",
                    "orig": "city",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "trip_id",
                    "orig": "trip_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{city}/transit/trips/{trip_id}",
              "parts": [
                "api",
                "v1",
                "live",
                "{live_id}",
                "transit",
                "trips",
                "{trip_id}"
              ],
              "rename": {
                "param": {
                  "city": "live_id"
                }
              },
              "select": {
                "exist": [
                  "live_id",
                  "trip_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "station",
                    "orig": "station",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{slug}/departures",
              "parts": [
                "api",
                "v1",
                "live",
                "{slug}",
                "departures"
              ],
              "select": {
                "$action": "departure",
                "exist": [
                  "slug",
                  "station"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{slug}/air",
              "parts": [
                "api",
                "v1",
                "live",
                "{slug}",
                "air"
              ],
              "select": {
                "$action": "air",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 4
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{slug}/air-uba",
              "parts": [
                "api",
                "v1",
                "live",
                "{slug}",
                "air-uba"
              ],
              "select": {
                "$action": "air_uba",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 5
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "city",
                    "orig": "city",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{city}/baustellen",
              "parts": [
                "api",
                "v1",
                "live",
                "{city}",
                "baustellen"
              ],
              "select": {
                "$action": "baustellen",
                "exist": [
                  "city"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 6
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "city",
                    "orig": "city",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{city}/ereignisse",
              "parts": [
                "api",
                "v1",
                "live",
                "{city}",
                "ereignisse"
              ],
              "select": {
                "$action": "ereignisse",
                "exist": [
                  "city"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 7
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{slug}/flood",
              "parts": [
                "api",
                "v1",
                "live",
                "{slug}",
                "flood"
              ],
              "select": {
                "$action": "flood",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 8
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{slug}/traffic",
              "parts": [
                "api",
                "v1",
                "live",
                "{slug}",
                "traffic"
              ],
              "select": {
                "$action": "traffic",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 9
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "city",
                    "orig": "city",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{city}/traffic-flow",
              "parts": [
                "api",
                "v1",
                "live",
                "{city}",
                "traffic-flow"
              ],
              "select": {
                "$action": "traffic_flow",
                "exist": [
                  "city"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 10
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{slug}/water-level",
              "parts": [
                "api",
                "v1",
                "live",
                "{slug}",
                "water-level"
              ],
              "select": {
                "$action": "water_level",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 11
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "slug",
                    "orig": "slug",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/{slug}/webcams",
              "parts": [
                "api",
                "v1",
                "live",
                "{slug}",
                "webcams"
              ],
              "select": {
                "$action": "webcam",
                "exist": [
                  "slug"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 12
            },
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": "Frankfurt (Main) Hauptbahnhof",
                    "kind": "query",
                    "name": "station",
                    "orig": "station",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/frankfurt-am-main/departures",
              "parts": [
                "api",
                "v1",
                "live",
                "frankfurt-am-main",
                "departures"
              ],
              "select": {
                "exist": [
                  "station"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 13
            },
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": "Hamburg Hauptbahnhof",
                    "kind": "query",
                    "name": "station",
                    "orig": "station",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/hamburg/departures",
              "parts": [
                "api",
                "v1",
                "live",
                "hamburg",
                "departures"
              ],
              "select": {
                "exist": [
                  "station"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 14
            },
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": "510",
                    "kind": "query",
                    "name": "stop_id",
                    "orig": "stop_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/live/nuernberg/departures",
              "parts": [
                "api",
                "v1",
                "live",
                "nuernberg",
                "departures"
              ],
              "select": {
                "exist": [
                  "stop_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 15
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/berlin/verkehrsmeldungen",
              "parts": [
                "api",
                "v1",
                "live",
                "berlin",
                "verkehrsmeldungen"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 16
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/dortmund/parking",
              "parts": [
                "api",
                "v1",
                "live",
                "dortmund",
                "parking"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 17
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/eround/charging",
              "parts": [
                "api",
                "v1",
                "live",
                "eround",
                "charging"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 18
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/frankfurt-am-main/parking",
              "parts": [
                "api",
                "v1",
                "live",
                "frankfurt-am-main",
                "parking"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 19
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/hamburg/verkehrslage",
              "parts": [
                "api",
                "v1",
                "live",
                "hamburg",
                "verkehrslage"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 20
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/hannover/verkehrsmeldungen",
              "parts": [
                "api",
                "v1",
                "live",
                "hannover",
                "verkehrsmeldungen"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 21
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/kiel/zaehlstellen",
              "parts": [
                "api",
                "v1",
                "live",
                "kiel",
                "zaehlstellen"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 22
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/koeln/umweltzone",
              "parts": [
                "api",
                "v1",
                "live",
                "koeln",
                "umweltzone"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 23
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/magdeburg/parking",
              "parts": [
                "api",
                "v1",
                "live",
                "magdeburg",
                "parking"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 24
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/live/wuppertal/parking",
              "parts": [
                "api",
                "v1",
                "live",
                "wuppertal",
                "parking"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 25
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": [
          [
            "live"
          ],
          [
            "live",
            "route"
          ],
          [
            "live",
            "trip"
          ]
        ]
      }
    },
    "meta": {
      "fields": [
        {
          "active": true,
          "name": "breaker_state",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "enabled",
          "req": true,
          "type": "`$BOOLEAN`",
          "index$": 1
        },
        {
          "active": true,
          "name": "source",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        }
      ],
      "name": "meta",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "header": [
                  {
                    "active": true,
                    "kind": "header",
                    "name": "if_none_match",
                    "orig": "if_none_match",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": 50,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 0,
                    "kind": "query",
                    "name": "offset",
                    "orig": "offset",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": "asc",
                    "kind": "query",
                    "name": "order",
                    "orig": "order",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/sources",
              "parts": [
                "api",
                "v1",
                "sources"
              ],
              "select": {
                "exist": [
                  "if_none_match",
                  "limit",
                  "offset",
                  "order",
                  "page",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meta`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/openapi.yaml",
              "parts": [
                "api",
                "v1",
                "openapi.yaml"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "station": {
      "fields": [
        {
          "active": true,
          "name": "data",
          "req": true,
          "type": "`$ANY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "meta",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 1
        }
      ],
      "name": "station",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "eva",
                    "orig": "eva",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/stations/{eva}/arrivals",
              "parts": [
                "api",
                "v1",
                "stations",
                "{eva}",
                "arrivals"
              ],
              "select": {
                "$action": "arrival",
                "exist": [
                  "eva"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "eva",
                    "orig": "eva",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/stations/{eva}/departures",
              "parts": [
                "api",
                "v1",
                "stations",
                "{eva}",
                "departures"
              ],
              "select": {
                "$action": "departure",
                "exist": [
                  "eva"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": [
          [
            "station"
          ]
        ]
      }
    }
  }
}


const config = new Config()

export {
  config
}

