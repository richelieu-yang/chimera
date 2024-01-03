## 在线Aes加密/解密
https://lovefree.cc/aes

## 加密模式
加密模式（ECB、CBC、CFB）

## 填充方式
填充方式（ZeroPadding、PKCS5Padding）[main.go](..%2F..%2F..%2F..%2F..%2FDownloads%2Fmain.go)

## key的长度
16字节: AES-128 
24字节: AES-192
32字节: AES-256

## !!!: 通用的 AES/CBC/PKCS7 加解密（Golang、Java、js）
#### Golang demo
```go
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// AES-CBC-PKCS7 加解密，支持传参 iv
// key 必须是 16 (AES-128)、24 (AES-192) 或 32 (AES-256) 字节的 AES 密钥
// iv 必须是 16 字节的初始化向量

// PKCS7Padding 填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding 去除填充
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

// AesEncrypt AES-CBC-PKCS7 加密
func AesEncrypt(data []byte, key []byte, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	data = PKCS7Padding(data, blockSize)
	crypted := make([]byte, len(data))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(crypted, data)
	return crypted
}

// AesDecrypt AES-CBC-PKCS7 解密
func AesDecrypt(data []byte, key []byte, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(data))
	blockMode.CryptBlocks(origData, data)
	origData = PKCS7UnPadding(origData)
	return origData
}

func main() {
	// 测试数据
	data := []byte("Hello, world!")
	key := []byte("00000yozo_config")
	iv := []byte("00000yozo_config")

	// 加密
	crypted := AesEncrypt(data, key, iv)
	fmt.Println("加密后的数据：", base64.StdEncoding.EncodeToString(crypted)) // 加密后的数据： Fe3xQJZVALMxovBw4qNGLA==

	// 解密
	origData := AesDecrypt(crypted, key, iv)
	fmt.Println("解密后的数据：", string(origData)) // 解密后的数据： Hello, world!
}
```

#### Java demo
```Java
import org.bouncycastle.jce.provider.BouncyCastleProvider;

import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.security.Security;
import java.util.Arrays;
import java.util.Base64;

public class Main
{
    public static void main(String[] args) throws Exception
    {
        //// 生成一个AES密钥
        //KeyGenerator keyGenerator = KeyGenerator.getInstance("AES");
        //keyGenerator.init(256); // 可以是128，192或256
        //SecretKey secretKey = keyGenerator.generateKey();


        //// 生成一个随机的初始向量
        //
        //SecureRandom random = new SecureRandom();
        //random.nextBytes(ivBytes);

        byte[] keyBytes = "00000yozo_config".getBytes();
        byte[] ivBytes = keyBytes;
        IvParameterSpec iv = new IvParameterSpec(ivBytes);

        // 要加密的字符串
        String originalStr = "Hello, world!";
        System.out.println("原始字符串: " + originalStr);    // 原始字符串: Hello, world!

        // 加密
        byte[] encryptedBytes = aesEncrypt(originalStr.getBytes(), keyBytes, iv);
        System.out.println("加密后的字节数组: " + Arrays.toString(encryptedBytes));

        // 使用Base64编码加密后的字节数组，方便显示和传输
        String encryptedStr = Base64.getEncoder().encodeToString(encryptedBytes);
        System.out.println("加密后的Base64字符串: " + encryptedStr);   // 加密后的Base64字符串: Fe3xQJZVALMxovBw4qNGLA==

        // 解密
        byte[] decryptedBytes = aesDecrypt(encryptedBytes, keyBytes, iv);
        System.out.println("解密后的字节数组: " + Arrays.toString(decryptedBytes));

        // 使用UTF-8编码解密后的字节数组，得到原始字符串
        String decryptedStr = new String(decryptedBytes, "UTF-8");
        System.out.println("解密后的字符串: " + decryptedStr);     // 解密后的字符串: Hello, world!
    }

    // 使用AES/CBC/PKCS7Padding加密数据
    public static byte[] aesEncrypt(byte[] data, byte[] key, IvParameterSpec iv) throws Exception
    {
        // 添加BouncyCastleProvider
        Security.addProvider(new BouncyCastleProvider());
        // 创建Cipher对象
        Cipher cipher = Cipher.getInstance("AES/CBC/PKCS7Padding");
        // 初始化Cipher对象，设置为加密模式
        cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(key, "AES"), iv);
        // 加密数据
        return cipher.doFinal(data);
    }

    // 使用AES/CBC/PKCS7Padding解密数据
    public static byte[] aesDecrypt(byte[] data, byte[] key, IvParameterSpec iv) throws Exception
    {
        // 添加BouncyCastleProvider
        Security.addProvider(new BouncyCastleProvider());
        // 创建Cipher对象
        Cipher cipher = Cipher.getInstance("AES/CBC/PKCS7Padding");
        // 初始化Cipher对象，设置为解密模式
        cipher.init(Cipher.DECRYPT_MODE, new SecretKeySpec(key, "AES"), iv);
        // 解密数据
        return cipher.doFinal(data);
    }
}
```

#### JavaScript demo
```js
/* 需要先导入 crypto-js.js */
// 密钥和偏移量
var key = CryptoJS.enc.Utf8.parse('00000yozo_config');
var iv = CryptoJS.enc.Utf8.parse('00000yozo_config');

// 测试加密解密函数
var word = "Hello, world!";
var ciphertext = encrypt(word);
var plaintext = decrypt(ciphertext);
console.log("明文:", word);           // 明文: Hello, world!
console.log("密文:", ciphertext);     // 密文: Fe3xQJZVALMxovBw4qNGLA==
console.log("解密:", plaintext);      // 解密: Hello, world!

// 定义加密函数，使用AES/cbc/pkcs7加密，返回base64编码的字符串
function encrypt(word) {
    // 将明文转换为Utf8编码的字节数组
    let srcs = CryptoJS.enc.Utf8.parse(word);
    // 使用AES/cbc/pkcs7加密
    let encrypted = CryptoJS.AES.encrypt(srcs, key, {
        iv: iv,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7,
    });
    // 将加密结果转换为base64编码的字符串
    return encrypted.toString();
}

// 定义解密函数，使用AES/cbc/pkcs7解密，返回Utf8编码的字符串
function decrypt(word) {
    // 使用AES/cbc/pkcs7解密
    let decrypt = CryptoJS.AES.decrypt(word, key, {
        iv: iv,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7,
    });
    // 将解密结果转换为Utf8编码的字符串
    return decrypt.toString(CryptoJS.enc.Utf8);
}
```

