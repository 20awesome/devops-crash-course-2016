#Homework 2

##TCP Window
The throughput of a communication is limited by two windows: the congestion window and the receive window. The congestion window tries not to exceed the capacity of the network (congestion control); the receive window tries not to exceed the capacity of the receiver to process data (flow control). The receiver may be overwhelmed by data if for example it is very busy (such as a Web server). Each TCP segment contains the current value of the receive window. If, for example, a sender receives an ack which acknowledges byte 4000 and specifies a receive window of 10000 (bytes), the sender will not send packets after byte 14000, even if the congestion window allows it.

##Calculate Bandwidth-delay Product and TCP buffer size
BDP ( Bits of data in transit between hosts) = bottleneck link capacity (BW) * RTT
throughput <= TCP buffer size / RTT

##Server IP Address
54.93.40.165
##TCP Window Size
RTT to masterofcode.com: 94.586 ms
1Gb Internet connection is assumed.
So optimal TCP window size is:
1000000000 * 0.094586 = 94586000 / 8 =  11823250 Bytes
