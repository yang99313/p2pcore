分布式的P2P消息传输库，目的是简化多局域网分布式系统的内部通信， 包括：用Id通信方式封装屏蔽IP端口和C-S的通信方式、封装屏蔽内网穿透开发、封装屏蔽分布式集群中Leader的选举。

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

* *Use [dog-tunnel](https://github.com/vzex/dog-tunnel) P2P tech to connect all clients for each other*

# 工作流程

https://github.com/metalwood/p2pcore/wiki/Home



# server/api.go 专门用来放server的api接口

func leaderChangedCallback(string uuid) // leader变化了就会触发这个回调


func pnServerStartup(listenAddr string, listen port, otherServerAddrList []string, otherServerAddrList []int, leaderChangedCallback cb) (error)

func pnServerCleanup()

func pnServerLookupInfo(uuid string) (role string, wanIp []string, lanIp []string, listenIp string, listen port int, error)

func pnServerLookupClients()(uuidList []string, error)



# client/api.go 专门用来放client的api接口

func pnClientStartup(listenAddr string, listen port, ServerAddrList []string, ServerPortList []int) (error)

func pnClientCleanup()

func pnClientLookupMasterId(uuid string, error)

func pnClientLookupServers()(connectedUuidList []string, unconnectedUuidList []string, error) //查询我已经连上和没有连上的Server的uuid

func pnClientLookupClients()(connectedUuidList []string, unconnectedUuidList []string, error) //查询我已经连上和没有连上的Client的uuid

func pnClientLookupMyId() (ClientId string) //查询自己的uuid

func pnClientSend(ClientId []string, data []byte) (error) //发送数据给指定uuid的Client

func pnClientSendMaster(data []byte) (error)

func pnClientRecv() (senderClientId string, senderRole string, data []byte, error)
