## Week02 
#### 作业题目：

我们在数据库操作的时候，
比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。
为什么，应该怎么做请写出代码？

#### 解答：
当 dao 层遇到一个 sql.ErrNoRows 的错误时，此时是没查到的错误。
dag 层没查到数据，具体怎么处理，应交由上层调用方的逻辑或顶层逻辑来处理。

那么需不需要 wrap 呢，我认为此时调用方是知道此类错误应该怎么处理的，所以无需 wrap，直接返回错误即可。

代码：
```go
package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	ID 		int
	Name    string
	Age     int
}

func (u *User) findUserById(id int) error {
	err := db.QueryRow("SELECT * FROM USERS WHERE ID = ?", id).Scan(u)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		} else {
			return errors.Wrap(err, "undefined error")
		}
	}
	return nil
}

```

## 学习笔记

#### Error 定义或封装：
1. Sentinel errors —— 预定义的 Error。
将会成为API的一部分，
在两个包之间创建了依赖。
2. Error types —— 实现了 error 接口的自定义类型。
因为是一个 type，所以调用者可以使用断言转换成这个类型，来获取更多的上下文信息。
3. Opaque errors —— 不透明的错误处理。只需返回错误即可。

#### Error Handling：

1. 遵循 error 只处理一次的原则；
2. 一般情况直接返回 error 即可，而不是在每个错误产生的地方都打上日志；
3. 在程序的顶部或者是工作的 goroutine 顶部(请求入口)，使用 %+v 把堆栈详情记录；
4. 使用 errors.Wrap 或者 errors.Wrapf 在程序逻辑中如果不适合立即处理 error ，并且还需要将调用栈信息以及一些自定义的msg记录下来。
5. 可以使用 Wrap error 对一般的基础库调用产生的 error 进行 wrap。
6. 对于特定的 error type 可以使用 errors.Is 或 errors.As
 来做错误检查。

#### Reference
1. [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)
2. [Error handling and Go](https://blog.golang.org/error-handling-and-go)
