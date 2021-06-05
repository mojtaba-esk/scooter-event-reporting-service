# NordSec Scootin

## Quick Setup:

To run the App simply copy the following code into your terminal:

```
git clone https://github.com/mojtaba-esk/NordSec-Scootin.git
cd NordSec-Scootin
sudo docker-compose up -d
```

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

