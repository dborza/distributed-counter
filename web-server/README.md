# web server

Running the server `go run server.go`

`GET http://localhost:8081/world`

Load testing - https://octoperf.com/blog/2017/10/26/how-to-install-jmeter-mac/

## Load testing

A list of load testing tools can be found [here](https://gist.github.com/denji/8333630).

From the list above I used [baton](https://github.com/americanexpress/baton) for load testing

I could install it with `go get -u github.com/americanexpress/baton` on MacOS.

I ran the following test locally

```
~/go/bin/baton -u http://localhost:8081/world -c 20 -r 200000
```

I could get the following results

```
=========================== Results ========================================

Total requests:                                200000
Time taken to complete requests:         4.367561177s
Requests per second:                            45792
Max response time (ms):                            59
Min response time (ms):                             0
Avg response time (ms):                          0.09

========= Percentage of responses by status code ==========================

Number of connection errors:                        0
Number of 1xx responses:                            0
Number of 2xx responses:                       200000
Number of 3xx responses:                            0
Number of 4xx responses:                            0
Number of 5xx responses:                            0

========= Percentage of responses received within a certain time (ms)======

        99% : 5 ms
        99% : 10 ms
        99% : 15 ms
        99% : 20 ms
        99% : 25 ms
        99% : 30 ms
        99% : 35 ms
        99% : 40 ms
        99% : 45 ms
       100% : 59 ms

===========================================================================
```
