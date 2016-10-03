package p2pcore

func masterChangedCallback(masterServerId string) // leader变化了就会触发这个回调

func p2pServerStartup(serverId string, serverName string, listenAddr string, otherServerAddr string, leaderChangedCallback cb) (err error)

func p2pServerCleanup()

func p2pServerLookupMyId() (serverId string)

func p2pServerLookupServers() (masterServerId string, serverIdList []string, err error)

func p2pServerLookupClients() (clientIdList []string, err error)

func p2pServerLookupClientInfo(clientId string) (wanIp []string, lanIp []string, listenIp string, listenPort int, err error)

func p2pServerSend(dstId string, data []byte) (err error)

func p2pServerSendMaster(data []byte) (err error)

func p2pServerRecv() (srcId string, srcRole string, data []byte, err error)
