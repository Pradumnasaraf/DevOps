## OSI Layer

<p align="center"><img alt="OSI" src="https://user-images.githubusercontent.com/51878265/206166710-cafe1502-ea85-433d-b4bd-6124f8110992.png"></p>

## TCP

TCP is a connection-oriented protocol. This means that it first establishes a link between the source and destination before it sends data. CP is a preferred protocol when data integrity is critical, such as in any transactional system. Eg: email and file transfer

## UDP

UDP in turn is not connection-oriented. UDP starts transmitting data immediately, without waiting for connection confirmation from the receiving side. Even though some data loss can happen, UDP is most often used in cases where speed is more important than perfect transmissions, such as in voice or video streaming.


## Ports

A virtual point where network connections start and end. So that multiple applications can communicate easily.

<p align="center"><img alt="Tcp port" src="https://user-images.githubusercontent.com/51878265/206188329-c5b10491-d39e-40ca-8369-1a9965559857.png"></p>


|Port Number| Process | Uses |
|:--:|:--:|:--:|
|80 | HTTP | |
|443| HTTPs| |
|3306| MySQl | |

To check which ports the system are using

```bash
netstat -a -b
```

## URL (Uniform Resource Locator)

A unique identifier is used to locate a resource on the Internet. Lilke HTML, JS files.

### URL Breakdown

<p align="center"><img src="https://user-images.githubusercontent.com/51878265/206189760-ea426560-0d3c-4c5f-a8c4-b4f7c9d6f106.png"></p>

- Generally, the URL dosen't contains a port number in the string because it is by default. For eg, `google.com` or `google.com:443`


## IP Address

<p align="center"><img alt="IP Address" src="https://user-images.githubusercontent.com/51878265/206245742-5b660f3d-5d22-421f-ab35-56faf05b0532.png"></p>


### IP Address classes


<p align="center"><img alt="IP Address classes" src="https://user-images.githubusercontent.com/51878265/206355697-45304b89-eaba-42ba-b55d-0385c271ec9b.png"></p>


<!--<p align="center"><img alt="" src=""></p> -->
