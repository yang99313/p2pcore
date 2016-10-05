# 目标

分布式的P2P消息传输库，目标是简化多局域网分布式系统的内部通信。包括：用Id～Id点到点通信模式简化取代IP端口式的CS通信模式；优先使用[dog-tunnel](https://github.com/vzex/dog-tunnel)内网穿透其次才使用Server转发节约Server流量、Server集群使用[raft](https://github.com/coreos/etcd/tree/master/raft)算法（etcd的实现版本）确保始终有一个Master（提供给用户在需要集群一致性的场合使用）.

PS:如果dog-tunnel不支持类似TCP的直连，则使用[kcp-tunnel](https://github.com/xtaci/kcptun)

#  架构

此系统分为Server和Peer两部分，Server又分普通Server和Master，所以总共三个角色

每个Server/Peer必须指定一个唯一Id来取代Ip和端口。每个Peer和Server都有一个Id和一个Name，通过api接口指定的。Id是角色代号＋随机数表示的编号，必须唯一，否则会带来错乱，Server的Id是"S"+"UUID", Peer的Id是"C"+"UUID"。名称是用户指定的字符串，方便人阅读和区分的名字

Master：系统唯一的节点，由所有Server节点们根据Raft算法自动选举出的，选举Master是为p2pcore的用户提供的，对p2pcore本身没有比普通Server多做任何工作，是个闲差

Server：协助Peer内网穿透以及在无法穿透时提供数据转发，部署时要求所有Peer都可以主动连上他，且流量便宜。系统中Server越多，系统越可靠。Server集群中每个Server都提供几乎完全相同的服务，相互直接基本没有数据同步，除了Server列表。集群数量>=1，之间全部两两互联，使用dog-tunnel直连

Peer：普通节点，提供Id－Id的通信模式，以接口方式提供上述功能。和其他Peer连接时，优先两两直连（即主动连其他Peer，也监听其他Peer的连接），如果无法直连则通过Server协助使用dog－tunnel内网穿透，如果无法穿透则通过Server转发

# 代码组织

尽量编写业务无关的可重用代码，放在https://github.com/metalwood/ezgo 中，以包方式引用进来，最大程度减少p2pcore项目的代码量，比如，dog-tunnel的常用api不够清晰明了，就可以单独封装成一个p2p-api.go

# 接口

请见源码目录，可按需调整

# Server工作流程

识别自己的网络信息（公网地址列表、内网地址列表、监听地址）

接入Server集群：接口中需要提供一个集群现有Server的地址，以他作介绍人，向此介绍人索要全部Server的地址, 连上全部Server，形成集群，任意两个Server之间都尽量互联

选举Master：按Raft算法投票，Master对p2pcore项目内部没有任何用途，是提供给p2pcore的用户的

同步数据：当Server集群有成员加入或者退出时，要进行广播（集群数量不会太大、广播可以接收，而Peer之间用广播进行数据同步的话，代价大），要减少不必要的广播。除此之外，Server间不共享/同步任何数据。这个数据集姑且命名为表C

等待客户端连接：把客户端上报的网络信息加入自己的Peer网络信息表A，把Peer上报的互连信息，加入自己的Peer互连信息表B

进入工作状态，和其他节点交互数据


# Peer工作流程

识别自己的网络信息（公网地址列表、内网地址列表、监听地址），姑且命名为表D

接入Server集群：连接任意一个Server，向Server查询全部Server列表，连上所有Server，保持连接

同步数据：向所有Server上报和刷新（更改时）自己的网络信息

同步数据：从所有Server上同步各个服务器所有在线Peer的网络信息（表E1~n）：Server会在Peer首次连接和数据更新后下发此信息，可以采用小技巧（比如拆解为增加、删除操作）减少同步对数据量

接入Peer集群：根据Server给的Peer网络信息，连接所有Peer。连接时，优先直接连接，然后是内网穿透。同时维护一张自己的连接信息表F，包括和哪些Peer连接，和哪些Server连接，是直连还是内网穿透连接

同步数据：把表F同步给所有Server

进入工作状态，允许用户通过api和其他Peer交换数据：Peer间交互数据时，优先查询查询表F尝试走P2P连接，如果和对方Id没有连接，则从表E查询走Server中转（一般可以中转的Serve不止一个，可随机选一个）



# 数据同步

Server之间需要同步的数据是表C，采用广播方式的同步

Server和Peer之间需要同步数据的是表D、表E、表F，采用主从方式同步（Server是主，Peer是从）

Peer之间没有需要同步的数据


# Server集群节点数变化时的处理

如果节点数量==1，则自己成为临时Master

如果节点数量==2，则根据进程启动到此刻的消耗毫秒数，先运行的为Master

如果节点数量>=3，则根据Raft选举Master


# 异常处理

如果Master挂了，由所有Server们选举出新的

如果Peer和部分Server之间掉线了，Peer不断尝试重练的同时完全可以正常工作，和其他Server通信

如果Peer和全部Server之间掉线了，Peer停止工作，直到和至少一个Server恢复连接

心跳和超时，如果超出一定时间还是没有数据传输则发送心跳，如果一直有数据传输，可以省略心跳