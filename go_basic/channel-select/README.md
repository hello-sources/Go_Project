# 应用场景
1. 协程间通信，即协程间数据传递
2. 并发场景下利用channel的阻塞机制，作为同步机制（类似队列）
3. 利用channel 关闭时发送广播的特性，作为协程退出通知

# channel 通过通讯共享内存
1. channel 的方向：读、写、读写
2. channel 协程间通信信道
3. channel 阻塞协程
4. channel 并发场景下同步机制
5. channel 通知协程退出
6. channel 的多路复用

# 注意
1. channel 用于协程间通讯，必须窜在读写双方，否则将造成死锁