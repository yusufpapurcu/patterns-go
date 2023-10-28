package module_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yusufpapurcu/patterns-go/dependency_injection/module"
	mocks "github.com/yusufpapurcu/patterns-go/dependency_injection/module/internal/mocks"
)

func Test_ProcessEmails(t *testing.T) {
	t.Run("run successfully", func(t *testing.T) {
		writerMock := mocks.WriterMock{}

		writerMock.WriteEmailsFunc = func(emails ...string) error {
			assert.Len(t, emails, 2)
			assert.Equal(t, "aa", emails[0])
			assert.Equal(t, "bb", emails[1])
			return nil
		}
		m := module.NewModule(&writerMock)
		m.ProcessEmails("aa", "bb")

		assert.Len(t, writerMock.WriteEmailsCalls(), 1)
	})
}
