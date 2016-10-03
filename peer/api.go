package p2pcore

func masterChangedCallback(masterServerId string)

func p2pPeerStartup(peerId string, peerName string, listenAddr string, serverAddr string, masterChangedCallback cb) (err error)

func p2pPeerCleanup()

func p2pPeerLookupMyId() (PeerId string)

func p2pPeerLookupMaster() (serverId string, err error)

func p2pPeerLookupServers() (connServerIdList []string, unConnServerIdList []string, err error)

func p2pPeerLookupServerInfo(serverId string) (role string, wanIp []string, lanIp []string, listenIp string, listenPort int, nat bool, err error)

func p2pPeerLookupPeers() (connPeerIdList []string, unConnPeerIdList []string, err error) //查询我已经连上和没有连上的Peer的Id, 未连上的其实就是无法内网穿透的客户端

func p2pPeerLookupPeerInfo(peerId string) (wanIp []string, lanIp []string, listenIp string, listenPort int, err error)

func p2pPeerSend(dstIdList []string, data []byte) (failIdList []string, err error) //发送数据给指定uuid的Peer，需要先查询是否直连，如果没有直连再查询哪些server支持此PeerId，然后从支持的server中随机选择。只有最终收到了，才能返回成功。

func p2pPeerSendMaster(data []byte) (err error) // 在数据包加上标记，如果发送过程中master改变了，哪怕数据送达了，这个接口也要报错

func p2pPeerRecv() (srcId string, srcRole string, data []byte, err error)
