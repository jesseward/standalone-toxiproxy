package ghtoxics

import (
	"encoding/hex"
	"io"
	"log"

	"github.com/Shopify/toxiproxy/v2/stream"
	"github.com/Shopify/toxiproxy/v2/toxics"
)

// DebugToxic prints bytes processed through pipe.
type DebugToxic struct{}

func (t *DebugToxic) Pipe(stub *toxics.ToxicStub) {
	buf := make([]byte, 32*1024)
	writer := stream.NewChanWriter(stub.Output)
	reader := stream.NewChanReader(stub.Input)
	reader.SetInterrupt(stub.Interrupt)
	for {
		n, err := reader.Read(buf)
		log.Printf("-- [DebugToxic] Processed %d bytes\n", n)
		if err == stream.ErrInterrupted {
			writer.Write(buf[:n])
			return
		} else if err == io.EOF {
			stub.Close()
			return
		}
		hex.Dump(buf[:n])
		writer.Write(buf[:n])
	}

}
