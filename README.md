Distributed P2P message communication network library.

# Features

* *Use [raft](https://github.com/hashicorp/raft) distributed consensus protocol between Servers*

* *Use [dog-tunnel](https://github.com/vzex/dog-tunnel) P2P tech to connect all clients for each other*

# 工作流程

https://github.com/metalwood/p2pcore/wiki/flow



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
