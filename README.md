# Mit6.824  基于golang1.9

## 完成lab1中五个part：
		1.1修改main/common_map ， common_reduce实现mapwork完成后nReduce个文件的分别写入与读出
		1.2修改main/wc.go 中的mapF，reduceF实现将输入文章中单词分离并统计
		1.3实现mapreduce/schedule.go 实现master的work调度，使用go程master可以实现并行处理
		1.4当1.3中schedule中call work出现异常时进行错误处理，使用chan在go程间通信
		1.5实现反向索引，修改main/ii.go中mapF,reduceF，注意去重复和统计数量
## 完成lab2的三个part：
		2.1初始化选举；Leader出现disconnect，然后重新连接；
		2.2从start()接收NewLog，并实现系统对日志的处理；
		可用性：是否能够在可接受时间内完成命令期望的commit；
		Followers断线重连的情况；
		多个start()并发；断线Leader重连；
		Leader与一部分Follower同时断线，分别重连；
		RPC数量限制；
		2.3机器故障从persist重启；提交关于之前的日志的commit；
		RPC出现随机不可靠始建；save persister
![图 1 ](./images/raft-图1.jpg) 
![图 2 ](./images/raft-图2.jpg) 
![图 3 ](./images/raft-图3.jpg) 

 
