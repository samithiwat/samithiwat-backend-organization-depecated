package client

import "gopkg.in/errgo.v2/errors"

func FormatErrors(errs []string) (result []error) {
	for _, err := range errs {
		result = append(result, errors.New(err))
	}

	return
}
