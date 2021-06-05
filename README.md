# NordSec Scootin

## Quick Setup

To run the App simply copy the following code into your terminal:

```
git clone https://github.com/mojtaba-esk/NordSec-Scootin.git
cd NordSec-Scootin
sudo docker-compose up -d
```
Note: _The database, tables, indices, etc will be created automatically and will be filled with some random data._

Once it is up, you can see the status of moving scooters in your browser: http://localhost:8080/

![Moving scooters with 50 Random clients](demo.gif "Moving scooters with 50 Random clients")



## Development
To activate the development mode you need to open `docker-compose.yml` file, under the desired service (e.g. server), change the target to development:

```
    build:
      context: ./client
      target: development  # development | test | production (default)
```

## Test
Please Change the target to `test` then build the container:

`sudo docker-compose up --build` 

You will see the test result in the terminal.

## Logs
See API server logs:
`sudo docker logs -f scootin-api-server`

See the dummy client logs:
`sudo docker logs -f scootin-dummy-client`


## ENV variables
### Server:
- `SERVING_ADDR`: Service address for the API server
- `STATIC_API_KEY`: Static API Key to Autheticate to the APIs

- `REDIS_USER`: Redis Username (_Please note that the current version of the App does not use redis_)
- `REDIS_PASSWORD`: Redis Password
- `REDIS_HOST`: Redis Host
- `REDIS_PORT`: Redis Port
- `REDIS_DB`: Redis Database number (_default:_ 0)

- `POSTGRES_DB`: PostgreSQL database name
- `POSTGRES_USER`: PostgreSQL username with correct autorizations
- `POSTGRES_PASSWORD`: PostgreSQL password
- `POSTGRES_PORT`: PostgreSQL port
- `POSTGRES_HOST`: PostgreSQL Hostname

### Dummy Client:
- `SCOOTIN_API_PATH`: The path to the API server
- `STATIC_API_KEY`: Static API Key to Autheticate to the APIs
- `NUM_OF_CLIENTS`: Number of dummy clients

If you change ech of these values in the `docker-compose.yml` file, you need to run the follwoing command in order to reflect the changes: `sudo docker-compose up -d`

PS: this command must be executed in the same directory where `docker-compose.yml` file is located (_e.g. NordSec-Scootin_).
