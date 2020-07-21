package basicKnowledge

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
重命名: os.Rename("1.go", "2.go")
是否相同:
 f1, _ := os.Stat("1.go")
 f2, _ := os.Stat("21.go")
 os.SameFile(f1, f2)
设置环境变量:
 os.Setenv("WD_PATH", "D:/golang")
返回你本地的系统temp目录:
 dir, _ := os.Getwd()
 path, _ := ioutil.TempDir(dir, "tmp")
 //这个返回的是系统temp
 temp := os.TempDir()
改变文件的f.Size()这个就改变了文件内容的长度
os.Truncate("1.go", 10)

os.Getwd() //当前的目录
os.Getenv("GOPATH") //环境变量
fmt.Println(os.Getegid())      windows -1  linux  0     //调用者的group的id
fmt.Println(os.Geteuid())     windows -1  linux  0     //用户的uid
fmt.Println(os.Getgid())      windows -1  linux  0     //调用者的gid的id
g, _ := os.Getgroups()
fmt.Println(g)                windows []  linux  []    //返回的是一个[]int的切片 显示调用者属于组的一系列id
fmt.Println(os.Getpagesize())  windows 4096linux  4096  //windows里边叫做虚拟内存 linux里边叫做swap
fmt.Println(os.Getppid())      windows -1  linux  8621  //调用者的组的进程id
fmt.Println(os.Getuid())    windows -1  linux  0  //调用者的数字用户id

os.Chdir("D:/test/src") //切换目录
获取文件的信息
 os.Stat("widuu.go")
os.Chmod()这个函数的原型是func Chmod(name string, mode FileMode) error改变文件的属性 譬如读写，linux上的0755这样大家可以理解了吧
os.Chtime()这个包，函数的原形是func Chtimes(name string, atime time.Time, mtime time.Time) error
输入string的文件的名称 访问时间 创建时间 返回的是error接口信息
os.Environ()的作用是获取系统的环境变量，函数原形是func Environ() []string返回是环境变量的[]string切片，说道这个就���和其他的一起说明了，那就是os.ClearEnv()清空环境变量
os.Exit()就是中断程序返回自定义的int类型的代码，函数运行是func Exit(code int)输入一个int的值就可以了
)函数os.Expand()这个其实就是一个回调函数替换的方法，函数的原形是func Expand(s string, mapping func(string) string) string 输入的是一个string。对应的是func(string)string的替换字符串的方法，如果没有字符就替换为空
mapping := func(s string) string {
  m := map[string]string{"widuu": "www.1.net", "xiaowei": "widuu"}
  return m[s]
 }
 data := "hello $xiaowei blog address $widuu"
 fmt.Printf("%s", os.Expand(data, mapping)) ////输出hello widuu blog address www.1.net}
)os.ExpandEnv()把字符串的s替换成环境变量的内容，函数的原形是func ExpandEnv(s string) string，输入的当然是要替换的字符，输出的当然还是字符串了
data := "GOBIN PATH $GOBIN"
 fmt.Println(os.ExpandEnv(data)) //输出我本地的环境变量的GOBIN的地址GOBIN PATH C:\Go\bin
os.Hostname()这个函数看字面的思意就懂了，是返回主机的HostName(),函数的原形是func Hostname() (name string, err error)返回主机名字和一个error的接口信息
 data, _ := os.Hostname()
 fmt.Println(data) //我是windows环境下返回我的win的主机名 eleven
