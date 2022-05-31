# Docker 命令


## 命令格式
---

```
~ docker --help
Usage:  docker [OPTIONS] COMMAND

Options:
      --config string      Location of client config files (default "/Users/miao/.docker")
  -c, --context string     Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with "docker context use")
  -D, --debug              Enable debug mode
  -H, --host list          Daemon socket(s) to connect to
  -l, --log-level string   Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")
      --tls                Use TLS; implied by --tlsverify
      --tlscacert string   Trust certs signed only by this CA (default "/Users/miao/.docker/ca.pem")
      --tlscert string     Path to TLS certificate file (default "/Users/miao/.docker/cert.pem")
      --tlskey string      Path to TLS key file (default "/Users/miao/.docker/key.pem")
      --tlsverify          Use TLS and verify the remote
  -v, --version            Print version information and quit

Management Commands:
  builder     Manage builds
  buildx*     Docker Buildx (Docker Inc., v0.8.2)
  checkpoint  Manage checkpoints
  compose*    Docker Compose (Docker Inc., v2.4.1)
  config      Manage Docker configs
  container   Manage containers
  context     Manage contexts
  image       Manage images
  manifest    Manage Docker image manifests and manifest lists
  network     Manage networks
  node        Manage Swarm nodes
  plugin      Manage plugins
  sbom*       View the packaged-based Software Bill Of Materials (SBOM) for an image (Anchore Inc., 0.6.0)
  scan*       Docker Scan (Docker Inc., v0.17.0)
  secret      Manage Docker secrets
  service     Manage services
  stack       Manage Docker stacks
  swarm       Manage Swarm
  system      Manage Docker
  trust       Manage trust on Docker images
  volume      Manage volumes

Commands:
  attach      Attach local standard input, output, and error streams to a running container
  build       Build an image from a Dockerfile
  commit      Create a new image from a container's changes
  cp          Copy files/folders between a container and the local filesystem
  create      Create a new container
  diff        Inspect changes to files or directories on a container's filesystem
  ...

```

- Management Commands : 在Docker 1.13 之后的命令，对命令进行了分组（推荐）
- Commands ：在Docker 1.12 之前，都是顶级命令
- Options ：命令选项

### 命令选项

命令选项有几种格式：
- 长选项
- 短选项
- 复合选项
- 无选项


```
# 长选项
docker container ls --all

# 短选项
docker container ls -a
docker container ls -q

# 复合选项
docker container ls -aq

# 无选项
docker container ls
```



## 容器命令
---

命令格式
```
# Management Commands
docker container run [OPTIONS] IMAGE [COMMAND [ARGS...]]

# 旧命令格式如下：
docker run [OPTIONS] IMAGE [COMMAND [ARGS...]]
```

一些常用的配置项为：
- -i或--interactive， 交互模式。
- -t或--tty， 分配一个pseudo-TTY，即伪终端。
- -d或--detach，在后台运行容器并输出容器的ID到终端。
- --rm在容器退出后自动移除。
- -p将容器的端口映射到主机。
- -v或--volume， 指定数据卷。


### 创建容器

命令格式
```
# Management Commands
docker container create [OPTIONS] IMAGE [COMMAND] [ARG...]

# 旧的命令格式如下：
docker create [OPTIONS] IMAGE [COMMAND] [ARG...]
```

新创建的容器的状态 (STATUS) 为Created。可以指定容器名，没指定就默认生成一个容器名。

容器唯一标识
- 长ID
- 短ID
- 容器名


常用参数
- -name指定一个容器名称，未指定时，会随机产生一个名字
- -hostname设置容器的主机名
- -mac-address设置MAC地址
- -ulimit设置ulimit选项


#### name与hostname

-hostname 和 -ip 类似，都是设置容器内部的hostname和ip，并不能作为其他容器连接的dns。
name 设置容器名


#### ulimit



### 启动容器

命令格式
```
# Management Commands
docker container start [OPTIONS] CONTAINER [CONTAINER...]

# 旧的命令格式如下：
docker start [OPTIONS] CONTAINER [CONTAINER...]
```


