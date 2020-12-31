package consts

type IsDeleted int

const (
	DELETED_FALSE IsDeleted = 0 // 未删除
	DELETED_TRUE  IsDeleted = 1
)

func (d IsDeleted) String() string {
	switch d {
	case DELETED_FALSE:
		return "deleted_false"
	case DELETED_TRUE:
		return "deleted_true"
	default:
		return "Unknown"
	}
}
