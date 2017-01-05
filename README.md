# autofill-pwn
A tool for testing if browsers are vulnerable to auto fill phishing

# What is this?
Based on This: https://github.com/anttiviljami/browser-autofill-phishing it would seem browsers will happily auto fill some information. But I wanted a way to test browsers without setting up a full stack. So here's a simple Go application. Just build your index.html file and pass it as an argument to the program:

```
./autofill-pwn -p 8080
2017/01/05 12:41:24 Starting server...
2017/01/05 12:41:29 Serving GET Request...
2017/01/05 12:41:30 Serving GET Request...
2017/01/05 12:41:31 Field address = []
2017/01/05 12:41:31 Field postal = []
2017/01/05 12:41:31 Field city = []
2017/01/05 12:41:31 Field country = []
2017/01/05 12:41:31 Field name = [abc]
2017/01/05 12:41:31 Field email = [abc@foo.com]
2017/01/05 12:41:31 Field phone = []
2017/01/05 12:41:31 Field organization = []
2017/01/05 12:41:31 Serving GET Request...
```

# Building
Is easy!
```
$ go build .
$ ./autofill-pwn
ERROR: Missing Argument
Usage:
    ./autofill-pwn -p 8080 [-f index.html]
```
