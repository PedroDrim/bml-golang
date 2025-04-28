package simpletableanalysis

import(
	"testing"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

type MockSimpleTableAnalysis struct{}

func (m *MockSimpleTableAnalysis) Analysis(userInfoList []userinfo.UserInfo) float64 {
	return 0
}

func TestAnalysis(t *testing.T) {
	t.Run("1. Integracao do 'analysis()'", func(t *testing.T) {
		mock := &MockSimpleTableAnalysis{}

		list := []userinfo.UserInfo{
			*userinfo.NewUserInfo("ua", "pa", 1),
			*userinfo.NewUserInfo("ub", "pb", 2),
			*userinfo.NewUserInfo("uc", "pc", 3),
		}		
		
		result := mock.Analysis(list)
		if result != 0 {
			t.Errorf("1. Esperava 0, recebeu %v", result)
		}
	})
}