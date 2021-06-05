
APIs Documentation
===================

### GET /serverReady
This is the very firts API that returns `true` if the API server is ready to be used. 

Why do we need it?
Becuse docker runs the containers in no particular order so we might end up running the `client` container before the `server` container 
and moreover server needs some initilization on its first launch, so it takes some time. Therefore, we ask `client` to wait for the __ready__ signal from the server.

#### Call Example:
`curl -X GET -H 'Content-Type: application/json' -i http://localhost:8080/serverReady`

__Output:__
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 05 Jun 2021 21:11:03 GMT
Content-Length: 4

true
```

------------

### GET /clients

This API retrieves all the clients with pagination. Please note that to call this API you need to send your API key in the header as well. 
For sake of simplicity, as requested, we use a Static API Key which by default is `Prefix.HashOfSomeSecretKey` and it is configurable by editing `.env` file in the project's root.

#### Call Example:
`curl -X GET -H 'Content-Type: application/json' -H 'X-API-KEY: Prefix.HashOfSomeSecretKey'  -i http://localhost:8080/clients`

__Output:__
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 05 Jun 2021 21:15:17 GMT
Transfer-Encoding: chunked

{
  "pagination": {
    "current_page": 1,
    "total_entries": 100,
    "total_pages": 1
  },
  "rows": [
    {
      "name": "Kerrie Stockstill",
      "uuid": "f8d2405d-ab7c-4e2a-8790-ab8e4d68d213"
    },
    {
      "name": "Stanley Caples",
      "uuid": "df48bf51-bf3f-4fbe-a4e0-5ff8cb98d2f8"
    },
    {
      "name": "Caryl Schmalz",
      "uuid": "1f61ac75-9183-4ad4-978e-1560842cd301"
    },
    ...
  ]
}
```

------------

### GET /scooters

This API retrieves all the scooters with pagination. 

#### Call Example:
`curl -X GET -H 'Content-Type: application/json' -H 'X-API-KEY: Prefix.HashOfSomeSecretKey'  -i http://localhost:8080/scooters`

__Output:__
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 05 Jun 2021 21:19:14 GMT
Transfer-Encoding: chunked

{
  "pagination": {
    "current_page": 1,
    "total_entries": 500,
    "total_pages": 3
  },
  "rows": [
    {
      "last_update": "2021-06-05T18:27:05.568591Z",
      "lat": 51.043430147739755,
      "lon": 13.769144774215633,
      "occupied": false,
      "uuid": "a6739cc5-a016-4d9b-81e8-dc32e90dc3a4"
    },
    {
      "last_update": "2021-06-05T18:27:05.569755Z",
      "lat": 51.02161991146944,
      "lon": 13.770331742066585,
      "occupied": false,
      "uuid": "3e8d8571-fbdd-4c1a-8d54-c956758c3b97"
    },
    ...
  ]
}
```

----------

### GET /scooters/:uuid

This API retrieves the scooter with the given __UUID__ and returnd _404_ error if not found. 

#### Call Example:
`curl -X GET -H 'Content-Type: application/json' -H 'X-API-KEY: Prefix.HashOfSomeSecretKey'  -i http://localhost:8080/scooters/a6739cc5-a016-4d9b-81e8-dc32e90dc3a4`

__Output:__
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 05 Jun 2021 21:21:57 GMT
Content-Length: 179

{
  "last_update": "2021-06-05T18:27:05.568591Z",
  "lat": 51.043430147739755,
  "lon": 13.769144774215633,
  "occupied": false,
  "uuid": "a6739cc5-a016-4d9b-81e8-dc32e90dc3a4"
}
```

----------


### GET /scooters/:uuid/location

This API retrieves the last known location of the scooter with the given __UUID__ and returnd _404_ error if not found.
Please note that, this location is for a moving scooter the stationary scooters' location can be retrieved with the previous API.

#### Call Example:
`curl -X GET -H 'Content-Type: application/json' -H 'X-API-KEY: Prefix.HashOfSomeSecretKey'  -i http://localhost:8080/scooters/cb7fb6f7-7dfd-4779-b266-003736f17e1a/location`

__Output:__
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 05 Jun 2021 21:31:21 GMT
Content-Length: 99

{
  "lat": 51.04386284695899,
  "lon": 13.75715747707119,
  "time": "2021-06-05T18:27:08.519826Z"
}
```

----------


### POST /scooters/:uuid/location

This API updates the location of a moving scooter with the given __UUID__. This API is meant to be called by the scooter itself every 3 seconds.

#### Input Format:
```
{
		"lat": <Float Number>,
		"lon": <Float Number>
}
```
#### Call Example:
`curl -X POST -H 'Content-Type: application/json' -H 'X-API-KEY: Prefix.HashOfSomeSecretKey'  -i http://localhost:8080/scooters/cb7fb6f7-7dfd-4779-b266-003736f17e1a/location --data '{ "lat": 51.05, "lon": 13.71}'`

