# 程序开发tips

## 关闭信号处理放在程序初始化最后
为保证程序的健壮性，一般会在程序启动时把所有信号处理都关闭。但若在程序初始化时过早的关闭信号处理程序，会导致程序在初始化时的错误无法暴露出来，导致程序在正常运行的假象。因此关闭信号一般放在程序初始化的最后，待程序所有初始化工作完成后再关闭信号处理。  
关闭信号处理的进程如同运行在“临界区”，应尽量使临界区变得尽可能的短。

## 流量控制

### 流量控制的几种方法

#### stop and wait

#### sliding window

### 流量控制的几种类型

#### open-loop flow control
The open-loop flow control mechanism is characterized by having no feedback between the receiver and the transmitter. 

#### closed-loop flow control
