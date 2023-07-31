# Capture Multi-Enviroment Dragon and Increase your Development Productivity

## 多环境的起源

### 隔离硬件的多环境
三代软件运行的硬件是相互隔离的
### 同一个硬件的多环境
同一个硬件系统中多个环境是相互隔离（四个环境）
### 开发网和生产网相隔离的多环境
开发网和生产网硬件环境相隔离（八个环境）

## 多环境的设计原则


## 不同环境特征
### 系统静态特征
四个目录
一个启动脚本
### 配置项特征
同一份代码/不同配置文件

### 运行时特征
不同环境的进程互不影响
pssf
pstree -p pidof pssf

基础设施
全局段的隔离：进程级；
数据库mariadb：databases级；
kafka：


## 多环境带来影响
### 概念的混乱
外部单位/仿真-->开发网/生产网-->四个环境-->多任务/单任务/单目标
### 环境的串杂
### 链接
qmake/makefile
更新
### 运行
已运行程序
调度内程序
### 配置文件
网络：组播；IP地址；

## 提高开发效率

### 多主机
tmux

### 代码/配置文件多
git

### 组播地址/网络
tcpdump