__Output:__
```
HTTP/1.1 200 OK
Date: Sat, 05 Jun 2021 21:34:10 GMT
Content-Length: 0
```

----------


### POST /search/freeScooters

This API receives two pairs of coordinations and searches in a semi-rectangular area for free (available) scooters.
Note: This API returns a specific maximum number of scooters (_default 200 but is configurable_).

#### Input Format:
```
{
	"start": {
		"lat": <Float Number>,
		"lon": <Float Number>
	},
	"end": {
		"lat": <Float Number>,
		"lon": <Float Number>
	}
}
```

#### Call Example:
`curl -X POST -H 'Content-Type: application/json' -H 'X-API-KEY: Prefix.HashOfSomeSecretKey'  -i http://localhost:8080/search/freeScooters --data '{"start": {"lat": 50.01,"lon": 12.11},"end": {"lat": 53.16,"lon": 14.67}}'`

__Output:__
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 05 Jun 2021 21:40:03 GMT
Transfer-Encoding: chunked

[
  {
    "lat": 51.043430147739755,
    "lon": 13.769144774215633,
    "uuid": "a6739cc5-a016-4d9b-81e8-dc32e90dc3a4"
  },
  {
    "lat": 51.04222290097839,
    "lon": 13.756148248464667,
    "uuid": "a7560273-2eac-401c-b051-a95d5e6eecaf"
  },
  {
    "lat": 51.02161991146944,
    "lon": 13.770331742066585,
    "uuid": "3e8d8571-fbdd-4c1a-8d54-c956758c3b97"
  },
  ...
]
```

----------


### GET /search/movingScooters

This API retrieves the last known location of all the moving scooters.

_Note 1: This API is used only for demo visulization purpose of this project._

_Note 2: This API does not use any API Key._

#### Call Example:
`curl -X GET -H 'Content-Type: application/json' -i http://localhost:8080/search/movingScooters`

__Output:__
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 05 Jun 2021 21:45:38 GMT
Transfer-Encoding: chunked

[
  {
    "lat": 51.03551597910139,
    "lon": 13.764477488309588,
    "time": "2021-06-05T21:45:37.975103Z",
    "uuid": "df5ab223-2307-4ec6-bf86-fe3208974560"
  },
  {
    "lat": 51.042582329008674,
    "lon": 13.759892094880684,
    "time": "2021-06-05T18:27:20.615891Z",
    "uuid": "fa990519-d77b-4dbd-a8fe-95cb4c8320a6"
  },
  ...
]
```

----------


### POST /tripStart

This API receives the __UUID__ of a particular scooter and the __UUID__ of the user (client) who is using it and starts a trip.
After successfullt starting a trip, this API genrates and returns a __UUID__ of the trip to be used in future for tracking the trip.

#### Input Format:
```
{
	"scooter_uuid": "<String>",
	"user_uuid": "<String>",
	"start": {
		"lat": <Float Number>,
		"lon": <Float Number>
	}
}
```

#### Call Example:
`curl -X POST -H 'Content-Type: application/json' -H 'X-API-KEY: Prefix.HashOfSomeSecretKey'  -i http://localhost:8080/tripStart --data '{"scooter_uuid": "a6739cc5-a016-4d9b-81e8-dc32e90dc3a4","user_uuid": "f8d2405d-ab7c-4e2a-8790-ab8e4d68d213","start": {"lat": 50.01,"lon": 12.11}}'`

__Output:__
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 05 Jun 2021 21:53:51 GMT
Content-Length: 38

"7cdbee24-271a-4ed6-b019-5c45dde76dcb"
```

----------


### POST /tripEnd

This API receives the __UUID__ of a particular scooter and the __UUID__ of the user (client) who is using it and ends the ongoing trip.
After successfullt ending the trip, it updates the stationary location of the scooter and marks it as free to use.

#### Input Format:
```
{
	"scooter_uuid": "<String>",
	"user_uuid": "<String>",
	"end": {
		"lat": <Float Number>,
		"lon": <Float Number>
	}
}
```

#### Call Example:
`curl -X POST -H 'Content-Type: application/json' -H 'X-API-KEY: Prefix.HashOfSomeSecretKey'  -i http://localhost:8080/tripEnd --data '{"scooter_uuid": "a6739cc5-a016-4d9b-81e8-dc32e90dc3a4","user_uuid": "f8d2405d-ab7c-4e2a-8790-ab8e4d68d213","end": {"lat": 50.01,"lon": 12.11}}'`

__Output:__
```
HTTP/1.1 200 OK
Date: Sat, 05 Jun 2021 21:59:11 GMT
Content-Length: 0
```

----------
