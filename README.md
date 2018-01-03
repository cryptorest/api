# REST API


# Old

Client:
 -- Firefox nightly with about:config network.http.spdy.enabled.http2draft set true
 -- Chrome: go to chrome://flags/#enable-spdy4, save and restart (button at bottom)

Make CA:
```
$ openssl genrsa -out rootCA.key 2048
$ openssl req -x509 -new -nodes -key rootCA.key -days 1024 -out rootCA.pem
```
... install that to Firefox

Make cert:
```
$ openssl genrsa -out server.key 2048
$ openssl req -new -key server.key -out server.csr
$ openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 123
```

# Current

CA:
```
$ openssl genrsa -out ./ca.key 4096
$ openssl req -new -x509 -days 123 -key ./ca.key -out ./ca.crt
```

Server:
```
$ openssl req -new -utf8 -nameopt multiline,utf8 -config ./server.cnf -newkey rsa:4096 -keyout ./server.key -nodes -out ./server.csr
$ openssl x509 -req -days 123 -in ./server.csr -CA ./ca.crt -CAkey ./ca.key -set_serial 01 -out ./server.crt -extfile ./server.cnf -extensions ext
```

Client:
```
$ openssl req -new -utf8 -nameopt multiline,utf8 -newkey rsa:4096 -nodes -keyout ./client.key -out ./client.csr
$ openssl x509 -req -days 123 -in ./client.csr -CA ./ca.crt -CAkey ./ca.key -set_serial 01 -out ./client.crt
$ openssl pkcs12 -export -out ./client.p12 -in ./client.crt -inkey ./client.pem
```

Checking:
```
$ curl -v --cacert ./ca.crt --cert ./client.crt --key ./client.key https://localhost:64443/
```
