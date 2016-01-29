# base65536
[![GoDoc](https://img.shields.io/badge/go-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/usk81/base65536)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/usk81/base65536/blob/master/LICENSE)
[![Build Status](http://img.shields.io/travis/usk81/base65536.svg?style=flat-square)](https://travis-ci.org/usk81/base65536)
[![Coverage Status](https://img.shields.io/coveralls/usk81/base65536.svg?style=flat-square)](https://coveralls.io/github/usk81/base65536?branch=master)
[![Gratipay User](https://img.shields.io/gratipay/user/YusukeKomatsu.svg?style=flat-square)](https://gratipay.com/YusukeKomatsu/)

An implementation of [ferno's base65536](https://github.com/ferno/base65536), implemented in GoLang.

## Description

This is a Go library for encoding and decoding.
If you show detail for base65536, see [ferno's base65536 README page](https://github.com/ferno/base65536/)

## Install

standard `go get`:

```
go get -u github.com/usk81/base65536
```

## Usage

encode:

```go
s := "hello world"
result := base65536.Encode(s)

Println(result)
// 驨ꍬ啯𒁷ꍲᕤ
```

decode:

```go
s := "驨ꍬ啯𒁷ꍲᕤ"

result := base65536.Decode(s)
Println(result)
// hello world
```


## Contribution

1. Fork it!
2. Create your feature branch: git checkout -b my-new-feature
3. Commit your changes: git commit -am 'Add some feature'
4. Push to the branch: git push origin my-new-feature
5. Submit a pull request :D

## Licence

[MIT](https://github.com/usk81/base65536/blob/master/LICENSE)

## Author

[Yusuke Komatsu](https://github.com/usk81)
