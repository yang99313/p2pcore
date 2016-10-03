func p2pClientStartup(clientId string, clientName string, listenAddr string, serverAddr string) (error)

func p2pClientCleanup()

func p2pClientLookupMyId() (ClientId string)

func p2pClientLookupMaster()(serverId string)

func p2pClientLookupServers()(connServerIdList []string, unConnServerIdList []string, error)

func p2pClientLookupClients()(connClientIdList []string, unConnClientIdList []string, error) //查询我已经连上和没有连上的Client的Id, 未连上的其实就是无法内网穿透的客户端

func p2pClientSend(dstId string, data []byte) (error) //发送数据给指定uuid的Client，需要先查询是否直连，如果没有直连再查询哪些server支持此ClientId，然后从支持的server中随机选择。只有最终收到了，才能返回成功。

func p2pClientRecv() (srcId string, data []byte, error)