# Memory Mapped Ping Pong

A `ping` application implemented using memory mapped files, compare-and-swap operations and memory barriers.

Implemented as an example for comparing different communication strategies between processes.

i.e.

```shell
./mmping PING
Two Processes: One Pings the other Pongs; using memory mapped files.
2023/07/04 00:05:11 PING 10M, 90.88/ns per ping, or 11.0M pings/sec
2023/07/04 00:05:12 PING 20M, 90.02/ns per ping, or 11.1M pings/sec
2023/07/04 00:05:13 PING 30M, 91.53/ns per ping, or 10.9M pings/sec
2023/07/04 00:05:14 PING 40M, 88.55/ns per ping, or 11.3M pings/sec
2023/07/04 00:05:14 PING 50M, 90.72/ns per ping, or 11.0M pings/sec
```

(on Intel(R) Core(TM) i7-8559U CPU @ 2.70GH, AMD/Ryzen 3800X is about 5.1M, Intel is better with compare-and-swap and
memory barriers)

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

0.024M vs 11.0M is about 450 times faster.