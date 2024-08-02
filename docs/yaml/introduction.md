---
sidebar_position: 1
title: YAML Introduction
---

## YAML - YAML Ain't Markup Language

YAML is a human-readable data serialization language. It is commonly used for configuration files and in applications where data is being stored or transmitted.

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

### Properties

- Similar to XML & JSON.
- We can't add commands.
- Strict syntax - (Indentation)
- Human readable data serialization language.

### Syntax

#### Key Value pair

```yaml
Name: "Pradumna Saraf"
1: "This a list"
```

#### List

```yaml
- apple
- mango
- Banana
- banana
```
or

```yaml
cities: [new delhi, patna,gujrat]
```

#### String and Variables

```yaml
name: Pradumna Saraf
fruit: "Mango"
job: 'Advocate'
age: 65
marks: 10.33
booleanValue: No, N, false, False, FALSE 
```

#### Multiline String

```yaml
Address: |
  01
  Delhi
  India
```
Single line in multiple line.

```yaml
message: >
  This all
  will be in a single
  line

```

#### Nested Mapping

```yaml
names: Pradumna
role:
  age: 22
  job: student
```  

#### Nested Sequence

```yaml
-
  - mango
  - apple
  - banana
-
  - marks
  - roll
```

#### Specify the data type

```yaml

# Integer
Zero: !!int 0
positiveNumber: !!int 45
negativeNumber: !!int -45
hexa: !!int 0x45

# Float
mark: !!float 56.55
infinity: !!float .inf
not a num: .nan
itNot: !!bool false

# String
string: !!str "hello"

# Null
surname: !!null #null or NULL ~ 
~: this a null key

# Exponential Numbers
myNum: 6.22ES56

# Dates and time
date: !!timestamp 2002-01-02
no Time zone: 2012-12-15T02:59:43
India Time: 2012-12-15T02:59:43 +5:30
```

### Usage

- Used in Kubernetes, Ansible, Docker, etc.
- Used to store data in key-value pairs.
- CI/CD tools like GitHub Actions, CircleCI, use YAML to create workflows.

### What's next?

- [Learning Resources](./learning-resources.md) - Learn more about YAML with these resources.
- [Tools](./tools.md) - Learn about the tools that you can use with YAML.