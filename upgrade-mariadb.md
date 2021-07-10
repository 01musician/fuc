# 升级mariadb并组建galera集群
由于工作关系，需要把生产环境的mariadb由5.5.56升级至10.1.20并组建galera集群，主要步骤分为如下几步：
## 备份原有mariadb里数据
一般mariadb升级版本数据库数据能得到保存，为避免万一数据丢失，可采用mysqldump来备份数据:  
1. 使用mysqldump备份原有数据库，备份数据库版本命令：  
`mysqldump –databases xx > xxx`  
升级后万一数据丢失，可用如下命令来导入备份的数据 `mysql < xxx`
## 升级mariadb软件
1.	登录老版本数据库，设置innodb_fast_shutdown为0：
`set GLOBAL innodb_fast_shutdown=0`
2.	使用`systemctl stop mariadb.service`停止mariadb数据库；
3.	使用 `yum remove mariadb-server`删除老版本数据库；
4.	确认yum源里包含新版本数据库后，使用`yum install mariadb-server`安装新版本数据库；
5.	使用`systemctl start mariadb.service`启动mariadb数据库；
6.	运行`mysql_upgrade`确认升级后数据库版本与原有数据库内容相容；
7.	若数据库中没包含原有数据，使用`mysqldump`导入备份的数据。

## 组建galera集群
组建galera集群主要包括如下几步：
1. 安装依赖包： 
	`yum install mariadb-server-galera galera`
2. 集群各节点安装好数据库并做好常规的初始化配置；
3. 修改主节点配置文件：`/etc/my.cnf.d/mariadb-server.conf` [mysqld]和[galera]配置如下  
```
[mysqld]
collation-server=utf8_general_ci
character-set-server=utf8
lower_case_table_names=1
skip_name_resolve=ON
server_id=160  # server-id每个节点唯一

[galera]
bind-address=0.0.0.0
wsrep_on=ON
wsrep_provider=/usr/lib64/galera/libgalera_smm.so
wsrep_cluster_address="gcomm://192.168.245.160,192.168.245.16"  #整个集群的IP地址
binlog_format=row
default_storage_engine=InnoDB
innodb_autoinc_lock_mode=2
wsrep_slave_threads=8
innodb_flush_log_at_trx_commit=0
```

4. 停止主节点服务：`systemctl stop mariadb`
5. 启动主机节点并创建集群，执行: `galera_new_cluster`
6. 在主机节点查询集群状态：
 	`mysql> show status like 'wsrep%';`
7. 修改其他节点的配置文件：`/etc/my.cnf.d/mariadb-server.conf`
	除了`server-id`外，其他配置都与主机节点相同
8. 修改完配置文件后重启服务：
	`systemctl stop mariadb.service`
	`systemctl start mariadb.service`

