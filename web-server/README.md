# web server

Running the server `go run server.go`

`GET http://localhost:8081/world`

Load testing - https://octoperf.com/blog/2017/10/26/how-to-install-jmeter-mac/

## Load testing

You can use [httperf](https://github.com/httperf/httperf) and use it like

```
httperf --server localhost --port 8081 --num-conns 80 --rate 10000 --timeout 1
```

I could get a default rate of

```
Connection rate: 8808.6 conn/s (0.1 ms/conn, <=26 concurrent connections)
```