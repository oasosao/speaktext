package utils

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/tosone/minimp3"
)

// 播放MP3音频
func PlayMp3(path string) {

	if strings.TrimSpace(path) == "" {
		slog.Error("path is empty")
		return
	}

	f, err := os.Open(path)
	if err != nil {
		slog.Error("Error opening audio file", err)
		return
	}
	defer f.Close()

	dec, err := minimp3.NewDecoder(f)
	if err != nil {
		slog.Error("Error decoding audio file", err)
		return
	}
	<-dec.Started()

	defer dec.Close()

	var op = oto.NewContextOptions{}

	op.SampleRate = dec.SampleRate
	op.ChannelCount = dec.Channels
	op.Format = oto.FormatSignedInt16LE

	// 创建一个 Oto 上下文
	otoCtx, readyChan, err := oto.NewContext(&op)
	if err != nil {
		slog.Error("oto.NewContext failed: " + err.Error())
		return
	}
	<-readyChan

	player := otoCtx.NewPlayer(dec)

	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	err = player.Close()
	if err != nil {
		slog.Error("player.Close failed: " + err.Error())
		return
	}
}
