## Homework №2

2.1. Answer what is 'window size'? How its calculated?
- TCP 'window size' is determines the largest TCP receive window that the system offers. The receive window is the number of bytes a sender can transmit without receiving an acknowledgment. This entry overrides TCP's negotiated maximum receive window size and replaces it with the value of this entry.
- TCP 'window size' is calculated as a product of bandwidth and round trip time - window size = bandwidth * RTT

2.2. Put here ip of your server
- Public IP - 52.59.34.190

2.3. Find optimal TCP window size to masterofcode.com

- Average RTT to masterofcode.com - 141.814 ms
- Assume that on I have 1Gb link on my server
- Optimal TCP window size is equal to - 1000000000 * 0.141814 = 141814000 / 8 = **17726750 Bytes**_
