![Build Status](http://35.246.226.178/api/badges/kalderasoft/go-auth/status.svg)

# Authentication Server

![Project Architecture](https://github.com/kalderasoft/go-auth/raw/master/static/architecture.png)

## Environment Variables

|Variable Name|Type|Required|Default|Description|
|---|---|---|---|---|
|DB_URL|string|true|nil|Database URL|
|DB_NAME|string|true|nil|Database Name|
|DB_COLLECTION|string|true|nil|Name of Mongo Collection of given Database|
|DB_USERNAME|string|false|nil|Username credential for Database Connection|
|DB_PASSWORD|string|false|nil|Password credential for Database Connection|
|JWT_SECRET|string|false|"secret-string"|JWT Secret for encryption or decryption|
|JWT_EXPIRE|int(minute)|false|5 minute|The period of validity of taken JWT|

## Example Scenarios

Firstly, build image.
```
docker build -t go-auth:build .
```

1. Local MongoDB server. For development or trying. No basic auth.

```
docker run --detach --publish 8000:80 \ 
-e DB_URL=localhost:27017 -e DB_NAME=auth \
-e DB_COLLECTION=users -e JWT_SECRET=myAwesomeSecret \
-e JWT_EXPIRE=60 go-auth:build
```

2. Local MongoDB server. For development or trying. With basic auth.
```
docker run --detach --publish 8000:80 \ 
-e DB_URL=localhost:27017 -e DB_NAME=auth \
-e DB_COLLECTION=users -e DB_USERNAME=admin \
-e DB_PASSWORD=admin -e JWT_SECRET=myAwesomeSecret \
-e JWT_EXPIRE=60 go-auth:build
```

## Possible problems
1. Dep install command. Maybe this command not works on the future. 
Because it gets a static file.

2. MongoDB Server access problems.
    * Check URL
    * Check MongoDB whitelist
    * Check MongoDB mongodb.conf: 
    ``sudo vim /etc/mongod.conf``
    Set bindIp to 0.0.0.0