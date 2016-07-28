**TCP Window Size**
-------------------
Determines the largest TCP receive window that the system offers. The receive window is the number of bytes a sender can transmit without receiving an acknowledgment. This entry overrides TCP's negotiated maximum receive window size and replaces it with the value of this entry.

TCP uses a receive window that is four times the size of the maximum TCP segment size (MSS) negotiated during connection setup, up to a maximum size of 64 KB.

**Bandwidth(bits/s) * Round trip latency(s) = TCP window size(bits)**

**Server details**
-------------------
54.244.0.127

**Optimal window size to masterofcode.com**
-------------------
round-trip min/avg/max/stddev = 129.012/184.389/295.647/65.982 ms

avg = 0.184389 s

internet connection = 100Mbit/s

**104857600(bit/s) * 0.184389(s) = 19334589(bit) = 2416824(byte)**
