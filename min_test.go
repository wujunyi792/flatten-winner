package flatten_winner

import (
	"fmt"
	"github.com/wujunyi792/flatten-winner/internal/service/scanFolder"
	"testing"
)

func TestWalk(T *testing.T) {
	e := scanFolder.LoadAllFolder()
	fmt.Println(e)
}
