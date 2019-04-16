# CENTOS终端常用命令
###  查看本机所有IP地址
`ifconfig -a`
### 开机即使能网卡
`sed -i "s/ONBOOT=no/ONBOOT=yes/" /etc/sysconfig/network-scripts/ifcfg-网卡名`  
网卡名可以通过`ifconfig -a`命令获得

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


