## 在线加解密
Aes加密/解密（密文类型: base64）
    https://lovefree.cc/aes
在线AES加密解密
    https://www.wushuangzl.com/encrypt/aes.html

## AES支持的模式
PS: 其中，使用最为普遍的是CBC模式，不建议使用的是ECB模式（不安全）
* ECB模式（The Electronic Codebook Mode）
* CBC模式（The Cipher Block Chaining Mode）
* CFB模式（The Cipher Feedback Mode）
* CTR模式（The Counter Mode）
* OFB模式（The Output Feedback Mode）

## 密钥KEY
AES标准规定区块长度只有一个值，固定为128Bit，对应的字节为16位。
AES算法规定密钥长度只有三个值，128Bit、192Bit、256Bit，对应的字节为16位、24位和32位，其中密钥KEY不能公开传输，用于加密解密数据.

#### !!!: key的长度
16字节: AES-128
24字节: AES-192
32字节: AES-256

## 初始化向量IV
该字段可以公开，用于将加密随机化。同样的明文被多次加密也会产生不同的密文，避免了较慢的重新产生密钥的过程，初始化向量与密钥相比有不同的安全性需求，因此IV通常无须保密。
然而在大多数情况中，不应当在使用同一密钥的情况下两次使用同一个IV，一般推荐初始化向量IV为16位的随机值.
