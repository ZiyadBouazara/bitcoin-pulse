# Testing

* Relevant tutorial : https://www.jetbrains.com/guide/go/tutorials/mock_testing_with_go/setting_up/
 
## Mocking with GoMock and Mockgen
We are opting for an automated approach to mocking. This method leverages code generation to create mock implementations of our interfaces, which can greatly simplify our testing process and reduce manual effort.

**Benefits of Using GoMock and mockgen:**

* Efficiency: Quickly generate mocks without writing boilerplate code.
* Maintainability: Automatically keep mocks in sync with interface changes.
* Advanced Features: Utilize GoMock's capabilities for setting expectations, specifying return values, and asserting call orders.
* Community Support: GoMock is widely used and supported in the Go community, providing resources and examples.

**GoMock and Mockgen are tools that we use to generate automatically mocks directly from files for us.**

You can call mockgen with three flags:

`-source` to define the source file to create mocks from

`-destination` to set the file name of the output file

`-package` to set the package to use for the resulting mock package

Here is an example for the `ports` package and `ports.go` file :

* You can run this command to generate the mocks for you :

```bash 
mockgen -source=path/to/ports.go -destination=path/to/mock_ports.go -package=path/to/mocks
```

Whenever interfaces change, regenerate your mocks to ensure they accurately reflect the new method signatures and behaviors.