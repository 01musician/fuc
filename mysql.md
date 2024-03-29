# MYSQL/Mariadb 相关

* 查看mariadb version

```sql
select version();
```

* The MariaDB Performance Schema is a feature for monitoring the performance of your MariaDB server.
* MariaDB comes pre-installed with a system database called mysql containing many important tables storing, in particular, grant and privilege information.


* 打印mariadb的默认设置
```bash
my_print_defaults
```

## Frequently use sql
* get distinct value number in a column
```sql
select count(distinct column_name) as distinct_count from table_name;
```


## 设置mariadb/mysql的用户权限

* 使用空用户名要注意 
The user name part of an account name is either a nonblank value that literally matches the user name for incoming connection attempts, or a blank value (empty string) that matches any user name.


## check mariadb tables size
```sql
SELECT table_name, round(((data_length + index_length) / 1024 / 1024), 2) AS "Size (MB)"
FROM information_schema.TABLES
WHERE table_schema = 'your_database' AND table_name = 'your_table';
```
or 

```
show table status like 'your_table';
```

## innodb related

### show innodb status
```sql
show engine innodb status
```

### change innodb_pool_size without restart mariadb
```sql
set global innodb_buffer_pool_size = 8G;
```

## check table's storage engine
```sql
SELECT TABLE_NAME, ENGINE
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = 'your_database_name'
AND TABLE_NAME = 'your_view_name';
```
## galera cluster
1. grastate.dat (Galera State File)

* The grastate.dat file is a simple text file that stores information about the state of the node in the Galera Cluster. It contains details such as the node's unique identifier (UUID), its current state (e.g., JOINER, DONOR, PRIMARY), and the sequence number of the last applied transaction.
* This file is mainly used for crash recovery and state reconciliation. When a node crashes and restarts, it checks the grastate.dat file to determine its previous state and the last transaction it successfully applied. This information helps the node catch up to the current state of the cluster.
* `grastate.dat` is typically located in the data directory of the MariaDB server, and its format is specific to Galera Cluster.

2. gvwstate.dat (Group Communication Service State File):
* The gvwstate.dat file stores information related to the group communication service in use by the Galera Cluster. It is used to keep track of the group members, their statuses, and the state of the communication group.
* The group communication service, such as the one provided by the underlying multicast or unicast implementation, is responsible for ensuring that all nodes in the cluster can communicate and coordinate effectively. This includes detecting node failures, electing a new primary node, and maintaining a consistent view of the cluster.
* `gvwstate.dat` helps the nodes maintain a consistent group view even if some nodes experience temporary network issues or failures. It contains information about the nodes' UUIDs, their view numbers, and the state of each node within the group.
