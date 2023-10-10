## TODO
(1) Go语言：识别全格式图片并执行相应的编码保存（支持webp） https://www.meiwen.com.cn/subject/gfiomctx.html
    image.Decode
(2) webp格式图片

## 第三方库davidbyttow/govips（libvips推荐的库）
GO语言高性能图片处理库govips
    https://www.bilibili.com/video/BV1L14y1i7iG/  
github(982 Star):
    https://github.com/davidbyttow/govips  

PS:
* libvips v8.3+ is required for GIF, PDF and SVG support.
* libvips v8.9+ is required for AVIF support. libheif compiled with a AVIF en-/decoder also needs to be present.

#### 要求!!!
* libvips 8.10+ 
* C compatible compiler such as gcc 4.6+ or clang 3.0+ 
* Go 1.14+ 

#### Mac环境安装
命令（安装 vips、pkg-config 和 gcc）: brew reinstall vips pkg-config gcc

## 第三方库h2non/bimg（u推荐）
「GoCN酷Go推荐」Go 语言高性能图像处理神器 h2non/bimg
    https://mp.weixin.qq.com/s/kAFZohzJo2DiKkxjnVti6A
github(2.4k Star): 
    https://github.com/h2non/bimg

* bimg 是一个小型且高效的库，支持常见的图像操作，如裁剪、缩放、旋转、缩放或水印。
* 它可以原生读取 JPEG、PNG、WEBP，如果 libvips@8.3+ 编译了适当的库绑定，还可以选择性地读取 TIFF、PDF、GIF 和 SVG 格式。
* bimg 能够将图像输出为 JPEG、PNG 和 WEBP 格式，包括在它们之间进行透明转换。
* h2non/bimg 提供以下出片处理 API：
    调整大小
    放大
    裁剪（包括智能裁剪支持，libvips 8.5+）
    旋转（根据 EXIF 方向自动旋转）
    翻转（具有基于EXIF元数据的自动翻转）
    翻转
    缩略图
    获取大小
    水印（使用文本或图像）
    高斯模糊效果
    自定义输出颜色空间（RGB，灰度...）
    格式转换以及压缩处理
    EXIF元数据（大小，Alpha通道，配置文件，方向...）修改
    修剪（libvips 8.6+）



## problems
将透明背景的PNG转换为JPG（或JPEG），默认背景色为黑色
    https://www.zongscan.com/demo333/95729.html
