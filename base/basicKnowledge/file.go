package basicKnowledge

import (
	"fmt"
	"log"
	"os"
)

/*
OpenFile:
OpenFile 既能打开一个已经存在的文件，也能创建并打开一个新文件。
func OpenFile(name string, flag int, perm FileMode) (*File, error)
OpenFile 是一个更一般性的文件打开函数，大多数调用者都应用 Open 或 Create 代替本函数。它会使用指定的选项（如 O_RDONLY 等）、指定的模式（如 0666 等）打开指定名称的文件。如果操作成功，返回的文件对象可用于 I/O。如果出错，错误底层类型是 *PathError。
要打开的文件由参数 name 指定，它可以是绝对路径或相对路径（相对于进程当前工作目录），也可以是一个符号链接（会对其进行解引用）。
位掩码参数 flag 用于指定文件的访问模式，可用的值在 os 中定义为常量（以下值并非所有操作系统都可用）：
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
其中，O_RDONLY、O_WRONLY、O_RDWR 应该只指定一个，剩下的通过 | 操作符来指定。该函数内部会给 flags 加上 syscall.O_CLOEXEC，在 fork 子进程时会关闭通过 OpenFile 打开的文件，即子进程不会重用该文件描述符。
注意：由于历史原因，O_RDONLY | O_WRONLY 并非等于 O_RDWR，它们的值一般是 0、1、2。
位掩码参数 perm 指定了文件的模式和权限位，类型是 os.FileMode，文件模式位常量定义在 os 中：

const (
    // 单字符是被 String 方法用于格式化的属性缩写。
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
    ModeAppend                                     // a: 只能写入，且只能写入到末尾
    ModeExclusive                                  // l: 用于执行
    ModeTemporary                                  // T: 临时文件（非备份文件）
    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
    ModeDevice                                     // D: 设备
    ModeNamedPipe                                  // p: 命名管道（FIFO）
    ModeSocket                                     // S: Unix 域 socket
    ModeSetuid                                     // u: 表示文件具有其创建者用户 id 权限
    ModeSetgid                                     // g: 表示文件具有其创建者组 id 的权限
    ModeCharDevice                                 // c: 字符设备，需已设置 ModeDevice
    ModeSticky                                     // t: 只有 root/ 创建者能删除 / 移动文件

    // 覆盖所有类型位（用于通过 & 获取类型位），对普通文件，所有这些位都不应被设置
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
    ModePerm FileMode = 0777 // 覆盖所有 Unix 权限位（用于通过 & 获取类型位）
)
改变文件时间戳
可以显式改变文件的访问时间和修改时间。
func Chtimes(name string, atime time.Time, mtime time.Time) error
文件属主:每个文件都有一个与之关联的用户 ID（UID）和组 ID（GID），籍此可以判定文件的属主和属组
func Chown(name string, uid, gid int) error
func Lchown(name string, uid, gid int) error
func (f *File) Chown(uid, gid int) error
文件权限
Owner（亦称为 user）：授予文件属主的权限。
Group：授予文件属组成员用户的权限。
Other：授予其他用户的权限。
可为每一类用户授予的权限如下：
Read：可阅读文件的内容。
Write：可更改文件的内容。
Execute：可以执行文件（如程序或脚本）。
Unix 中表示：rwxrwxrwx
目录权限
目录与文件拥有相同的权限方案，只是对 3 种权限的含义另有所指。
读权限：可列出（比如，通过 ls 命令）目录之下的内容（即目录下的文件名）。
写权限：可在目录内创建、删除文件。注意，要删除文件，对文件本身无需有任何权限。
可执行权限：可访问目录中的文件。因此，有时也将对目录的执行权限称为 search（搜索）权限。
访问文件时，需要拥有对路径名所列所有目录的执行权限。例如，想读取文件 /home/studygolang/abc，
则需拥有对目录 /、/home 以及 /home/studygolang 的执行权限（还要有对文件 abc 自身的读权限）。
在文件相关操作报错时，可以通过 os.IsPermission 检查是否是权限的问题。
func IsPermission(err error) bool

*/

func OpenFile() {
	file, err := os.Open("file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//如果我们要获取 FileInfo 接口没法直接返回的信息，比如想获取文件的上次访问时间，示例如下：
	fileInfo, err := os.Stat("test.log")
	if err != nil {
		log.Fatal(err)
	}

	//sys := fileInfo.Sys()
	//stat := sys.(*syscall.Stat_t)
	//fmt.Println(time.Unix(stat.Atimespec.Unix()))
	fmt.Println(fileInfo.Name())
}

func chMod() {
	file, err := os.Create("golang.txt")
	if err != nil {
		log.Fatal("err:", err)
	}
	defer file.Close()
	fileMode := getFileMode(file)
	log.Println("file mode:", fileMode)
	file.Chmod(fileMode | os.ModeSticky)
	log.Println("change after, file mode:", getFileMode(file))
	// Output:
	// 2016/06/18 15:59:06 file mode: -rw-rw-r--
	// 2016/06/18 15:59:06 change after, file mode: trw-rw-r--
	// ls -l 看到的 golang.tx 是：-rw-rw-r-T
	// 当然这里是给文件设置了 sticky 位，对权限不起作用。系统会忽略它。
}
func getFileMode(file *os.File) os.FileMode {

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("file stat error:", err)
	}
	return fileInfo.Mode()
}
