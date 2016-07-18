TCP Window Size
---
**TCP window size** is the amount of bytes that can be transmitted without receiving an acknowledgement from the other side.
**TCP window size** is calculated as a product of bandwidth and round trip time: 
**_TCPWindowSize = Bandwidth * RTT_**

Server IP Address
---
54.93.56.17

Optimal TCP Window Size
---
Average RTT to masterofcode.com: 96.183 ms
1Gb Internet connection is assumed.
So optimal TCP window size is:
_1000000000 * 0.096183 = 96183000 / 8 = **12022875 Bytes**_
