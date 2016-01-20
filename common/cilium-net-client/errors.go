package cilium_net_client

import (
	"errors"
	"fmt"
	"io"

	"github.com/noironetworks/cilium-net/common/types"

	"github.com/docker/docker/vendor/src/github.com/jfrazelle/go/canonical/json"
)

var ErrConnectionFailed = errors.New("Cannot connect to the cilium-net-daemon. Is the cilium-net-daemon running on this host?")

func processErrorBody(serverResp io.ReadCloser, i interface{}) error {
	d := json.NewDecoder(serverResp)
	var sErr types.ServerError
	if err := d.Decode(&sErr); err != nil {
		fmt.Errorf("error retrieving server body response: %s", err)
	}
	return fmt.Errorf("server error for endpoint: '%+v', (%d) %s", i, sErr.Code, sErr.Text)
}
