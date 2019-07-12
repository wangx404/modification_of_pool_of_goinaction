# modification_of_pool_of_goinaction
Pool example in "GO In action" has a small problem to fix.

The output of ths pool example would look like below:
`
2019/07/12 14:59:35 Acquire: New Resource
2019/07/12 14:59:35 Create: New Connection 1
2019/07/12 14:59:35 Acquire: New Resource
2019/07/12 14:59:35 Create: New Connection 2
2019/07/12 14:59:35 Query: QID[3] CID[1]
2019/07/12 14:59:35 Release: In Queue
2019/07/12 14:59:35 Acquire: Shared Resource
2019/07/12 14:59:36 Query: QID[1] CID[2]
2019/07/12 14:59:36 Release: In Queue
2019/07/12 14:59:36 Acquire: Shared Resource
2019/07/12 14:59:36 Query: QID[24] CID[1]
2019/07/12 14:59:36 Release: In Queue
2019/07/12 14:59:36 Acquire: Shared Resource
2019/07/12 14:59:36 Query: QID[0] CID[2]
2019/07/12 14:59:36 Release: In Queue
2019/07/12 14:59:36 Acquire: Shared Resource
2019/07/12 14:59:36 Query: QID[14] CID[1]
2019/07/12 14:59:36 Release: In Queue
2019/07/12 14:59:36 Acquire: Shared Resource
2019/07/12 14:59:36 Query: QID[15] CID[2]
2019/07/12 14:59:36 Release: In Queue
2019/07/12 14:59:36 Acquire: Shared Resource
2019/07/12 14:59:36 Query: QID[19] CID[1]
2019/07/12 14:59:36 Release: In Queue
2019/07/12 14:59:36 Acquire: Shared Resource
2019/07/12 14:59:37 Query: QID[20] CID[2]
2019/07/12 14:59:37 Release: In Queue
2019/07/12 14:59:37 Acquire: Shared Resource
2019/07/12 14:59:37 Query: QID[16] CID[1]
2019/07/12 14:59:37 Release: In Queue
2019/07/12 14:59:37 Acquire: Shared Resource
2019/07/12 14:59:37 Query: QID[21] CID[2]
2019/07/12 14:59:37 Release: In Queue
2019/07/12 14:59:37 Acquire: Shared Resource
2019/07/12 14:59:38 Query: QID[17] CID[1]
2019/07/12 14:59:38 Release: In Queue
2019/07/12 14:59:38 Acquire: Shared Resource
2019/07/12 14:59:38 Query: QID[22] CID[2]
2019/07/12 14:59:38 Release: In Queue
2019/07/12 14:59:38 Acquire: Shared Resource
2019/07/12 14:59:38 Query: QID[23] CID[2]
2019/07/12 14:59:38 Release: In Queue
2019/07/12 14:59:38 Acquire: Shared Resource
2019/07/12 14:59:38 Query: QID[18] CID[1]
2019/07/12 14:59:38 Release: In Queue
2019/07/12 14:59:38 Acquire: Shared Resource
2019/07/12 14:59:38 Query: QID[2] CID[1]
2019/07/12 14:59:38 Release: In Queue
2019/07/12 14:59:38 Acquire: Shared Resource
2019/07/12 14:59:38 Query: QID[4] CID[1]
2019/07/12 14:59:38 Release: In Queue
2019/07/12 14:59:38 Acquire: Shared Resource
2019/07/12 14:59:38 Query: QID[8] CID[2]
2019/07/12 14:59:38 Release: In Queue
2019/07/12 14:59:38 Acquire: Shared Resource
2019/07/12 14:59:39 Query: QID[5] CID[2]
2019/07/12 14:59:39 Release: In Queue
2019/07/12 14:59:39 Acquire: Shared Resource
2019/07/12 14:59:39 Query: QID[10] CID[1]
2019/07/12 14:59:39 Release: In Queue
2019/07/12 14:59:39 Acquire: Shared Resource
2019/07/12 14:59:39 Query: QID[6] CID[2]
2019/07/12 14:59:39 Release: In Queue
2019/07/12 14:59:39 Acquire: Shared Resource
2019/07/12 14:59:39 Query: QID[9] CID[1]
2019/07/12 14:59:39 Release: In Queue
2019/07/12 14:59:39 Acquire: Shared Resource
2019/07/12 14:59:39 Query: QID[7] CID[2]
2019/07/12 14:59:39 Release: In Queue
2019/07/12 14:59:39 Acquire: Shared Resource
2019/07/12 14:59:39 Query: QID[12] CID[2]
2019/07/12 14:59:39 Release: In Queue
2019/07/12 14:59:39 Acquire: Shared Resource
2019/07/12 14:59:39 Query: QID[13] CID[2]
2019/07/12 14:59:39 Release: In Queue
2019/07/12 14:59:40 Query: QID[11] CID[1]
2019/07/12 14:59:40 Release: In Queue
2019/07/12 14:59:40 Shutdown Program.
2019/07/12 14:59:40 Close: Connection 2
2019/07/12 14:59:40 Close: Connection 1
`
