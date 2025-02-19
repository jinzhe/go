package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Get(url string, headers ...map[string]string) (string, error) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	if len(headers) > 0 {
		for key, val := range headers[0] {
			request.Header.Set(key, val)
		}
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return string(body), fmt.Errorf("server returned status code: %d", response.StatusCode)
	}

	return string(body), nil
}

func Post(u string, params interface{}, headers ...map[string]string) (string, error) {
	var body io.Reader

	switch v := params.(type) {
	case map[string]string:
		values := url.Values{}
		for key, val := range v {
			values.Set(key, val)
		}
		body = strings.NewReader(values.Encode())

	case string:
		body = strings.NewReader(v)

	case []byte:
		body = bytes.NewReader(v)

	default:
		jsonData, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("序列化参数失败: %w", err)
		}
		body = bytes.NewReader(jsonData)
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       100,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: true,
		},
	}

	request, err := http.NewRequest("POST", u, body)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %w", err)
	}

	if len(headers) > 0 {
		for key, val := range headers[0] {
			request.Header.Set(key, val)
		}
	} else {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	request.Header.Set("Accept", "*/*")
	request.Header.Set("Connection", "keep-alive")

	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}
	defer response.Body.Close()


	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}


	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return string(data), fmt.Errorf("服务器返回错误状态码: %d", response.StatusCode)
	}

	return string(data), nil
}

func PostJSON(u string, params any, headers ...map[string]string) (string, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return "", fmt.Errorf("JSON 序列化失败: %w", err)
	}

	request, err := http.NewRequest("POST", u, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	if len(headers) > 0 {
		for key, val := range headers[0] {
			request.Header.Set(key, val)
		}
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return string(data), fmt.Errorf("服务器返回错误状态码: %d", response.StatusCode)
	}
	return string(data), nil
}
func Put(u string, params interface{}, headers ...map[string]string) (string, error) {
    var body io.Reader

    switch v := params.(type) {
    case map[string]string:
        values := url.Values{}
        for key, val := range v {
            values.Set(key, val)
        }
        body = strings.NewReader(values.Encode())

    case string:
        body = strings.NewReader(v)

    case []byte:
        body = bytes.NewReader(v)

    default:
        jsonData, err := json.Marshal(v)
        if err != nil {
            return "", fmt.Errorf("序列化参数失败: %w", err)
        }
        body = bytes.NewReader(jsonData)
    }

    // 创建请求
    request, err := http.NewRequest("PUT", u, body)
    if err != nil {
        return "", fmt.Errorf("创建请求失败: %w", err)
    }

    // 设置请求头
    if len(headers) > 0 {
        for key, val := range headers[0] {
            request.Header.Set(key, val)
        }
    } else {
        request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    }

    // 创建客户端
    client := &http.Client{
        Timeout: 30 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns:       100,
            IdleConnTimeout:    90 * time.Second,
            DisableCompression: true,
        },
    }

    // 发送请求
    response, err := client.Do(request)
    if err != nil {
        return "", fmt.Errorf("请求失败: %w", err)
    }
    defer response.Body.Close()

    // 读取响应
    data, err := io.ReadAll(response.Body)
    if err != nil {
        return "", fmt.Errorf("读取响应失败: %w", err)
    }

    // 检查状态码
    if response.StatusCode < 200 || response.StatusCode >= 300 {
        return string(data), fmt.Errorf("服务器返回错误状态码: %d", response.StatusCode)
    }

    return string(data), nil
}
func Delete(u string, headers ...map[string]string) (string, error) {
    // 创建请求
    request, err := http.NewRequest("DELETE", u, nil)
    if err != nil {
        return "", fmt.Errorf("创建请求失败: %w", err)
    }

    // 设置请求头
    if len(headers) > 0 {
        for key, val := range headers[0] {
            request.Header.Set(key, val)
        }
    }

    // 创建带超时的客户端
    client := &http.Client{
        Timeout: 30 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns:       100,
            IdleConnTimeout:    90 * time.Second,
            DisableCompression: true,
        },
    }

    // 发送请求
    response, err := client.Do(request)
    if err != nil {
        return "", fmt.Errorf("请求失败: %w", err)
    }
    defer response.Body.Close()

    // 读取响应
    data, err := io.ReadAll(response.Body)
    if err != nil {
        return "", fmt.Errorf("读取响应失败: %w", err)
    }

    // 检查状态码
    if response.StatusCode < 200 || response.StatusCode >= 300 {
        return string(data), fmt.Errorf("服务器返回错误状态码: %d", response.StatusCode)
    }

    return string(data), nil
}
