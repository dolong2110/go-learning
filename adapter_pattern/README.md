# Adapter Pattern with Golang
### 1. What is adapter pattern?
##### - Adapter pattern is used to make two interfaces with different principle, can communicate. For this we can add an intermediate adapter, which will be in charge of converting from one interface to another.
### 2. Code
##### - first create a struct represent a bicycle in bicycle.go file.
##### - Then add another struct to represent a car in car.go file.
##### - And create an interface for communicate the same order to this structs and this two structs understand us and execute the order in communicate.go file.
##### - And a factory for get the right adapter in factory.go file.
##### - With this we will have the possibility for communicate the same order to this different structs ignoring how this works inside because now we have a “translator” can speak with both of this.

### 3. Conclusion
##### - So with this Design pattern we can make a translator for communicate with different structs they not necessary works same internally.
