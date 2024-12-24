---
sidebar_position: 1
title: WebAssembly Introduction
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

## WebAssembly Modules

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

