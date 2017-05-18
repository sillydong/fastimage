# fastimage

Golang implementation of [fastimage](https://pypi.python.org/pypi/fastimage/0.2.1).

Finds the type and/or size of an image given its uri by fetching as little as needed.

**RECREATED** and **IMPROVED** fastimage for golang.

Started from Ruben Fonseca (@[rubenfonseca](http://twitter.com/rubenfonseca)) 's fastimage [https://github.com/rubenfonseca/fastimage](https://github.com/rubenfonseca/fastimage)

Included TIFF and WEBP support from [https://github.com/golang/image](https://github.com/golang/image)

## How?

fastimage parses the image data as it is downloaded. As soon as it finds out
the size and type of the image, it stops the download.

## Install

    $ go get github.com/sillydong/fastimage

## Usage

For instance, this is a big 10MB JPEG image on wikipedia:

**Method1**

```go
customHeaders := map[string]string{
    "X-SECRET-HEADER": "your-header-value"
}
instance := fastimage.NewFastImage(2, customHeaders)
//leave it to nil to use default header settings
//eg. 
//instance := fastimage.NewFastImage(2, nil)

url1 := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"
imagetype1, size1, err1 := instance.Detect(url1)
fmt.Printf("%+v\t%+v\t%+v\n", imagetype1, size1, err1)

url2 := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"
imagetype2, size2, err2 := instance.Detect(url2)
fmt.Printf("%+v\t%+v\t%+v\n", imagetype2, size2, err2)
```

**Method2**

```go
url1 := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"
imagetype1, size1, err1 := fastimage.GetImageSize(url1)
fmt.Printf("%+v\t%+v\t%+v\n", imagetype1, size1, err1)

url2 := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"
imagetype2, size2, err2 := fastimage.GetImageSize(url2)
fmt.Printf("%+v\t%+v\t%+v\n", imagetype2, size2, err2)
```

**Notice**

Method1 is better to detect multiple images because it shares one *http.Client.

## Supported file types

| File type | Can detect type? | Can detect size? |
|-----------|:----------------:|:----------------:|
| PNG       | Yes              | Yes              |
| JPEG      | Yes              | Yes              |
| GIF       | Yes              | Yes              |
| BMP       | Yes              | Yes              |
| TIFF      | Yes              | Yes              |
| WEBP      | Yes              | Yes              |

**Notice**

Some webp images' Content-Type in header is `text/plain`. Here in this library we only parse response data when Content-Type has keyword **"image"** in.


## TODO

I'm thinking about using [`valyala/fasthttp`](https://github.com/valyala/fasthttp) to replace `net/http` to get better network performance.

### License

fastimage is under MIT license. See the [LICENSE](https://github.com/sillydong/fastimage/blob/master/LICENSE) file for details.
