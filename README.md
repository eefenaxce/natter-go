# NatterGo

NatterGo 是一个用 Go 语言实现的高性能网络通信工具。

## 项目简介

- 采用 Go 1.24 及以上版本
- 支持 UPnP、STUN、DTLS 等协议
- 依赖全部为宽松开源协议（MIT/BSD-2-Clause/BSD-3-Clause）

## 安装方法

```bash
git clone https://github.com/eefenaxce/natter-go.git
cd natter-go
go build ./cmd/...
```

## 使用方法

### 1. 编译所有命令行工具

```bash
go build -o bin ./cmd/...
```
编译后所有可执行文件会在 `bin/` 目录下。

### 2. 运行某个子命令

假设有一个子命令叫 `natter-server`，你可以这样运行：

```bash
./bin/natter-server --help
```

### 3. 常见命令与参数

以 `natter-server` 和 `natter-client` 为例（请根据你实际的cmd目录内容替换）：

#### 启动服务端
```bash
./bin/natter-server --listen :9000 --log-level debug
```
- `--listen`：监听地址和端口（如`:9000`）
- `--log-level`：日志级别（如`info`、`debug`）

#### 启动客户端
```bash
./bin/natter-client --server 127.0.0.1:9000 --user alice
```
- `--server`：服务端地址
- `--user`：用户名

### 4. 查看所有可用命令

```bash
ls ./cmd
```
每个子目录对应一个可执行程序，编译后在 `bin/` 目录下。

### 5. 查看命令帮助

```bash
./bin/xxx --help
```
会显示该命令支持的所有参数和用法。

### 6. 示例：UPnP端口映射

假如有 `natter-upnp` 工具：
```bash
./bin/natter-upnp --add --port 12345 --desc "NatterGo Test"
```

---

> 具体命令和参数请以 `cmd/` 目录下实际子命令和 `--help` 输出为准。

## 依赖说明

主要依赖如下：

- [github.com/huin/goupnp](https://github.com/huin/goupnp) (BSD-2-Clause)
- [github.com/pion/stun](https://github.com/pion/stun) (MIT)
- [github.com/pion/dtls/v2](https://github.com/pion/dtls) (MIT)
- [github.com/pion/logging](https://github.com/pion/logging) (MIT)
- [github.com/pion/transport/v2, v3](https://github.com/pion/transport) (MIT)
- [github.com/wlynxg/anet](https://github.com/wlynxg/anet) (MIT)
- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) (BSD-3-Clause)
- [golang.org/x/net](https://pkg.go.dev/golang.org/x/net) (BSD-3-Clause)
- [golang.org/x/sync](https://pkg.go.dev/golang.org/x/sync) (BSD-3-Clause)
- [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) (BSD-3-Clause)

所有依赖均为宽松许可证（MIT/BSD-2-Clause/BSD-3-Clause），可安全用于闭源或商业项目。

## 目录结构

```
NatterGo/
├── cmd/         # 各种命令行子程序入口
├── internal/    # 内部实现代码
├── go.mod       # Go module 文件
├── go.sum       # 依赖校验文件
```

## 许可证 & 贡献

本项目采用 https://choosealicense.com/licenses/apache-2.0/ 许可证。
贡献者需签署 https://cla-assistant.io/eefenaxce/natter-go（贡献者协议），流程自动完成，不影响日常 PR。

```
                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity exercising
      permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but not
      limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship covered by this License,
      whether in Source or Object form, made available under the License,
      as indicated by a copyright notice that is included in or attached
      to the work.

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based upon (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and derivative works thereof.

      "Contribution" shall mean any design or work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control
      systems, and issue tracking systems that are managed by, or on behalf
      of, the Licensor for the purpose of discussing and improving the Work,
      but excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to use, reproduce, modify, merge, publish,
      distribute, sublicense, and/or sell copies of the Work, and to
      permit persons to whom the Work is furnished to do so, subject to
      the following conditions:

      The above copyright notice and this permission notice shall be
      included in all copies or substantial portions of the Work.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, trademark, and attribution
          notices from the Source form of the Work, excluding those notices
          that do not pertain to any part of the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices alongside those of the Work or Derivative Works.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. When redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.

   END OF TERMS AND CONDITIONS

   APPENDIX: How to apply the Apache License to your work.

      To apply the Apache License to your work, attach the following
      boilerplate notice, with the fields enclosed by brackets "[]"
      replaced with your own identifying information. (Don't include
      the brackets!)  The text should be enclosed in the appropriate
      comment syntax for the file format. We also recommend that a
      file or class name and description of purpose be included on the
      same "printed page" as the copyright notice for easier
      identification within third-party archives.

   Copyright 2025 eefenaxce

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```

> 本项目所有依赖均为 MIT/BSD-2-Clause/BSD-3-Clause 等宽松许可证，无 GPL/LGPL 等传染性条款，可放心用于闭源或商业项目。 
