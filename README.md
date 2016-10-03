分布式的P2P消息传输库，目的是简化多局域网分布式系统的内部通信。包括：用Id－Id的点到点通信模式简化取代IP端口的CS通信模式、内部尽量使用内网穿透其次使用Server转发、Server集群使用Raft算法确保始终有一个Master。

#  架构简介

此系统分为Server和Client两部分，Server端为数据转发节点，Client为工作节点，每个Server/Client有个唯一的Id（随机生成的UUID）。

Server的数量大于等于1，可以自动组成一个集群，Server之间全部两两互联，使用dog-tunnel直连。

Client的数量不限制，先通过dog－tunnel尽量两两互联直接交互数据，如果无法连接则通过Server转发数据。

# 代码组织

尽量编写业务无关的可重用代码，放在https://github.com/metalwood/ezgo 中，以包方式引用进来，最大程度减少p2pcore项目的代码量。
server/服务端的业务代码
client/客户的的业务代码
test/编写带main函数的测试程序，验证系统工作可靠

# Features

* *Use [raft](https://github.com/hashicorp/raft) distributed consensus protocol between Servers*

* *Use [dog-tunnel](https://github.com/vzex/dog-tunnel) P2P tech to try to connect all clients for each other*


# 接口设计

客户端接口：client/api.go

服务端接口：server/api.go

可根据实际情况调整


# Server工作流程

识别自己的网络状态：包括公网Ip(列表)、内网Ip(列表)、监听Ip和端口。

连接其他Server，索要Master的地址：如果配置文件中指定了其他Server的地址，则连接其他Server，并获取Master的地址

连接Master，索要全部Server的地址。

连接全部Server，组成集群：任意两个Server之间都尽量互联。

选举Master：Raft算法，投票给连接数最多的，以便Master进行Server间的数据同步。

Server向Master上报和更新自己的网络状态。

维护所有Server的网络信息C，由Master负责同步，从表B中同步来的。

接收客户端连接：把客户端上报的网络信息加入自己的Client网络信息表A。

维护所有连着的Client的网络信息和连接信息表（只是当前Server，不是所有Server）A。

如果是Master，还要额外维护所有Server的网络信息表B（就是所有Server的网络地址而已），并负责此数据在Server集群中的同步。


# Client工作流程

识别自己的网络状态：包括公网Ip列表、内网Ip列表、监听Ip和端口。

连接Server集群：连接任意一个Server，查询全部Server并全部连上，保持连接。

向所有Server上报和刷新（更改时）自己的网络信息：Client上报自己的Id和网络状态，并实时刷新。

从Server接收该Server的连接所有Client的Id列表和网络信息：Server会在Client首次连接和数据更新后下发连接它的Client列表。

连接其他Client：根据Server给的Client列表，首先是直接连接，然后是内网穿透，尝试连接尽可能多的Client，只要不超出配置文件中配置的最大值。Client直接交互数据时，首先是直接交互，如果直连列表中有对方Id则直连，查询哪些Server的列表中有他，然后随机选一个，把消息转给Server，如果找不到，则无法发送。

维护一张自己的连接信息表A：包括和哪些UUID连接，是直连还是内网穿透连接。

维护所有Server下发的Server的连接信息表B(1~n)

发送数据给其他Client：结合表A和B(1~n)，根据Client Id，先看是否直连，有就直接发送，再看哪些Server可以代发，有就随机选一个代发，否则返回找不到目的地。


# Server集群节点数变化时的处理

如果节点数量为1，则自己成为临时Master。

如果节点数量为2，则根据进程启动到此刻的消耗毫秒数决定，先运行的为Master。

如果节点数量大于等于3，则根据Raft选举Master。


# Id和名称

Id是角色代号＋随机数表示的唯一编号，Server的Id是"S-$UUID", Client的Id是"C-$UUID"
名称是用户在配置中指定的字符串，方便人阅读和区分的名字。


# 角色

Master 系统唯一的节点，由所有Server节点们根据Raft算法自动选举出的。Master也做Server的工作，只是额外多了全局信息统计监控的工作。

Server 数据中继，由配置文件人为手工指定，一般是HighId（如果需要的话）、流量便宜的节点，此系统中Server越多，系统越可靠。

Client 普通节点。


# 异常处理

如果Master挂了，由所有Server们选举出新的。

如果某个Client和部分Server之间掉线了，Client不管，尽可和其他Server通信。

如果某个Client和全部Server之间掉线了，则Client停止工作，直到和Server恢复连接。


# 注意

允许部分Server在内网中，Client在公网吗？不考虑这种可能，如果这样部署，是不符合系统设计的配置。