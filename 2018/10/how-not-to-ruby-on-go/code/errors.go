package dao

//START OMIT
import "github.com/pkg/errors"

func A() {
	_, err := B()
	if err != nil {
		errors.Wrap(err, "B failed")
	}
// END OMIT
}

func B() (bytes []byte, error error) {
	bytes =  make([]byte, 0)
	return
}