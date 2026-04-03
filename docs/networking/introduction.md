---
sidebar_position: 1
title: Networking Introduction
description: Learn about the basics of Networking.
tags: ["Networking", "OSI", "TCP", "UDP", "Ports", "URL", "IP Address"]
keywords: ["Networking", "OSI", "TCP", "UDP", "Ports", "URL", "IP Address"]
slug: "/networking"

---

Networking basics matter in DevOps because almost every system depends on moving data between services, hosts, and users. The goal is not to memorize everything at once, but to understand the main concepts well enough to reason about how applications communicate.

### OSI Layer

<p align="center"><img alt="OSI" src="https://user-images.githubusercontent.com/51878265/206166710-cafe1502-ea85-433d-b4bd-6124f8110992.png"></img></p>

### TCP

TCP is a connection-oriented protocol. It establishes a connection before sending data, which makes it a good fit when reliability matters more than raw speed. Common examples include web traffic, email, and file transfers.

### UDP

UDP is not connection-oriented. It sends data without waiting for connection confirmation from the receiver. Some packets may be lost, but UDP is useful when low latency matters more than perfect delivery, such as in voice, video, or gaming traffic.


### Ports

A port is a logical endpoint used by applications to send and receive network traffic. Ports allow multiple services to communicate on the same machine without getting mixed together.

<p align="center"><img alt="Tcp port" src="https://user-images.githubusercontent.com/51878265/206188329-c5b10491-d39e-40ca-8369-1a9965559857.png"></img></p>


|Port Number| Process | Uses |
|:--:|:--:|:--:|
|80 | HTTP | |
|443| HTTPS| |
|3306| MySQL | |

To check which ports the system is using:

```bash
netstat -a -b
```

### URL (Uniform Resource Locator)

A URL is a unique identifier used to locate a resource on the internet, such as an HTML page, image, or API endpoint.

### URL Breakdown

<p align="center"><img src="https://user-images.githubusercontent.com/51878265/206189760-ea426560-0d3c-4c5f-a8c4-b4f7c9d6f106.png"></img></p>

- A URL often omits the port when the default port is implied. For example, `https://google.com` usually means port `443`.


### IP Address

<p align="center"><img alt="IP Address" src="https://user-images.githubusercontent.com/51878265/206245742-5b660f3d-5d22-421f-ab35-56faf05b0532.png"></img></p>


### IP Address classes


<p align="center"><img alt="IP Address classes" src="https://user-images.githubusercontent.com/51878265/206355697-45304b89-eaba-42ba-b55d-0385c271ec9b.png"></img></p>

### What's next?

- [Networking Commands](./commands.md) - Learn about the commands that you can use with Networking.
- [Learning Resources](./learning-resources.md) - Learn more about Networking with these resources.
