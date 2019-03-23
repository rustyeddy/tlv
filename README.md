# Type Length Vector (TLV)

A simple and efficient Real Time Communication Protocol

TLV is a simple (Type Length Vector) application protocol for
communicating applicaitons.  There are other options, that include 
ad-hoc formatting, where the package formats and handling are all
specific to a single piece of communication.  This requires anit-DRY 
coding practices and introduces bugs and stuff.

By sticking to a simple, but well defined library, many of the 
required inter-system communications can be handled seamlessly.

We could go with something like, JSON (not a bad choice), or 
RPC (hrmm), or XML (please God, No!).  But each of these choices are 
a little heavy weight in terms of Bandwidth, Memory and processing 
required.

For real time communications (except for streaming media) most 
communications are short and often bursty.  However, the messages are
typically sort but need to be communicated rapidly.

In this case it makes sense to use a protocol that is very simple, 
yet does require a "tight coupleing" between the applicaitons that are 
communicating.

The type of message will determine how long it is, and how much data
it can pass around.

Short Code:

```
+----+----+
|1        |  ~ 1 - byte compact message
+----+----+
      Type Only

+----+----+----+----+
|01       |         | ~ 2 - byte short message
+----+----+----+----+
      Type & Length 

+----+----+----+----+----+----+
|00       |         |     ....  ~ 2 - 256bytes (2^7)
+----+----+----+----+----+----+
      Type Length & Data
```
