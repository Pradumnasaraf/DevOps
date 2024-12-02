---
sidebar_position: 1
title: YAML Введение
---

## YAML - YAML это не язык разметки

YAML(рекурсивный акроним англ. "YAML Ain't Markup Language" — "YAML — не язык разметки"".
YAML - это человеко-читабельный язык серилизации данных. Часто используется для конфигурационных файлов и
приложений где данные сохраняются или передаются.




Eg:

```yaml
name: Pradumana
age: 21
address:
  city: Bangalore
  state: Karnataka
  country: India
  phone:
    office: 0000000
    home:
        - 1111111
        - 2222222
```

### Свойства

- Похож на  XML и JSON.
- Мы не можем добавить комманды.
- Строгий синтаксис - Отступы(Indentation)
- Человеко читабельный язык серилизации данных.

### Синтаксис

#### Пара - ключ, значение

```yaml
Name: "Pradumna Saraf"
1: "This a list"
```

#### Список

```yaml
- apple
- mango
- Banana
- banana
```
или

```yaml
cities: [new delhi, patna,gujrat]
```

#### Строки и переменнные

```yaml
name: Pradumna Saraf
fruit: "Mango"
job: 'Advocate'
age: 65
marks: 10.33
booleanValue: No, N, false, False, FALSE
```

#### Многострочный текст

```yaml
Address: |
  01
  Delhi
  India
```
Отнострок в многострочном тексте.

```yaml
message: >
  This all
  will be in a single
  line

```

#### Вложеннные типы

```yaml
names: Pradumna
role:
  age: 22
  job: student
```  

#### Вложенные последовательности

```yaml
-
  - mango
  - apple
  - banana
-
  - marks
  - roll
```

#### Специализированные типы данных

```yaml

# Integer (Целочисленный)
Zero: !!int 0
positiveNumber: !!int 45
negativeNumber: !!int -45
hexa: !!int 0x45

# Float (Дробные, с плавающей запятой)
mark: !!float 56.55
infinity: !!float .inf
not a num: .nan
itNot: !!bool false

# String (Строки)
string: !!str "hello"

# Null (Нулевой тип)
surname: !!null #null or NULL ~
~: this a null key

# Exponential Numbers(Экспонента)
myNum: 6.22ES56

# Dates and time (Дата и время)
date: !!timestamp 2002-01-02
no Time zone: 2012-12-15T02:59:43
India Time: 2012-12-15T02:59:43 +5:30
```

### Использование

- Применяются в Kubernetes, Ansible, Docker, и так далее.
- Используется для хранения данных типа ключ, значение.
- CI/CD инструменты такие как GitLab GitHub Actions, CircleCI, используют YAML для создания процесса(трубопровода)

### Что дальше?

- [Список обучающих ресурсов](./learning-resources.md) - Узнать больше о YAML на этих ресурсах.
- [Инструменты](./tools.md) - Узнать больше об инструментах которые используют YAML.
