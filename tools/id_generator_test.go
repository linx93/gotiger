package tools

import "testing"

func TestUUID_GetStrId(t *testing.T) {
	var uuid IdGenerator = &UUID{}
	id := uuid.GetStrId()
	t.Logf("id: %v", id)
	t.Logf("idgen: %v", uuid.GetNumberId())
}

func TestSnowFlake_GetStrId(t *testing.T) {
	var snowFlake IdGenerator = &SnowFlake{}
	t.Logf("id: %v", snowFlake.GetStrId())
	t.Logf("id: %v", snowFlake.GetNumberId())

}
