/* 需要先导入 crypto-js.js */
// 密钥和偏移量
var key = CryptoJS.enc.Utf8.parse('0123456789abcdef');
var iv = CryptoJS.enc.Utf8.parse('0123456789abcdef');

// 测试加密解密函数
var word = "test测试~！@#￥%……&*（）——+-=";
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


