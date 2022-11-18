# Key Value pair
Name: "Pradumna Saraf"
1: "This a list"

---

#lists
- apple
- mango
- Banana
- banana

---
# List of cities
cities:
  - new delhi
  - patna
  - gujrat 

---

cities: [new delhi, patna,gujrat]

---

{mango: "Yellow fruit", age: 56}
...

# 3 ways to represnt string
name: Pradumna Saraf
fruit: "Apple"
job: 'swe'

age: 45
marks: 10.332
booleanValue: No # n, N, fasle, Fasle, FALSE 

##### Specify the data type

# Integer
Zero: !!int 0
postiveNumber: !!int 45
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
surnmae: !!null #null or NULL ~ 
~: this a null key

# Exponential Numbers
myNum: 6.22ES56

# Dates and time
date: !!timestamp 2002-01-02
no Time zone: 2012-12-15T02:59:43
India Time: 2012-12-15T02:59:43 +5:30

---

##### Multiline

Address: |
  01
  Patna, India
  800008

---
##### Write a single lines in multiple lines

message: >
  This all
  will be in a single
  line


########### Sequence Data

students: !!seq
  - marks
  - name
  - roll

# Some of the keys of the seq will be empty - Sparse Seq

Sparse seq:
  - hey
  - how
  - 
  - null

---

########### Nested sequence

-
  - mango
  - apple
  - banana
-
  - marks
  - roll

# key: value pairs are called maps
#!!map
---
####### Nested Mapping

names: Pradumna
role:
  age: 22
  job: student
---

# Same as
names: Pradumna
role: {  age: 22,job: student}

########## Pair: keys may have duplicate !!pairs

pair Example: !!pairs # Array of hashtables
  - job: student
  - job: mentor

---
# Same as
pair Example: !!pairs [ job: student, job: mentor]

---

######### Set: will allow you to have unique value

names: !!set
  ? Pradumna
  ? Mary
  ? Liem
  ? Ana
  ? Eddie
  ? Sara
  ? Allie
  ? Alyse


########## Dictionary !!omap

People: !!omap
  - Pradumna:
      name: Pradumna
      age: 78
      height: 1446
  - Rahul:
      name: Rahul
      age: 20

############# Reusing some properties using anchors @base
likings of Person: &likes
  fav fruits: mango
  dislike: grapes

Person1:
  name: Pradumna Saraf
  <<: *likes

Person2:
  name: Pradumna Saraf
  <<: *likes
  dislike: berries #Override
  
  Schools:
- name: dps
  principle: someone
  students:
    - rno: 12
    - name: Pradumna
    - marks: 988




  
  


























