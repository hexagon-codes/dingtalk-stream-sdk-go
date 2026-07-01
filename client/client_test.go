package client

import (
	"os"
	"strings"
	"testing"
)

func TestProcessLoopDoesNotSendOnClosedSignalChannels(t *testing.T) {
	src, err := os.ReadFile("client.go")
	if err != nil {
		t.Fatalf("read client.go: %v", err)
	}
	text := string(src)
	for _, banned := range []string{"closeChan", "close(pongChan)", "close(readChan)"} {
		if strings.Contains(text, banned) {
			t.Fatalf("processLoop still contains unsafe channel-close pattern %q", banned)
		}
	}
	for _, required := range []string{"doneOnce", "signalDone", "case <-done:"} {
		if !strings.Contains(text, required) {
			t.Fatalf("processLoop missing single-owner done signal marker %q", required)
		}
	}
}

/**
 * @Author linya.jj
 * @Date 2023/3/22 14:23
 */

func TestNewDingtalkOpenStreamClient(t *testing.T) {
}

func TestDingtalkOpenStreamClient_Start(t *testing.T) {
}

func TestDingtalkOpenStreamClient_processDataFrame(t *testing.T) {
}

func TestDingtalkOpenStreamClient_Close(t *testing.T) {

}

func TestDingtalkOpenStreamClient_reconnect(t *testing.T) {
}

func TestDingtalkOpenStreamClient_GetHandler(t *testing.T) {
}

func TestDingtalkOpenStreamClient_CheckConfigValid(t *testing.T) {
}

func TestDingtalkOpenStreamClient_GetConnectionEndpoint(t *testing.T) {

}

func TestDingtalkOpenStreamClient_OnDisconnect(t *testing.T) {

}

func TestDingtalkOpenStreamClient_OnPing(t *testing.T) {
}

func TestDingtalkOpenStreamClient_SendDataFrameResponse(t *testing.T) {
}

func TestDingtalkOpenStreamClient_SendErrorResponse(t *testing.T) {
}
