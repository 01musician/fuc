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

## show innodb status
```sql
show engine innodb status
```
