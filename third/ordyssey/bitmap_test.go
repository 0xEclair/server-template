package ordyssey

import (
	"fmt"
	"github.com/joho/godotenv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	godotenv.Load()
}

func TestAllBitmaps(t *testing.T) {
	res, err := AllBitmaps()
	assert.NoError(t, err)

	fmt.Println(res)
}
