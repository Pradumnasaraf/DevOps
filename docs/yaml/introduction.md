---
sidebar_position: 1
title: YAML Introduction
description: Learn about YAML (YAML Ain't Markup Language) and its properties.
tags: ["YAML", "Data Serialization", "Configuration"]
slug: "/yaml"
---

## YAML - YAML Ain't Markup Language

YAML is a human-readable data serialization language. It is commonly used for configuration files and for describing structured data in tools such as Kubernetes, GitHub Actions, and Docker Compose.

Example:

```yaml
name: Pradumna
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

- Similar in purpose to JSON and XML.
- Designed for data, not executable commands.
- Sensitive to indentation and spacing.
- Easy for humans to read when formatted clearly.

### Syntax

#### Key-value pair

```yaml
Name: "Pradumna Saraf"
1: "This is a value"
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
cities: [new delhi, patna, gujarat]
```

#### String and Variables

```yaml
name: Pradumna Saraf
fruit: "Mango"
job: 'Advocate'
age: 65
marks: 10.33
booleanValue: false
```

#### Multiline String

```yaml
Address: |
  01
  Delhi
  India
```
Folded style joins multiple lines into a single string.

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
myNum: 6.22E56

# Dates and time
date: !!timestamp 2002-01-02
no Time zone: 2012-12-15T02:59:43
India Time: 2012-12-15T02:59:43 +5:30
```

### Usage

- Used in Kubernetes, Ansible, Docker, and many other tools.
- Useful for storing data in key-value form.
- Common in CI/CD systems like GitHub Actions and CircleCI.

### What's next?

- [Learning Resources](./learning-resources.md) - Learn more about YAML with these resources.
- [Tools](./tools.md) - Learn about the tools that you can use with YAML.
