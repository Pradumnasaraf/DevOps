---
sidebar_position: 1
title: WebAssembly Introduction
description: WebAssembly (Wasm) is a binary instruction format that allows you to write code in languages like C, C++, and Rust, and run it in the browser at near-native speed.
tags: ["WebAssembly", "Wasm", "JavaScript", "Programming"]
keywords: ["WebAssembly", "Wasm", "JavaScript", "Programming"]
slug: /webassembly
---

In simple terms WebAssembly (Wasm) is a binary instruction format. With WebAssembly, you can write code in languages like C, C++, and Rust, and run it in the browser at near-native speed. It is design to run alongside JavaScript, allowing both to work together.

## Why WebAssembly?

As JavaScript is the only language that runs in the browser. It comes with its own limitations and challenges like:

- Interpreted language: it's translated on-the-fly by the browser and different browsers have different JavaScript engines. Like V8 in Chrome, SpiderMonkey in Firefox, etc. And it's still generally slower than compiled languages.
- Performance: Applications which require high performance like games, video editing, real-time graphics intensive computations, etc. often encounter performance issues.
- Memory limitations: JavaScript has memory limitations. The garbage collector can introduce a pause in the application execution and doesn't allow direct memory access.
-  Concurrency: JavaScript is single-threaded and doesn't support multi-threading. It can't take advantage of multi-core processors. Yes, it can efficiently handle asynchronous operations using callbacks, promises, async/await, etc. But can't truly run multiple tasks in parallel. Becomes more pain point with real-time simulations or data intensive applications.

WebAssembly was designed to address these limitations. With it's low-level binary format, it can be executed at near-native speed. Additional benefits include:
- Language agnostic: It's not tied to any specific language. You can write code in any language that compiles to WebAssembly.
- Secure: It runs in a sandboxed environment and doesn't have direct access to the host system.
- Portable: We can WASM application be it smartphones, desktops, or servers.
- Concurrency and Parallelism: WebAssembly can take advantage of multi-core processors and run multiple tasks in parallel.
- Compactness and Speed: WebAssembly is designed to be fast to decode and execute. It's compact and can be downloaded quickly. Web application which uses WebAssembly can load faster and run more efficiently. The speed is achieved by the binary format which is smaller and and download faster and binary decoding is faster than parsing text.

## Use Cases

- **Machine Learning and AI**: With this app can make real-time predictions, image recognition, etc without sending data to the server and latency in processing.
- **Multimedia Applications**: Video editing, audio processing, real-time graphics, etc. Rendering and video/image filters need high performance, with WASM it can be smoother and faster.
- **Real-time Data Visualization**: Real-time simulations, data processing, etc. can be done more efficiently. Sector like finance, logistics, etc can benefit from this.
- **3D Games and Simulations**: Games which requires intensive graphics operations can benefit from WASM. 3D games and simulations can run more smoothly and efficiently.

### Real-world examples

