# GO Coding Guide

[English](README.md) | [中文](README_ZH_CN.md)

> go 语言官方编码规范指导,https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters
>
> 参考ashleymcnamara的ppt分享

> Following these conventions will make your source code easier to read, easier to maintain, and easier for someone else to understand. 
> Write Code Like the Go Team
> - how to organize your code into packages, and what those packages should contain;
> - code pattern and conventions that are prevalent in the standard library;
> - how to write your code to be more clear and understandable;
> - unwritten Go conventions that go beyond "go fmt" and make you look like a veteran Go contributor wrote it ;

## Outline
> - Package
> - Naming Conventions
> - Sources Code Conventions

## Package Code Organization 
### a. library 
There are two key areas of code organization in Go that will make a huge impact on the usability,testability,and functionality of you code;
- Package Naming
- Package Organization

Packages contain code that has single purpose
- `archive`  `cmd` `crypto` `errors` `go` `index` `matl`

when a group of packages provides a common set of functionalities with different implementations, they're organized under aparent.
Look at the contents of package encoding:
- `ascii85`  `base32` `binary` `encoding.go` `hex` 
- `asn1`     `base64` `csv`   `gob`         `json`

Some commonalities:
- Packages names describe their purpose
- It's very easy to see what a package does by looking at the name
- Names are generally short
- When necessary, use a descriptive parent package and several children implementing the functionality -- like the encoding package
 ```
  ├─encoding
  │  ├─charmap
  │  ├─htmlindex
  │  ├─ianaindex
  │  ├─internal
  │  │  └─identifier
  │  ├─japanese
  │  ├─korean
  │  ├─simplifiedchinese
  │  ├─testdata
  │  ├─traditionalchinese
  │  └─unicode
  │  ├─encoding.go // is the interface to encoding
  
 ```
The packages we've seen are all libraries. They're intended to be imported and used by some executable program like
a service or command line tool.

### b. application
What should the organization of your executable applications look like?

When you have an application, the package organization is subtly different. The difference is the command,the executable
that ties all of those packages together. 

Application package organization has a huge impact on the testability and functionality of your system.

When writing an application your goal should be to write code that is easy to understand, easy to refactor, and
easy for someone else to maintain.

Most libraries focus on providing a singularly scoped function; logging, ending, network access.

You application will tie all of those libraries together to create a tool or service. That tool or service will be much larger in scope.

When you're building an application, you should organize your code into packages, but those packages should be centered on two categories:
- Domain Types
- Services

Domain Types are types that model your business functionality and objects.

Services are packages that operate on or with the domain types.

[sample](https://medium.com/@benbjohnson/standaard-package-layout-7cdbc8391fc1)

The package containing your domain types should also define the interfaces between your domain types and the rest of the world. These interfaces define the things you want to do tith your domain types.
- ProductService
- SupplierService
- AuthenticationService
- EmployeeStorage
- RoleStorage 

Your domain type package should be the root of your application repository. This makes it clear to anyone opening the codebase what types are being used, and what operations will be performed on those types.

The domain type package, or root package of your application should not have any external dependencies.
- It exists for the sole purpose of describing your types and their behaviors.

The implementations of your domain interfaces should be in separate packages, organized by dependency. 

Dependencies include:
- External data sources
- Transport logic (http, RPC)

You should have one package per dependency. 

Why one package per dependency?
- simple testing
- Easy substitution/ replacement
- No circular dependencies
              
## Naming Conventions
> there are two hard thing in computer science:cache invalidation, naming thing, and off-by-one errors
> - Every developer on  Twitter 

Naming things is hard, but putting some thought into your type, function, and package names will make your code more readable.

### a. packages 
A package name should have the following characteristics:
- short
  - Prefer "transport" over "transportmechanisms"
- clear
  - Name for clarity based on function:"bytes"
  - Name to describe implementation of external dependency: "postgres"

Packages should provide functionality for one and only one purpose. Avoid catchall packages:
- util
- helpers
- etc.

Frequently they're a sign that you're missing an interface somewhere.
`util.ConvertOtherToThing()` should probably be refactored into a Thinger interface
catchall packages are always the first place you'll run into problems with testing and circular dependencies, too. 

### b. Variables

Some common conventions for variable names:
- user cameCase not snake_case
- use single letter variables to represent indexes
  - `for i:=0; i < 10; i++ {}`
- use short but descriptive variable names for other things
  - var count int
  - var cust Customer
There are no bonus points in Go for obfuscating your code by using unnecessarily short variables.

Use the scope of the variable as your guide. The farther away from declaration you use it, the longer the name should be. 

- use repeated letters to represent a collection/slice/array
  - `var tt []*Thing`
- inside a loop/range, use the single letter
  - `for i, t := range tt{}`
These conventions are common inside Go's own source code.   

### c. function and methods

Avoid a package-level function name that repeats the package name.
- GOOD: `log.Info()`
- BAD: `log.LogInfo()`

The package name already declares the purpose of the package, so there's no need to repeat it.

Go code does't have setters and getters.
- GOOD: `custSvc.Customer()`
- BAD:  `custSvc.GetCustomer()`

If your interface has only one function, append '-er' to the function name:
````go
type Stringer interface{
	String() string
}
````

If your interface has more than one function, use a name to represent its functionality:
````go
type CustomerStorage interface{
	Customer(id int)(*Customer, error)
	Save(c *Customer) error
	Delete(id int) error
}
````

Inside apackage separate code into logical concerns.

If the package deals with multiple types, keep the logic for each type in its own source file:
````
package: postgres

orders.go
suppliers.go
products.go
````

## Sources Code Conventions
In the package that defines your domain objects, define the types and interfaces for each object in the same source file:
````
package: inventory

orders.go
-- contains Orders type and OrderStorage interface
````

Make comments in full sentences, always.
````go
// An Order represents an order from a customer. 
type Order struct{}
````

Use `goimports` to manage your imports, and they'll always be in canonical order. Standard lib first, external next. 

Avoid the else clause. Especially in error handling. 
````go
if err != nil {
	// error handling
	return  // or continue, etc.
}
// normal code
````

