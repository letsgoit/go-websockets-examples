# go-websockets-example
Understanding usage of Websockets with Golang

#### server1 : Simple `echo` server
```
$ go run server1/main.go
# Open client1/index.html in a web browser
```

#### server2 : Stock Market Intraday Price Server

Data from https://www.kaggle.com/nickdl/snp-500-intraday-data/data#
Extract the `zip` file downloaded from above to server2/data

```
$ go run server1/main.go
# This can be tested with client1/index.html similar to server1
``` 
