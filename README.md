
* *Use raft distributed consensus protocol*

# msg-agent
distributed named pipe for forward data, allowed multi read write clients.

send message with peer id instead of TCP address.

stringArray dp_lookup(string clientId);
dp_read(string pipeName, data, size);
dp_write(string pipeName, data, size);

如果有公网节点，就可能遇到

topic广播机制：接收方主动查询topic的情况

服务或节点发现、点对点传输、中转传输、消息队列，把四者封装成对用户透明化

p2pQueue

sharedIndex: dbName queueName storePoint

distribCollection:


put(long key, bin data, long size);
get(uuid storePoint, long key);
push(string collName, bin data, long size);
pop(uuid storePoint, string collName);

kv-collection
//不需要知道 collection: sub, key: topicMd5, val: clientIdList
collection: pub, key: topicMd5, val: clientIdList

api.go
bool p2pNetStartup(string peerAddr, int peerPort); // 连接任意节点，即可加入网络，整个网络需要隐式维护一个节点列表，任意两个节点间都要互联，如果实在无法连通，则进行中转，但对api而言，中转是透明的。
string p2pNetLookupMyId();
string p2pNetLookupClientId();
bool p2pNetSend(dstUuid, );
bool p2pNetRecv(dstUuid, );
void p2pNetQuit();