- [TensorFlow.js](https://www.tensorflow.org/js): TensorFlow.js integrates with WebAssembly to enhance computation performance. Making real-time browser predictions possible.
- [Figma](https://www.figma.com/): Intricate design operations and rendering processes are fluid and responsive, enhancing user experience.
- [AutoDesk](https://www.autodesk.com/): AutoDesk uses WebAssembly to run complex CAD operations in the browser.
- [Unity](https://unity.com/): Unity supports WASM enabling developers to create high-performance 3D browser games.
- [Shopify](https://www.shopify.com/): Shopify uses WASM to optimize its theme editor, allowing merchants to design and preview their online stores in real-time.

## WebAssembly vs JavaScript

WebAssembly and JavaScript are not competing technologies. They are designed to work together. WebAssembly is not a replacement for JavaScript. Both have it's own strengths and weaknesses. On one hand, JavaScript is high-level, dynamic, ubiquity and ecosystem. With few line of code we can create a form or button, etc. On the other hand, WebAssembly give us flexibility to write code in any language and run it in the browser. And comes with the performance benefits.

For better perspective how they work together. Take an example of image processing web application. The UI and user interactions, like button, sliders and dropdown menus or the event handling (select a filter, adjust a slider) or feedback (display notification, progress bar, etc) can be done in JavaScript. On the other hand, heavy lifting like image processing, that need computational power can be done in WebAssembly. This way we can take advantage of both technologies. Like tha actual image processing can be done in WebAssembly ensuring that filters are applied quickly and efficiently. Or complex algo ported from existing C or C++ libraries can be complied to WASM and use to process the image data.

## Building blocks of WebAssembly

### Binary format

Unlike human readable text, the binary format is designed for machines, ensuring that it can be quickly decoded and executed. For example, the following code in C:

```c
int add(int a, int b) {
  return a + b;
}
```

When compiled to WebAssembly, might looks like this:

```
00 61 73 6d 01 00 00 00 01 07 01 60 02 7f 7f 01
7f 03 02 01 00 07 0a 01 06 61 64 64 00 00 0a 09
01 07 00 20 00 20 01 6a 0b
```

### Stack-based virtual machine

WebAssembly is a stack-based virtual machine. It means that it uses a stack to store and retrieve data. The stack is a data structure that follows the Last In First Out (LIFO) principle.

![Stack-based virtual machine](https://github.com/user-attachments/assets/86a18a37-7352-4633-8968-f542e4ca6787)

### Linear memory

WebAssembly has a linear memory model. It allows the WebAssembly module to allocate and access memory in a linear fashion. It's like a big array that can be accessed by the WebAssembly module. It starts from 0 and growing up. It can be accessed via two primary instructions: `load` and `store`. The `load` instruction reads data from memory, while the `store` instruction writes data to memory. 

The instructions are suffixed with the data type and size. Example `i32`, `i64`, `f32`, `f64`. It describes the type of data and the size of the data in memory. So, in `i32` the `i` stands for integer and `32` stands for 32-bit or 4 bytes. As 8 bits make 1 byte, 32 bits make 4 bytes. Similarly, `f32` stands for floating point number with 32-bit or 4 bytes.

For example if we want to store number 7 in the 5th slot of the memory. We have to use `i32` type. To get the byte address of the 5th slot, we have to multiply the `slot number` * `size per slot`. So, 4 * 4 = 16. As slot starts from 0, instead of multiplying by 5, we have to multiply by 4. So, we store our number 7 in the 16th byte address. 

![linear memory](https://github.com/user-attachments/assets/55cc953d-6d1a-4037-8515-1fc758d869e6)

### Modules

Like other programming languages, WebAssembly code is organized into modules. A module is a collection of functions, types, tables, memories, and globals. Each module can import and export entities. It can also define its own entities. The module is the basic unit of code in WebAssembly.

![module](https://github.com/user-attachments/assets/4c4a978b-c105-4cd1-b8e8-4ab6d07f9693)

For example, the following code defines a simple module with a single function:

```wasm
(module
  (func $add (param $a i32) (param $b i32) (result i32)
    get_local $a
    get_local $b
    i32.add)
)
```

Here, the module defines a function called `add` that takes two 32-bit integers as input and returns a 32-bit integer as output.

### Security

WASM is designed to be secure as it run in a sandboxed environment. Which means it doesn't have direct access to the host system. It can't access the file system, network, or other resources. It can only interact with the host environment through the WebAssembly System Interface (WASI). WASI provides a set of APIs that allow WebAssembly modules to interact with the host environment in a secure and controlled manner. 

If WebAssembly program tries to access a memory location that it hasn't been granted access to, the browser will halt the execution and throw an error. This ensures that WebAssembly programs can't access memory and safety of user's data and system.

![security](https://github.com/user-attachments/assets/b8741862-2bcd-4c72-b8da-c3cf9651a3ab)

### Interoperability With JavaScript

They both can work hand-in-hand. We can call WebAssembly functions from JavaScript and vice versa. 

For example we have complied a add function in WebAssembly. In JavaScript we can call it like this:

```js
const result = wasmModule.instance.exports.add(2, 3);
console.log(result); // 5
```

![interoperability](https://github.com/user-attachments/assets/15b07717-3045-4f59-8818-76d108583926)

### Text Format

Whilst WebAssembly primary format is binary, it also has a human-readable text format. This text format is called WebAssembly Text Format (WAT). It's a simple, verbose and readable representation of the binary format. It's useful for debugging and understanding the structure of the WebAssembly module.

![text format](https://github.com/user-attachments/assets/0e3c9e90-08da-4d72-9908-b0d75cc92d13)

## Modules

Like any other high level programming language, WebAssembly code is organized into modules. A module is a collection of functions, types, tables, memories, and globals. It's backbones of WebAssembly. 

### Functions

They are defined block of code that can be called from other parts of the program or even from the JavaScript code. Functions can take parameters and return values. They can be imported from other modules or exported to be used by other modules.

For example this a function that takes two parameters and returns the sum of them:

```wasm
(module
  (func $add (param $a i32) (param $b i32) (result i32)
    get_local $a // fetch value
    get_local $b
    i32.add) // Addition
)
```

Here:
- `(func $add ...)` defines a function called `add`. We use `$` to define the name of the function.
- `(param $a i32) (param $b i32)` defines two parameters of type `i32`. That is 32-bit integer.
- `(result i32)` defines the return type of the function. Same with the parameters name, we use `$` to define the name of the result.
- `get_local $a` and `get_local $b` are instructions that get the values of the parameters.
- `i32.add` is an instruction that adds the two values together.
- The function returns the result of the addition.

In WASM stack-base model, the result of the last expression in the function body is implicitly returned. So, we don't need to use `return` keyword. 

Talking about value stack it will change in this order: 

empty (locals with be a, b) -> a (locals with be b) -> a, b (locals with be empty) -> (a + b) (locals with be empty)

#### Essence of WebAssembly Modules

It not juts about function and using the functionality in our code. It about how these functions and along with other components comes to form a complete executable module. It's more like a container that holds all the necessary components for a specific functionality. 

For example for image processing module, it can have functions for applying filters, resizing images, etc. It can have memories to store the image data, tables to store the filter data, etc. It can have globals to store the image dimensions, etc.

![module](https://github.com/user-attachments/assets/a8ea0ea4-c60d-4d5e-a05c-cbc5191b8eb5)

### Beyond Functions

Apart from functions modules can have other components like:

**Global Variables**: These are variables that are accessible from anywhere in the module. They can be imported from other modules or defined within the module. They can be mutable or immutable. They can be used to store constants, configuration values, etc.

**Memory**: It's a linear array of bytes that can be accessed by the WebAssembly module. It can be used to store data that needs to be shared between functions. It can be used to store images, audio, video, etc. It can be imported from other modules or defined within the module.

**Tables**: They are arrays of elements that can be accessed by index. They can be used to store function pointers, objects, etc. They can be imported from other modules or defined within the module.

**Types**: They are used to define the types of functions, tables, memories, and globals. They can be imported from other modules or defined within the module.

![beyond functions](https://github.com/user-attachments/assets/4dec5e7a-de30-4c5f-b9fe-769d7c953747)

### Importing and Exporting

Modules can declare dependencies, pulling in necessary functions, memories, tables, and globals from other modules. This is done using the `import` keyword. For example if a module requires the current date from the JavaScript, it would import the function like:

```wasm
(module
  (import "js" "getCurrentDate" (func $currentDate (result i32)))
)
```

Conversely, to provide its functionality to outside world, a module can export specific component like functions, memories, tables, and globals using the `export` keyword. For example, to export a function that adds two numbers:

```wasm
(module
  (func $fibonacci (param $n i32) (result i32) ...)
  (export "fibonacci" (func $fibonacci))
)
```

### Security Considerations

Modules are designed to be secure. As they are encapsulated nature, they operate in a confined environment, preventing unintended access to the host system. 

For eg: If a module has a function to access to Database it can't directly access the database. It has go though specific Web APIs, ensuring data integrity and security.

![security considerations](https://github.com/user-attachments/assets/4a34f53d-77d5-41fc-92aa-fa8fcf42c290)

### Portability and Universality

WebAssembly offers cross-platform compatibility. Modules are platform-agnostic can run on any device that supports WebAssembly. For example a physics simulation module once created can be run on PC's browser, a smartphone, or even a smart TV ensuring consistent user experience, behavior, and performance.

![portability and universality](https://github.com/user-attachments/assets/186de57e-4541-4b8c-894d-3a2491d39e5a)

## Instructions and Data Types

### Why we need instructions and data types?

It's more like as the verbs of a language, dictating what actions to perform. And data types are like nouns, specifying what kind of data to work with. Together they form the building blocks of WebAssembly.

### Instructions

Like we give instructions verbally, we give instructions in WebAssembly. For example, to add two numbers, we use the `i32.add` instruction. It's a binary instruction that tells the WebAssembly engine to add two 32-bit integers.

As WebAssembly instructions are low-level the execution is fast.

Instructions - References:

- `.add`: Add two numbers
- `.sub`: Subtract two numbers
- `.mul`: Multiply two numbers
- `.div`: Divide two numbers
- `.eq`: Check if two numbers are equal
- `.lt`: Check if one number is less than another
- `.gt`: Check if one number is greater than another
- `.block`: Define a block of code
- `.loop`: Define a loop
- `.br`: Branch to a label
- `.load`: Load data from memory
- `.store`: Store data in memory

![instructions and data types](https://github.com/user-attachments/assets/253d9410-4b43-4fbd-b64a-d78e6ba32437)

### Data Types

Data types define the kind of data we can work with. It help set the rule and how the data is stored and processed. For example in a weather app the temperature in once city is a whole number like 25°C, while in another city it's a decimal number like 25.5°C. So, we represent using `i32` (32-bit integer) and `f32` (32-bit floating point number) respectively.

Data types are very essential to make the code is predictable by defining the kind of data and how it's used, we avoid errors and ensure efficient use of memory. For example imagine pouring water into a glass. Each glass (data type) can hold a specific amount and shape of water (data). If we try to pour too much water or wrong type of liquid, it either won't fit or won't function as intended. Same goes with data types, ensuring data fits well and works as intended.

Data Types - References:

- `i32`: 32-bit integer
- `i64`: 64-bit integer
- `f32`: 32-bit floating point number
- `f64`: 64-bit floating point number
- `v128`: 128-bit vector
- `funcref`: Function reference
- `externref`: External reference
- `anyref`: Any reference

![instructions and data types](https://github.com/user-attachments/assets/ff3cd2c3-9264-4ef1-8531-344cbe592756)

### Symbiosis of Instructions and Data Types

Instructions and data types work together to perform specific tasks. For example, to add two numbers, we use the `i32.add` instruction. It's a binary instruction that tells the WebAssembly engine to add two 32-bit integers. Here, `i32` is the data type that specifies the kind of data we are working with, and `add` is the instruction that tells the engine what action to perform.

## Memory and Tables

### Why we need Memory and Tables?

Think it more like a building a city (which is a program) the buildings are the data and they plot of land that's the memory. The roads are the tables. And the city services are the functions and those functions needs a directory to be organized for that we have tables. Without memory and tables, the city (program) can't function properly and can be chaotic.

![memory and tables](https://github.com/user-attachments/assets/f8071a47-0a88-4919-a796-ae7e9d7ec816)

### Memory

When we so something like add two number say 3 and 5, these number are stored in what's called WebAssembly memory. After the addition it stores the result back in the memory, that is the result 8. And by doing this it make isolated from the rest of the computer's memory making it secure and if something goes wrong it won't affect the rest of the computer. 

Memory can grow as needed. It can be accessed by the WebAssembly module as well as the JavaScript code. 

![memory](https://github.com/user-attachments/assets/772d45c9-639a-4075-8ced-d6cb8d0762b8)

### Tables

They are more like directories and lookup tables. They don't store actual data but store references to functions. For example, if we have a function that calculates the square of a number, we can store a reference to that function in a table. When we need to calculate the square of a number, we can look up the function in the table and call it.

Not only it stores references to functions, it can store references global variables, objects, 

![tables](https://github.com/user-attachments/assets/e5578a4b-5cc2-46b8-843b-d2b315bfcf5b)


## Tools and Ecosystem

### Why we need Tools and Ecosystem?

For anything having right sets of tools can make the difference between a smooth process and a bumpy ride. For WASM, the tool and the surrounding ecosystem simplify the development, testing and deployment process. So that dev can focus on the core functionality of the application rather than worrying about the underlying complexities.

### CNCF - WASM Landscape

Cloud Native Computing Foundation (CNCF) has taken the lead and introduced the WebAssembly Landscape. It provides a bird's eye view of tools, libraries, and projects that are part of the WebAssembly worlds. Some core areas:

**Programming Languages**: The real beauty of WASM chimed in when it comes to programming languages. It's not tied to any specific language. You can write code in any language that compiles to WebAssembly. Some popular languages are C, C++, Rust, Go, etc. On the other hand we managed languages like Kotlin, Dart, etc, rely on WASM by compiling their interpreters into it. Also new languages like Moonbit, and Grain, optimized for WASM.

**Runtimes**: Onces are code complied into WASM bytecode, it needs a runtime to execute it. Some popular runtimes are WasmEdge, WasmTime, WebAssembly Micro Runtime (WAMR), etc. These runtimes provide the necessary environment to execute the WASM code. They all offer safety, speed and portability to the code.

**Application Frameworks**: Dev can lean on the libraries and frameworks. WasmEdge, for instance, supports advance POSIX APIs, allowing Rust and JS app framework to run seamlessly.

Tailored Framework: A framework for building WebAssembly microservices, and WasmCloud, which simplifies the development for distributed applications.

**Edge/Bare Metal**: As cross-platform compatibility is one of the key features of WASM, that means it can run across different operating systems and architectures, including in edge and IOT computing. 

**AI inference**: As AI is becoming more prevalent, and AI becoming staple in data centers. Runtimes like WasmTime and WasmEdge are integrating with native AI/ML libraries like TensorFlow, PyTorch, etc, to facilitate AI inference in the language like Rust.

**Embedded Functions**: WASM can also be executed as embedded functions in various software products. For example, database like LibSql and OpenGauss have integrated WASM to run user-defined functions.

**Tooling** Tools like cargo, LLVM and Binaryen are essential for compiling code to WASM.

**Deployment**: We can easily take benefit of the cloud-native ecosystem. And deploy on Docker, Kubernetes, etc.

**Debugging and Observability**: Debugging and observability are essential for any application in production. Tools like WASI logging are making it easier to debug and monitor WASM applications.

**Artifacts** It's really important for software supply chain. Repositories like DockerHub and Harbor are stepping to store, and track WASM packages.

## Working with WebAssembly

As we know WebAssembly comes in Binary and Text format. The binary format with `.wasm` extension. It serves as a universal compilation target for high-level languages. The text format with `.wat` extension. It's a human-readable representation of the binary format. It's useful for debugging and understanding the structure of the WebAssembly module.

### Binary Format

To compile a C++ code to WebAssembly binary format, we can use the Emscripten compiler. It's a toolchain for compiling to WebAssembly. It can compile C and C++ code to WebAssembly. It can also compile other languages like Rust, Go, etc.

For heads up we compile our C++ code to WebAssembly like this using Emscripten:

```bash
emcc hello.cpp -o hello.js
```

This above command compiles the `hello.cpp` file to WebAssembly and generates the `hello.js` along with the `hello.wasm` file.

We will have a deeper dive into the compilation process in the next section.

#### Structure of Wasm Binary

The Binary file is organized into modules. Each module starts with a specific structure that includes a magic number and a version number. The magic number is `00 61 73 6d` which spells out `wasm` in ASCII. The version number is `01 00 00 00` which indicates the version of the WebAssembly binary format.

And each module can have sections. Each section has a specific purpose. For example, the `type` section defines the types of functions, the `function` section defines the functions, the `memory` section defines the memory, etc. So, Type Section (Section ID: 1) defines the function signature, Function Section (Section ID: 3) defines the function bodies, Memory Section (Section ID: 5) defines the memory, etc.

All sections and their purpose:

- Section 0: Custom Section: It's used to define custom data.
- Section 1: Type Section: It used to define the function signature.
- Section 2: Import Section: It used to import functions, memories, tables, and globals from other modules.
- Section 3: Function Section: It used to define the function bodies.
- Section 4: Table Section: It used to define the tables. It defines the table types, including the element type and the limits.
- Section 5: Memory Section: It used to define the memory.
- Section 6: Global Section: It used to define the global variables.
- Section 7: Export Section: It used to export functions, memories, tables, and globals to other modules.

Let's take an example of Table Section:

```wasm
04  ;; Section ID: Table Section
06  ;; Section Length: 8 bytes

01  ;; Number of tables: 1
70 00 01  ;; Table type: Element type (anyfunc), Initial size (1)
```

So, here 
- `04` is the Section ID for Table Section for the Memory Section.
- `06` is the Section Length in bytes. Indicating that Memory Section is 6 bytes long.
- `00 01` is the limit of the table. It's a 32-bit integer. The first byte `00` indicates the minimum size of the table, and the second byte `01` indicates the maximum size of the table. Here, the table has a minimum size of 0 and a maximum size of 1.
  
Another example of Global Section

```wasm
06  ;; Section ID: Global Section
19  ;; Section Length: 25 bytes.

03 ;; Number of Global Variables
7F 01 41 0B  ;; Global variable 1: Type (i32), Mutable, Initialization(i32.const 11)
```

So, here
- `06` is the Section ID for Global Section.
- `19` is the Section Length in bytes. Indicating that Global Section is 25 bytes long.
- `03` is the Number of Global Variables.
- `7F` is the Global variable type. Here, it's `i32`.
- `01` is the mutability of the global variable. Here, it's mutable.
- `41 0B` is the initialization value of the global variable. Here, it's `i32.const 11`.

It's not necessary to understand the binary format in detail. But it's good to know how it's structured and how it works. Knowing it helps us achieve performance optimization, debugging, security auditing, adv. features, etc. Otherwise dev may work at high level of abstraction using languages like Rust, C+++, or JavaScript that compiles to WebAssembly.

### Text Format

As we know WASM binary format is not human readable. To make it human readable, we have WebAssembly Text Format (WAT). It's a simple, verbose, and readable representation of the binary format. It's useful for debugging and understanding the structure of the WebAssembly module.

For example, the following binary code:

Example:

```Wasm
(module
;; Define a function working
  (func $add (param $a i32) (param $b i32) (result i32) 
    get_local $a
    get_local $b
    i32.add)
  (export "add" (func $add))
)
```

So, here:

- `(module...)` : Module's definition. Everything inside this will help construct the module.
- `(func $add (param $a i32) (param $b i32) (result i32)..)` : Function's definition. Here, we define a function called `add`.
- `(param $a i32) (param $b i32)` : Parameters of the function. Here, we define two parameters of type `i32`.
- `(result i32)` : Return type of the function. Here, we define the return type of the function as `i32`.
- `get_local $a` and `get_local $b` : Instructions that get the values of the parameters.
- `i32.add` : Instruction that adds the two values together.
- `(export "add" (func $add))` : Export the function. Here, we export the `add` function.
- `i32.add` : Instruction that adds the two values together.

Seeing from a different pair of eyes, WASM a language. It has Grammar, Whitespace and structure. So, in above parenthesis defining scope and structure. While `func`, `param` and `result` specify the code's functionality. White space and indentation are used to make the code more readable.

Let's break down the above code and see how it works behind the scenes:

- `get_local $a` gets the value and pushes it onto the stack.
- `get_local $b` gets the value and pushes it onto the stack.
- `i32.add` pops the two values from the stack, adds them together, and pushes the result back onto the stack.
- The function returns the result.

Comments are also supported in the WebAssembly Text Format. They are enclosed in `;;` and are ignored by the WebAssembly engine. They are used to add notes, explanations, and documentation to the code.


#### Lexical Structure

If we look at Lexical Structure of WAT, it's more like a grammar of the language. It defines how the code is written and how it's structured. So, in above example we have identifiers like `add`, `a`, `b`, etc. We have keywords like `module`, `func`, `param`, `result`, etc. 

#### Types

In WAT, we have different types like `i32`, `i64`, `f32`, `f64`, etc. They define the kind of data we are working with. For example, `i32` is a 32-bit integer, `i64` is a 64-bit integer, `f32` is a 32-bit floating point number, and `f64` is a 64-bit floating point number. It also supports vector types like `v128`, reference types like `funcref`, `externref`, `anyref`, etc.

An example of vector type:

```wasm
(module
;; .... Previous definition ...
  (func $add (param $a v128) (param $b v128) (result v128)
    get_local $a
    get_local $b
    v128.add)

  ;; .... other functions ...

  (export "add" (func $add))

  ;; .... other exports ...
)
```

An example of reference type:

```wasm
(module
;; .... Previous definition ...
  (func $add (param $a funcref) (param $b externref) (result anyref)
    get_local $a
    get_local $b
    anyref.add)

  ;; .... other functions ...

  (export "add" (func $add))

  ;; .... other exports ...
)
```

#### Instructions

They define the actions that the WebAssembly engine can perform. For example, `i32.add` is an instruction that adds two 32-bit integers. 

##### if-else block

```wasm
(module
;; .... Previous definition ...
  (func $add (param $a i32) (param $b i32) (result i32)
    get_local $a
    f32.const 0
    f32.lt ;; Check if $a is less than 0
    if (result i32)
      f32.const 0 ;; If $a is less than 0, return 0
    else
      get_local $a
      get_local $b
      f32.add ;; If $a is greater than 0, add $a and $b
    end)
  ;; .... other functions ...
)
```

Here `if` and `else` are instructions that define a conditional block. The Bloc can be nested within each other for complex conditions.


##### Loop block

```wasm
(module
;; .... Previous definition ...
  (func $add (param $a i32) (param $b i32) (result i32)
    get_local $a
    get_local $b
    loop (result i32)
      i32.add
      get_local $a
      i32.const 1
      i32.sub
      tee_local $a
      i32.const 0
      i32.eq
      br_if 0
    end)
  ;; .... other functions ...
)
```

Here `loop` is an instruction that defines a loop block. The loop block can have a condition that determines when to exit the loop. The `br_if` instruction is used to break out of the loop if the condition is met

So here they way it's working is:

- `loop` starts the loop block.
- `i32.add` adds the two values together.
- `get_local $a` gets the value of `$a`.
- `i32.const 1` pushes the value `1` onto the stack.
- `i32.sub` subtracts `1` from `$a`.
- `tee_local $a` duplicates the value of `$a` and stores it back in `$a`.
- `i32.const 0` pushes the value `0` onto the stack.
- `i32.eq` checks if `$a` is equal to `0`.
- `br_if 0` breaks out of the loop if `$a` is equal to `0`.
- `end` ends the loop block.

## WASI - WebAssembly System Interface

WASI (WebAssembly System Interface) is a system interface for WebAssembly. It provides a set of APIs that allow WebAssembly modules to interact with the host environment in a secure and controlled manner. It allows WebAssembly modules to access system resources like files, network, and environment variables.

![WASI](https://github.com/user-attachments/assets/e701c4dc-bcaf-4f74-b4f3-a7d4940a4535)

Without WASI WASM modules can't access the file system, network, or other resources. It's isn't a flaw, it's by design. A couple of more important points to understand WASI's design:

- It ensures whether a person is using the application in Japan, France, or the US, on a variety of devices, the application will work the same way.
- It only unlock specific permissions. For example, if a WebAssembly module needs to read a file, it can only access the file system and not the network or other resources.
- WASI is giving dynamism to the static web pages.
- Incorporating the specific functionalities without burdening the application with unnecessary features.
  
![WASI](https://github.com/user-attachments/assets/a271f1fb-32f5-4356-997d-6297037d2ac6)

### WASI Functions

WASI provides a set of functions that allow WebAssembly modules to interact with the host environment. These functions are divided into different categories like:

![WASI Functions](https://github.com/user-attachments/assets/6e324ae0-e74b-4b53-a72b-9932be008004)

- **File Operations**: Functions like `fd_read`, `fd_write`, `fd_seek`, etc. allow WebAssembly modules to read and write files.
- **Network Activities**: Functions like `sock_send`, `sock_recv`, `sock_shutdown`, etc. allow WebAssembly modules to send and receive data over the network.
- **System Information**: Functions like `args_sizes_get`, `args_get`, `environ_sizes_get`, `environ_get`, etc. allow WebAssembly modules to access system information like command-line arguments, environment variables, etc.
- **Time and Clock**: Functions like `clock_time_get`, `clock_res_get`, etc. allow WebAssembly modules to access time and clock information.

An example of `fd_read` function:

```wasm
(module
  ;; Import the fd_read function from the WASI module
  (import "wasi_snapshot_preview1" "fd_read" (func $fd_read (param i32 i32 i32 i32) (result i32)))

  ;; Memory declaration and other necessary code...

  ;; A function to read data from a file
  (func $read_file
    ;; Assuming file descriptor is stored in memory at offset 0
    (i32.const 0) ;; File descriptor
    (i32.const data_offset) ;; pointer to memory location where data will be stored
    (i32.const data_length) ;; length of the data to read
    (i32.const result_offset) ;; pointer to memory location where the result will be stored
    (call $fd_read) ;; Call the fd_read function

    ;; Handle the read data as necessary
  )
  ;; Export our read_file function
  (export "read_file" (func $read_file))
)
```

Here, we import the `fd_read` function from the WASI module. We then define a function called `read_file` that reads data from a file using the `fd_read` function. We pass the file descriptor, memory location to store the data, data length, and memory location to store the result to the `fd_read` function.

## Creating a WebAssembly Module

As we know modules are the core of WASM. We will se how we can convert code written in various languages like C, C++, Rust, etc. to WebAssembly module.


### Compiling C++/C Code to WebAssembly

To compile C++/C code to WebAssembly, we can use the Emscripten compiler. It's a toolchain for compiling to WebAssembly. It can compile C and C++ code to WebAssembly. It can also compile other languages like Rust, Go, etc.

For example, to compile a C code to WebAssembly, we can use the following command:

```c
// hello.c
#include <stdio.h>

int add(int a, int b) {
  return a + b;
}
```

```bash
emcc -03 -s WASM=1 -o hello.wasm hello.c
```

Here:

- `emcc` is the Emscripten compiler.
- `-03` is the optimization level. It's the highest level of optimization. The range is from `-O0` to `-O3`. It optimizes the code for size and speed.
- `-s WASM=1` tells the compiler to generate WebAssembly output. Here `s` stands for setting.
- `-o hello.wasm` specifies the output file name.
- `hello.c` is the input C file.

Together, these commands compile the `hello.c` file to WebAssembly and generate the `hello.wasm` file. The `hello.wasm` is the WASM binary file or we can call it as a WASM module. This a compact version

### WebAssembly in the Browser

To bring the WASM module to life in the browser, we use JavaScript. JavaScript fetches the WASM module using the `fetch` API. Once the module is fetched, it's compiled and instantiated using the `WebAssembly.instantiateStreaming` method. The module is then executed using the exported functions. Once the object is created, we can call the exported functions like any other JavaScript function.

```js
// index.js
fetch('hello.wasm')
  .then(response => response.arrayBuffer())
  .then(bytes => WebAssembly.instantiate(bytes))
  .then(results => {
    const instance = results.instance;
    const add = instance.exports.add;
    console.log(add(2, 3)); // 5
  });
```

Some older browsers may not support the `WebAssembly.instantiateStreaming` method. In that case, we can use the `WebAssembly.instantiate` method. So, this was just an example, we have keep the nuances in mind while building the application to cater to the different browsers.

#### Linear Memory

Memory is like a vast storage room where data is kept. When we talk about WASM and JavaScript sharing memory, it's more like a shared storage room. Where each section having its own shelves to store data. For WebAssembly, memory is a linear array of bytes. It's a continuous block of memory, much like a long shelf, where data is stored sequentially. This is the primary way for WebAssembly to store and access data.

Now the way it works is, when JavaScript want to send data to WASM, it stores the data on this shared shelf. And when WASM wants to send data back to JavaScript, it stores the data on this shared shelf. This way both JavaScript and WASM can access and exchange data. For example, JavaScript want to convert 25°C to Fahrenheit, it stores the value 25 on the shelf. WASM then reads the value, converts it to Fahrenheit, and stores the result back on the shelf. JavaScript then reads the result and displays it to the user.

Benefits of Linear Memory:

- It's fast and efficient. As data is stored sequentially, it's easy to access and process.
- Eliminates the need of complex data transfer mechanisms.
- Ensures data integrity and security. As data is stored in a controlled environment, it's safe from unauthorized access.

#### WASM Browser Runtime

So, every modern browser comes equipped with a JavaScript engine, like V8 in Chrome or SpiderMonkey in Firefox. These engines are responsible for compiling and executing JavaScript code to machine code, and run it. 

When it comes to WASM, as it is a different technology. The WASM doesn't need an separate engine. It's the JavaScript engine that compiles and executes the WASM code. JavaScript engines have been extended to handle WASM code. JavaScript engine takes the charge, but with a different set of tools tailored for WASM. Most of browser support WASM, like Chrome, Firefox, Safari, Edge, etc.

#### WASM JavaScript API

The WebAssembly JavaScript API provides a way to interact with WebAssembly modules from JavaScript. It provides methods to compile, instantiate, and execute WebAssembly modules. It also provides methods to access and manipulate linear memory, tables, and globals.

Some of the key methods of the WebAssembly JavaScript API are:

- `WebAssembly.instantiateStreaming`: Compiles and instantiates a WebAssembly module from a streamed response.
- `WebAssembly.instantiate`: Compiles and instantiates a WebAssembly module from an array buffer.
- `WebAssembly.Module`: Represents a WebAssembly module. It creates a new WebAssembly module from a binary buffer.
- `WebAssembly.Instance`: Represents an instance of a WebAssembly module. It creates a new instance of a WebAssembly module.
- `WebAssembly.Memory`: Represents a linear memory instance. It creates a new linear memory instance. It can be used to store and access data.
- `WebAssembly.Table`: Represents a table instance. It creates a new table instance. It can be used to store function pointers, objects, etc.
- `WebAssembly.CompileError`: Represents a compile error. It's thrown when there is a compile error in the WebAssembly module.
- `WebAssembly.LinkError`: Represents a link error. It's thrown when there is a link error in the WebAssembly module.
- `WebAssembly.RuntimeError`: Represents a runtime error. It's thrown when there is a runtime error in the WebAssembly module.

WebAssembly.Memory - Example:

```js
const memory = new WebAssembly.Memory({ initial: 10, maximum: 100 });
```

Above example creates a new linear memory instance with an initial size of 256 pages and a maximum size of 256 pages. It can be used to store and access data.

For reference:
1 KiB = 1024 bytes
1 WASM Memory Page = 64 KiB or 65536 bytes

So, in above example, the initial size of the memory is 10 pages, that comes to 10 * 64 KiB = 640 KiB. The maximum size of the memory is 100 pages, that comes to 100 * 64 KiB = 6400 KiB or 6.25 MiB (6400 / 1024).

Memory Pages are the unit of memory allocation in WebAssembly. They are fixed-size blocks of memory that can be allocated and deallocated. They are used to store and access data in WebAssembly modules.

### Running WASM in Browser

Before running any WASM module in the browser, we need to compile the code to WebAssembly. We can use the `Emscripten` compiler to compile C and C++ code to WebAssembly. Once the code is compiled, we can load the WebAssembly module in the browser using JavaScript. We have already seen how to compile C code to WebAssembly.

We will taken an example of a C program as see step by step how to compile and run it in the browser.

Create a C program called `hello.c`:

```c
hello.c
#include <stdio.h>

int main() {
    printf("Hello, WebAssembly!\n");
    return 0;
}
``` 

Compile the `hello.c` file to WebAssembly using the `Emscripten` compiler:

```bash
emcc hello.c -o hello.html
```

It will generate the `hello.html` file along with the `hello.wasm` and `hello.js` files. We can open the `hello.html` file in the browser to run the WebAssembly module.

If you look closely at the `hello.html` file, you will see that it loads the `hello.js` file, which in turn loads the `hello.wasm` file. The `hello.js` file contains the JavaScript code to load and run the WebAssembly module. In this way everything is connected and the WASM module is executed in the browser.