gojson:

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
//追加内容到文件末尾的办法
//Seek()查到文件末尾的偏移量
//WriteAt()则从偏移量开始写入
func appendToFile(fileName string, content string) error{
	// 以只写的模式，打开文件
	f,err := os.OpenFile(fileName, os.O_WRONLY,0644)
	if err != nil{
		fmt.Println("create file failed error :", err.Error())
		return err
	}
	defer f.Close()
	// 查找文件末尾的偏移量
	n,_:= f.Seek(0,os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	_,err =  f.WriteAt([]byte(content), n)
	return err

}

func WriteFile(){
	f, err := os.OpenFile("file2.txt",os.O_RDWR | os.O_CREATE | os.O_APPEND, 0x644)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	wint,err := f.WriteString("hello world")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%d\n",wint)
	_,err = f.Seek(0,0)
	if err != nil{
		panic(err)
	}
	bs:=make([]byte,100)
	rint, err := f.Read(bs)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%d, %s \n", rint,bs)
}

func TestImage(){
	f1,err := os.Open("1.jpg")
	if err != nil{
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Open("2.jpg")
	if err != nil{
		panic(err)
	}
	defer f2.Close()
	f3,err := os.Open("3.jpg")
	if err != nil{
		panic(err)
	}
	defer f3.Close()
	m1,err := jpeg.Decode(f1)
	if err!=nil{
		panic(err)
	}
	bounds := m1.Bounds()
	m2,err:= jpeg.Decode(f2)
	if err !=nil{
		panic(err)
	}
	m:= image.NewRGBA(bounds)
	white := color.RGBA{255,255,255,255}
	draw.Draw(m,bounds,&image.Uniform{white},image.ZP,draw.Src)
	draw.Draw(m,bounds,m1,image.ZP,draw.Src)
	draw.Draw(m,image.Rect(100,200,300,600),m2,image.Pt(250,60),draw.Src)
	err = jpeg.Encode(f3,m,&jpeg.Options{90})
	if err != nil{
		panic(err)
	}
	fmt.Printf("ok\n")
}

func OsRead(){
	b := make([]byte,100) //设置读取的字节数
	f,_:= os.Open("base.go")
	n,_:= f.Read(b)
	fmt.Println(n)
	fmt.Println(string(b[:n]))//输出内容 为什么是n而不直接输入100呢？底层这样实现的
	/*
	  n, e := f.read(b)
	     if n < 0 {
	       n = 0
	     }
	  if n == 0 && len(b) > 0 && e == nil {
	     return 0, io.EOF
	    }
	*/
	//所以字节不足100就读取n
}

// 加入下标，可以自定义读取多少
func OsReadAt(){
	f,_:= os.Open("base.go")
	b := make([]byte,20)
	n,_:= f.ReadAt(b,15)
	fmt.Println(n)
	fmt.Println(string(b[:n]))
}

//打开一个文件夹，然后设置读取文件夹文件的个数，返回的是文件的fileinfo信息
func OsReaddir()  {
	f,err:= os.Open("src") //打开一个目录
	if err != nil{
		fmt.Println(err)
	}
	defer  f.Close()
	ff,_:= f.Readdir(10)//设置读取的数量 <=0是读取所有的文件 返回的[]fileinfo
	for i, fi := range ff {
		fmt.Printf("filename %d: %+v\n",i,fi.Name()) //输出文件的名称
	}
}

//返回的是文件名 []string的slice
func OsReaddirnames()  {
	f, _ := os.Open("bin")
	names, err := f.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
	}
	for i, name := range names {
		fmt.Printf("filename %d: %s\n", i, name)
	}
}

//(f *File).Seek()这个函数大家一看就懂了，就是偏移指针的地址，函数的原型是
//func (f *File) Seek(offset int64, whence int) (ret int64, err error)
//其中offset是文件指针的位置 whence为0时代表相对文件开始的位置，
//1代表相对当前位置，2代表相对文件结尾的位置 ret返回的是现在指针的位置
func OsReadSeek(){
	b := make([]byte, 10)
	f, _ := os.Open("1.go")
	defer f.Close()
	f.Seek(1, 0)    //相当于开始位置偏移1
	n, _ := f.Read(b)
	fmt.Println(string(b[:n]))  //原字符package 输出ackage
}

//返回的是n写入的字节数
func OsReadAppend() {
	f, _ := os.OpenFile("1.go", os.O_RDWR|os.O_APPEND, 0755) //以追加和读写的方式去打开文件
	n, _ := f.Write([]byte("helloword"))                     //我们写入hellword
	fmt.Println(n)                                           //打印写入的字节数
	b := make([]byte, 20)
	f.Seek(0, 0)            //指针返回到0
	data, _ := f.Read(b)
	fmt.Println(string(b[:data]))        //输出了packagehelloword
}

//在偏移位置多少的地方写入
func OsWriteAt() {
	f, _ := os.OpenFile("1.go", os.O_RDWR, os.ModePerm)
	f.WriteAt([]byte("widuu"), 10) //在偏移10的地方写入
	b := make([]byte, 20)
	d, _ := f.ReadAt(b, 10)    //偏移10的地方开始读取
	fmt.Println(string(b[:d])) //widuudhellowordhello
}

//写入字符串函数
func OsWriteString() {
	f, _ := os.OpenFile("2.go", os.O_RDWR, os.ModePerm)
	n, _ := f.WriteString("hello word widuu") //写入字符串
	fmt.Println(n)
	b := make([]byte, n)
	f.Seek(0, 0)    //一定要把偏移地址归0否则就一直在写入的结尾处
	c, _ := f.Read(b)
	fmt.Println(string(b[:c])) //返回hello word widuu
}

//创建目录
func OsCreateFolder() {
	var path string
	if os.IsPathSeparator('\\') {  //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	fmt.Println(path)
	dir, _ := os.Getwd()  //当前的目录
	err := os.Mkdir(dir+path+"md", os.ModePerm)  //在当前目录下生成md目录
	//err := os.MkdirAll(dir+"/a/b/c", os.ModePerm)  //生成多级目录
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建目录" + dir + path + "md成功")
}
// if err := compress(`gopkg`, `gopkg.zip`); err != nil {
//    fmt.Println(err)
//  }
// 参数frm可以是文件或目录，不会给dst添加.zip扩展名
func compress(frm, dst string) error {
	buf := bytes.NewBuffer(make([]byte, 0, 10*1024*1024)) // 创建一个读写缓冲
	myzip := zip.NewWriter(buf)              // 用压缩器包装该缓冲
	// 用Walk方法来将所有目录下的文件写入zip
	err := filepath.Walk(frm, func(path string, info os.FileInfo, err error) error {
		var file []byte
		if err != nil {
			return filepath.SkipDir
		}
		header, err := zip.FileInfoHeader(info) // 转换为zip格式的文件信息
		if err != nil {
			return filepath.SkipDir
		}
		header.Name, _ = filepath.Rel(filepath.Dir(frm), path)
		if !info.IsDir() {
			// 确定采用的压缩算法（这个是内建注册的deflate）
			header.Method = 8
			file, err = ioutil.ReadFile(path) // 获取文件内容
			if err != nil {
				return filepath.SkipDir
			}
		} else {
			file = nil
		}
		// 上面的部分如果出错都返回filepath.SkipDir
		// 下面的部分如果出错都直接返回该错误
		// 目的是尽可能的压缩目录下的文件，同时保证zip文件格式正确
		w, err := myzip.CreateHeader(header) // 创建一条记录并写入文件信息
		if err != nil {
			return err
		}
		_, err = w.Write(file) // 非目录文件会写入数据，目录不会写入数据
		if err != nil {    // 因为目录的内容可能会修改
			return err     // 最关键的是我不知道咋获得目录文件的内容
		}
		return nil
	})
	if err != nil {
		return err
	}
	myzip.Close()        // 关闭压缩器，让压缩器缓冲中的数据写入buf
	file, err := os.Create(dst) // 建立zip文件
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = buf.WriteTo(file) // 将buf中的数据写入文件
	if err != nil {
		return err
	}
	return nil
}

const (
	DataRoot   = "./tmp/" // 存放封面图的根目录
	TimeoutLimit = 10    // 设置超时时间
	PageUrl   = "http://api.lovebizhi.com/macos_v4.php?a=category&spdy=1&tid=3&order=hot&color_id=3&device=105&uuid=436e4ddc389027ba3aef863a27f6e6f9&mode=0&retina=0&client_id=1008&device_id=31547324&model_id=105&size_id=0&channel_id=70001&screen_width=1920&screen_height=1200&bizhi_width=1920&bizhi_height=1200&version_code=19&language=zh-Hans&jailbreak=0&mac=&p={pid}"
)

// 壁纸类型，有编号，长宽和URL
type Wallpaper struct {
	Pid   int
	Url   string
	Width  int
	Height  int
}
// 将图片下载并保存到本地
func SaveImage(paper *Wallpaper) {
	res, err := http.Get(paper.Url)
	defer res.Body.Close()
	if err != nil {
		fmt.Printf("%d HTTP ERROR:%s", paper.Pid, err)
		return
	}
	//按分辨率目录保存图片
	Dirname := DataRoot + strconv.Itoa(paper.Width) + "x" + strconv.Itoa(paper.Height) + "/"
	if ! isDirExist(Dirname) {
		os.Mkdir(Dirname, 0755);
		fmt.Printf("dir %s created\n", Dirname)
	}
	//根据URL文件名创建文件
	filename := filepath.Base(paper.Url)
	dst, err := os.Create(Dirname + filename)
	if err != nil {
		fmt.Printf("%d HTTP ERROR:%s", paper.Pid, err)
		return
	}
	// 写入文件
	io.Copy(dst, res.Body)
}
func isDirExist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}
func TestDownloadImage() {
	//检查并创建临时目录
	if ! isDirExist(DataRoot) {
		os.Mkdir(DataRoot, 0755);
		fmt.Printf("dir %s created", DataRoot)
	}
	//生成一个数据序列，用来获取分页
	pow := make([]int, 2)
	for i := range pow {
		if (i > 0) {
			url := strings.Replace(PageUrl, "{pid}", strconv.Itoa(i), -1);
			fmt.Println(i, url);
			response, err := http.Get(url)
			if( err != nil) {
				fmt.Println(err)
				continue
			}
			body, _ := ioutil.ReadAll(response.Body)
			js, err := simplejson.NewJson(body)
			//遍历data下的所有数据
			data := js.Get("data").MustArray()
			for _, v := range data {
				v := v.(map[string]interface{})
				for kk, vv := range v {
					if(kk == "file_id") {
						//这里 vv 是一个[]interface{} json.Number，不知道怎么取出值，这里用了比较傻的Sprintf
						vv := fmt.Sprintf("%s", vv)
						imgid,_ := strconv.Atoi(vv)
						url := fmt.Sprintf("http://s.qdcdn.com/c/%d,1920,1200.jpg", imgid)
						fmt.Println(kk, imgid, url);
						paper := &Wallpaper{imgid, url, 1920, 1200}
						SaveImage(paper);
					}
				}
			}
		}
	}
	fmt.Println("oh yes, all job done.")
}