package rights

import (
	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context, operations FileOperations) bool {
	r := Read(c)
	w := Write(c)
	d := Delete(c)
	p := r | w | d

	if p&operations == R|W|D {
		return true
	} else if p&operations == R|W {
		return true
	} else if p&operations == R|D {
		return true
	} else if p&operations == W|D {
		return true
	} else if p&operations == R {
		return true
	} else if p&operations == W {
		return true
	} else if p&operations == D {
		return true
	} else {
		return false
	}
}
