# InfranodeOpenData SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "InfranodeOpenData",
            "slug": "infranode-open-data",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
          "factor": 2,
          "maxDelay": 2000,
          "minDelay": 50,
          "retries": 2,
          "statuses": [
            408,
            425,
            429,
            500,
            502,
            503,
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://infranode.dev",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "city": {},
                "compare": {},
                "health": {},
                "live": {},
                "meta": {},
                "station": {},
            },
        },
        "entity": {
      "city": {
        "fields": [
          {
            "name": "data",
            "title": "Data",
            "type": "`$ANY`",
            "req": True,
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
          },
          {
            "name": "meta",
            "title": "Meta",
            "type": "`$OBJECT`",
            "req": True,
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "city",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/public-tenders",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "public-tenders",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "public-tenders",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 50,
                    },
                    {
                      "name": "match",
                      "orig": "match",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 0,
                    },
                    {
                      "name": "q",
                      "orig": "q",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "since",
                      "orig": "since",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "status",
                      "orig": "status",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "$action": "public_tender",
                  "exist": [
                    "limit",
                    "match",
                    "offset",
                    "q",
                    "since",
                    "slug",
                    "status",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/council-papers",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "council-papers",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "council-papers",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 50,
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 0,
                    },
                    {
                      "name": "paper_type",
                      "orig": "paper_type",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "q",
                      "orig": "q",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "since",
                      "orig": "since",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "$action": "council_paper",
                  "exist": [
                    "limit",
                    "offset",
                    "paper_type",
                    "q",
                    "since",
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/transit",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "transit",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "transit",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "near",
                      "orig": "near",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "page",
                      "orig": "page",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "q",
                      "orig": "q",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "radius_m",
                      "orig": "radius_m",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 1000,
                    },
                  ],
                },
                "select": {
                  "$action": "transit",
                  "exist": [
                    "limit",
                    "near",
                    "page",
                    "q",
                    "radius_m",
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/tenders",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "tenders",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "tenders",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "query": [
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 50,
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 0,
                    },
                    {
                      "name": "q",
                      "orig": "q",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "since",
                      "orig": "since",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "status",
                      "orig": "status",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "limit",
                    "offset",
                    "q",
                    "since",
                    "status",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/stations",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "stations",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "stations",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "q",
                      "orig": "q",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "$action": "station",
                  "exist": [
                    "limit",
                    "q",
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/traffic",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "traffic",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "traffic",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "full",
                      "orig": "full",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                    {
                      "name": "include",
                      "orig": "include",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "$action": "traffic",
                  "exist": [
                    "full",
                    "include",
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/pois",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "pois",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "pois",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "type",
                      "orig": "type",
                      "type": "`$STRING`",
                      "kind": "query",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "poi",
                  "exist": [
                    "slug",
                    "type",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "slug": "id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/accidents",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "accidents",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "accidents",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "accident",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/air",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "air",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "air",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "air",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/air-uba",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "air-uba",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "air-uba",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "air_uba",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/base",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "base",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "base",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "base",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/bathing-water",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "bathing-water",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "bathing-water",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "bathing_water",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/bike-counts",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "bike-counts",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "bike-counts",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "bike_count",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/business-registrations",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "business-registrations",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "business-registrations",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "business_registration",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/charging",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "charging",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "charging",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "charging",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/charging-status",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "charging-status",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "charging-status",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "charging_status",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/civil-protection-warnings",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "civil-protection-warnings",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "civil-protection-warnings",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "civil_protection_warning",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/construction",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "construction",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "construction",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "construction",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/crime-stats",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "crime-stats",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "crime-stats",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "crime_stat",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/demographics",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "demographics",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "demographics",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "demographic",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/district-heating",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "district-heating",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "district-heating",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "district_heating",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/drinking-water",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "drinking-water",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "drinking-water",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "drinking_water",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/education",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "education",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "education",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "education",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/election",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "election",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "election",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "election",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/energy",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "energy",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "energy",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "energy",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/events",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "events",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "events",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "event",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/fire-danger",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "fire-danger",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "fire-danger",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "fire_danger",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/flood",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "flood",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "flood",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "flood",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/fuel-prices",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "fuel-prices",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "fuel-prices",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "fuel_price",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/geo",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "geo",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "geo",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "geo",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/government-offices",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "government-offices",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "government-offices",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "government_office",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/health",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "health",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "health",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "health",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/heritage",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "heritage",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "heritage",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "heritage",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/holidays",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "holidays",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "holidays",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "holiday",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/hospitals-atlas",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "hospitals-atlas",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "hospitals-atlas",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "hospitals_atla",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/icu-live",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "icu-live",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "icu-live",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "icu_live",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/indicators",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "indicators",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "indicators",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "indicator",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/insolvencies",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "insolvencies",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "insolvencies",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "insolvency",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/land-values",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "land-values",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "land-values",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "land_value",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/markets",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "markets",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "markets",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "market",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/office-wait-times",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "office-wait-times",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "office-wait-times",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "office_wait_time",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/overview",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "overview",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "overview",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "overview",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/parcel-lockers",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "parcel-lockers",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "parcel-lockers",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "parcel_locker",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/parking",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "parking",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "parking",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "parking",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/playgrounds",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "playgrounds",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "playgrounds",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "playground",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/pollen-uv",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "pollen-uv",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "pollen-uv",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "pollen_uv",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/population-density",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "population-density",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "population-density",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "population_density",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/post-boxes",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "post-boxes",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "post-boxes",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "post_box",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/post-offices",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "post-offices",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "post-offices",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "post_office",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/power-load",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "power-load",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "power-load",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "power_load",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/power-price",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "power-price",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "power-price",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "power_price",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/public-toilets",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "public-toilets",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "public-toilets",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "public_toilet",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/public-wifi",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "public-wifi",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "public-wifi",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "public_wifi",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/recycling-centres",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "recycling-centres",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "recycling-centres",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "recycling_centre",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/road-events",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "road-events",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "road-events",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "road_event",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/sharing",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "sharing",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "sharing",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "sharing",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/solar",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "solar",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "solar",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "solar",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/solar-roofs",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "solar-roofs",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "solar-roofs",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "solar_roof",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/station-arrivals",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "station-arrivals",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "station-arrivals",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "station_arrival",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/station-departures",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "station-departures",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "station-departures",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "station_departure",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/station-facilities",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "station-facilities",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "station-facilities",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "station_facility",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/tax-rates",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "tax-rates",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "tax-rates",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "tax_rate",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/tourism",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "tourism",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "tourism",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "tourism",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/tree-cadastre",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "tree-cadastre",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "tree-cadastre",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "tree_cadastre",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/unemployment",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "unemployment",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "unemployment",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "unemployment",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/vehicle-registrations",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "vehicle-registrations",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "vehicle-registrations",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "vehicle_registration",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/water-level",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "water-level",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "water-level",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "water_level",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/weather",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "weather",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "weather",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "weather",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/weather-warnings",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "weather-warnings",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "weather-warnings",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "weather_warning",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/cities/{slug}/webcams",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "cities",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "webcams",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "cities",
                  "{slug}",
                  "webcams",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "webcam",
                  "exist": [
                    "slug",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "compare": {
        "fields": [
          {
            "name": "city",
            "title": "City",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "data",
            "title": "Data",
            "type": "`$OBJECT`",
          },
          {
            "name": "source_status",
            "title": "Source Status",
            "type": "`$STRING`",
            "req": True,
          },
        ],
        "name": "compare",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/compare",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "compare",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "compare",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "header": [
                    {
                      "name": "if_none_match",
                      "orig": "if_none_match",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "city",
                      "orig": "city",
                      "type": "`$STRING`",
                      "kind": "query",
                      "reqd": True,
                      "example": "berlin,koeln,hamburg",
                    },
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 50,
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 0,
                    },
                    {
                      "name": "order",
                      "orig": "order",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "asc",
                    },
                    {
                      "name": "page",
                      "orig": "page",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 1,
                    },
                    {
                      "name": "resource",
                      "orig": "resource",
                      "type": "`$STRING`",
                      "kind": "query",
                      "reqd": True,
                    },
                    {
                      "name": "sort",
                      "orig": "sort",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "city",
                    "if_none_match",
                    "limit",
                    "offset",
                    "order",
                    "page",
                    "resource",
                    "sort",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "health": {
        "fields": [
          {
            "name": "redis",
            "title": "Redis",
            "type": "`$BOOLEAN`",
            "req": True,
            "short": "true wenn Redis erreichbar (Ping erfolgreich)",
          },
          {
            "name": "status",
            "title": "Status",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "version",
            "title": "Version",
            "type": "`$STRING`",
            "req": True,
          },
        ],
        "name": "health",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/health",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "health",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "health",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "live": {
        "fields": [
          {
            "name": "data",
            "title": "Data",
            "type": "`$ANY`",
            "req": True,
          },
          {
            "name": "meta",
            "title": "Meta",
            "type": "`$OBJECT`",
            "req": True,
          },
        ],
        "name": "live",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{city}/transit/routes/{route_id}/status",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "live_id",
                  },
                  {
                    "lit": "transit",
                  },
                  {
                    "lit": "routes",
                  },
                  {
                    "var": "route_id",
                  },
                  {
                    "lit": "status",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{live_id}",
                  "transit",
                  "routes",
                  "{route_id}",
                  "status",
                ],
                "rename": {
                  "param": {
                    "city": "live_id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "live_id",
                      "orig": "city",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                    {
                      "name": "route_id",
                      "orig": "route_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "live_id",
                    "route_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{city}/transit/departures",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "live_id",
                  },
                  {
                    "lit": "transit",
                  },
                  {
                    "lit": "departures",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{live_id}",
                  "transit",
                  "departures",
                ],
                "rename": {
                  "param": {
                    "city": "live_id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "live_id",
                      "orig": "city",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "stop_id",
                      "orig": "stop_id",
                      "type": "`$STRING`",
                      "kind": "query",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "transit_departure",
                  "exist": [
                    "live_id",
                    "stop_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{city}/transit/trips/{trip_id}",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "live_id",
                  },
                  {
                    "lit": "transit",
                  },
                  {
                    "lit": "trips",
                  },
                  {
                    "var": "trip_id",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{live_id}",
                  "transit",
                  "trips",
                  "{trip_id}",
                ],
                "rename": {
                  "param": {
                    "city": "live_id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "live_id",
                      "orig": "city",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                    {
                      "name": "trip_id",
                      "orig": "trip_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "live_id",
                    "trip_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{slug}/departures",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "departures",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{slug}",
                  "departures",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                  "query": [
                    {
                      "name": "station",
                      "orig": "station",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "$action": "departure",
                  "exist": [
                    "slug",
                    "station",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{slug}/air",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "air",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{slug}",
                  "air",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "air",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{slug}/air-uba",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "air-uba",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{slug}",
                  "air-uba",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "air_uba",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{city}/baustellen",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "city",
                  },
                  {
                    "lit": "baustellen",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{city}",
                  "baustellen",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "city",
                      "orig": "city",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "baustellen",
                  "exist": [
                    "city",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{city}/ereignisse",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "city",
                  },
                  {
                    "lit": "ereignisse",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{city}",
                  "ereignisse",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "city",
                      "orig": "city",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "ereignisse",
                  "exist": [
                    "city",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{slug}/flood",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "flood",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{slug}",
                  "flood",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "flood",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{slug}/traffic",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "traffic",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{slug}",
                  "traffic",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "traffic",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{city}/traffic-flow",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "city",
                  },
                  {
                    "lit": "traffic-flow",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{city}",
                  "traffic-flow",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "city",
                      "orig": "city",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "traffic_flow",
                  "exist": [
                    "city",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{slug}/water-level",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "water-level",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{slug}",
                  "water-level",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "water_level",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/{slug}/webcams",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "var": "slug",
                  },
                  {
                    "lit": "webcams",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "{slug}",
                  "webcams",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "slug",
                      "orig": "slug",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "webcam",
                  "exist": [
                    "slug",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/frankfurt-am-main/departures",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "frankfurt-am-main",
                  },
                  {
                    "lit": "departures",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "frankfurt-am-main",
                  "departures",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "query": [
                    {
                      "name": "station",
                      "orig": "station",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "Frankfurt (Main) Hauptbahnhof",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "station",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/hamburg/departures",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "hamburg",
                  },
                  {
                    "lit": "departures",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "hamburg",
                  "departures",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "query": [
                    {
                      "name": "station",
                      "orig": "station",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "Hamburg Hauptbahnhof",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "station",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/nuernberg/departures",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "nuernberg",
                  },
                  {
                    "lit": "departures",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "nuernberg",
                  "departures",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "query": [
                    {
                      "name": "stop_id",
                      "orig": "stop_id",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "510",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "stop_id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/berlin/verkehrsmeldungen",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "berlin",
                  },
                  {
                    "lit": "verkehrsmeldungen",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "berlin",
                  "verkehrsmeldungen",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/dortmund/parking",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "dortmund",
                  },
                  {
                    "lit": "parking",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "dortmund",
                  "parking",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/eround/charging",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "eround",
                  },
                  {
                    "lit": "charging",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "eround",
                  "charging",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/frankfurt-am-main/parking",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "frankfurt-am-main",
                  },
                  {
                    "lit": "parking",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "frankfurt-am-main",
                  "parking",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/hamburg/verkehrslage",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "hamburg",
                  },
                  {
                    "lit": "verkehrslage",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "hamburg",
                  "verkehrslage",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/hannover/verkehrsmeldungen",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "hannover",
                  },
                  {
                    "lit": "verkehrsmeldungen",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "hannover",
                  "verkehrsmeldungen",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/kiel/zaehlstellen",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "kiel",
                  },
                  {
                    "lit": "zaehlstellen",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "kiel",
                  "zaehlstellen",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/koeln/umweltzone",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "koeln",
                  },
                  {
                    "lit": "umweltzone",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "koeln",
                  "umweltzone",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/magdeburg/parking",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "magdeburg",
                  },
                  {
                    "lit": "parking",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "magdeburg",
                  "parking",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/live/wuppertal/parking",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "live",
                  },
                  {
                    "lit": "wuppertal",
                  },
                  {
                    "lit": "parking",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "live",
                  "wuppertal",
                  "parking",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "meta": {
        "fields": [
          {
            "name": "breaker_state",
            "title": "Breaker State",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "enabled",
            "title": "Enabled",
            "type": "`$BOOLEAN`",
            "req": True,
          },
          {
            "name": "source",
            "title": "Source",
            "type": "`$STRING`",
            "req": True,
          },
        ],
        "name": "meta",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/sources",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "sources",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "sources",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meta`",
                },
                "args": {
                  "header": [
                    {
                      "name": "if_none_match",
                      "orig": "if_none_match",
                      "type": "`$STRING`",
                      "kind": "header",
                    },
                  ],
                  "query": [
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 50,
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 0,
                    },
                    {
                      "name": "order",
                      "orig": "order",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "asc",
                    },
                    {
                      "name": "page",
                      "orig": "page",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "example": 1,
                    },
                    {
                      "name": "sort",
                      "orig": "sort",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "if_none_match",
                    "limit",
                    "offset",
                    "order",
                    "page",
                    "sort",
                  ],
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/openapi.yaml",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "openapi.yaml",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "openapi.yaml",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "station": {
        "fields": [],
        "name": "station",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/stations/{eva}/arrivals",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "stations",
                  },
                  {
                    "var": "eva",
                  },
                  {
                    "lit": "arrivals",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "stations",
                  "{eva}",
                  "arrivals",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "eva",
                      "orig": "eva",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "arrival",
                  "exist": [
                    "eva",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "GET",
                "orig": "/api/v1/stations/{eva}/departures",
                "segments": [
                  {
                    "lit": "api",
                  },
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "stations",
                  },
                  {
                    "var": "eva",
                  },
                  {
                    "lit": "departures",
                  },
                ],
                "parts": [
                  "api",
                  "v1",
                  "stations",
                  "{eva}",
                  "departures",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "eva",
                      "orig": "eva",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "departure",
                  "exist": [
                    "eva",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
