你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
<!-- ![Build Status](http://34.74.167.3/api/badges/erkanzileli/go-auth/status.svg) -->

# Authentication Server

![Project Architecture](https://github.com/erkanzileli/go-auth/raw/master/static/architecture.png)

## Environment Variables

| Variable Name | Type        | Required | Default         | Description                                 |
| ------------- | ----------- | -------- | --------------- | ------------------------------------------- |
| DB_URL        | string      | true     | nil             | Database URL                                |
| DB_NAME       | string      | true     | nil             | Database Name                               |
| DB_COLLECTION | string      | true     | nil             | Name of Mongo Collection of given Database  |
| DB_USERNAME   | string      | false    | nil             | Username credential for Database Connection |
| DB_PASSWORD   | string      | false    | nil             | Password credential for Database Connection |
| JWT_SECRET    | string      | false    | "secret-string" | JWT Secret for encryption or decryption     |
| JWT_EXPIRE    | int(minute) | false    | 5 minute        | The period of validity of taken JWT         |

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
   - Check URL
   - Check MongoDB whitelist
   - Check MongoDB mongodb.conf:
     `sudo vim /etc/mongod.conf`
     Set bindIp to 0.0.0.0
