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


### `cut` related
To cut a specific field from a text file or output in Linux terminal using a semicolon (;) as the separator, you can use the cut command with the -d (delimiter) option and the -f (fields) option.
`cut -d ';' -f N <filename>`


### 读取光盘内容 
* 先用blkid命令查找/dev目录下对应cdrom的文件，一般为/dev/sr0，再用mount命令把cdrom附载到指定目录（需要管理员权限) 
`mount /dev/sr0 制定目录`

### `tcpdump`相关
* 抓住某个网卡上某个IP地址的包  
`tcpdump -i 网卡名 host ip地址`
* 显示网络包包头和数据区内容    
`tcpdump -i 网卡名 -xx`
* 显示网络包包头和指定长度的数据区内容  
`tcpdump -i 网卡名 -xx -s 长度`
* 把某个网卡加入组播（然后可以在该网卡上捕捉组播包）  
`ip addr add 组播地址/网络掩码 dev 网卡名 autojoin`
* 按照任意制定位置过滤包（filter for given bytes on given position）
> `expr relop expr`
>       True if the relation holds, where relop is one of >, <, >=, <=,  =,  !=,  and
>       expr  is an arithmetic expression composed of integer constants (expressed in
>       standard C syntax), the normal binary operators [+, -, *, /, &, |, <<, >>], a
>       length  operator,  and  special packet data accessors.  Note that all compar-
>       isons are unsigned, so that, for example, 0x80000000 and 0xffffffff are >  0.
>       To access data inside the packet, use the following syntax:
>            proto [ expr : size ]
>       Proto  is  one of ether, fddi, tr, wlan, ppp, slip, link, ip, arp, rarp, tcp,
>       udp, icmp, ip6 or radio, and indicates the protocol layer for the index oper-
>       ation.   (ether,  fddi,  wlan,  tr,  ppp, slip and link all refer to the link
>       layer. radio refers to the "radio header" added  to  some  802.11  captures.)
>       Note  that  tcp, udp and other upper-layer protocol types only apply to IPv4,
>       not IPv6 (this will be fixed in the future).  The byte  offset,  relative  to
>       the  indicated  protocol layer, is given by expr.  Size is optional and indi-
>       cates the number of bytes in the field of interest; it  can  be  either  one,
>       two,  or  four,  and  defaults to one.  The length operator, indicated by the
>       keyword len, gives the length of the packet.

用这个方法来抓udp包特别合适，因为udp包包头为定长的八个字节，很容易用偏移量来抓包。
* tcpdump中抓包时当包的长度比较小时，后面会自动填充0 
因为ethernet要求最小的包长度为64，在不包括最后4个字节的Frame Check Sequence情况下，发送方会在最后补全字节0.



### 图形化登录相关
* 获取当前默认登录方式
`systemctl get-default`
* 禁止图形化登录（需要管理员权限,重启后生效）  
`systemctl set-default multi-user.target`
* 恢复图形化登录（需要管理员权限,重启后生效）  
`systemctl set-default graphical.target`

### 查找文件中特定内容
* 查找当前目录下所有文件内容中是否包含“abc”  
`find . -type f | xargs grep -n "abc"`
* 查找当前目录下所有.h、.c文件内容中是否包含“abc”  
`find . -name "*.[hc]" | xargs grep -n "abc"` 
* 查找当前目录下所有.h、.c和.cpp文件内容中是否包含“abc”  
`find . -name "*.[hc]" -o -name "*.cpp"| xargs grep -n "abc"` 
* 查找当前目录下30分钟前更新的文件
`find . ! -cmin -30` 

### 查找并替换文件中特定内容
* 查找当前目录下所有文件内容中否包含“abc”的文件，并把“abc”替换成“xyz”  
`find . -type f -print0 | xargs -0 sed -i 's/abc/xyz/g'`

### rename a serise of file whose name contain pattern "a" to name containing pattern "b"
`find . -type f -name "*a*" -exec bash -c 'mv "$1" "${1//a/b}"' _ {} \;`

## 十进制与十六进制互换
* 使用bc：  
`echo "obase=16; 34" | bc`
* 使用printf  
`printf "%x\n" 34`
### 文件编码相关
* 查看一个文件(比如a.c)的编码  
`file a.c`
* 猜测一个文件(比如a.c)的编码（更好的方式）  
`uchardet a.c` 或者 `encguess a.c`
* 把所有.h .c文件由gb2312编码转换为utf8编码  
`find . -name "*.[hc]" -print0 -exec iconv -f gb2312 -t utf8 -o {}.converted {} \; -exec mv {}.converted {} \;`
* 把当前目录下所有.h, .c, .cpp和makefile文件由gb2312编码转换为utf8编码  
`find . -name "*.[hc]" -o -name "*.cpp" -o -name "makefile" -print0 -exec iconv -f gb2312 -t utf8 -o {}.converted {} \; -exec mv {}.converted {} \;`

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
* vsftpd服务器改变默认目录  
在/etc/vsftpd/vsftpd.conf中添加`local_root=/想要的目录`
* vsftpd服务器不能list目录内容 \
`setsebool -P ftpd_full_access=on`


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
* ssh主机登录慢
一般为主机开启了

### ansible相关
* 在ansible的输出中显示主机名，而不是ip地址  
在ansible的主机定义文件中使用格式`hostname ansible_ssh_host=ip`


### vi/vim相关
* 删除所有空白行
`:g/^$/d`或`:v/./d`


### Network Related
* Sending a mulitcast UDP packet does not need gateway 
I ping a local machine without receiving a reply, but i can receive its multicast traffic. Then i found its gateway does not function well. Sending a multicast packet does not need gateway.

### OverlayFS example
* `mkdir lower upper merged`
* `sudo mount -t overlay overlay -o lowerdir=lower,upperdir=upper,workdir=merged merged`
* `sudo umount merged`
