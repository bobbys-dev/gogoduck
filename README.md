# gogoduck
Super simple blockchain with Go language.

The Block data structure holds
i) data about the current transaction
ii) the SHA256 hash of the previous block
iii) a SHA256 hash computation of current data hash and the previous block's hash

The simulation starts with a genesis transaction and block. After each subsequent transaction, the blockchain will be printed.

```
----Transactions to be run----
index: 0 value: Alice sent Bob 10.00
index: 1 value: Bob received 2.18 from Charlie
index: 2 value: Charlie sent Alice 13.00
----Start of simulation----
The genesis block hash:
f04542972735b200b9923f7bcb981b980c35fbf751a515d9fe86cc1489ff3f3e
The genesis block transaction:
This is a dummy genesis transaction: Alpha created 10^100^100 units!

Transaction# 0 : Block generated: {
   Prev:f04542972735b200b9923f7bcb981b980c35fbf751a515d9fe86cc1489ff3f3e
   Alice sent Bob 10.00
   Cur:0bc3494d9afb35259c38e523452853d5947ddc119400eadb3f72ef7e6655765a
}
Transaction# 1 : Block generated: {
   Prev:0bc3494d9afb35259c38e523452853d5947ddc119400eadb3f72ef7e6655765a
   Bob received 2.18 from Charlie
   Cur:13ef8aa7c185600c96e23fbeaae50dac34e165bc23e3033dde56de0a0a6f97f2
}
Transaction# 2 : Block generated: {
   Prev:13ef8aa7c185600c96e23fbeaae50dac34e165bc23e3033dde56de0a0a6f97f2
   Charlie sent Alice 13.00
   Cur:6611baa6def1404d7702f01f2a6c7e76fdc1988a1872ebc72e19b5e2908a19ac
}

```

The next repo update will show that modification to a transaction will cause subsequent hashes to be different. 

The complete implementation will have:
- Raft-like consensus when different processes have computed different blocks