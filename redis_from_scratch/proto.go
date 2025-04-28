package main

import (
	"bytes"
	"fmt"
)

// The formatting that is done here using the Sprintf is defined in the documentation of the resp protocol for redis

func respWriteMap(m map[string]string) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString("%" + fmt.Sprintf("%d\r\n", len(m)))
	// rw := resp.NewWriter(buf)
	for k, v := range m {
		buf.WriteString(k)
		buf.WriteString(":" + v + "\r\n")
	}
	return buf.Bytes()
}
