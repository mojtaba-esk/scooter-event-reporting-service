
APIs Documentation
===================

### /serverReady
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

### /clients

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

### /scooters

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




`curl -X POST -H 'Content-Type: application/json' -H 'X-API-KEY: myKey' -i http://localhost:8080/clients --data '{ "lat" :555}'`
