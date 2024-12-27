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

In WASM stack0base model, the result of the last expression in the function body is implicitly returned. So, we don't need to use `return` keyword. 

Talking about value stack it will change in this order: 

empty (locals with be a, b) -> a (locals with be b) -> a, b (locals with be empty) -> (a + b) (locals with be empty)

#### Essence of WebAssembly Modules

It not juts about function and using the functionality in our code. It about how these functions and along with other components comes to form a complete executable module. It's more like a container that holds all the necessary components for a specific functionality. 

For example for image processing module, it can have functions for applying filters, resizing images, etc. It can have memories to store the image data, tables to store the filter data, etc. It can have globals to store the image dimensions, etc.

![module](https://github.com/user-attachments/assets/a8ea0ea4-c60d-4d5e-a05c-cbc5191b8eb5)

### Beyond Functions

Apart from functions modules can have other components like:

**Global Variables**: They are variables that are accessible from anywhere in the module. They can be imported from other modules or defined within the module. They can be mutable or immutable. They can be used to store constants, configuration values, etc.

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

Data types are very essential to make the code is predictable by defining the kind of data and hoe it's used, we avoid errors and ensure efficient use of memory. For example imagine pouring water into a glass. Each glass (data type) can hold a specific amount and shape of water (data). If we try to pour too much water or wrong type of liquid, it either won't fit or won't function as intended. Same goes with data types, ensuring data fits well and works as intended.

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