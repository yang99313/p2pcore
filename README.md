分布式的P2P消息传输库，目的是简化多局域网分布式系统的内部通信， 包括：用Id通信方式封装屏蔽IP端口和C-S的通信方式、封装屏蔽内网穿透、封装屏蔽集群中Raft算法选举Leader的行为。

#  架构简介

此系统分为Server和Client两部分，Server端为数据转发节点，Client为工作节点，每个Server/Client有个唯一的Id（随机生成的UUID）。

Server的数量大于等于1，可以自动组成一个集群，Server之间全部两两互联，使用dog-tunnel直连。

Client的数量不限制，先通过dog－tunnel尽量两两互联直接交互数据，如果无法连接则通过Server转发数据。

# 代码组织

尽量编写业务无关的可重用代码，放在https://github.com/metalwood/ezgo 中，以包方式引用进来，最大程度减少p2pcore项目的代码量。
server/服务端的业务代码
client/客户的的业务代码
test/编写带main函数的测试程序，验证系统工作可靠

# Features

* *Use [raft](https://github.com/hashicorp/raft) distributed consensus protocol between Servers*

* *Use [dog-tunnel](https://github.com/vzex/dog-tunnel) P2P tech to try to connect all clients for each other*

# 工作流程

https://github.com/metalwood/p2pcore/wiki/Home



# server/api.go

专门用来放server的api接口, p2pcore应用到其他系统时，server端功能和client不同，server是不需要任何修改的，只是提供了这些编程接口用于查询而已。

func leaderChangedCallback(string uuid) // leader变化了就会触发这个回调

func p2pServerStartup(listenAddr string, listen port, otherServerAddrList []string, otherServerAddrList []int, leaderChangedCallback cb) (error)

func p2pServerCleanup()

func p2pServerLookupMyId()(serverId string)

func p2pServerLookupServers()(masterId string, serverIdList []string, error)

func p2pServerLookupClientInfo(clientId string) (wanIp []string, lanIp []string, listenIp string, listenPort int, error)

func p2pServerLookupClients()(clientIdList []string, error)

func p2pServerSend(dstId string, data []byte)(error)

func p2pServerRecv()(srcId string, data []byte)




# client/api.go

专门用来放client的api接口

func p2pClientStartup(listenAddr string, listen port, serverAddrList []string, serverPortList []int) (error)

func p2pClientCleanup()

func p2pClientLookupMyId() (ClientId string) //查询自己的uuid

func p2pClientLookupServers()(masterServerId string, connServerIdList []string, unConnServerIdList []string, error)

func p2pClientLookupClients()(connClientIdList []string, unConnClientIdList []string, error) //查询我已经连上和没有连上的Client的Id, 未连上的其实就是无法内网穿透的客户端

func p2pClientSend(dstId string, data []byte) (error) //发送数据给指定uuid的Client，需要先查询是否直连，如果没有直连再查询哪些server支持此ClientId，然后从支持的server中随机选择。只有最终收到了，才能返回成功。

func p2pClientRecv() (srcId string, data []byte, error)
