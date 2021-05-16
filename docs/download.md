## climan
`climan` is command line tool manager for cloud native technologies.
| arch/os | macos | linux | windows |
|:-:|:-:|:-:|:-:|
| 386 |  | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-linux-386) | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-windows-386.exe) |
| amd64 | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-darwin-amd64) | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-linux-amd64) | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-windows-amd64.exe) |
| arm |  | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-linux-arm) | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-windows-arm.exe) |
| arm64 | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-darwin-arm64) | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-linux-arm64) |  |
| ppc64le |  | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-linux-ppc64le) |  |
| s390x |  | [Download](https://github.com/ghokun/climan/releases/latest/download/climan-linux-s390x) |  |

<sup>[Checksums](https://github.com/ghokun/climan/releases/latest/download/climan-checksums.txt)</sup>
```bash
# Download binary and make it executable
chmod +x climan

# Install
export PATH=$HOME/.climan/bin:$PATH # add to PATH
./climan install climan

# Uninstall
# export PATH=$HOME/.climan/bin:$PATH < remove from  PATH
rm -rf ~/.climan
```
## commands
```bash
# list      : Lists tools or versions of a tool
climan list
climan list kubectl

# install   : Installs the latest or selected version of a tool
climan install kubectl
climan install kubectl v1.18.0

# uninstall : Uninstalls the latest or selected version of a tool
climan uninstall kubectl
climan uninstall kubectl v1.17.0

# default   : Sets default version of a tool
climan default kubectl v1.16.0

# prune     : Uninstalls non-default versions of a tool
climan prune kubectl
```
## tools
