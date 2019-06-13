# Database load balancer 
## General overview 
The call of the app is to route requests between nodes that performs requests,
processes them and responds with processed data. All these nodes have an API interface
for performing DB requests and the main goal of the Database load balancer is 
determinate which node is the most stress-free and route request to it 
otherwise the load balancer should send a request to the service which 
will build a new node's instance up and then the balancer will route 
the request to the new node. Other words, it's kind of [nginx](https://nginx.org/) 
but for database 
request. Also, the balancer uses the database, it stores there the info about 
current node's allocations and the list of the available nodes.
## Configurations
The database load balancer is a configurable application. Here's the default 
config file which locates at the app's root directory:

***config.json***
```json
{
  "database": {
    "vendor" : "postgres",
    "host": "localhost",
    "port": 5432,
    "user": "test",
    "password": "test",
    "db_name": "querier_manager",
    "ssl_mode": "disable"
  },
  "web_server": {
    "port": 8080
  },
  "instances_master": {
    "ip_address": "10.0.0.1",
    "port": 80
  }
}
```
The config file includes three general sections: `database`, `web_server` and 
`instances_master`.

The database's section includes the general info about how to connect to the 
database such as: 
* `vendor` - means the database's vendor
* `host` - the ip address on which run the database
* `port` - the database port
* `user` and `password` - the database's credentials
* `db_name` - the name of the database which will be used for storing current
 node's allocations and the list of available nodes
 
The web server's section have only one configurable parameter - `port`, 
this option is determinate on which port will be spun web server up.
 
The instances master's section have two parameters: `ip_address` and `port`.
This section uses for configuring the connection to service which deploys
a new nodes

## Building and spinning the application up

To download the app perform the following command:
```bash
go get github.com/Andry75/PG_querier_manager
```

Then navigate to the app's dir:

```bash
cd $GOPATH/src/github.com/Andry75/PG_querier_manager
```

To build the app perform the following command:

```bash
go build main.go
```

To start the app perform the following command:

```bash
./main
```

