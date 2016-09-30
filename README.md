Distributed P2P communication network. 

# Features

* *Use [raft](https://github.com/hashicorp/raft) distributed consensus protocol*

* *Use [dog-tunnel](https://github.com/vzex/dog-tunnel) P2P tech to connect all peers for each other*

# api.go 专门用来放api接口

func leaderChangedCallback(string peerId) // leader变化了就会触发这个回调

// 如果是刚启动的节点，peerAddr和peerPort可以留空，等别人来加入它等网络
func p2pNetStartup(listenAddr string, listen port, peerAddr string, peerPort int, callback leaderChangedCallback) (error) // 连接已有网络中的任意角色节点，即可加入该网络，任意两个节点间都要尽量直连（如果存在实在无法连通，则进行中转，但对api而言，中转是透明的； 如果不存在无法连通的情况更好）。每个实例自动生成一个随机的uuid作为peerId并保存在相同程序文件相同路径下配置文件中带密码加密的字段中，不允许随便改。
由leader维护一个节点信息列表，包括peerId、角色、公网地址、内网地址等关键信息。
刚接入网络、进程刚启动、leader更换时，peer在内部需要隐式自动上报自己的节点信息，以便leader知晓和提供查询

func p2pNetMyId() (peerId string) //查询自己的peerId

func p2pNetLookupRole(raftRole string) (peerId []string, error) //向leader查询

func p2pNetLookupPeer(peerId string) (raftRole string, wanIp []string, lanIp []string, error) //向leader查询

func p2pNetLookupMyLanPeers()(peerId []string, error) //查询和我在同一个局域网的所有节点的peerId

func p2pNetSend(peerId []string, data []byte) (error) //发送数据给指定uuid的peer

func p2pNetSendLeader(data []byte) (error)

func p2pNetRecv() (senderPeerId string, senderRole string, data []byte, error)

func p2pNetCleanup()
