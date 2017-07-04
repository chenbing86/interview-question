# 面试题
### 魔镜面试题,答题源代码
    开发环境 : Mac
    包管理工具 : gb 
    引用第三方包 : 
    	字符串通配符配 : https://github.com/ryanuber/go-glob

### 引入包
    gb vendor fetch https://github.com/ryanuber/go-glob
### 编译
    gb build
### 查看参数
    ./bin/main --help
运行结果

    Usage of ./bin/main:
    -f string 
    	忽略哪些目录、文件，支持通配符
    -gr int
    	并发数量.默认10 (default 10)
    -o string
    	输出目录树的文件,默认tree.txt (default "tree.txt")
    -p string
    	遍历的文件夹路径，默认为当前路径 (default "./")
	   
### 运行例子
例1 : 

    ./bin/main -p /Users/chenbing/Downloads/html
运行结果

    appstyle.css,4d87b3fc257dbff0c83b86655d8af20b1e322446,4802 
    .DS_Store,1ee6a9948e5245be2cc5351140d848d0c6093e46,12292 
    bank.css,70e4ef3301d52d83e1256d4693824ee58a9f9b8d,7829 
    bootstrap-datepicker3.min.css,29cb7d62899700071eae93f717fffd18801749b7,21268 
    bootstrap.css.map,16506513c5f3d95982e73fb6820cf3c4c58d6897,390518 
    bootstrap-theme.css,ecf2245dd39ea3b17ccb4bea42cda46356376078,26132 
    bootstrap.min.css,08df9a96752852f2cbd310c30facd934e348c2c5,122540 
    icomoon.svg,528a622257ddc1859a695a738b54bb9317c0a3e3,47276 
    bootstrap-datetimepicker.min.css,1cd861cc00e5b380a3665fb05b1cef97ad7003e7,11258 
    bootstrap.css,6987e3bdad7a3a5d143ddf2453e29782dbd99c29,147430 
    datepicker.css,57b4adee5da840195be58fb392d9eb3bae8d3e05,10225 
    fancybox.css,8f7895c622430e104f710a47fa1ee59a0ca183d9,8578 
	
例2 :

    ./bin/main -p /Users/chenbing/Downloads/html -o tree1.txt -f "*/images/*|*.css" -gr 5
运行结果

    bootstrap-theme.css.map,abc2e3a1163de52378c71bb92d007c22abb0a393,47721 
    icomoon.eot,c255d348d0c6015713d63b3d48e197fa49ba812b,15648 
    bootstrap.css.map,16506513c5f3d95982e73fb6820cf3c4c58d6897,390518 
    icomoon.svg,528a622257ddc1859a695a738b54bb9317c0a3e3,47276 
    .DS_Store,7abb0403e070175ce244dfe5dc991654a3c8a47c,6148 
    icomoon.woff,60051617cde2e0f29152b5e0a386a1c8e116edb5,15560 
    icomoon.ttf,b22a8655f9b05ce30ff204a58863eb1cd16e7134,15484 
    glyphicons-halflings-regular.eot,86b6f62b7853e67d3e635f6512a5a5efc58ea3c3,20127 
    glyphicons-halflings-regular.woff,278e49a86e634da6f2a02f3b47dd9d2a8f26210f,23424 
    .DS_Store,8dbf5e77bec92f9f9ebefe3a8194c37216fd225a,6148 
    glyphicons-halflings-regular.ttf,44bc1850f570972267b169ae18f1cb06b611ffa2,45404 
    glyphicons-halflings-regular.woff2,ca35b697d99cae4d1b60f2d60fcd37771987eb07,18028 
    glyphicons-halflings-regular.svg,de51a8494180a6db074af2dee2383f0a363c5b08,108738 
    chrome.jpg,be203de4964425856c0a837b3532e9359b65421b,11231     