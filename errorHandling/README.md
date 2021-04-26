# Error Handling in Golang
##### Errors are one of the most important aspects of a programming language. The way you handle errors impacts the performance of the application in many ways such as,
##### - Consistency
##### - Traceability
##### - Debuggability
##### - Maintainability
##### How errors are defined in Golang is a bit different from languages like Java. In Go, errors are values. For example, you can assign an error to a variable in the same way you assign an integer to a variable. Consider the following example,
##### config, err := ioutil.ReadFile("config.json")
      if err!=nil{
      fmt.Print("Error:",err)
      }
##### We can handle Go errors mainly in two ways,
##### - We can panic (Will crash the application)
##### - We can handle gracefully. (Can log the error and return)

### 1. When to panic
##### When some unexpected issue happens, panic can be used. Mostly panic is being used to fail the application in case of any issue which interrupts the normal operation of the application. For example, we can think of a program which uses a MySQL database to store data. Usually, the application would try to establish a connection with the MySQL database when initializing. But if the application fails to establish the connection with the database, the application can’t continue to function properly. So in this kind of scenarios, the application should panic.
##### db, err := sql.Open("mysql", "username:pw@tcp(localhost:3306)/db")
      if err != nil {
      panic(err.Error())
      }
##### Panic will result in a stack trace which will allow us to trace the error. Here is an example,
##### goroutine 11 [running]:
      testing.tRunner.func1(0xc420092690)
      /usr/local/go/src/testing/testing.go:711 +0x2d2
      panic(0x53f820, 0x594da0)
      /usr/local/go/src/runtime/panic.go:491 +0x283
      github.com/yourbasic/bit.(*Set).Max(0xc42000a940, 0x0)
      ../src/github.com/bit/set_math_bits.go:137 +0x89
      github.com/yourbasic/bit.TestMax(0xc420092690)
      ../src/github.com/bit/set_test.go:165 +0x337
      testing.tRunner(0xc420092690, 0x57f5e8)
      /usr/local/go/src/testing/testing.go:746 +0xd0
      created by testing.(*T).Run
      /usr/local/go/src/testing/testing.go:789 +0x2de
### 2. When not to panic
##### But consider an application which allows users to login. What if a user tries to login with an email which doesn’t exist in the database. In this kind of scenarios, we can’t panic. We have to handle the error gracefully. We can log the error with the login details that the user entered and return an error response to the user.
There is no hard and fast rule for handling errors gracefully in Golang. So here is a better way to handle errors in your application gracefully.
### 3. Handling errors gracefully
##### We will be using a custom-defined struct in order to wrap the error rather than passing the plain error. Let’s define an errors package inside our application.
##### Error struct contains:
##### - Operation field is meant to contain the names of the functions, where the errors were wrapped.
##### - Type field contains the type of the error. For example, an error can be of type NotFoundError. This error type is very useful when we are translating the error into an error code or an HTTP error response.
##### - Error will contain the error object before wrapping it using the struct.
##### - Severity will contain log level, which makes it very easy when we are logging the error in our application.
