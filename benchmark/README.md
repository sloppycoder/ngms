## benchmark the API using locust

Use [Locust.io](http://locust.io) to benchmark APIs. For my testing Locust can't generate more than 2000 TPS of loan even with 2 slaves, so I switched to use [Boomer](https://github.com/myzhan/boomer) load generator, which can generate tons of requests using 1 single process.

### Install Locust
```
 conda install -c conda-forge locust

```

### Run Locust for testing REST only
```
locust --host http://127.0.0.1:3000

# for bigger load use master slave mode
locust --host http://127.0.0.1:3000 --master

# open a new shell for each slave
locust --slave

# open browser to http://localhost:8089 for Locust UI

```

### Run Locust master and Boomer as slave
```

# build slave binary
go build -o slave main.go

# start locust master, content of locusfile.py will not be used by Boomer
locust --master

# start slave, 1 single process is enough for more than 20000TPS of load
# use the -proto flag to control rest of grpc endpoint to test

./slave -proto rest 
or
./slave -proto grpc

use --help flag to see all flags supported by Boomer

```


