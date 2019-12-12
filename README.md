### 以不动产登记业务为例，使用超级账本搭建政务数据区块链原型

- ### 关于业务流程
- 区域链（基于证书准入的联盟链）网络
![网络拓扑](http://xujf000.tk:28888/img/topo.png "网络拓扑")
- 业务流程时序
![业务流程时序图](http://xujf000.tk:28888/img/flow.png "业务流程时序图")
- ### 关于超级账本
关于区块链、超级账本、智能合约等概念，请自行谷哥和度娘。建议直接参考官网文档 https://hyperledger-fabric.readthedocs.io/en/release-1.4/ 

- ### 关于本例说明
- 本例仅限搭建原型用于验证技术可行性目的，远达不到产品级可用；
- 本例采用了3个组织，每组织2个节点，用于不动产、房管、税务三个部门做背书节点；3份智能合约（网签合同、纳税凭证、不动产权证书）跑在1个通道；状态数据库采用CouchDB;
- 本例DEMO:  http://xujf000.tk:28888/   （VPS,1CPU 1G MEM)

- ### 如何搭建区块链并部署运行智能合约
本例采用1.4.3版本。以下步骤在centos7上完成，并适用于ubuntu/MACOS/WINDOWS等
1. 安装环境（go1.3以上，docker-ce,docker-compose,git)
```bash
cd /opt
wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
tar zxvf go*.gz
yum install -y yum-utils   device-mapper-persistent-data   lvm2
wget -O /etc/yum.repos.d/docker-ce.repo https://download.docker.com/linux/centos/docker-ce.repo
yum install -y docker-ce docker-compose git
```
2. 下载超级账本官方超级账本网络示例
```bash
curl -sSL http://bit.ly/2ysbOFE | bash -s -- 1.4.3 1.4.3 0.4.15
```
将会在当前/opt目录下生成fabric-samples目录，将自动下载命令工具和镜像。
3. 将fabric-samples/bin和/opt/go/bin 加入本地PATH

4. 下载本示例
```bash
cd /opt/fabric-samples
git clone https://github.com/fengzxu/fcc.git
cd fcc
```
5. 启动示例网络，创建区块链网络（2个组织，每组织2个节点）
```bash
chmod +x *.sh
./1.startNetwork.sh
```
完成后，结果显示：
```bash
========= All GOOD, BYFN execution completed =========== 
```
6. 加入第3个组织，2个节点
```bash
./2.addOrg3.sh
```
完成后，结果显示：
```bash
========= Org3 is now halfway onto your first network =========
```
7. 部署政务智能合约，并实例化
```bash
./3-1.installNetcon.sh   #合约：网签合同备案
./3-2.installEstateBook.sh   #合约：不动产权证书
./3-3.installEstateTax.sh    #合约：不动产业务缴税
```
完成后，结果显示：
```bash
Get instantiated chaincodes on channel mychannel:
Name: estatebook, Version: 1.0, Path: github.com/chaincode/estatebook, Escc: escc, Vscc: vscc
Name: estatetax, Version: 1.0, Path: github.com/chaincode/estatetax, Escc: escc, Vscc: vscc
Name: netcon, Version: 1.0, Path: github.com/chaincode/netcon, Escc: escc, Vscc: vscc
```
8. 编译后台。 代码位于appcode/fccserver/src 可自行编译，或直接使用已编译完成的可执行文件。
```bash
chmod +x appcode/fccserver/src/fccserver
```
启动后台容器
```bash
./4.startAppcli.sh
docker logs -f appcli
```
如果启动正常，会显示：
```bash
[fcc-server] 2019/12/12 03:03:55 system db initiated successfully.
[fcc-server] 2019/12/12 03:03:56 Chaincode client initialed successfully.
[fcc-server] 2019/12/12 03:03:56 Server started on  :1206
```
9. 编译和部署前端。 前端采用VUE，也可使用其它前端框架或HTML。使用GNINX或其它WEB服务器部署编译后的前端代码。注：当前未使用登录和权限设置。

![DEMO](http://xujf000.tk:28888/img/demo.png "DEMO")

第一次操作数据上链时，区块链网络后端会根据背书节点和合约数量创建镜像并启动容器（本例为3*2) 大约耗时30-60秒，之后每次上链操作约1秒，查询小于1秒。


- ### 产品化需要注意的几个问题
- 背书节点与查询节点：涉及合约数据的变更操作（新建、修改、删除）需要多个背书节点（AND、OR策略），查询合约和节点可在同一通道或不同通道；
- 共识：需要采用KAFKA+ZK集群或ETCD集群；
- 证书与权限：通过CA集群和MSP来管理发放不同权限的证书，考虑用证书或其它方式实现基于RBAC的权限控制；
- 如果可以，采用国产加密算法；

