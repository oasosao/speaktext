# 文本转语音 

这是一个利用微软翻译API的文本转语音的免费go程序

### 程序下载地址
- **Github 下载：**
<https://github.com/oasosao/speaktext/releases/tag/v0.0.1-alpha>。
- **Gitee 下载：**  <https://gitee.com/oasosao/speaktext/releases/tag/v0.0.1-alpha>

### 源码安装

```sh
$ git clone github.com/oasosao/speaktext.git
$ cd ./speaktext
$ go mod tidy
$ go build
$ ./speaktext -text="你好，世界！"
```


## 使用方法:

```sh
$ speaktext -text="你好，世界！"  -rate="20%"  -voice="zh-CN-YunxiNeural"
$ # 创建一个文本文件， speak.txt 在文本中写上一些内容。
$ speaktext -text="./speak.txt" -rate="20%" -voice="zh-CN-YunxiNeural"
$
```

**参数参考**

```sh
-autoPlay
    	是否自动播放音频 (default true)
  -cache string
    	保存音频目录路径 (default "./SpeakCache")
  -pitch string
    	语调 如: -pitch=50% 或者 -pitch=-50%
  -rate string
    	语速 如: -rate=20% 或者 -rate=-50%
  -text string
    	需要转语音的文本可以是文本文件 如 text='翻译文本' 或者 text='./text.txt'
  -voice string
    	 eg: -voice='zh-CN-YunyangNeural' 声音角色如下: 
    	zh-CN-XiaoxiaoNeural
    	zh-CN-XiaoyiNeural
    	zh-CN-YunjianNeural
    	zh-CN-YunxiaNeural
    	zh-CN-YunxiNeural
    	zh-CN-YunyangNeural
    	zh-CN-liaoning-XiaobeiNeural
    	zh-CN-shaanxi-XiaoniNeural
    	zh-HK-HiuMaanNeural
    	zh-HK-WanLungNeural
    	zh-HK-HiuGaaiNeural
    	zh-TW-HsiaoChenNeural
    	zh-TW-YunJheNeural
    	zh-TW-HsiaoYuNeural
    	 (default "zh-CN-YunXiNeural")
```
