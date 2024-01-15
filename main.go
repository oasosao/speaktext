package main

import (
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"

	"github.com/oasosao/speaktext/api/bing"
	"github.com/oasosao/speaktext/utils"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

const voiceStr = `
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
`

func main() {
	var text, cacheDir, voice, pitch, rate string
	var autoPlay bool
	flag.StringVar(&text, "text", "", "需要转语音的文本可以是文本文件 如 text='翻译文本' 或者 text='./text.txt'")
	flag.StringVar(&cacheDir, "cache", "./SpeakCache", "保存音频目录路径")
	flag.StringVar(&voice, "voice", "zh-CN-YunXiNeural", fmt.Sprintf(" eg: -voice='zh-CN-YunyangNeural' 声音角色如下: %s", voiceStr))
	flag.StringVar(&pitch, "pitch", "", "语调 如: -pitch=50% 或者 -pitch=-50%")
	flag.StringVar(&rate, "rate", "", "语速 如: -rate=20% 或者 -rate=-50%")
	flag.BoolVar(&autoPlay, "autoPlay", true, "是否自动播放音频")
	flag.Parse()

	println()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		return
	}

	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		slog.Error("创建目录失败", err)
		return
	}

	var zText []byte

	if _, err := os.Stat(text); os.IsNotExist(err) {
		zText = []byte(text)
	} else {
		fb, err := os.ReadFile(text)
		if err != nil {
			slog.Error("打开文件失败", err)
			return
		}
		zText = fb
	}

	zFB := bytes.Buffer{}
	zFB.Write(zText)
	zFB.WriteString(voice)
	zFB.WriteString(pitch)
	zFB.WriteString(rate)

	fileName := crateFileName(zFB.Bytes(), cacheDir, voice)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if !saveAudio(zText, fileName, bing.Ops{VoiceName: voice, ProsodyPitch: pitch, ProsodyRate: rate}) {
			slog.Error("保存音频失败", err)
			return
		}
		fmt.Println("音频已经保存在:", fileName)
	} else {
		fmt.Println("已经音频文件存在:", fileName)
	}
	println()

	if autoPlay {
		utils.PlayMp3(fileName)
	}
}

func saveAudio(text []byte, fileName string, ops bing.Ops) bool {
	audioBytes := bing.TextToAudioBytes(text, ops)
	if audioBytes == nil {
		slog.Error("转语音失败, 数据为空。")
		return false
	}

	if err := os.WriteFile(fileName, audioBytes, 0775); err != nil {
		slog.Error("转语音失败", err)
		return false
	}

	return true
}

func crateFileName(text []byte, saveDir, voice string) string {
	hash := md5.New()
	hash.Write(text)
	fileName := path.Join(saveDir, fmt.Sprintf("%s_%x.mp3", voice, hash.Sum(nil)))
	return fileName
}
