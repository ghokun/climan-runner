## climan <a class="github-button" href="https://github.com/ghokun/climan" data-icon="octicon-star" data-size="large" data-show-count="true" aria-label="Star ghokun/climan on GitHub">Star</a> <a class="github-button" href="https://github.com/ghokun/climan/issues" data-icon="octicon-issue-opened" data-size="large" data-show-count="true" aria-label="Issue ghokun/climan on GitHub">Issue</a>
`climan` is command line tool manager for cloud native technologies.

<script async defer src="https://buttons.github.io/buttons.js"></script>

<table>
    <thead>
        <tr>
            <th style="text-align:center">arch/os</th>
            <th style="text-align:center">macos</th>
            <th style="text-align:center">linux</th>
            <th style="text-align:center">windows</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align:center">386</td>
            <td style="text-align:center"></td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-linux-386">Download</a></td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-windows-386.exe">Download</a>
            </td>
        </tr>
        <tr>
            <td style="text-align:center">amd64</td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-darwin-amd64">Download</a>
            </td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-linux-amd64">Download</a>
            </td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-windows-amd64.exe">Download</a>
            </td>
        </tr>
        <tr>
            <td style="text-align:center">arm</td>
            <td style="text-align:center"></td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-linux-arm">Download</a></td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-windows-arm.exe">Download</a>
            </td>
        </tr>
        <tr>
            <td style="text-align:center">arm64</td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-darwin-arm64">Download</a>
            </td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-linux-arm64">Download</a>
            </td>
            <td style="text-align:center"></td>
        </tr>
        <tr>
            <td style="text-align:center">ppc64le</td>
            <td style="text-align:center"></td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-linux-ppc64le">Download</a>
            </td>
            <td style="text-align:center"></td>
        </tr>
        <tr>
            <td style="text-align:center">s390x</td>
            <td style="text-align:center"></td>
            <td style="text-align:center"><a
                    href="https://github.com/ghokun/climan/releases/latest/download/climan-linux-s390x">Download</a>
            </td>
            <td style="text-align:center"></td>
        </tr>
    </tbody>
</table>

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