docker container run 是create 和 start命令二合一，同时还支持一些额外操作。

例如 --rm 在容器运行退出后，完成docker container rm 操作

通过输入Ctrl+P+Q把前台容器放入后台运行。


### 停止容器
命令格式
```
# Management Commands
docker container stop CONTAINER [CONTAINER...]

# 旧的命令格式如下：
docker stop CONTAINER [CONTAINER...]
```

### 重启容器
命令格式
```
# Management Commands
docker container restart CONTAINER [CONTAINER...]

# 旧的命令格式如下：
docker restart CONTAINER [CONTAINER...]
```


### 暂停进程
暂停容器中的进程

命令格式
```
# Management Commands
docker container pause CONTAINER [CONTAINER...]

# 旧的命令格式如下：
docker pause [OPTIONS] CONTAINER [CONTAINER...]
```

### 恢复进程
命令格式
```
# Management Commands
docker container unpause CONTAINER [CONTAINER...]

# 旧的命令格式如下：
docker unpause [OPTIONS] CONTAINER [CONTAINER...]
```


### 查看容器列表
命令格式
```
# Management Commands
docker container ls [OPTIONS]

# 旧的命令格式如下：
docker ps [OPTIONS]
```

可选配置项
- a 显示所有的容器。
- q 仅显示ID。
- s 显示总的文件大小。


### 连接到正在运行的容器
命令格式
```
# Management Commands
docker container attach [OPTIONS] CONTAINER

# 旧的命令格式如下：
docker attach [OPTIONS] CONTAINER
```

#### attach与exec

- attach 是通过连接stdin，连接到容器内输入输出流，会在输入exit后终止进程。
- docker exec -it DOCKER_ID /bin/bash 会创建一个伪终端，不会因为exit而终止进程。


### 查看容器元数据
命令格式
```
# Management Commands
docker container inspect [OPTIONS] CONTAINER [CONTAINER...]

# 旧的命令格式如下：
docker inspect [OPTIONS] CONTAINER [CONTAINER...]
```

### 容器日志管理
命令格式
```
# Management Commands
docker container logs [OPTIONS] CONTAINER

# 旧的命令格式如下：
docker logs [OPTIONS] CONTAINER
```

常见配置项
- -t或--timestamps显示时间戳
- -f实时输出，类似于tail -f


### 显示容器中的进程

命令格式

```
# Management Commands
docker container top CONTAINER

# 旧的命令格式如下：
docker top CONTAINER
```

### 查看文件修改
命令格式

```
# Management Commands
docker container diff CONTAINER

# 旧的命令格式如下：
docker diff CONTAINER
```

查看相对于镜像的文件系统来说，容器中做了哪些改变。


### 删除容器
命令格式

```
# Management Commands
docker container rm [OPTIONS] CONTAINER [CONTAINER...]

# 旧的命令格式如下：
docker rm [OPTIONS] CONTAINER [CONTAINER...]
```


删除所有停止的容器

```
docker container prune [OPTIONS]
```


常见配置项
- -f或--force跳过警告提示
- --filter执行过滤删除


## 镜像管理
---


### 镜像列表
命令格式

```
# Management Commands
docker image ls

# 旧的命令格式如下：
docker images
```


### 查看镜像详情
命令格式

```
# Management Commands
docker image inspect ubuntu

# 旧的命令格式如下：
docker inspect ubuntu
```


### 搜索镜像

```
docker search ubuntu
```


### 拉取镜像

```
# Management Commands
docker image pull [OPTIONS] NAME[:TAG|@DIGEST]

# 旧的命令格式如下：
docker pull [OPTIONS] NAME[:TAG|@DIGEST]
```


### 构建镜像
在修改容器的基础上构建镜像

```
# Management Commands
docker container commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]

# 旧的命令格式如下：
docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]
```


### BUILD镜像
从一个 Dockerfile 文件中自动读取指令构建一个新的镜像。

```
docker image build [OPTIONS] PATH | URL
```


### 删除镜像

```
# Management Commands
docker image rm ubuntu:latest

# 旧的命令格式如下：
docker rmi ubuntu:latest
```




