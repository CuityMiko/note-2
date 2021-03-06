### 语法： grep 参数 文件
```
P/E    解释模式作为Perl的正则表达式匹配,P-perl正则，E-egrep命令       cat abc.txt | grep -P 'ABC(DEF)G[a-z]{0,3}XYZ'
o    仅显示匹配的行，即匹配的内容，如搜索 abc ，匹配成功则只显示abc内容。cat abc.txt | grep -o 'abc'
i    忽略大小写         cat abc.txt | grep -i | grep 'Abc' 
n    输出结果时同时输出该行的行号  cat abc.txt | grep -n 'abc'
l    从多文件查找时，只输出找到匹配内容的文件名	cat abc.txt | grep -l 'abc'
h    从多个文件查找时，只输出匹配的内容，不显示文件名  cat abc.txt | grep -h 'abc'
c    只输出匹配内容的总行数	cat abc.txt | grep -c 'abc'
v    反向查找，输出匹配内容以外的内容（未匹配的到的内容）		cat abc.txt | grep -v 'abc'
R/r  按目录递归查找 cat abc.txt | grep -r 'abc' * 或  /home/work/code grep -R --include="*.conf" "关键字" /home/wiki 
F    将范本样式视为固定字符串的列表。	cat abc.txt | grep -F 'abc'

grep -r -P -n --exclude='a.php' 'abc|1242' work   #在搜索结果中排除a.php文件检索,这里的文件名可以用正则
即：包含用 --include='xxx.file'参数，不包含用exclude='xxx.file参数，需配合目录使用

注：不加参数默认是输出匹配的行，grep支持正则表达式(加参数P即可)
```
### 目录文件中查找
	grep -R --include="*.conf" "关键字" /home/wiki 
	
### 查找当前目录下的所有china关键字
	grep -R 'china' * 
	grep 'china' *
	grep -R 'china' /home/work/code
	grep 'china' /home/liguibing/php

### 在目录中递归搜索关键字
	grep -R "关键字" /home/wiki/
	grep -F '77308535' commitsql.20140820 | grep -F '2014-08-20 14'
### 搜索app关键词
	grep -o -P "APP" /home/work/liguibing/test/a.txt 
	grep -in "APP" /home/work/liguibing/test/a.txt 
 
### 多次查找，即在第一次结果中继续查找
	grep -in "APP" a.txt | grep -i 'abc'  

### 显示tmp目录下包含abc关键字的所有文件名
	grep -l "abc" /home/users/liguibing/tmp/*  

### 多文件查找,a，b文件中包含abc关键字
	grep 'abc' b.txt c.txt         

### 行首匹配
	grep "^22200632206206" abc.txt  匹配行首以 22200632206206 的行
    
### 匹配行首匹配第几行
	grep "^....AB" abc.txt  匹配行首第5、6列的值为：A B 的行

### 行尾匹配：
	grep "\.7$" abc.txt  匹配行尾为 .7 的所有行
### 配合常用的正则表达式查找,范围匹配
	grep '/9[3-5]' a.txt  匹配／93之后的字符

### 次数匹配 C至少出现两次的行
	grep 'c\{2,\}' a.txt

### 查找字母e重复出现2次且以p结尾的行
	grep 'e\{2\}p' a.txt
	
### 次数匹配 C出现两次10次的行
	grep 'c\{2,10\}' a.txt

### 匹配.txt结尾的文件
	grep '\.txt$'

### 或匹配，用参数E
	grep -E 'abc|xy'  a.txt

### 匹配空行数
	grep -c '^$' abc.txt

grep有一个很大的家族：扩展grep、fgrep、egrep、agrep等

