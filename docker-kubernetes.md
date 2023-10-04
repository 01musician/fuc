# Docker & Kubernetes

## Docker
### Docker Network
#### Docker Network Overview

 Container networking refers to the ability for containers to connect to and communicate with each other, or to non-Docker workloads.

 A container has no information about what kind of network it's attached to, or whether their peers are also Docker workloads or not. A container only sees a network interface with an IP address, a gateway, a routing table, DNS services, and other networking details. That is, unless the container uses the none network driver.

#### Published ports
#### Ip address and hostname

By default, the container gets an IP address for every Docker network it attaches to. A container receives an IP address out of the IP subnet of the network. The Docker daemon performs dynamic subnetting and IP address allocation for containers. Each network also has a default subnet mask and gateway.
 
