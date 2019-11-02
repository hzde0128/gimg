# Gimg

- - -
Gimg是[zimg](https://github.com/buaazp/zimg)的golang版本。

完全兼容zimg的文件目录储存格式，支持文件和类Redis协议(SSDB)储存。

环境要求：

* 操作系统： ubuntu/debian/osx
* golang版本： >= 1.4
* ImageMagick版本： 6.9.1-7 - 6.9.10-71



# Install

- - -
## Ubuntu/Debian

- - -
	wget https://imagemagick.org/download/ImageMagick-6.9.10-71.tar.gz
	tar zxvf ImageMagick-6.8.10-71.tar.gz
	cd ImageMagick-6.9.10-71/
	./configure
	make & make install
	ldconfig /usr/local/lib
- - -
## OSX

- - -
	brew install ImageMagick
	

- - -
## CentOS/RedHat


- - -
        cat << EOF > /etc/yum.repos.d/imagemagick.repo
[imagemagick]
name=image magick
baseurl=https://imagemagick.org/download/linux/CentOS/x86_64/
enabled=1
gpgcheck=0
EOF
        yum -y install ImageMagick ImageMagick-libs ImageMagick-devel

## 安装
- - -
        make build	
	cd build
	./gimg --config=./conf/config.ini
	
	
## Demo
---
[http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93](http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93)

[http://182.92.189.64:8081/info?md5=a258607b53444f32208e864f44a06b93](http://182.92.189.64:8081/info?md5=a258607b53444f32208e864f44a06b93)

[http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93?w=100&h=100&x=-1&y=-1](http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93?w=100&h=100&x=-1&y=-1)

[http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93?r=45](http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93?r=45)

[http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93?g=1](http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93?g=1)

[http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93?f=png](http://182.92.189.64:8081/a258607b53444f32208e864f44a06b93?f=png)
