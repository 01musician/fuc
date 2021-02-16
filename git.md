# Git

## 版本控制的历史

Local Version Control Systems --> Centralized Version Control Systems(svn) --> Distributed Version Control Systems(git).

In DVCS, clients don't just check out the latest snapshot of the files: they fully mirror the repository.

The three states:

 ![logo](images/local-operation.png "local operations")

## git工作流程

git管理文件的状态变迁
![logo](images/file-lifecycle.png "file status lifecycle")


## 分支（branch）


## 有趣的问题

### 如何用远程分支覆盖本地文件
```git fetch --all```
```git reset --hard origin/master```
