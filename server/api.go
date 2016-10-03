func leaderChangedCallback(serverId string) // leader变化了就会触发这个回调

func p2pServerStartup(serverId string, servername string, listenAddr string, otherServerAddr string, leaderChangedCallback cb) (error)

func p2pServerCleanup()

func p2pServerLookupMyId()(serverId string)

func p2pServerLookupServers()(masterId string, serverIdList []string, error)

func p2pServerLookupClientInfo(clientId string) (wanIp []string, lanIp []string, listenIp string, listenPort int, error)

func p2pServerLookupClients()(clientIdList []string, error)

func p2pServerSend(dstId string, data []byte)(error)

func p2pServerRecv()(srcId string, data []byte)