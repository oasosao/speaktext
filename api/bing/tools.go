package bing

import (
	"bytes"
)

type Ops struct {
	VoiceName    string // 文本转语音输出的语音角色
	ProsodyPitch string // 语调
	ProsodyRate  string // 语速
}

// 文本转语音
func TextToAudioBytes(text []byte, ops Ops) []byte {
	strBuf := bytes.Buffer{}
	strBuf.Write(text)

	bing := NewBing()

	if ops.VoiceName != "" {
		bing.VoiceName = ops.VoiceName
	}

	if ops.ProsodyPitch != "" {
		bing.ProsodyPitch = ops.ProsodyPitch
	}

	if ops.ProsodyRate != "" {
		bing.ProsodyRate = ops.ProsodyRate
	}

	return bing.TextToAudio(strBuf)
}

// 翻译
func TranslateTextBytes(text, from, to string) []byte {
	bing := NewBing()
	return bing.Translate(text, from, to)

}
