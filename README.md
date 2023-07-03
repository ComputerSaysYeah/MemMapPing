# Memory Mapped Ping Pong

A `ping` application implemented using memory mapped files, compare-and-swap operations and memory barriers.

Implemented as an example for comparing different communication strategies between processes.

i.e.

```shell
./mmping PING
Two Processes: One Pings the other Pongs; using memory mapped files.
2023/07/03 23:59:01 PING 10M, 194.24/ns per ping, or 5.1M pings/sec
2023/07/03 23:59:03 PING 20M, 194.19/ns per ping, or 5.1M pings/sec
2023/07/03 23:59:05 PING 30M, 194.18/ns per ping, or 5.1M pings/sec
2023/07/03 23:59:07 PING 40M, 194.18/ns per ping, or 5.1M pings/sec
2023/07/03 23:59:09 PING 50M, 194.19/ns per ping, or 5.1M pings/sec
```

(Ryzen 3800X, on Intel Architectures Compare-And-Swap operations are twice as fast, you can achieve about 10M pings/sec)

A normal ping would take:

```shell 
ping -s 16 localhost 
PING localhost (127.0.0.1) 16(44) bytes of data.
24 bytes from localhost (127.0.0.1): icmp_seq=1 ttl=64 time=0.041 ms
24 bytes from localhost (127.0.0.1): icmp_seq=2 ttl=64 time=0.023 ms
24 bytes from localhost (127.0.0.1): icmp_seq=3 ttl=64 time=0.034 ms
24 bytes from localhost (127.0.0.1): icmp_seq=4 ttl=64 time=0.033 ms
```

0.041ms is about 410us or 410000ns. 0.024M (24K) pings/sec.

0.024M vs 5.1M is about 212 times faster.