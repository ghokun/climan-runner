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
```
█▀▀ █░░ █ █▀▄▀█ ▄▀█ █▄░█  Version : 0.0.4
█▄▄ █▄▄ █ █░▀░█ █▀█ █░▀█  Commit  : 2bd3b3a043440baf9666c66b4847d36445a00a0d

NAME      LATEST  DESCRIPTION                                      
argocd    v2.0.1  Declarative continuous deployment for Kubernetes 
arkade    0.7.15  Open Source Kubernetes Marketplace               
climan    v0.0.4  Cloud tools cli manager                          
crc       1.26.0  Local single node Openshift                      
helm      v3.5.4  The Kubernetes Package Manager                   
k3d       v4.4.3  k3s in Docker                                    
kind      v0.10.0 Kubernetes in Docker                             
kn        v0.22.0 Knative cli                                      
kubectl   v1.21.1 Kubernetes command line tool                     
kustomize v4.1.2  Customization of kubernetes YAML configurations  
minikube  v1.20.0 Run Kubernetes locally                           
odo       v2.2.0  Developer-focused cli for OpenShift              
tkn       v0.18.0 Cli for interacting with Tekton                  
```
