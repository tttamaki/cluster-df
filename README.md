# CLUSTER-DF

This is inspired by [cluster-smi](https://github.com/PatWie/cluster-smi) for GPUs and [cluster-top](https://github.com/PatWie/cluster-top) for CPUs.


The same as `df` but for multiple machines at the same time.

It only supports Linux (tested under Ubuntu)!
Ext4 filesystems are shown.


Output should be something like

```
Tue Dec 4 14:03:42 2018 (http://github.com/tttamaki/cluster-df)
+---------------------+-----------------------------+---------+---------+---------+------+----------------+
| Node                | Filesystem                  | Total   | Used    | Avail   | Use% | mount          |
+---------------------+-----------------------------+---------+---------+---------+------+----------------+
| your host 1         | /dev/sdb1                   | 9240 GB |  559 GB | 8214 GB |  6.1 | /mnt/HDD10TB-1 |
|                     | /dev/sda1                   | 9240 GB |   39 MB | 8774 GB |  0.0 | /mnt/HDD10TB-2 |
|                     | /dev/mapper/ubuntu--vg-root | 1831 GB |  159 GB | 1579 GB |  8.7 | /              |
+---------------------+-----------------------------+---------+---------+---------+------+----------------+
| your host 2         | /dev/sdb1                   | 9239 GB |  534 GB | 8238 GB |  5.8 | /mnt/HDD10TB-2 |
|                     | /dev/sda1                   | 9239 GB |  165 GB | 8608 GB |  1.8 | /mnt/HDD10TB-1 |
|                     | /dev/mapper/ubuntu--vg-root | 1831 GB |  144 GB | 1594 GB |  7.9 | /              |
+---------------------+-----------------------------+---------+---------+---------+------+----------------+
```

Install and usage are the same with [cluster-top](https://github.com/PatWie/cluster-top).
