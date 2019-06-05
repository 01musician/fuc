# CENTOS终端常用命令
###  查看本机所有IP地址
`ifconfig -a`
### 开机即使能网卡
`sed -i "s/ONBOOT=no/ONBOOT=yes/" /etc/sysconfig/network-scripts/ifcfg-网卡名`  
网卡名可以通过`ifconfig -a`命令获得
### 网络抓包相关  
* 通过命令行`tcpdump`（需要管理员权限)  
`tcpdump -i 网卡名`  
网卡名可以通过`ifconfig -a`命令获得
* 通过图形化界面 `wireshark`（需要管理员权限)  
先安装图形化抓包工具wireshark  
`yum install -y wireshark-gnome.x86_64`  
再在命令终端节目输入`wireshark`启动图形化界面（需要管理员权限）
### hostname相关
* 暂时设置主机名（只对当前终端有效）（需要管理员权限）  
`hostname XXX` `XXX`表示主机名
* 永久设置主机名（需要管理员权限）  
`hostnamectl set-hostname XXX`

### `tcpdump`相关
* 抓住某个网卡上某个IP地址的包  
`tcpdump -i 网卡名 host ip地址`
* 显示网络包包头和数据区内容    
`tcpdump -i 网卡名 -xx`
* 显示网络包包头和指定长度的数据区内容  
`tcpdump -i 网卡名 -xx -s 长度`
* 把某个网卡加入组播（然后可以在该网卡上捕捉组播包）  
`ip addr add 组播地址/网络掩码 dev 网卡名 autojoin`


### 图形化登录相关
* 禁止图形化登录（需要管理员权限）  
`systemctl set-default multi-user.target`
* 恢复图形化登录（需要管理员权限）  
`systemctl set-default graphical.target`

### 查找文件中特定内容
* 查找当前目录下所有文件内容中是否包含“abc”  
`find . -type f | xargs grep -n "abc"`
* 查找当前目录下所有.h、.c文件内容中是否包含“abc”  
`find . -name "*.[hc]" | xargs grep -n "abc"` 
* 查找当前目录下所有.h、.c和.cpp文件内容中是否包含“abc”  
`find . -name "*.[hc]" -o -name "*.cpp"| xargs grep -n "abc"` 
### 十进制与十六进制互换
* 使用bc：  
`echo "obase=16; 34" | bc`
* 使用printf  
`printf "%x\n" 34`
### 文件编码相关
* 查看一个文件(比如a.c)的编码  
`file a.c`
* 把所有.h .c文件由gb2312编码转换为utf8编码  
`find . -name "*.[hc]" -print0 -exec iconv -f gb2312 -o {}.converted {} \; -exec mv {}.converted {} \;`
* 把当前目录下所有.h, .c, .cpp和makefile文件由gb2312编码转换为utf8编码  
`find . -name "*.[hc]" -o -name "*.cpp" -o -name "makefile" -print0 -exec iconv -f gb2312 -o {}.converted {} \; -exec mv {}.converted {} \;`

### 防火墙相关
* 关闭防火墙（需要管理员权限)  
`systemctl stop firewalld`
* 在系统重启时防止自动启动防火墙（需要管理员权限)  
`systemctl disable firewalld`
* 查看防火墙状态  
`systemctl status firewalld`
* 启动防火墙（需要管理员权限)    
`systemctl enable firewalld`
### 系统服务相关
* 启动telnet服务  
`systemctl start telnet.socket`
* 启动ftp服务器  
`systemctl start vsftpd`
### Yum安装软件相关
* 配置Yum源  
编辑`/etc/yum.repos.d/###-Base.repo`文件，修改`baseurl`字段值，指向Yum源地址
* 查询Yum源中是否包含某软件安装包(比如ftp)  
`yum search ftp`
* 安装软件：查询某软件安装包（ftp）是否存在后，可以通过如下命令安装该软件包:  
`yum install -y ftp`
### SSH相关
* 安装ssh客户端  
`yum install -y ssh`
* 无密码ssh登录(scp拷贝文件)至远程主机  
假定存在两台主机（A，B），两主机均存在账户user，要从A机无密码ssh登录至B机（或者无密码scp文件至B机）:
  1. 先在A机运行如下命令：  
`ssh-keygen -t rsa`  
该命令会在A机`/home/user/.ssh`产生`id_rsa`和`id_rsa.pub`两个文件;
  2. 把`id_rsa.pub`文件中的内容追加至B机`/home/user/.ssh/authorized_keys`文件（若该文件不存在则创建该文件）;  
  3. 在B机使用命令`chmod 600 /home/user/.ssh/authorized_keys`修改文件权限。  
至此，在A机可以无密码登录（无密码scp拷贝文件）至B机







