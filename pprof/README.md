# Go 性能分析
1. CPU
2. 内存
3. IO
4. goroutine
5. 死锁

# pprof 与单元测试 注意事项
1. 线上环境性能分析，选择流量较少的时间段
2. 通过测试或指标，或通过测试环境模拟生产环境
3. 如果在完全没有流量的情况获取指标数据，没有意义

# pprof 内置工具，数据采集方式
1. web网页采集
2. 基准测试采集数据
3. 通过硬编码测试

# 指标
1. Flat:函数自身运行资源小号
2. Sum:指的就每一行的flat与上面所有行flat%的总和
3. Cum:当前函数与其调用栈

通过Web网页，以图片或网页查看性能情况需要安装插件
地址：https://graphviz.gitlab.io/download/

# 命令行分析
1. top 按资源的消耗排序
2. list 展示资源消耗的源码
3. web 以 svc 图片展示连线图

# web 应用
1. go tool pprof -http=:8080 <source>
2. go tool pprof -web <source>