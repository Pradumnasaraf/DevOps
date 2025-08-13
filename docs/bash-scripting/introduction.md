---
sidebar_position: 1
title: Введение в Bash скриптинг
description: Глубокое погружение в Bash скриптинг, язык скриптинга для Unix оболочки.
tags: ["Bash", "Scripting", "Unix", "Shell"]
keywords: ["Bash", "Scripting", "Unix", "Shell"]
slug: "/bash"
---

Мы начинаем с создания файла с расширением `.sh`. Например, `script.sh`. Затем мы пишем в нем скрипт. Например:

Базовый скрипт

```bash
#!/bin/bash

echo "Hello World"
```    

Мы можем запустить его с помощью `bash script.sh` или `./script.sh`. Но второй будет работать только если скрипт исполняемый (разрешение на выполнение). Мы можем сделать его исполняемым с помощью `chmod +x script.sh`. Теперь мы можем запустить его с помощью `./script.sh`.

### Shebang

Первая строка bash скрипта называется shebang. Она говорит системе, какой интерпретатор использовать для запуска скрипта. Shebang для bash - это `#!/bin/bash`. Shebang для python - это `#!/usr/bin/env python`. Это варьируется от языка к языку.

### Переменные

Это заполнитель для значения. Как и в любом другом языке программирования. Мы используем $ для доступа к переменной. Например: `$NAME`. Мы также можем использовать `${NAME}`. Фигурные скобки гарантируют, что переменная не будет принята за команду.

```bash
#!/bin/bash

# Объявление переменной
NAME="John"

# Использование переменной

echo "My name is $NAME"
echo "My name is ${NAME}"
```

ПРИМЕЧАНИЕ: Мы можем создавать переменные с помощью `NAME="John"` через CLI; мы не можем использовать это в скрипте, потому что она не экспортирована. Мы можем экспортировать ее с помощью `export NAME="John"`. Теперь мы можем использовать ее в скрипте.

Но здесь есть еще одна загвоздка. Если мы выйдем из терминала и откроем новый, переменная исчезнет. Чтобы сделать ее постоянной, мы можем добавить ее в файл `.bashrc`. Это скрытый файл в домашней директории. Мы можем открыть его с помощью `vi ~/.bashrc` или любого другого редактора. Мы можем добавить переменную в файл. Например: `export NAME="John"`.

### Пользовательский ввод

Мы можем принимать ввод от пользователя с помощью команды `read`.

```bash
#!/bin/bash

echo "Enter your name: "
read NAME
echo "Hello $NAME, nice to meet you!"
```

### Аргументы

Мы можем передавать аргументы в скрипт. Аргументы хранятся в `$1`, `$2`, `$3` и так далее. `$0` - это имя скрипта.

Например: `bash script.sh arg1 arg2`

```bash
#!/bin/bash

echo "First Argument: $1"
```

### Арифметические операции

Мы можем выполнять арифметические операции в bash. Мы используем `(( ))` для выполнения арифметических операций.

```bash
#!/bin/bash

echo $(( 5 + 5 ))
```

#### Арифметические операторы

- `+` - Сложение
- `-` - Вычитание
- `*` - Умножение
- `/` - Деление
- `%` - Модуль

### Условные операторы

Мы можем использовать оператор `if` для проверки условия. Синтаксис:

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
