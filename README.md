[![PkgGoDev](https://pkg.go.dev/badge/github.com/CortexFoundation/rosetta-cortex)](https://pkg.go.dev/github.com/CortexFoundation/rosetta-cortex)
[![license](https://img.shields.io/github/license/CortexFoundation/rosetta-cortex.svg)](https://github.com/CortexFoundation/rosetta-cortex/master/LICENSE)
[![LoC](https://tokei.rs/b1/github.com/CortexFoundation/rosetta-cortex)](https://github.com/CortexFoundation/rosetta-cortex)

# Cortex Rosetta Shared Lib

Shared lib used in Cortex-SDK. This includes the libraries that are used by
different versions of the SDK, including Launchpad and Stargate.

Cortex full node

https://github.com/CortexFoundation/CortexTheseus
```
./cortex --rpc --rpcapi "net,web3,debug,admin,personal,ctxc"
```

Starting the proxy service
```
go run cmd/main.go
```

Do testing

```
./test-insall.sh
./test-run.sh
```
Test result
```
2021/04/19 20:54:07 increasing syncer concurrency to 59 (projected new cache size: 0.142355 MB)
2021/04/19 20:54:07 increasing syncer concurrency to 60 (projected new cache size: 0.144768 MB)
2021/04/19 20:54:07 increasing syncer concurrency to 61 (projected new cache size: 0.147181 MB)
2021/04/19 20:54:07 increasing syncer concurrency to 62 (projected new cache size: 0.149593 MB)
2021/04/19 20:54:07 increasing syncer concurrency to 63 (projected new cache size: 0.152006 MB)
2021/04/19 20:54:07 increasing syncer concurrency to 64 (projected new cache size: 0.154419 MB)
[MEMORY] Heap: 4148.303467MB Stack: 1.500000MB System: 4302.548340MB GCs: 2
[STATS] Blocks: 264425 (Orphaned: 100) Transactions: 2 Operations: 3 Accounts: 8892 Reconciliations: 0 (Inactive: 0, Exempt: 0, Skipped: 570839, Coverage: 0.000000%)
[PROGRESS] Blocks Synced: 264330/3661132 (Completed: 7.219898%, Rate: 121.811982/second) Time Remaining: 7h44m45s Reconciler Queue: 0 (Last Index Checked: -1)
2021/04/19 20:54:22 successful value log garbage collection (6.879729026s)
[MEMORY] Heap: 5346.595734MB Stack: 1.531250MB System: 5587.370018MB GCs: 2
[STATS] Blocks: 266087 (Orphaned: 100) Transactions: 2 Operations: 3 Accounts: 8892 Reconciliations: 0 (Inactive: 0, Exempt: 0, Skipped: 570839, Coverage: 0.000000%)
[PROGRESS] Blocks Synced: 265989/3661133 (Completed: 7.265210%, Rate: 122.013761/second) Time Remaining: 7h43m45s Reconciler Queue: 0 (Last Index Checked: -1)
[MEMORY] Heap: 5941.703407MB Stack: 1.531250MB System: 6262.950294MB GCs: 2
[STATS] Blocks: 267880 (Orphaned: 100) Transactions: 2 Operations: 3 Accounts: 8892 Reconciliations: 0 (Inactive: 0, Exempt: 0, Skipped: 573966, Coverage: 0.000000%)
[PROGRESS] Blocks Synced: 267786/3661133 (Completed: 7.314293%, Rate: 122.277169/second) Time Remaining: 7h42m31s Reconciler Queue: 0 (Last Index Checked: -1)

+--------------------------+--------------------------------+-----------+
|     CHECK:DATA STATS     |          DESCRIPTION           |   VALUE   |
+--------------------------+--------------------------------+-----------+
| Blocks                   | # of blocks synced             |    268017 |
+--------------------------+--------------------------------+-----------+
| Orphans                  | # of blocks orphaned           |       100 |
+--------------------------+--------------------------------+-----------+
| Transactions             | # of transaction processed     |         2 |
+--------------------------+--------------------------------+-----------+
| Operations               | # of operations processed      |         3 |
+--------------------------+--------------------------------+-----------+
| Accounts                 | # of accounts seen             |      8892 |
+--------------------------+--------------------------------+-----------+
| Active Reconciliations   | # of reconciliations performed |         0 |
|                          | after seeing an account in a   |           |
|                          | block                          |           |
+--------------------------+--------------------------------+-----------+
| Inactive Reconciliations | # of reconciliations performed |         0 |
|                          | on randomly selected accounts  |           |
+--------------------------+--------------------------------+-----------+
| Exempt Reconciliations   | # of reconciliation failures   |         0 |
|                          | considered exempt              |           |
+--------------------------+--------------------------------+-----------+
| Failed Reconciliations   | # of reconciliation failures   |         0 |
+--------------------------+--------------------------------+-----------+
| Skipped Reconciliations  | # of reconciliations skipped   |    573966 |
+--------------------------+--------------------------------+-----------+
| Reconciliation Coverage  | % of accounts that have been   | 0.000000% |
|                          | reconciled                     |           |
+--------------------------+--------------------------------+-----------+

2021/04/19 20:54:37 successful value log garbage collection (4.730788902s)
badger 2021/04/19 20:54:37 INFO: Storing value log head: {Fid:2 Len:34 Offset:17914198}
badger 2021/04/19 20:54:37 INFO: [Compactor: 173] Running compaction: {level:0 score:1.73 dropPrefixes:[]} for level: 0
badger 2021/04/19 20:54:38 INFO: LOG Compact 0->1, del 2 tables, add 1 tables, took 1.217551533s
badger 2021/04/19 20:54:38 INFO: [Compactor: 173] Compaction for level: 0 DONE
badger 2021/04/19 20:54:38 INFO: Force compaction on level 0 done
```

