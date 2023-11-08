
# go-js

Hi, this is my first published project (i hope). Go-js is mini package that aim to help JavaScript developer transition to Golang.


### 1. Installation
- Go get
```
go get github.com/bimonugraraga/go-js
```
- Import
```
import (
    "github.com/bimonugraraga/go-js"
)
```
### 2. Function
- Ternary
    ```
    exm := gojs.Ternary(condition, result1, result2)
    ```
    ##### 1. condition is ``` bool``` can be form of ```x == x && ...```
    ##### 2. result1 and result2 is ```interface```, if condition is ```true``` then Ternary will return ```result1``` and vice versa
- IsFalsy 
    ```
    exmIn := nil
    exmOut := gojs.IsFalsy(exmIn)
    ```
    IsFalsy will return true as long as input is ```nil, "", error, and 0```



---
You can email bimonugraraga@gmail.com for feedback or complaint. Thank you. 

