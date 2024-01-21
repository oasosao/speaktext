package bing

import (
	"bytes"
	"fmt"
	"html"
	"log/slog"
	"net/url"
	"regexp"
	"strings"

	"github.com/oasosao/speaktext/utils"
)

type Options struct {
	// 请求API所需的参数
	IG    string
	Token string
	Key   string
	IID   string

	// 生成的语音配置项
	VoiceName    string // 语音输出的语音角色
	ProsodyPitch string // 语音输出基线音节
	ProsodyRate  string // 语音输出的讲出速率

	// 翻译
	FromLang string
	ToLang   string
}

var (
	ParamsApi       = "https://cn.bing.com/translator"   // 获取参数的API
	AudioApi        = "https://cn.bing.com/tfettts"      // 转语音API
	TranslateApi    = "https://cn.bing.com/ttranslatev3" // 翻译API
	TranslateApiExt = "https://cn.bing.com/tlookupv3"    // 翻译扩展API
	UserAgent       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
	ReqFormType     = "application/x-www-form-urlencoded"

	// 定义从页面数据中获取IG参数值的正则
	IgReg = regexp.MustCompile("IG:\"[0-9a-zA-Z]*?\"")
	RegM  = regexp.MustCompile("\".*\"")

	//  定义从页面数据中获取token 和 key 参数值的正则
	TkDataReg = regexp.MustCompile("(params_AbusePreventionHelper).*3600000];")
	TkReg     = regexp.MustCompile(`\[.*\]`)

	// 设置 post 请求header
	postHeaderMap = map[string]string{
		"User-Agent":   UserAgent,
		"Content-Type": ReqFormType,
	}

	// 设置 get 请求header
	getHeaderMap = map[string]string{"User-Agent": UserAgent}
)

// 获取所需参数
func NewBing() *Options {
	resBytes := utils.NewReq("GET", ParamsApi, nil, getHeaderMap)
	if resBytes == nil {
		slog.Error("请求出错")
		return nil
	}

	options := &Options{}

	// 从页面中获取IG参数的值
	igStr := IgReg.Find(resBytes)

	options.IG = strings.Trim(RegM.FindString(string(igStr)), "\"")
	options.IID = "translator.5024"

	tkData := TkDataReg.Find(resBytes)
	tkStr := TkReg.FindString(string(tkData))

	value := strings.TrimSuffix(strings.TrimPrefix(tkStr, "["), "]")
	tkArr := strings.Split(value, ",")
	if len(tkArr) < 3 {
		slog.Error("获取到的数据不符合要求", "tkStr", tkStr)
		return nil
	}

	options.Key = tkArr[0]
	options.Token = strings.Trim(RegM.FindString(tkArr[1]), "\"")

	// 文本转语音参数
	options.VoiceName = "zh-CN-YunXiNeural" // 文本转语音输出的语音角色
	options.ProsodyPitch = "0%"             // 指示文本的基线音节。
	options.ProsodyRate = "0%"              // 指示文本的讲出速率。

	// 翻译
	options.FromLang = "zh-Hans"
	options.ToLang = "en"

	return options
}

// 文本转语音
func (b *Options) TextToAudio(text bytes.Buffer) (audioByte []byte) {

	cont := strings.ReplaceAll(text.String(), "\n\n", "\n")
	cont = strings.ReplaceAll(cont, "\r\n", "\n")
	strArr := strings.Split(cont, "\n")

	audioBytes := bytes.Buffer{}

	for _, strText := range strArr {

		if strings.TrimSpace(strText) == "" {
			continue
		}

		strText = html.EscapeString(strText)

		ssmlText := fmt.Sprintf(`<speak version="1.0" xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="https://www.w3.org/2001/mstts" xml:lang="zh-CN"><voice name="%s"><prosody pitch="%s" rate="%s">%s</prosody></voice></speak>`,
			b.VoiceName,
			b.ProsodyPitch,
			b.ProsodyRate,
			strText,
		)

		if _, err := audioBytes.Write(b.GetTextAudioData(ssmlText)); err != nil {
			slog.Error(err.Error())
			return
		}
	}

	return audioBytes.Bytes()
}

// 获取文本生成的音频数据
func (b *Options) GetTextAudioData(ssml string) []byte {

	formValues := url.Values{}
	formValues.Set("ssml", ssml)
	formValues.Set("token", b.Token)
	formValues.Set("key", b.Key)

	reqBody := strings.NewReader(formValues.Encode())

	reqUrl, err := url.Parse(AudioApi)
	if err != nil {
		slog.Error("解析URL出错", "err", err)
		return nil
	}

	tq := reqUrl.Query()
	tq.Set("isVertical", "1")
	tq.Set("IG", b.IG)
	tq.Set("IID", b.IID+".2")

	reqUrl.RawQuery = tq.Encode()
	bodyByte := utils.NewReq("POST", reqUrl.String(), reqBody, postHeaderMap)

	return bodyByte
}

// 翻译文本
func (b *Options) Translate(text, from, to string) []byte {
	buf := bytes.Buffer{}
	t1 := b.transBase(TranslateApi, text, from, to)
	if t1 != nil {
		buf.Write(t1)
	}

	// 翻译扩展
	// buf.WriteString("\n")
	// b.IID = b.IID + ".5"
	// t2 := b.transBase(TranslateApiExt, text, from, to)
	// if t2 != nil {
	// 	buf.Write(t2)
	// }
	return buf.Bytes()
}

// 获取翻译数据
func (b *Options) transBase(uri, text, fromLang, to string) []byte {

	if fromLang == "" {
		b.FromLang = "zh-Hans"
	}

	if to == "" {
		b.ToLang = "en"
	}

	formValues := url.Values{
		"text":     {text},
		"from":     {b.FromLang},
		"fromLang": {b.FromLang},
		"to":       {b.ToLang},
		"token":    {b.Token},
		"key":      {b.Key},
	}

	reqBody := strings.NewReader(formValues.Encode())

	reqUrl, err := url.Parse(uri)
	if err != nil {
		slog.Error("解析URL出错", "err", err)
		return nil
	}

	urlParams := url.Values{
		"isVertical": {"1"},
		"IG":         {b.IG},
		"IID":        {b.IID},
	}

	reqUrl.RawQuery = urlParams.Encode()

	resBody := utils.NewReq("POST", reqUrl.String(), reqBody, postHeaderMap)

	return resBody
}
