# NordSec Scootin

## Quick Setup:

To run the App simply copy the following code into your terminal:

```
git clone https://github.com/mojtaba-esk/NordSec-Scootin.git
cd NordSec-Scootin
sudo docker-compose up -d
```
Note: _The database, tables, indices, etc will be created automatically and will be filled with some random data._

Once it is up, you can see the status of moving scooters in your browser: http://localhost:8080/

![Moving scooters with 50 Random clients](demo.gif "Moving scooters with 50 Random clients")

## Build from source code:

```
git clone https://github.com/mojtaba-esk/NordSec-Scootin.git
cd NordSec-Scootin
sudo docker-compose up -d --build
```

## Development:
To activate the development mode you need to open `docker-compose.yml` file, under the desired service (e.g. server), change the target to development:

```
    build:
      context: ./client
      target: development  # development | test | production (default)
```

## Test:
Please Change the target to `test` then build the container:

`sudo docker-compose up --build` 

You will see the test result in the terminal.

