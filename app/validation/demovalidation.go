package validation

import (
    "errors"
    "github.com/carefreex-io/example/proto"
)

func DemoService_HelloValidate(request *proto.DemoHelloRequest) (err error) {
    if request.Name == "" {
        err = errors.New("request name is required")
        return err
    }

    return nil
}
