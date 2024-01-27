### Description
Simple url shortner written in go with plug n play support for multiple db


### Create Databases
```bash
cd db
docker compose up
```

## Run the service

### Use Latest release
```bash
cd urlshortner
```


## Build
```bash
cd urlshortner
mkdir -p ../releases/$(date +%F) 
go build -ldflags="-s -w"  -o ../releases/$(date +%F)/urlshortner .

```

## Run
```bash
cd urlshortner

go run .

```