package httpKit

import "net/http"

// RespondIcon
/**
 * 将 二进制形式的icon图片 响应给前端.
 *
 * @param writer e.g. Gin中的ctx.Writer
 */
func RespondIcon(writer http.ResponseWriter, iconData []byte) (int, error) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "image/x-icon")
	length, err := writer.Write(iconData)

	//writer.(http.Flusher).Flush()
	if flusher, ok := writer.(http.Flusher); ok {
		flusher.Flush()
	}

	return length, err
}

func RespondHtml(statusCode int, writer http.ResponseWriter, htmlData []byte) (int, error) {
	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "text/html")
	writer.Header().Add("Content-Type", "charset=utf-8")
	length, err := writer.Write(htmlData)

	//writer.(http.Flusher).Flush()
	if flusher, ok := writer.(http.Flusher); ok {
		flusher.Flush()
	}

	return length, err
}
