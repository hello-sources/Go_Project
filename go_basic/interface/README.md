# interface 隐式实现
1. golang 对象实现 interface 无需任何关键词，只需要对该对象的方法集中包含接口定义的所有方法且方法签名一致
2. 对象实现接口可以借助 struct 内嵌的特性，实现接口的默认实现
3. 类型 T 方法集包含全部 receiver T 方法: 类型 * T 方法集包含 receiver T + *T 方法
4. 类型 T 实例 value 或 pointer 可以调用全部的方法，编译器会自动转换
5. 类型 T 接口，不管是 T 还是 *T 都实现了该接口
6. 类型 *T 实现接口，只有 T 类型的指针实现了该杰阔

# 应用场景
1. 面向接口编程