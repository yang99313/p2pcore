Distributed P2P communication network. 

# Features

* *Use [raft](https://github.com/hashicorp/raft) distributed consensus protocol*

* *Use [dog-tunnel](https://github.com/vzex/dog-tunnel) P2P tech to connect all peers for each other*

# api.go 专门用来放api接口

func leaderChangedCallback(string peerId) // leader变化了就会触发这个回调

func p2pNetStartup(peerAddr string, peerPort int, callback leaderChangedCallback) (error) // 连接已有网络中的任意角色节点，即可加入该网络，任意两个节点间都要尽量直连，如果实在无法连通，则进行中转，但对api而言，中转是透明的。每个实例自动生成一个随机的uuid作为peerId并保存在相同程序文件相同路径下配置文件中带密码加密的字段中（不允许改）。
由leader维护一个节点信息列表，包括peerId、角色、公网地址、内网地址等关键信息, 如果leader更换了，所有peer在内部需要隐式自动上报节点信息，以便leader知晓和提供查询

func p2pNetMyId() (string) //查询自己的peerId

func p2pNetLookupRole(raftRole string) (peerId []string, error) //向leader查询

func p2pNetLookupPeer(peerId string) (raftRole string, wanIp []string, lanIp []string, error) //向leader查询

func p2pNetSend(peerId string, data []byte) (error) //发送数据给指定uuid的peer

func p2pNetSendLeader(data []byte) (error)

func p2pNetRecv() (senderPeerId string, senderRole string, data []byte, error)

func p2pNetCleanup()
