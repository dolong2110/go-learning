# Inheritance in Go
##### - Go doesn’t have a concept of class or have a class-based inheritance. Go is a bit different than the regular object-oriented languages.
### 1. How to reuse the code if inheritance is not there?
##### Go has a unique feature called embedded fields in the struct (More like composition). Using those we can reuse the code, it is pretty much similar to what Inheritance provides in other languages.
The Embedded field is a special field, a field without a field name. The example is animal.go file.
##### In this example, Animal is an embedded field in the parrot struct. The Animal struct is behaving like a base struct and Parrot struct has inherited all the fields and methods of it.
Let’s have a look at our main function:
line 41: We are calling a method Eat which is overridden by parrot struct.
line 43: We are calling a method Sleep which is a method of the animal struct(The parent struct).
line 45: We are using field(Weight) of the Animal as it is of Parrot field.
### 2. Benefit of using embedded fields
##### Many times we require our code to be more flexible and change its behavior at the run time. Inheritance is a very tightly coupled feature, as it is defined at compile time we can’t change the class behavior at run time. With composition, we can always change the class behavior with dependency injection or setters.
##### So favor composition over inheritance
##### - We could have an embedded an interface in the Parrot struct. Embedded feature is available for both struct and interface.
##### - Example in animalInterface.go file.
##### Now our error struct implementation is complete. Let’s see how we can use this custom error effectively,
##### - When an error occurs, we can wrap the error using our custom error implementation,
##### const operation errors.Operation = "UserStore.GetUserById"
      db := dbConn()
      selDB, err := db.Query("SELECT * FROM User ORDER BY id DESC")
      if err!=nil{
        return nil, errors.NewError(operation, errors.Unexpected, err, logrus.ErrorLevel)
      }
##### - While passing the error to the upper level, pass it with the operation name
##### const operation errors.Operation = "UserService.GetUser"
      User, err := userStore.getUserById(id)
      if err != nil {
        return nil, err.WithOperation(operation)
      }
##### - We can log the error using the log level set in the error and we can log the operations array as well. It will result in something like,
##### ["UserStore.GetUserById", "UserService.GetUser"]
##### Operations array is much cleaner and parsable when debugging errors. We can trace the error using the operations array.
