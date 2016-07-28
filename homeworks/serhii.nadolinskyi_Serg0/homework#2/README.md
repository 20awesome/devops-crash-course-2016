* 'window size' - max length of data which can be transmitted in one certain package
* server ip: 52.28.184.172
* Optimal window size:  
1,000,000,000 bps * 0.091 seconds = 91,000,000 bits / 8 = 11,375,000 Bytes ~ 11 Mb
* Optimal window size formula:  
Bandwidth-in-bits-per-second * Round-trip-latency-in-seconds = TCP window size in bits / 8 = TCP window size in bytes
