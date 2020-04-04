package fox

import (
	"lzr"
	"encoding/hex"
	"strings"
)

// Handshake implements the lzr.Handshake interface
type HandshakeMod struct {
}

const (
    // ORIGINAL_QUERY is the hex encoding of the query that will be sent to each server.
    ORIGINAL_QUERY = "666f7820612031202d3120666f782068656c6c6f0a7b0a" +
        "666f782e76657273696f6e3d733a312e300a69643d693a310a686f73744e" +
        "616d653d733a7870766d2d306f6d64633031786d790a686f737441646472" +
        "6573733d733a3139322e3136382e312e3132350a6170702e6e616d653d73" +
        "3a576f726b62656e63680a6170702e76657273696f6e3d733a332e372e34" +
        "340a766d2e6e616d653d733a4a61766120486f7453706f7428544d292053" +
        "657276657220564d0a766d2e76657273696f6e3d733a32302e342d623032" +
        "0a6f732e6e616d653d733a57696e646f77732058500a6f732e7665727369" +
        "6f6e3d733a352e310a6c616e673d733a656e0a74696d655a6f6e653d733a" +
        "416d65726963612f4c6f735f416e67656c65733b2d32383830303030303b" +
        "333630303030303b30323a30303a30302e3030302c77616c6c2c6d617263" +
        "682c382c6f6e206f722061667465722c73756e6461792c756e646566696e" +
        "65643b30323a30303a30302e3030302c77616c6c2c6e6f76656d6265722c" +
        "312c6f6e206f722061667465722c73756e6461792c756e646566696e6564" +
        "0a686f737449643d733a57696e2d393943422d443439442d353434322d30" +
        "3742420a766d557569643d733a38623533306263382d373663352d343133" +
        "392d613265612d3066616264333934643330350a6272616e6449643d733a" +
        "76796b6f6e0a7d3b3b0a"
    // RESPONSE_PREFIX is the prefix that will identify a Fox service.
    RESPONSE_PREFIX = "fox a 0 -1 fox hello"
)

func (h *HandshakeMod) GetData( dst string ) []byte {
	data, _ := hex.DecodeString(ORIGINAL_QUERY)
    return []byte(data)
}

func (h *HandshakeMod) Verify( data string ) string {
    if strings.HasPrefix( data, RESPONSE_PREFIX ) {
		return "fox"
	}
    return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake( "x11", &h )
}

