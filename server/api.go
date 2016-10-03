package p2pcore

func masterChangedCallback(masterServerId string)

// sponsorServerAddr是介绍人的地址，介绍此进程加入Server集群
func p2pServerStartup(serverId string, serverName string, listenAddr string, sponsorServerAddr string, masterChangedCallback cb) (err error)

func p2pServerCleanup()

func p2pServerLookupMyId() (serverId string)

func p2pServerLookupMaster() (serverId string, err error)

func p2pServerLookupServers() (masterServerId string, serverIdList []string, err error)

func p2pServerLookupServerInfo(serverId string) (role string, wanIp []string, lanIp []string, listenIp string, listenPort int, nat bool, err error)

func p2pServerLookupPeers() (peerIdList []string, myLanPeerIdList []string, err error)

func p2pServerLookupPeerInfo(peerId string) (wanIp []string, lanIp []string, listenIp string, listenPort int, nat bool, err error)

func p2pServerSend(dstIdList []string, data []byte) (failIdList []string, err error)

func p2pServerSendMaster(data []byte) (err error) // 在数据包加上标记，如果发送过程中master改变了，哪怕数据送达了，这个接口也要报错

func p2pServerRecv() (srcId string, srcRole string, data []byte, err error)
