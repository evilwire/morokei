package app

type external struct {

}

var __external__ *external = nil

func External() *external {
	if __external__ == nil {
		// create _external_
		__external__ = &external {

		}
	}

	return __external__
}
