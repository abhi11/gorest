### gorest
A simple rest api interface to query and post logs


## Current APIs:
* GET APIs
  * /logs : Fetches all logs
  * /logs/after/{timestamp} : Fetches all lgs after the timestamp (unix time)
  * /logs/before/{timestamp} : Fetches all lgs before the timestamp (unix time)

## Running on localhost
You would need mongodb to run gorest.

```
$ mongod

$ go get github.com/abhi11/gorest
$ ./gorest
```
## TODO
* Insertion of data
* Expose the post apis
* write a script to start gorest [assuming mongodb is present]
* Adding docker support
