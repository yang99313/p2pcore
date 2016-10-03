分布式的P2P消息传输库，目的是简化多局域网分布式系统的内部通信。包括：用Id－Id的点到点通信模式简化取代IP端口的CS通信模式、内部优先使用[dog-tunnel](https://github.com/vzex/dog-tunnel)内网穿透其次才使用Server转发、Server集群使用[raft](https://github.com/coreos/etcd/tree/master/raft)算法确保始终有一个Master。

#  架构简介

此系统分为Server和Peer两部分，Server端为数据转发节点，Peer为工作节点，每个Server/Peer必须指定一个唯一Id。

Server的用途是给Peer提供内网穿透协助，无法穿透则提供数据转发。Server集群中每个Server都提供相同的服务，相互直接基本没有数据同步，除了Server列表。集群数量>=1，之间全部两两互联，使用dog-tunnel直连。

Peer的用途是提供Id－Id的通信模式，并可以和Server直接交换数据，然后以接口方式提供上述功能。客户端数量不限制，先两两直连，如果无法直连则通过Server协助使用dog－tunnel内网穿透，如果无法穿透则通过Server转发。

# 代码组织

尽量编写业务无关的可重用代码，放在https://github.com/metalwood/ezgo 中，以包方式引用进来，最大程度减少p2pcore项目的代码量，比如，dog-tunnel的常用api不够清晰明了，就可以单独封装成一个dog-tunnel-api.go。

# 接口设计

客户端接口：peer/api.go

服务端接口：server/api.go

可根据实际情况调整


# Server工作流程

识别自己的网络信息（公网地址列表、内网地址列表、监听地址、是否有局域网NAT防火墙）

连接其他介绍人Server，索要Master的地址

连接Master，索要全部Server的地址

连接全部Server，加入集群，任意两个Server之间都尽量互联

选举Master：按Raft算法投票，Master对p2pcore项目内部没有任何贡献和用途，是提供给p2pcore的用户的。

等待客户端连接：把客户端上报的网络信息加入自己的Peer网络信息表A，把Peer上报的互连信息，加入自己的Peer互连信息表B。

当Server集群有成员加入或者退出时，要进行广播（集群数量不会太大、广播可以接收，而Peer之间用广播进行数据同步的话，代价就大了），当然，可以采取适当的技巧（比如记住其他server的广播历史）减少不必要的广播。除此之外，Server间不共享/同步任何数据。这个数据集姑且命名为表C。


# Peer工作流程

识别自己的网络信息（公网地址列表、内网地址列表、监听地址、是否有局域网NAT防火墙），姑且命名为表D

连接Server集群：连接任意一个Server，向Server查询全部Server列表，连上所有Server，保持连接

向所有Server上报和刷新（更改时）自己的网络信息

从所有Server上同步各个服务器所有在线Peer的网络信息（表E1~n）：Server会在Peer首次连接和数据更新后下发此信息，可以采用小技巧（比如拆解为增加、删除操作）减少同步对数据量。

连接其他Peer：根据Server给的Peer网络信息，连接所有Peer。连接时，优先直接连接，然后是内网穿透。根据连接，维护一张自己的连接信息表F，包括和哪些Peer连接，和哪些Server连接，是直连还是内网穿透连接，然后同步给所有Server。

和其他Peer交换数据（这个操作由用户调用了，提供接口而已）：Peer间交互数据时，如果和对方Id有连接最好（查询表F），如果和对方Id没有连接，则（从表E）查询哪些Server的列表中有他，有就从可以代发的中间随机选一个代发，否则返回找不到目的地错误。



# 数据同步

Server之间需要同步的数据是表C

Server和Peer之间需要同步数据的是表D、表E、表F

Peer之间没有需要同步对数据


# Server集群节点数变化时的处理

如果节点数量==1，则自己成为临时Master。

如果节点数量==2，则根据进程启动到此刻的消耗毫秒数，先运行的为Master。

如果节点数量>=3，则根据Raft选举Master。


# Id和Name

每个Peer和Server都有一个Id和一个Name，通过api接口指定的。

Id是角色代号＋随机数表示的编号，必须唯一，否则会带来错乱，Server的Id是"S"+"UUID", Peer的Id是"C"+"UUID"

名称是用户指定的字符串，方便人阅读和区分的名字。


# 角色

Master 系统唯一的节点，由所有Server节点们根据Raft算法自动选举出的，选举Master是为p2pcore的用户提供的，对p2pcore本身没有比普通Server多做任何工作，是个闲差。

Server 数据中继，一般是HighId（如果需要的话）、流量便宜的节点，此系统中Server越多，系统越可靠。

Peer 普通节点，和其他Peer连接时，同时启用了Peer和Server模式的通信。


# 异常处理

如果Master挂了，由所有Server们选举出新的。

如果Peer和部分Server之间掉线了，Peer不断尝试重练的同时完全可以正常工作，和其他Server通信。

如果Peer和全部Server之间掉线了，Peer停止工作，直到和至少一个Server恢复连接。


# 注意

允许部分Server在内网中，Peer在公网吗？不考虑这种可能，如果这样部署，是不符合系统设计的配置。