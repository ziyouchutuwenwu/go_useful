package encoder

import (
	"github.com/axgle/mahonia"
)

func ConvertString(content string, srcEncoding string, destEncoding string) string {
	srcEncoder := mahonia.NewDecoder(srcEncoding)
	contentResult := srcEncoder.ConvertString(content)

	destEncoder := mahonia.NewDecoder(destEncoding)
	_, translatedContent, _ := destEncoder.Translate([]byte(contentResult), true)
	result := string(translatedContent)

	return result
}