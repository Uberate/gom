# 构建 GOM

GOM 提供了多种二进制能力

- [ ] 控制台工具
- [ ] 网络运行工具
    - [ ] `Docker` 镜像
    - [ ] `Kubernetes` - `deployment.app` 的 `YAML` 配置文件

## 1. 从源码构建

拉取源代码:

```shell
git clone git@github.com:Uberate/gom.git

cd gom
```

### 1.1. 构建控制台工具

#### 1.1.1. 构建

通过 makefile 执行构建脚本:

> 在仓库的根目录下执行

```shell
make gom-all-arch

# 如果提示：
# fatal: No names found, cannot describe anything. 
# 为正常现象，不影响构建，但是构建的产物 version 信息会出现异常，请确认是否拉取到正确的 git-tag 信息。
```

上述命令执行后，在仓库的根目录下会出现 `output` 文件夹，包含了所有 `gom` 支持的平台、架构的二进制文件。

#### 1.1.2. 清理

通过 makefile 执行清理脚本:

> 在仓库的根目录下执行

```shell
make clean 

# 如果提示：
# fatal: No names found, cannot describe anything.
```
