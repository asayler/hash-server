Hash Server
-----------

**_A simple hashing microservice_**

Hashes (and optionally salts) passwords using one or more rounds of
SHA512. Inserts an arbitrary delay into the hashing path to ease
concurrency analysis.

### Authorship ###

By [Andy Sayler](https://www.andysayler.com)

### Dependencies ###

+ https://github.com/braintree/manners

### Setup ###

```
$ git clone https://github.com/asayler/hash-server.git
$ go get https://github.com/braintree/manners
```

### Run ###

_By default, the server runs on port 8080 and uses a 5 second delay_

```
$ go run serv.go [port]
```

### Client Request Examples ###

_Data is returned base64 encoded. When salts are used, salt is
returned first and separated from hash output via `|`._

**Hash password using SHA512**
```
$ curl -X POST --data "password=Testing" http://[host]:[port]
ZPAml8zRwK50HZ4ib5VxJ9p6YU1qGPVfnycm0gJ/qsHpXmGdrFQX60iY/WqfuK65zdAF6RPIDldFTK5Lb8bl1g==
```

**Hash password w/ Random 32-byte Salt**
```
$ curl -X POST --data "password=Testing&salt=" http://[host]:[port]
bdYeFlCj27OB+0LL20v4JtQTvM4gYC98cLdNq8qntpo=|fMq58SFQbvmleJtXXDlJtFKIHvCsP6qbHexi/FKlpRrqQ0AfjJfURM7X0LNnyqb0frfRS2eKhN8OmzkSBYSd3Q==
```

**Hash password w/ User Specified Salt**
```
$ curl -X POST --data "password=Testing&salt=abcdefg" http://[host]:[port]
abcdefg|pBgplaz997cuKFfDHSfNLnvi4vjpspFVGGcj7/vIMUsquVFi+F/RAuzUQ5p235ZsW8Iiv49AeBiaHN/E+B1wsQ==
```

**Hash password w/ Multiple Rounds of SHA512**
```
$ curl -X POST --data "password=Testing&rounds=100" http://[host]:[port]
zs069hQAtAo3XavUSwhstTbN6T2AM8LGDrw5nzJn/9m+mlKXlKWYPgPyvE9Sv+M14l6ieXoq/lSP8InbKPfS8w==
```

**Stop accepting requests and shutdown server once in-flight requests complete**
```
$ curl -X POST --data "shutdown=" http://[host]:[port]
Shutting down
```

**Input can also be passed a multipart from data**
```
$ curl -X POST -F "password=Testing" -F "salt=" http://localhost:8080
gnHMrKpNhBkbsrCmqGPkpz7VTLYeKrmENHP8Ql/5sYo=|J1/HjFoZvfLPbGyGjEHepjA+zp7yP8GgUmsc8i9ZbfAA9G7gn8jUSRXiD5vq34XcZCRRqaLNO6hOc0kxFzcczw==
```
