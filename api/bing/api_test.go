package bing

import (
	"bytes"
	"log/slog"
	"os"
	"testing"
)

func TestGetParams(t *testing.T) {
	// ta := NewBing()
	// // 翻译
	// a := ta.Translate("不知不觉已经沦为弃子的邻居少年，日子倒是依旧过得优哉游哉，成天带着他的贴身丫鬟，在小镇内外逛荡，一年到头游手好闲，也从来不曾为银子发过愁。", "", "")
	// fmt.Printf("%+s", a)

	b := NewBing()
	// 生成语音
	b.ProsodyRate = "+20%"
	str := bytes.Buffer{}
	str.WriteString("不知不觉已经沦为弃子的邻居少年，日子倒是依旧过得优哉游哉，成天带着他的贴身丫鬟，在小镇内外逛荡，一年到头游手好闲，也从来不曾为银子发过愁。")

	audioByte := b.TextToAudio(str)

	if err := os.WriteFile("demo.mp3", audioByte, 0666); err != nil {
		slog.Error(err.Error())
		return
	}
}
