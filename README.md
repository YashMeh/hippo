# hippo
![Kafka](https://img.shields.io/badge/kafka-2.3.0-green)        ![InfluxDB](https://img.shields.io/badge/influxdb-1.8-green)               ![Zookeeper](https://img.shields.io/badge/zookeeper-3.6.2-green) ![Docker](https://img.shields.io/badge/docker-20.10.2-blue)  ![Go](https://img.shields.io/badge/go-1.15.6-blue) 

Hippo is a data ingestor service for gRPC and REST based clients. It publishes your messages on a kafka queue and eventually saves them to influxDB, it is build to be easily scalable and prevent SPOF for your mission-critical data collection.
## Why ? Why not ?
While making this project I learnt so many things and it gave me an excuse to solve one of the problems that I face a lot, dealing with server logs. 
## <u>Architecture</u>
Insert Picture here
## <u>Alternative Methods</u>
- Use a 3rd party paid service for managing your data collection pipeline <b>but</b>, in the long run you will become dependent and as the data grows your cost will increase manifold.
- Create a proxy and start dumping everything on cache(Redis etc.) and create a worker to feed that to the database <b>but</b>, cache is expensive to maintain and there is only so much that you can store on the cache.
- Directly feed to the database <b>but</b>, you will be bounded by query/secs and will not receive reliable gRPC support for major databases.
## <u>Configurations </u>
In order to run hippo, you need to configure the following- 
1. `kafka-pixy/kafka-pixy.yaml`
Point to your kafka broker(s)
```yaml
kafka:
seed_peers:

- kafka-server1:9092
```   
Point to your zookeeper
```yaml
zoo_keeper:
seed_peers:
- zookeeper:2181
```
2. `vector/vector.toml`
This `.toml` file will be used by [timberio/vector](https://github.com/timberio/vector) (which is :heart: btw) to configure your source (apache-kafka) and your sink(influxDB)
## <u>Usage </u>
The proxy ([kafka-pixy](https://github.com/mailgun/kafka-pixy)) opens up 2 ports for the communication.
| Port | Protocol |  
| ----------- | ----------- |  
| 19091 | gRPC |  
| 19092 | HTTP |

The [official repository](https://github.com/mailgun/kafka-pixy) does a great job :heart: in explaining how to connect with the proxy server but for the sake of simplicity, let me explain here.
#### Using REST based clients
It can be done like this (here foo is the topic name and bar is the group_id)
1. Produce a message on a topic
```bash
curl -X POST localhost:19092/topics/foo/messages?sync -d msg='sample message'
```
2. Consume a message from a topic
```bash
curl -G localhost:19092/topics/foo/messages?group=bar
```
### Using gRPC based clients
The [official repository](https://github.com/mailgun/kafka-pixy)(which is awesome, btw) provides the `.proto` file for the client, just produce the `stub` for the language of your choice(they have helper libraries for `go` and `python`) and you will be ready to go.
For the sake of simplicity, I have provided a [sample client](https://github.com/YashMeh/hippo/blob/main/gRPC-clients/go/client.go).

## <u>Running</u>
Simply run `docker-compose up` to setup the whole thing
### How to use in production :fire:
Some things that I can think of on top of my head - 
1. Add a reverse-proxy(a load balancer maybe (?) ) that can communicate over to the pixy instances in different AZs  
2. Configure `TLS` in the pixy , add password protection to `zookeeper` and configure proper policies for `influxdb`
3. Use any container orchestration tool (kubernetes or docker-swarm and all the fancy jazz) to scale up the servers whenever required.

## <u>Usecases</u>
Hippo can be used to handle high volume streaming data, some of the usecases that I can think of are -
1. Logging server logs produced by different microservices 
2. IoT devices that generate large amount of data

## TBD (v2 maybe (?))
- [ ] Right now, the messages sent do not enforce any schema, this feature will be very useful while querying the influxdb (Add an enrichment layer to the connector itself, maybe (?))  
- [ ] Provide benchmarks

### Yash Mehrotra

![GitHub followers](https://img.shields.io/github/followers/YashMeh?label=Follow&style=social) ![Twitter URL](https://img.shields.io/twitter/follow/YashMeh29715504?label=Follow&style=social)

---

```C++
if(repo.isAwesome || repo.isHelpful){
    StarRepo();
}
```