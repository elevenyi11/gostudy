状态码导出
目的：用于给策划翻译成其他语言
操作方式：
1.将ResultStatus.proto文件放在本目录
2.执行python daochu.py
3.将生成的code.txt文件给策划让其翻译(只可替换中文部分，不可以改动文件结构)


状态码导入
目的：将策划翻译后的状态码提示导回游戏内
操作方式：
1.将策划返回回来的文件放入本目录，并且重命名为code.txt
2.将游戏内的状态码文件(ResultStatus.proto)放入本目录
3.执行python daoru.py
4.将生成的ResultStatus_new.proto文件内容拷贝覆盖游戏内的状态码文件内容