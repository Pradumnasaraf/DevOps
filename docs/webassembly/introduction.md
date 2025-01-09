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

## Closer look at WASM Compilation

If we first look how JavaScript is compiled. It's uses JIT (Just-In-Time) compilation. It compiles the code on the fly, as it's executed. Previously it use to be interpreted line by line. But with JIT compilation, it compiles the code into machine code, and then executes it. This makes the code run faster and more efficiently.

It was good, but WASM made it even better. WASM two-staged compilation process. Ahead-of-Time (AOT) compilation and Just-In-Time (JIT) compilation. So, the code is first compiled to WASM binary format using AOT compilation. One the WASM module is loaded, it's compiled to machine code using JIT compilation. This makes the code run even faster and more efficiently.

### Compilation Process

WebAssembly takes 2-stage compilation process. 

The first stage is Ahead-of-Time (AOT) compilation. In this stage, the code is compiled to WebAssembly binary format. This is done using a WebAssembly compiler like Emscripten. The WebAssembly binary format is a compact and efficient representation of the code. It's optimized for size and speed.

Once the WASM Binary is loaded, the second stage is Just-In-Time (JIT) compilation. In this stage, the WASM module is compiled to machine code. This is done by the JavaScript engine. The machine code is then executed by the CPU. This makes the code run faster and more efficiently. 

**C++/C Code -> AOT Compilation -> WASM Binary -> JIT Compilation -> Machine Code**

In JIT compilation, WASM use advance techniques to compile the code. It uses a technique called **Tiered Compilation**. Here's how it works:

**Baseline Compiler**: It's a sprinter. It compiles the code quickly, but the code is not optimized. It's used to get the code up and running quickly. Gives the speed, but not the efficiency. 

**Optimizing Compiler**: It runs in parallel with the Baseline Compiler. It takes time to compile the code, but the code is highly optimized. It's used to optimize the code for size and speed. It make sure every operation is optimized for the peak performance.

**Streaming Compilation**: While Optimizing Compiler is refining, that is need that user doesn't have to wait. This is where Streaming Compilation comes in. It compiles while the code is being downloaded. Till the complete binary is downloaded, the code is already compiled a good chunk of it. In this way it ensure that there is no lag in downloading anf running the code.

To bridging all these together, WASM uses a technique called **Tiered Compilation**. It uses a combination of Baseline Compiler, Optimizing Compiler, and Streaming Compilation to compile the code. It ensure that the application starts with the speed of Baseline Compiler, as the Optimizing Compiler produces more refine code it seamlessly transitions to it. And while the Streaming Compilation ensures that everything happen in real-tme providing the user a smooth experience. In-short Tiered Compilation works is to Baseline Compiler and transition to optimized machine code.

A high-level overview of the compilation process:

