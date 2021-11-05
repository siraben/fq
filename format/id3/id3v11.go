package id3

import (
	"github.com/wader/fq/format"
	"github.com/wader/fq/format/registry"
	"github.com/wader/fq/pkg/decode"
)

func init() {
	registry.MustRegister(&decode.Format{
		Name:        format.ID3V11,
		Description: "ID3v1.1 metadata",
		DecodeFn:    id3v11Decode,
	})
}

func id3v11Decode(d *decode.D, in interface{}) interface{} {
	d.AssertAtLeastBitsLeft(128 * 8)
	d.FieldUTF8("magic", 4, d.AssertStr("TAG+"))
	d.FieldUTF8("title", 60)
	d.FieldUTF8("artist", 60)
	d.FieldUTF8("album", 60)
	d.FieldU8("speed", d.MapUToStr(decode.UToStr{
		0: "unset",
		1: "slow",
		2: "medium",
		3: "fast",
		4: "hardcore",
	}))
	d.FieldUTF8("genre", 30)
	d.FieldUTF8("start", 6)
	d.FieldUTF8("stop", 6)

	return nil
}
