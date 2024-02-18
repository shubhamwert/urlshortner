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
## Run as docker
```bash
docker run -it -p 9080:9080 dockerofshubham/go-urlshortner 

```
## Mount Config
```bash
docker run -it -v <path-to-config>:/configs dockerofshubham/go-urlshortner 

```


## Build
```bash
cd urlshortner
mkdir -p ../releases/size-reduced-$(date +%F) 
go build -ldflags="-s -w" -gcflags=all="-1 -B"   -o ../releases/size-reduced-$(date +%F)/urlshortner .

```
## Run with prebuild 
```bash
./urlshortner --configPath <config-path> --configName <config-name>
```


## Run
```bash
cd urlshortner

go run .

```
## Local Development (DB setup)
Docker Compose
```bash
cd db
docker compose up
```


Local DynamoDb devlopment
```
aws dynamodb create-table \  
  --table-name UrlDb \
  --attribute-definitions \
      AttributeName=encodedURL,AttributeType=S \
      AttributeName=userId,AttributeType=S \
  --key-schema \
      AttributeName=encodedURL,KeyType=HASH \
      AttributeName=userId,KeyType=RANGE \
  --billing-mode PROVISIONED \
  --provisioned-throughput \
      ReadCapacityUnits=10,WriteCapacityUnits=10 \
  --endpoint-url=http://localhost:8000

```
