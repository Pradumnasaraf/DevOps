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

Unlike human readable text, the binary format is designed for machines, ensuring that it can be quickly decoded and executed. FOr example, the following code in C:

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

WebAssembly is a stack-based virtual machine. It means that it uses a stack to store and retrieve data. The stack is a data structure that follows the Last In First Out (LIFO) principle

![Stack-based virtual machine](https://github.com/user-attachments/assets/b0e7f220-edeb-41d3-ae9f-f2dcaac00b01)

### Linear memory

WebAssembly has a linear memory model. It allows the WebAssembly module to allocate and access memory in a linear fashion. It's like a big array that can be accessed by the WebAssembly module. It starts from 0 and growing up. It can be accessed via two primary instructions: `load` and `store`. The `load` instruction reads data from memory, while the `store` instruction writes data to memory. 

The instructions are suffixed with the data type and size. Example i32, i64, f32, f64. It describes the type of data and the size of the data in memory. So, in `i32` the `i` stands for integer and `32` stands for 32-bit or 4 bytes. As 8 bits make 1 byte, 32 bits make 4 bytes. Similarly, `f32` stands for floating point number with 32-bit or 4 bytes.

For example if we want to store number 7 in the 5th slot of the memory. We have to use `i32` type. To get the byte address of the 5th slot, we have to multiply the slot number by the size of the data type. So, 4 * 4 = 16. As slot starts from 0, instead of multiplying by 5, we have to multiply by 4.

![Screenshot 2024-12-20 at 4 21 07 PM](https://github.com/user-attachments/assets/a8838314-1871-4b6f-a413-063ef5394eda)

### Modules

Like other programming languages, WebAssembly code is organized into modules. A module is a collection of functions, types, tables, memories, and globals. Each module can import and export entities. It can also define its own entities. The module is the basic unit of code in WebAssembly.

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

WASM is designed to 

