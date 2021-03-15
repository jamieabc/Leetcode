package main

//     Note: This is a companion problem to the System Design problem: Design TinyURL.
//
// TinyURL is a URL shortening service where you enter a URL such as https://leetcode.com/problems/design-tinyurl and it returns a short URL such as http://tinyurl.com/4e9iAk.
//
// Design the encode and decode methods for the TinyURL service. There is no restriction on how your encode/decode algorithm should work. You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.

type Codec struct {
	Table map[string]string
	Size  int
}

func Constructor() Codec {
	return Codec{
		Table: make(map[string]string),
		Size:  6,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	encoded := []byte{0}
	for i := range longUrl {
		encoded = append(encoded, byte(longUrl[i])^byte(encoded[len(encoded)-1]))
		if len(encoded) > this.Size {
			encoded = encoded[1:]
		}
	}
	this.Table[string(encoded)] = longUrl

	return string(encoded)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.Table[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
