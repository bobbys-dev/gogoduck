# gogoduck
Super simple blockchain with Go language.

The simulation starts with a genesis transaction and block. After each subsequent transaction, the blockchain will be printed.

For now, the hash is just a sum of the previous and the string length of a transaction statement.

The next repo update will show that modification to a transaction will cause subsequent hashes to be different. 

The complete implementation will have:
- At least an XOR or checksum or more advanced hash,
- Raft-like consensus when different processes have computed different blocks