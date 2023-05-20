# go test 两种模式
1. 本地模式
2. 列表模式

# go test 参数详解
- -bench 运行匹配的性能测试用例
- -benchtime 运行次数或时间 1s 1m 100x
- -count 执行测试次数，相当于多次运行 go test
- -cover 启用覆盖率分析，表示多少代码被测试
- -cpu 1,2,4 go test 会执行3次，其中runtime.GOMAXPROCS分别是1,2,4
- -fuzz 运行匹配的模糊测试用例，只能一个一个的运行
- -fuzz 模糊测试时长，默认一直运行
- -list 列出匹配的顶层测试用例
- -parallel 指定性能测试时，并行CPU数量
- -run 运行匹配的功能测试的测试用例
- -short 标识是否缩短运行时间
- -timeout go test 超时时间，默认十分钟，设置为0表示禁用
- -v 打印所有输出
- -benchmem 打印内存分配统计信息
- -blockprofile 将阻塞数据写入到指定的文件
- -coverprofile 将覆盖率数据写入指定文件
- -cpuprofile 将CPU使用数据写入指定文件
- -memprofile 将内存数据写入到指定文件
- -mutexprofile 将互斥锁信息写入到指定文件
- -outputdir 指定输出目录
- -trace 将执行跟踪信息写入到指定文件