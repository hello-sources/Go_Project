# 验证框架tag之间符号
1. 逗号(,):多个验证标签的分隔符
2. 横线(-):跳过该字段不验证
3. 竖线(|):使用多个验证标签，只需要满足其中一个即可
4. omitempty:如果字段未设置，则忽略该符号后面的校验
5. dive:表示深入一层验证，用于切片和集合