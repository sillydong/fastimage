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


	url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"
	
	imagetype, size, err := fastimage.DetectImageType(url)
	if err != nil {
		// Something went wrong, http failed? not an image?
		panic(err)
	}

	switch imagetype {
	case fastimage.JPEG:
		log.Printf("JPEG")
	case fastimage.PNG:
		log.Printf("PNG")
	case fastimage.GIF:
		log.Printf("GIF")
	case fastimage.BMP:
		log.Printf("BMP")
	case fastimage.WEBP:
		log.Printf("WEBP")
	case fastimage.TIFF:
		log.Printf("TIFF")
	}

	log.Printf("Image type: %s", imagetype.String())
	log.Printf("Image size: %v", size)


## Supported file types


| File type | Can detect type? | Can detect size? |
|-----------|:----------------:|:----------------:|
| PNG       | Yes              | Yes              |
| JPEG      | Yes              | Yes              |
| GIF       | Yes              | Yes              |
| BMP       | Yes              | Yes              |
| TIFF      | Yes              | Yes              |
| WEBP      | Yes              | Yes              |


# Project details

### License

fastimage is under MIT license. See the [LICENSE][license] file for details.

[license]: https://github.com/sillydong/fastimage/blob/master/LICENSE
