# Go-Utils

Go-Utils是一个集成的Golang工具库，旨在提供常用的功能和工具函数，以便加速和简化Golang应用程序的开发过程。

## 功能特点

- 提供了丰富的工具函数和实用工具，涵盖了各种常见的编程任务和功能。
- 模块化设计，简单易用，方便集成到现有的Golang项目中。
- 经过精心编写和优化，具有高性能和高效的特点。
- 持续更新和维护，以适应不断变化的开发需求。

## 安装

使用go get命令来安装Go-Utils：


```bash
go get -u github.com/cc-hearts/go-utils
```

## 使用示例

以下是一些使用Go-Utils的示例代码：

```go
package main

import "github.com/cc-hearts/go-utils/lib/validator"

func main() {
	if validator.IsValidJson(`{"name": "cc heart"}`) == true {
		// ...
    }
}

```

## 许可证

[MIT](./License)