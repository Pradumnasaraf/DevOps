## Bash Scripting

### Resources

- [Chmod Calculator](https://chmodcommand.com/)

We start by creating a file with the `.sh` extension. For example `script.sh`. Then we write the script in it. For example:

Basic Script

```bash
#!/bin/bash

echo "Hello World"
```    

We can run this by `bash script.sh` or `./script.sh`. But the second will only work if the script is executable (permission to execute). We can make it executable by `chmod +x script.sh`. Now we can run it by `./script.sh`.

### Shebang

The first line of a bash script is called the shebang. It tells the system which interpreter to use to run the script. The shebang for bash is `#!/bin/bash`. The shebang for python is `#!/usr/bin/env python`. It vary from the langauge to language.

### Variables

It is placeholder for a value. Just like any other programming language. We use $ to access the variable. Eg: `$NAME`. We can also use `${NAME}`. The braces are used to make sure that the variable is not mistaken for a command.

```bash
#!/bin/bash

# Variable Declaration
NAME="John"

# Variable Usage

echo "My name is $NAME"
echo "My name is ${NAME}"
```

### User Input

We can take input from the user using the `read` command.

```bash
#!/bin/bash

echo "Enter your name: "
read NAME
echo "Hello $NAME, nice to meet you!"
```

### Arguments

We can pass arguments to the script. The arguments are stored in the `$1`, `$2`, `$3` and so on. `$0` is the name of the script.

Eg: `bash script.sh arg1 arg2`

```bash
#!/bin/bash

echo "First Argument: $1"
```

### If Statement

```bash
#!/bin/bash

if [ "$1" == "John" ]
then
    echo "Hello John"
elif [ "$1" == "Doe" ]
then
    echo "Hello Doe"
else
    echo "I don't know you"
fi
```





