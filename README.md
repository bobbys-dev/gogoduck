# gogoduck
Super simple blockchain with Go language.

The simulation starts with a genesis transaction and block. After each subsequent transaction, the blockchain will be printed.

For now, the hash is just a sum of the previous and the string length of a transaction statement.
```
----Transactions to be run----

index: 0 value: Alice sent Bob 10.00

index: 1 value: Bob received 2.15 from Charlie

index: 2 value: Charlie sent Alice 13.00

----Start of simulation----

The genesis block is: {0 This is a placeholder for the imaginary genesis transaction. 60}

Transaction# 0 : Block generated:  {60 Alice sent Bob 10.00 80}

Transaction# 1 : Block generated:  {80 Bob received 2.15 from Charlie 110}

Transaction# 2 : Block generated:  {110 Charlie sent Alice 13.00 134}
```

The next repo update will show that modification to a transaction will cause subsequent hashes to be different. 

The complete implementation will have:
- At least an XOR or checksum or more advanced hash,
- Raft-like consensus when different processes have computed different blocks