![Compilation Process](https://github.com/user-attachments/assets/5bbccebb-7428-49eb-8fbb-238df7f49811)


### List of WASM Compilers

There are several compilers available that can compile code to WebAssembly. Some of the popular compilers are:

- **LLVM**: It's a popular compiler infrastructure that can compile C, C++, Rust, and other languages to WebAssembly. It's a good at optimizing code and generating efficient machine code.
- **Binaryen**: It's not just a compiler, but a toolchain for WebAssembly. It can optimize and generate WebAssembly code.
- **Emscripten**: Emscripten is a toolchain for compiling to WebAssembly. It can compile C and C++ code to WebAssembly. It can also compile other languages like Rust, Go, etc. It's famous for taking complex code and run it in the browser. IT uses LLVM and Binaryen under the hood.
- **wasm-pack**: It's a tool for building and working with Rust and WebAssembly. It can compile Rust code to WebAssembly and generate the necessary JavaScript bindings. It works well with the Rust Package Manager (Cargo) and the Rust ecosystem.
- **AssemblyScript**: It's a TypeScript-like language that compiles to WebAssembly. It's designed to be a drop-in replacement for JavaScript. It's easy to learn and use, and it can be compiled to WebAssembly using the AssemblyScript compiler. The good part about it is, it's very well optimized for WebAssembly. It's good a memory management.

Some online tools for learning and experimenting with WebAssembly:

- **WasmFiddle**: It's an online tool for experimenting with WebAssembly. It allows you to write, compile, and run WebAssembly code in the browser. It's a good way to learn and understand how WebAssembly works.
- **WebAssembly Explorer**: It's an online tool for exploring WebAssembly. It allows you to compile C and C++ code to WebAssembly and see the generated WebAssembly code. It's a good way to understand how C and C++ code is compiled to WebAssembly.
- **WABT Online**: It's an online tool for working with WebAssembly. It allows you to assemble, disassemble, and run WebAssembly code in the browser. It's a good way to learn and experiment with WebAssembly.
- **WasmCloud Playground**: It's an online tool for building and running WebAssembly modules. It allows you to write, compile, and run WebAssembly code in the browser. It's a good way to learn and experiment with WebAssembly.

## Emscripten Compiler

To keep it simple Emscripten is a cool that that helps coders written in C, C++, Rust, etc. and bring to web. Many games and game engines like Unity, Nebula3, Geogram, etc, or graphics program like OpenGL ES 2.0, IMGUI has successfully ported to web using Emscripten.

![Emscripten](https://github.com/user-attachments/assets/e2cb88e0-5bab-4afc-80da-4de2964a5890)

It uses LLVM, Closure, Clang, Binaryen, etc. to compile the code to WebAssembly. It's a good at optimizing code and generating efficient machine code. It's famous for taking complex code and run it in the browser. It uses LLVM and Binaryen under the hood. It converts code in JavaScript and WebAssembly. The JavaScript code can be run both in the browser and on the server.

To compile a C code to WebAssembly using Emscripten, we can use the following command:

```bash
emcc hello.c
```

If your node version is old and you are don't support wasm, you can use the following command to compile the code:

```bash
emcc hello.c -s WASM=0
```

This will only generate the JavaScript code and not the WebAssembly code. The `-s WASM=0` flag tells the compiler not to generate WebAssembly output. Here `s` stands for setting.

Emscripten also make it convenient to test the code in the browser by generating the HTML file along with the JavaScript and WebAssembly files. It's a good way to test and run the code in the browser.

We can also specify the output file name using the `-o` flag. For example, to generate the `hello.html` file, we can use the following command:

```bash
emcc hello.c -o hello.html
```

So, this will generate a total of three files: `hello.html`, `hello.js`, and `hello.wasm`. The `hello.html` file is the main HTML file that loads the `hello.js` file, which in turn loads the `hello.wasm` file. The `hello.js` file contains the JavaScript code to load and run the WebAssembly module.

### Strict Mode

Strict mode is a way to opt-in to a restricted variant of JavaScript. Helps catch common errors, also ensure code compatibility and future-proofing. 

Some of the key features of strict mode are:

![strict mode](https://github.com/user-attachments/assets/43edb431-5a3d-4446-bafa-6d973f98709e)

For example if we are using `malloc` in our C++ code, which is the old way of allocating memory. But in strict mode, it's not allowed. And it will throw an error.

### Exporting Functions

When compiling the default behavior is not to export any functions. The reason being, it's more secure, optimization and enable encapsulation. However, there will be cases where exporting function will be necessary.

To do that first we need to define the function using the `EMSCRIPTEN_KEEPALIVE` macro. This macro tells the compiler to keep the function alive and not optimize it out.

```c
#include <emscripten.h>

EMSCRIPTEN_KEEPALIVE
int add(int a, int b) {
  return a + b;
}
```

Then we need to export the function using the `EXPORTED_FUNCTIONS` flag. This flag tells the compiler to export the specified functions.

```bash
emcc hello.c -o hello.html -s EXPORTED_FUNCTIONS="['_add']"
```

Then we call the exported function in the JavaScript code like this:

```js
const add = Module._add(2, 3); // 5

console.log(add);
```

### Using file system

As JavaScript doesn't have access to the file system and runs in a sandboxed environment. Emscripten provides a way to work with files in the browser. It provides a virtual file system that allows you to read and write files in the browser. The way it works is Emscripten preloads the file into the virtual file system. We use the `--preload-file` flag to preload the file.

For eg our C++ program open and reads a file called `input.txt` and we want to compile it to WebAssembly. We can use the following command:

```cpp
#include <stdio.h>

int main() {
  FILE *file = fopen("input.txt", "r");
  if (file == NULL) {
    printf("Error opening file\n");
    return 1;
  }

  char buffer[256];
  while (fgets(buffer, sizeof(buffer), file) != NULL) {
    printf("%s", buffer);
  }

  fclose(file);
  return 0;
}
```

```bash
emcc hello.c -o hello.html --preload-file input.txt
```

### Code Optimizations

Emscripten provides a way to optimize the code using the `-O` flag (it's a capital letter O). The range is from `-O0` to `-O3`. The `-O0` flag disables optimization. The `-O1` flag enables basic optimization. The `-O2` flag enables more optimization. The `-O3` flag enables the highest level of optimization. This will a game changer in terms of performance.

```bash
emcc hello.c -o hello.html -O3
```

There are more option in Emscripten compiler we can use `--help` flag to see all the options.

```bash
emcc --help
```

## Optimizing WebAssembly Code

Optimizing WebAssembly code is important for improving performance, reducing the size, reducing the resources usages, and most importantly the execution time (speed). There are several ways to optimize WebAssembly code:

- **Dead Code Elimination**: Remove unused code from the WebAssembly module. It reduces the size of the module and improves performance.
- **Wasm-opt**: Use the `wasm-opt` tool to optimize the WebAssembly module. It can optimize the code, reduce the size, and improve performance.
- **Compiler support**: Use a compiler that supports WebAssembly optimization. Compilers like Emscripten, Binaryen, etc. have built-in optimization features.
- **Efficient memory management**: Optimize memory usage in the WebAssembly module. Use linear memory efficiently and avoid unnecessary memory allocations.
- **Runtime**: Use a runtime like WasmTime that provides advanced optimization features. It can optimize the code, reduce the size, and improve performance.
- **AOT Compilation**: By pre-compiling the code to WebAssembly, it can optimize the code and reduce the size. It can also improve performance by compiling the code ahead of time.
- **Wizer for Pre-installation**: The module launch time can be reduced up to 60ms by using Wizer.
- **Programming Language**: Choose a programming language that compiles to efficient WebAssembly code. Languages like Rust, C++, etc. are optimized for WebAssembly.

This parameters become more crucial when we are building a large scale application. Also, it's more like a necessity because it ensures no matter where the application is running, it's running efficiently.

## Security Considerations

Security is a major concern when working with WebAssembly. As WebAssembly runs in a sandboxed environment, it's important to ensure that the code is secure and doesn't pose a security risk.

There are already some out of the box security features in WebAssembly:

- **Memory Safety**: Ensure that the code is memory safe and doesn't have buffer overflows, memory leaks, etc. Use tools like `wasm-opt` to optimize the code and ensure memory safety.
- **Digital Signatures**: Use digital signatures to verify the authenticity of the WebAssembly module. It ensures that the code hasn't been tampered with and is safe to run.
- **Sandboxing**: Run the WebAssembly code in a sandboxed environment. It isolates the code from the host environment and prevents it from accessing sensitive resources.
- **Same Origin Policy**: Use the same-origin policy to restrict the WebAssembly code from accessing resources from other origins. It ensures that the code can only access resources from the same origin.

### Best Practices on developer side

As a developer we can take some measures to ensure the security of the WebAssembly code:

- **Verify the input**: Ensure that the input is validated and type-checked before passing it to the WebAssembly module.
- **Verify memory boundaries**: Ensure that the memory accesses are within the bounds of the linear memory. Ensure we are using only allocated memory.
- **Incorporated Functions**: Ensure that the imported functions are safe and don't pose a security risk. Use only trusted functions. It can be done by limiting access and monitor incoming and outgoing data.
- **System Calls**: Limit the system calls to only necessary functions. Avoid unnecessary system calls that can pose a security risk as it's directly communicating with the OS or the host environment. Instead use WASI for system calls and libraries.

## WebAssembly Text Format (WAT)

WebAssembly Text Format (WAT) is a human-readable representation of the WebAssembly binary format. It's a simple, verbose, and readable format that allows you to write and read WebAssembly code. It's useful for debugging and understanding the structure of the WebAssembly module.

To convert a WAT file to a WASM file, we can use the `wat2wasm` tool. It's a tool that converts a WAT file to a WASM file. It takes a WAT file as input and generates a WASM file as output.

For example, to convert a WAT file called `hello.wat` to a WASM file, we can use the following command:

```bash
wat2wasm hello.wat -o hello.wasm
```

## Beyond Browser

We know WASM run effectively in the browser almost at native speed. But it's not just limited to the browser. We can create a similar environment outside of browse.

### Server-side Runtime

Runtimes like Wasmtime, Wasmer, WAVM, Lucet, etc. can run WebAssembly code on the server. With a dedicated runtime, app can process data faster, handle more requests, and provide a smoother user experience. 

For example like we did in by running the C code in the browser, we can run the same code on the server.

```c
// hello.c
#include <stdio.h>

int main() {
    printf("Hello, WebAssembly!\n");
    return 0;
}
```

```bash

emcc hello.c -o hello.js
```

Then we can run the `hello.wasm` file using node.js:

```bash
node hello.wasm
```

And it will output `Hello, WebAssembly!` to the console.

### Cloud Environment

Cloud environment works differently than the traditional server. WASM can run seamlessly in the cloud environment. It can be used to build serverless functions, microservices, APIs, etc. It can be used to build scalable, efficient, and secure cloud applications.

So from our previous example let's modify the JS file to make it work with the cloud environment.

```js
const express = require('express');
const app = express();
const wasmModule = require('./hello.js');

app.get('hello', (req, res) => {
  wasmModule.onRuntimeInitialized = () => {
    const hello = wasmModule.cwrap('hello', string, []);
    res.send(hello());
  };
});
app.listen(3000, () => console.log('Server running on port 3000'));
```

Now in the cloud, endpoints like this become the gateways to interact with our services. Now to make it cloud-compatible we will Dockerize the application.

```Dockerfile
FROM node:20
WORKDIR /app
COPY . .
RUN npm install express
CMD ["node", "server.js"]
```

Then we can build the Docker image and run it in the cloud environment.

```bash
docker build -t hello-wasm-node-server .
docker run -p 3000:3000 hello-wasm-node-server
```

And the server will be running on port 3000. If we visit `http://localhost:3000/hello`, it will output `Hello, WebAssembly!`.

Now if we deploy this Docker Image to the cloud, it will be running in the cloud environment. And we can access the service will a url provided instead of `localhost`.

### On the Edge

There are many edge computing platforms like AWS WaveLength, Azure Edge Zones, etc. As wasm app getting more containerized, it's becoming more easier to deploy on the edge. 

## WASM Runtimes

It's one of the crucial pieces of the WebAssembly ecosystem. The idea of run time is similar to Java's WORA (Write Once Run Anywhere), that was Java biggest promise and WOW factor. It takes the Java Code an turn into something machine can understand. No matter where you are running the code, be it on the server, browser, cloud, edge, etc. it will run the same way, until and unless it has JVM. And WASM is bit like Java's byte code, but for the web (any beyond).

![JVM](https://github.com/user-attachments/assets/c44f7800-eab4-4851-9c4b-6f2152a6a008)

So these WASM runtimes are similar to JVM for Java but designed for WebAssembly. They take WASM code and execute it on different platforms. 

![WASM Runtimes](https://github.com/user-attachments/assets/207a5808-23db-4600-9373-a2d715da06e5)

### Working of WASM Runtimes

For example, we have a WASM module (binary) created from the high-level language like C, C++, Rust, etc. Now we want to run this module on the runtime (or server). This server can be anything like a browser, cloud, edge, etc. Now this runtime will take this module, unpack it, readying it for execution. 

Next the runtime will translate this module into something the host machine can understand. More like translating a foreign language to the native language. Now depending on the runtime, this can be done Ahead-of-Time (AOT) like preparing to cook a meal before the guest arrives. Or Just-In-Time (JIT) - on the fly, like cooking the meal as the guest arrives.

Once the translation is done, the instructions are ready to be executed. Runtimes ensure these instructions are executed efficiently and correctly. During this time the runtime manges resources, memory, etc. to ensure the WASM code has what it needs while keeping it secure and isolated from the host environment. 

If module needs to interact with outside world, like calling a JavaScript function in a web browser or accessing manges these interactions. It's lik a mediator ensuring the communication is smooth and secure.

After all the instructions are executed, the runtime concludes the process. The result of module's execution is sent back to the host environment or end user.

To sum up words, it's like a stage manager in a play. It sets the stage, ensures everything is in place for performance. It manages the actors, props, etc. during the performance. And once the performance is over, it wraps up everything and sends the audience home. WASM runtime ensures that the WASM code not only runs correctly but also delivers the expected results wherever you want it to.

![WASM Runtimes](https://github.com/user-attachments/assets/4d9aad72-52d7-4c59-b52b-6597845d62d5)

### Various WASM Runtimes

There are several WebAssembly runtimes available that can run WebAssembly code on different platforms. We can choose based on our need and what kind of environment we are working in. Some of the popular WebAssembly runtimes are:

- **Bytecode Alliance - Wasmtime**: Known for adaptability. It can be embedded in language like Rust, Python and C. Works well with small IoT devices to large cloud servers. Offers pre-compilation and runtime interpretation.

- **Wasmer**: Specialized in light weight container for WebAssembly. Known for universal compatibility. It can be deployed on various platforms from cloud to desktop and IoT devices. Also features seamless integration with other languages. Also, it's known for it's speed and efficiency.

- **Bytecode Alliance - Lucet**: This is designed with security in mind. Most prominent for running risky WASM program securely. Excels handling the WASI in environment like Fastly edge cloud. Maintain efficient runtime without compromising on security.

- **Wasm3**: Known for it's small footprint. It's designed to run on small devices like microcontrollers. It's fast and efficient. It's designed to run on resource-constrained devices.

- **Bytecode Alliance - WAMR**: Known for it's minimalistic design makes it an ideal choice for resource-constrained devices. At it core it features the `iwasm` VM support JIT and AOT compilation as well as WASM interpreter. Despite it size, it adheres to W3C WASM MVP standard, ensuring consistency, reliability and performance.

Each and every runtime has it's own set of features and capabilities. It's important to choose the right runtime based on the requirements of the application. So, for general purpose application Wasmer or Wasmtime is a good choice. On other hand Web3 is about performance and portability. Lucet is about security and efficiency. Where as Emscripten is about bringing the code to the web.

![WASM Runtimes](https://github.com/user-attachments/assets/bd16cf74-1c4a-4715-9a84-674a969f0529)

## Cloud-Native WebAssembly

Taking best of both worlds, WebAssembly and Cloud-Native. It's not just about deploying WASM code in the cloud, it's about building cloud-native applications using WebAssembly. It's about leveraging the power of WebAssembly to build scalable, efficient, and secure cloud applications.

It comes with all the benefit of cloud-native offers. There in an Alliance of Fastly, Red hat, Mozilla, Intel, etc. to bring the best of WebAssembly to the cloud. The Alliance is constantly building tools and runtimes to run WASM out of the browser. They are working on building a secure, efficient, and scalable cloud-native platform for WebAssembly.

### Complimenting Cloud Native core pillars

As we know Cloud-Native is more than just a buzzword, it focuses on building and running scalable dynamic environments like public, private and hybrid cloud. WASM compliments the core pillars and play a pivotal role in enhancing this philosophy.

As we know in Cloud-Native, the applications are built using microservices. And each service handles specific function and cen be developed, deployed, and scaled independently. 
WASM enhances this by offering a consistent runtime for these services. As container are lightweight and standalone packages that can contain code, dependencies, and runtime. WASM complements that by providing a binary format that can be embedded in the container. It ensures that the code runs fast, secure, and efficiently that WASM offers.

![Cloud Native](https://github.com/user-attachments/assets/d13e4085-330b-4787-9b24-bb18189b9def)

#### Serverless Computing

Serverless is all about running code without managing the infrastructure. With WASM, serverless function can be written in any language that compiles to WebAssembly. It can be run on any WebAssembly runtime. It ensures that the code runs fast, secure, and efficiently. It's a good way to build scalable, efficient, and secure serverless applications.

#### CI/CD Pipelines

For any production level application CI/CD is important so that app can be deployed to prod with manual intervention. WASM fit seamlessly in CI/CD pipelines. It can be compiled in WASM binary, tested in consistent environment, and deployed to the cloud. This ensure that app benefit rapid iteration of CI?CD while also leveraging the potability and performance of WASM.

## Docker and WASM

Docker has changed the way we build, ship, and run applications. It make sure application run same in different environment, like on a developer's laptop or a big server in the cloud.It do comes with own set of challenges like it's dependence on the host OS. That is Docker image that is build for one kind of machine might not work on another, leading to the famous "It works on my machine" problem.

WASM works differently, it doesn't need the whole file system or OS to run. It focuses on just using what's necessary when the code is running. This is where WASI role comes in. It only access the necessary resources and nothing more. This makes WASM more portable and efficient than Docker. So docker is like a architecture-centric, where WASM is platform-centric.

### Docker VS WASM?

We will see how Docker and WASM are different and how they can be used together.

Let understand with an example:

If we take simple Rust program to create a web server.

```rust
use warp::Filter;

#[tokio::main]

async fn main() {
  let hello = warp::path!("hello" / String).map(|name| format!("Hello, {}!", name));
  warp::serve(hello).run(([127, 0, 0, 1], 3030)).await;
}
```

First, let's containerize the application using Docker. We will create a Dockerfile to build the application and run it in a container.

```Dockerfile
FROM rust:latest
WORKDIR /app
COPY . .
RUN cargo install --path .
CMD ["app"]
```

Then we can build the Docker image.

```bash
docker build -t rust-hello-world .
```

Next, let's compile this Rust program to WebAssembly. This involves setting up Rust toolchain for WASM compilation. We will use the `wasm-pack` tool to compile the Rust program to WebAssembly.

```bash
wasm-pack build
```
**Size Comparison**

Now if we compare the Docker image and WASM binary. The Docker image size is between 100-200 MB, where as the WASM binary is around 1-10 MB. This is because Docker image contains the whole file system and OS, where as WASM binary contains only the necessary code and resources.

![Size and Startup Time](https://github.com/user-attachments/assets/533d501b-51d4-4188-8cd5-24538481466f)

**Startup Time**

The startup time of Docker image is around 1-2 seconds, where as the startup time of WASM binary is around 100-200 ms. This is because Docker image needs to setup the virtual environment, where as WASM binary is ready to run.

![Size and Startup Time](https://github.com/user-attachments/assets/6ee39a58-7fc0-4ff2-a42c-f7fcf931105a)

Docker is pretty fast and behaves a lot like software running on the host machine. But it can affected by how big the image and container and the complexity of the application. WASM on the other hand is designed to be fast and efficient. It's particularly good for small, lightweight applications that need to run quickly and efficiently.

**Web Browser**

Docker wasn't made for working inside web browser. It's more about running on the server or standalone machine. WASM on the other hand is designed to run in the browser. It's fast, secure, and efficient. It's a good way to run code in the browser. Like Docker it can also run on server, Desktop, Cloud, Edge, etc. Unlike Docker* it can run on mobile devices.

![Web Browser](https://github.com/user-attachments/assets/a16dcacd-7864-47cc-a248-17f8fe54f4e1)

**Interaction with Host**

Docker interacts with the host machine through the Docker daemon. It has more direct relationship with the host machine. It can interact with system resources and OS-level operations. WASM on the other hand operates at a higher level abstraction. Its interaction with the system is mediated though WASI, which limits its direct access to system resources, which provide security benefits but can limit its capability for certain types of system interactions.

![Interaction with Host](https://github.com/user-attachments/assets/11c1979f-7468-46b0-a10b-7dca6ca2dc36)

### Docker and WASM Together

It's not WASM Vs Docker. It's more like bringing these both technologies together. From Docker side we can take the revolutionized the way we bundle and deploy applications. Along with portability. From WASM side we can take the speed, security, and efficiency. And it's ability to run code written in languages like C, C++, Rust, etc. in the browser with near-native performance and rapid startup time.

And Docker is supporting this evolution. That's we can package our native application code within WASM container and then share it a Docker image. This way we can take advantage of both technologies. With this we will able to WASM containers alongside of traditional Linux and Windows containers. 

![Docker and WASM](https://github.com/user-attachments/assets/c1fdd157-157a-4caf-ae73-aaaa034dec1f)

As Docker uses Docker Engine at it core, which processes all the container related requests. When we wun `docker run` command, the Engine uses a runtime like `runc` to start the container process. Post this it uses the containerd-shim to manage the container process, overseeing aspects like log and process management.

In case of WASM, it uses a special runtime called WasmEdge. To integrate this with Docker, a new `containerd-sim` named `containerd-wasm-shim`, has been developed. This shim ensure that WASM containers can be managed just like any other container in Docker.

![Docker and WASM](https://github.com/user-attachments/assets/4336dd56-7876-4aa3-acf1-8e7a95b1cbc4)

Now with this we can deploy app that combines the efficiency and performance of WASM with the universality, ease and portability of Docker. Now no matter whether you run this image on a microservices architecture, cloud, edge, or on a developer's laptop.

![Docker and WASM](https://github.com/user-attachments/assets/c40d0a50-b706-49e5-a921-8828579da3a7)

## Tools and Frameworks

### Wasm Frameworks

Frameworks are the building blocks of WebAssembly applications. They provide a set of tools, libraries, and conventions for building, deploying, and managing WebAssembly applications. Here are some pain points that frameworks solve:

- Without a framework, WASM devs find themselves **manually register services for load balancing, service discovery, etc**. Frameworks make it easier to build, deploy, and manage WebAssembly applications.

- **Effective communication** between WASM services is crucial. Without a framework, the communication channels can be become complex. 
  
- Another pain becomes the **security**, without a dedicated framework, devs find themselves implementing security feature like auth and encryption manually.

- **Integration with other platforms** can be cumbersome, leading to platform-specific code. And that become bottleneck to change the platform.

- Writing a lot of **boilerplate code** and make the application bloated and harder to maintain. It reduces modularity and scalability.

Now to sort out these pain points, there are frameworks available that can help in building, deploying, and managing WebAssembly applications. Some of the popular WebAssembly frameworks are:

![Wasm Frameworks](https://github.com/user-attachments/assets/c452bea9-17d4-436b-a86a-37b643b59795)

**SPIN**: It's an Open Source framework specifically designed for building WebAssembly applications. Its focus is on building, deploying, and running microservices that are fast, secure, and composable. It also aligns wih the latest advancements in WebAssembly and and supports various WASM runtimes.

![spin](https://github.com/user-attachments/assets/fc3abb46-20b2-4aa5-a7a3-fe963ac3d5bc)

**WasmCloud**: It's a versatile runtime capable of running app not just in cloud but also at the edge, within browsers, on compact devices, and virtually anywhere. Game changer in distributed computing.

![wasmcloud](https://github.com/user-attachments/assets/3b05ac61-1797-4110-bfa6-9572349489ff)

> Image credit: All the images and graphics are taken from KodeKloud course on WebAssembly. Shoutout to them for creating such wonderful graphics.
