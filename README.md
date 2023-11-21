# wxhelper-cli

![GitHub](https://img.shields.io/github/license/wuxs/wxhelper-cli)
![GitHub stars](https://img.shields.io/github/stars/wuxs/wxhelper-cli)
![GitHub forks](https://img.shields.io/github/forks/wuxs/wxhelper-cli)

`wxhelper-cli` 是一个命令行版本的 wxhelper HTTP API 客户端，它允许您通过命令行与 wxhelper 服务进行交互。您可以使用 `WXHELPER_SERVER` 环境变量来指定 wxhelper 服务的地址。

## 功能特点

- 通过命令行与 wxhelper 服务进行交互
- 支持通过环境变量配置 wxhelper 服务地址
- 提供丰富的命令行选项和参数，方便使用各种 wxhelper API 功能

## 安装

您可以通过以下方式安装 `wxhelper-cli`：

```shell
$ go install github.com/wuxs/wxhelper-cli
```

## 使用方法

使用 `wxhelper-cli` 非常简单。只需在命令行中输入相应的命令和参数即可。

```shell
$ wxhelper-cli [options] [command] [arguments]
```

### 配置 wxhelper 服务地址

您可以通过 `WXHELPER_SERVER` 环境变量来配置 wxhelper 服务的地址。例如：

```shell
$ export WXHELPER_SERVER=http://127.0.0.1:19088
```

### 示例命令

下面是一些示例命令的用法：

- 获取当前登录用户信息：

  ```shell
  $ wxhelper-cli userinfo
  ```

- 发送文本消息：

  ```shell
  $ wxhelper-cli send <content> --user <wxid> [--at <wxids>] 
  ```

- 查看帮助信息：

  ```shell
  $ wxhelper-cli --help
  ```

## 贡献

如果您发现任何问题或有改进建议，请随时提交 issue 或 pull request。我们欢迎并感谢您的贡献！

## 许可证

该项目基于 [MIT 许可证](LICENSE)。请查阅许可证文件以获取更多信息。

---

以上是一个简单的 README 示例，您可以根据您的项目实际情况进行修改和扩展。希望对您有帮助！如果您有任何其他问题，请随时提问。