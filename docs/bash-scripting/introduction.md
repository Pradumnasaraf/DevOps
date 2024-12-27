---
sidebar_position: 1
title: Bash Scripting Introduction
description: A deep dive into Bash Scripting, a scripting language for the Unix shell.
tags: ["Bash", "Scripting", "Unix", "Shell"]
keywords: ["Bash", "Scripting", "Unix", "Shell"]
slug: "/bash"
---

We start by creating a file with the `.sh` extension. For example, `script.sh`. Then we write the script in it. For example:

Basic Script

```bash
#!/bin/bash

echo "Hello World"
```    

We can run this by `bash script.sh` or `./script.sh`. But the second will only work if the script is executable (permission to execute). We can make it executable by `chmod +x script.sh`. Now we can run it by `./script.sh`.

### Shebang

The first line of a bash script is called the shebang. It tells the system which interpreter to use to run the script. The shebang for bash is `#!/bin/bash`. The shebang for python is `#!/usr/bin/env python`. It varies from language to language.

### Variables

It is a placeholder for a value. Just like any other programming language. We use $ to access the variable. Eg: `$NAME`. We can also use `${NAME}`. The braces ensure the variable is not mistaken for a command.

```bash
#!/bin/bash

# Variable Declaration
NAME="John"

# Variable Usage

echo "My name is $NAME"
echo "My name is ${NAME}"
```

NOTE: We can create variables by `NAME="John"` through CLI; we can't use it in the script because it is not exported. We can export it by `export NAME="John"`. Now we can use it in the script.

But here is one more catch. If we exit the terminal and open a new one, the variable will be gone. To make it permanent, we can add it to the `.bashrc` file. It is a hidden file in the home directory. We can open it by `vi ~/.bashrc` or any other editor. We can add the variable to the file. Eg: `export NAME="John"`.

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

### Arithmetic Operations

We can do Arithmetic operations in bash. We use the `(( ))` to do Arithmetic operations.

```bash
#!/bin/bash

echo $(( 5 + 5 ))
```

#### Arithmetic Operators

- `+` - Addition
- `-` - Subtraction
- `*` - Multiplication
- `/` - Division
- `%` - Modulus

### Conditional Statements

We can use the `if` statement to check for a condition. The syntax is:

```bash
#!/bin/bash

if [ "$1" == "John" ]
then
    echo "Hello John"
    exit 1
elif [ "$1" == "Doe" ]
then
    echo "Hello Doe"
else
    echo "I don't know you"
fi
```

```bash
if [$1 == "hello"], then echo "Hello World", fi
```

#### Comparison Operators

- `==` - Equal to
- `>` - Greater than
- `<` - Less than
- `>=` - Greater than or equal to
- `<=` - Less than or equal to
- `!=` - Not equal to
- `-ge` - Greater than or equal to
- `-le` - Less than or equal to

#### Boolean Operators

- `-a` - And
- `-o` - Or
- `!` - Not

### Loops

We can use loops to repeat a set of commands. There are two types of loops in bash. `for` and `while`. The body is enclosed in `do` and `done`.

#### For Loop

```bash
#!/bin/bash

for i in 1 2 3 4 5
do
    echo $i
done
```

#### While Loop

```bash

#!/bin/bash

i=1
while [ $i -le 5 ]
do
    echo $i
    (( i++ ))
done
```

#### Break and Continue

We can use `break` and `continue` in loops. `break` will break the loop and `continue` will skip the current iteration.

```bash
#!/bin/bash

for i in 1 2 3 4 5
do
    if [ $i -eq 3 ]
    then
        continue
    fi
    echo $i
done
```

### Functions  

We can create functions in bash. The syntax is:

```bash
#!/bin/bash

function sayHello() {
    echo "Hello World"
}

sayHello
```


- `exit 1` - Exit the script with an error (non-zero exit code).
- $RANDOM gives a random number between 0 and 32767.
- $SHELL gives the path of the shell.
- $USER gives the username of the user.
- $HOSTNAME gives the hostname of the machine.


### jq

jq is a command-line JSON processor. It is used to parse JSON. It is used to extract data from JSON. It is used to transform JSON. It is used to generate JSON.

#### Installation

```bash
sudo apt install jq
```

#### Usage

The format in JSON.

```bash
echo '{"name": "John", "age": 30}' | jq
```

It will print out the specified key.

```bash
echo '{"name": "John", "age": 30}' | jq '.name'
```

### What's next?

- [Learning Resources](./learning-resources.md) - Learn more about Bash Scripting with these resources.
- [Other Resources](./other-resources.md) - Explore more about Bash Scripting with these resources.
- [Tools](./tools.md) - Explore the tools used in Bash Scripting.
