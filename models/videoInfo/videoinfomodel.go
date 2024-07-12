package videoInfo

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VideoinfoModel = (*customVideoinfoModel)(nil)

type (
	// VideoinfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoinfoModel.
	VideoinfoModel interface {
		videoinfoModel
		withSession(session sqlx.Session) VideoinfoModel
	}

	customVideoinfoModel struct {
		*defaultVideoinfoModel
	}
)

// NewVideoinfoModel returns a model for the database table.
func NewVideoinfoModel(conn sqlx.SqlConn) VideoinfoModel {
	return &customVideoinfoModel{
		defaultVideoinfoModel: newVideoinfoModel(conn),
	}
}

func (m *customVideoinfoModel) withSession(session sqlx.Session) VideoinfoModel {
	return NewVideoinfoModel(sqlx.NewSqlConnFromSession(session))
}
