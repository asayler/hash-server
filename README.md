Hash Server
-----------

**_A simple hashing microservice_**

Hashes (and optionally salts) passwords using one or more rounds of
SHA512.

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

__By default, the server runs on port 8080__

```
$ go run serv.go [port]
```

### Client Request Examples ###

_Data is returned base64 encoded. When Salts are used, salt is
returned first and separated from data in return via `|`._

**Hash password using SHA512**
```
$ curl --data "password=Testing" http://[host]:[port]
```

**Hash password w/ Random 32-byte Salt**
```
$ curl --data "password=Testing&salt=" http://[host]:[port]
```

**Hash password w/ User Specified Salt**
```
$ curl --data "password=Testing&salt=abcdefg" http://[host]:[port]
```

**Hash password w/ Multiple Rounds of SHA512**
```
$ curl --data "password=Testing&rounds=100" http://[host]:[port]
```

**Stop accepting requests and shutdown server**
```
$ curl --data "shutdown=" http://[host]:[port]
```
