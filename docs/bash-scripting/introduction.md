---
sidebar_position: 1
title: Bash Scripting Introduction
description: A deep dive into Bash Scripting, a scripting language for the Unix shell.
tags: ["Bash", "Scripting", "Unix", "Shell"]
keywords: ["Bash", "Scripting", "Unix", "Shell"]
slug: "/bash"
---

We usually start by creating a file with the `.sh` extension, for example `script.sh`, and then add Bash commands to it.

Basic Script

```bash
#!/bin/bash

echo "Hello World"
```    

You can run this script with `bash script.sh` or `./script.sh`. The second form only works if the file is executable. To make it executable, run `chmod +x script.sh`.

### Shebang

The first line of a Bash script is called the shebang. It tells the system which interpreter should run the file. For Bash, a common shebang is `#!/bin/bash`.

### Variables

A variable is a named value. In Bash, we use `$` to read it, for example `$NAME`. You can also write `${NAME}` when you want to make the variable boundary explicit.

```bash
#!/bin/bash

# Variable Declaration
NAME="John"

# Variable Usage

echo "My name is $NAME"
echo "My name is ${NAME}"
```

NOTE: If you set `NAME="John"` in the shell, it is only a shell variable. To make it available to child processes, export it with `export NAME="John"`.

If you close the terminal, that exported variable is gone. To make it persistent for future shell sessions, add it to `~/.bashrc` or your shell startup file.

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

Example: `bash script.sh arg1 arg2`

```bash
#!/bin/bash

echo "First Argument: $1"
```

### Arithmetic Operations

We can do arithmetic operations in Bash with `(( ))`.

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

The compact one-line form exists, but the block form is easier to read and debug.

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

We can create functions in Bash like this:

```bash
#!/bin/bash

function sayHello() {
    echo "Hello World"
}

sayHello
```


- `exit 1` exits the script with a non-zero status code.
- `$RANDOM` gives a random number between `0` and `32767`.
- `$SHELL` gives the path of the shell.
- `$USER` gives the username of the current user.
- `$HOSTNAME` gives the hostname of the machine.


### jq

`jq` is a command-line JSON processor. It is useful for reading, filtering, transforming, and generating JSON from shell scripts.

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
