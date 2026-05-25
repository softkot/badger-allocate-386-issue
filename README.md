```
lsb_release -a
No LSB modules are available.
Distributor ID:	Ubuntu
Description:	Ubuntu 22.04.5 LTS
Release:	22.04
Codename:	jammy
```
```
go version
go version go1.26.2 linux/amd64
```

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go run .
```

run just fine

```
badger 2026/05/25 18:47:24 INFO: All 0 tables opened in 0s
badger 2026/05/25 18:47:24 INFO: Discard stats nextEmptySlot: 0
badger 2026/05/25 18:47:24 INFO: Set nextTxnTs to 0
badger 2026/05/25 18:47:35 INFO: Lifetime L0 stalled for: 0s
badger 2026/05/25 18:47:35 INFO: 
Level 0 [ ]: NumTables: 04. Size: 231 KiB of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
```

but

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=386 go run .
```

outputs:

```
badger 2026/05/25 18:45:18 INFO: All 0 tables opened in 0s
badger 2026/05/25 18:45:18 INFO: Discard stats nextEmptySlot: 0
badger 2026/05/25 18:45:18 INFO: Set nextTxnTs to 0
badger 2026/05/25 18:45:27 ERROR: writeRequests: while opening file: /tmp/59afbd42-0850-4222-b5f9-5a185fcbfd81/000348.vlog err: cannot allocate memory
while mmapping /tmp/59afbd42-0850-4222-b5f9-5a185fcbfd81/000348.vlog with size: 8388608
2026/05/25 18:45:27 ERROR update error error="while opening file: /tmp/59afbd42-0850-4222-b5f9-5a185fcbfd81/000348.vlog err: cannot allocate memory\nwhile mmapping /tmp/59afbd42-0850-4222-b5f9-5a185fcbfd81/000348.vlog with size: 8388608"
```
