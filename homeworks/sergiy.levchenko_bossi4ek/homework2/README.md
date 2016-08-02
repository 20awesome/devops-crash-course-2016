## Lection 2
#### Answer what is 'window size'? How its calculated?
- Window size - Determines the largest TCP receive window that the system offers. The receive window is the number of bytes a sender can transmit without receiving an acknowledgment.

- Formula to calculate the optimal TCP window size:
  Bandwidth-in-bits-per-second * Round-trip-latency-in-seconds = TCP window size in bits / 8 = TCP window size in bytes
  
#### Put here ip of your server.

Public IP - 52.59.47.100

#### Find optimal TCP window size to masterofcode.com.

ping masterofcode.com
│PING masterofcode.com (104.236.105.130) 56(84) bytes of data.
│64 bytes from masterofcode.com (104.236.105.130): icmp_seq=1 ttl=54 time=97.5 ms
│64 bytes from masterofcode.com (104.236.105.130): icmp_seq=2 ttl=54 time=97.5 ms
│64 bytes from masterofcode.com (104.236.105.130): icmp_seq=3 ttl=54 time=97.7 ms
│64 bytes from masterofcode.com (104.236.105.130): icmp_seq=4 ttl=52 time=91.0 ms
│^C
│--- masterofcode.com ping statistics ---
│4 packets transmitted, 4 received, 0% packet loss, time 3001ms
│rtt min/avg/max/mdev = 91.001/95.951/97.748/2.859 ms

1000000000 bps * 0.0096 seconds / 8 = 12000000 bytes