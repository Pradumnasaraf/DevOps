
### Ping

Ping is a command that is used to test the connectivity between two devices. It sends ICMP echo requests to the target device and waits for ICMP echo replies. If the target device replies, it means that the connection is working.

```bash
ping google.com
```

### Traceroute

Traceroute is a command that is used to find the path taken by the packets to reach the target device. It sends ICMP echo requests to the target device and waits for ICMP time exceeded messages. If the target device replies, it means that the connection is working.

```bash
traceroute google.com
```

### Nslookup

Nslookup is a command that is used to find the IP address of a domain name. It sends a DNS query to the DNS server and waits for a response.

```bash
nslookup google.com
```

### Netstat

Netstat is a command that is used to display the network connections and routing tables. It can be used to find the open ports on a device.

```bash
netstat -tulpn
```

### Hostname

Hostname is a command that is used to display the name of the current host.

```bash
hostname
```

