# 文本转语音 

这是一个利用微软翻译API的文本转语音的免费go程序

使用方法:

```sh
$ go install github.com/oasosao/speaktext@latest
$ speaktext -text "你好，世界！" 
```

### 参数参考

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
