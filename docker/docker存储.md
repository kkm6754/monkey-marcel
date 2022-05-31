# Docker 存储
---

Docker 镜像层不可写，容器层可写，容器层的数据不能方便移动、管理。

Docker可以将数据从主机挂载到容器中，并提供了三种方式：
1. 卷 volumes
   - 完全由Docker管理，一般存储在主机的特定目录下（/var/lib/docker/volumes/）
2. 绑定挂载 bind mount
    - 主机上的指定文件或目录
3. 临时文件系统 tmpfs
   - 存储在主机内存中



## 存储挂载使用

通过 docker container run 的两个参数来实现挂载

- -v 或 --volume
  1. 由三个由冒号（:）分隔的字段组成，[HOST-DIR:]CONTAINER-DIR[:OPTIONS]。
  2. HOSt-DIR 可以是卷，也可以是主机的目录（文件）; 不指定时，默认会创建一个匿名卷。
  3. CONTAINER-DIR 容器内的目录或文件，必须要指定
  4. OPTIONS 配置，可以多个，通过逗号分隔
     1. 表示读写权限 ro rw
     2. 表示卷是否多容器可使用 Z：当前容器使用  z: 可以多个容器使用
- --mount
  1. 由多个键值对组成，键值对之间由逗号分隔。例如：type=volume,source=volume1,destination=/volume1,ro=true。
  2. type 指定存储类型，可以是 volume、 bind、 tmpfs
  3. source（或src） 指定挂载源，可以指定卷名，可以指定主机目录；类型是volume或tmpfs时，可以不指定。
  4. destination(或dst 或target) 指定容器中目录或文件
  5. ro 配置项，


```
docker run -d \
  --name=nginxtest \
  --mount source=nginx-vol,destination=/usr/share/nginx/html \
  nginx:latest

docker run -d \
  --name=nginxtest \
  -v nginx-vol:/usr/share/nginx/html \
  nginx:latest
```




## volumes

- 有名称的卷
- 匿名的卷


### 本地卷列表

```
docker volume ls
```


### 卷创建

```
docker volume create [VOLUME_NAME]
```

不指定卷名称，会生成一个匿名卷，匿名卷会自动生成一个随机值ID作为唯一标识。


## bind mount

- 如果绑定挂载时指定的容器目录是非空的，则该目录中的内容将会被覆盖。
- 如果主机上的目录不存在，会自动创建该目录。


### 文件写问题

挂载文件，同时在容器中和主机写，可能会出现一些特殊情况。





## tmpfs

```
docker run \
    -it \
    --mount type=tmpfs,target=/test \
    --name shiyanlou008 \
    --rm ubuntu bash
```




# 数据卷容器
---

