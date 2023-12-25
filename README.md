# simpleFileServer

## 简介

### 起步

这个项目只是我熟悉golang过程中的产物
启动后会通过环境变量<mark>WORK_HOME</mark>获取工作目录；<mark>PASSWD</mark>获取管理员密码;<mark>ROOT_USER</mark>获取管理员用户名；
如果这些都没有指定默认值如下
<mark>WORK_HOME</mark>:./
<mark>PASSWD</mark>:123456
<mark>ROOT_USER</mark>:root

### webdav

这块的思路来源是[fungaren/gin-webdav: WebDAV server for gin-gonic (github.com)](https://github.com/fungaren/gin-webdav)

感谢大佬提供的方案

webdav的路径是 `http://127.0.0.1:8080/webdav` ~~可以使用windows的资源管理器访问~~ （由于windows默认只支持https的webdav身份认证，所有有需要请自行修改注册表来支持)

推荐使用[https://www.raidrive.com/]() 这个软件连接

## 相关接口

### 获取文件夹或文件二进制流

```shell
curl --location '127.0.0.1:8080/api/你的实际文件路径'
```

如果请求的路径是文件夹会返回json，如果是文件会返回二进制流

```json
{
    "Data": {
        "RootPath": "./",
        "Parent": "./下载/",
        "NowPath": "./下载/02_SourceHanSerif-VF",
        "FileItems": [
            {
                "Name": "LICENSE.txt",
                "Size": 4463,
                "Mode": 438,
                "ModTime": "2023-08-17T08:53:16+08:00",
                "IsDir": false
            },
            {
                "Name": "Variable",
                "Size": 0,
                "Mode": 2147484159,
                "ModTime": "2023-08-18T11:09:20+08:00",
                "IsDir": true
            }
        ]
    }
}
```

### 上传文件与创建文件夹

```shell
curl --location '127.0.0.1:8080/api/需要创建文件或文件夹的夫级路径' \
--form 'path="新建文件夹2"' \
--form 'file=@"/D:/FFOutput/66.bmp"'
```

返回值为文件列表

```json
{
    "Data": {
        "RootPath": "/tmp/test/",
        "Parent": "/tmp/",
        "NowPath": "/tmp/test/",
        "FileItems": [
            {
                "Name": "新建文件夹1",
                "Size": 60,
                "Mode": 2147484141,
                "ModTime": "2023-12-21T01:13:13.237728674+08:00",
                "IsDir": true
            },
            {
                "Name": "新建文件夹2",
                "Size": 40,
                "Mode": 2147484141,
                "ModTime": "2023-12-20T22:58:22.495875604+08:00",
                "IsDir": true
            },
            {
                "Name": "66.bmp",
                "Size": 1314430,
                "Mode": 420,
                "ModTime": "2023-12-21T16:37:26.877759761+08:00",
                "IsDir": false
            }
        ]
    }
}
```

### 删除文件或文件夹

`cm9vdDoxMjM0NTY3`这个东西其实就是`root:1234567`这个字符串用base64加密后得到的

```shell
curl --location --request DELETE '127.0.0.1:8080/api/需要删除的文件路径' \
--header 'Authorization: Basic cm9vdDoxMjM0NTY3'
```

返回值也是json

```json
{
    "Data": "ok"
}
```

## TODO

- [ ] docker镜像构建

- [ ] web的前端页面

- [ ] 用户权限管理

- [ ] 读写改删权限分离

- [ ] 配置文件剥离
