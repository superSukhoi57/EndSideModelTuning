package types

type AnswerResp struct {
	Base   Base   `json:"base"`
	Answer string `json:"answer"`
}

type Base struct {
	Code      int32  `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type QuestionReq struct {
	Question string `json:"question"`
}

type MachineCreateReq struct {
	IP     string `json:"ip"`
	Pwd    string `json:"pwd,omitempty"`
	UserID int64  `json:"userid"`
	Core   *int   `json:"core,omitempty"`
	RAM    *int   `json:"ram,omitempty"`
	Memory *int   `json:"memory,omitempty"`
	OS     string `json:"os,omitempty"`
	Desc   string `json:"desc,omitempty"`
}

type MachineUpdateReq struct {
	ID       int64  `json:"id"`
	IP       string `json:"ip,omitempty"`
	Pwd      string `json:"pwd,omitempty"`
	IsFinsh  *int8  `json:"isfinsh,omitempty"`
	ResultID *int64 `json:"resultid,omitempty"`
	Core     *int   `json:"core,omitempty"`
	RAM      *int   `json:"ram,omitempty"`
	Memory   *int   `json:"memory,omitempty"`
	OS       string `json:"os,omitempty"`
	Desc     string `json:"desc,omitempty"`
}

type MachineListReq struct {
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"pageSize,default=10"`
	IP       string `form:"ip,optional"`
	UserID   int64  `form:"userid,optional"`
	IsFinsh  *int8  `form:"isfinsh,optional"`
}

type MachineResp struct {
	ID       int64  `json:"id"`
	IP       string `json:"ip"`
	UserID   int64  `json:"userid"`
	IsFinsh  int8   `json:"isfinsh"`
	ResultID *int64 `json:"resultid"`
	Core     *int   `json:"core"`
	RAM      *int   `json:"ram"`
	Memory   *int   `json:"memory"`
	OS       string `json:"os"`
	Desc     string `json:"desc"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type PageResp struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

type ParameterCreateReq struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"userid"`
	Parameters string `json:"parameters,omitempty"`
	Script     string `json:"script"`
	Desc       string `json:"desc,omitempty"`
}

type ParameterUpdateReq struct {
	ID         int64  `json:"id"`
	Parameters string `json:"parameters,omitempty"`
	Script     string `json:"script,omitempty"`
	Desc       string `json:"desc,omitempty"`
}

type ParameterListReq struct {
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"pageSize,default=10"`
	UserID   int64  `form:"userid,optional"`
	Desc     string `form:"desc,optional"`
}

type ParameterResp struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"userid"`
	Parameters string `json:"parameters"`
	Script     string `json:"script"`
	Desc       string `json:"desc"`
	CreateAt   string `json:"createAt"`
	UpdateAt   string `json:"updateAt"`
}

type TaskCreateReq struct {
	ID          int64  `json:"id"`
	ParameterID int64  `json:"paramterid"`
	UserID      int64  `json:"userid"`
	Desc        string `json:"desc,omitempty"`
}

type TaskUpdateReq struct {
	ID   int64  `json:"id"`
	Desc string `json:"desc,omitempty"`
}

type TaskListReq struct {
	Page        int    `form:"page,default=1"`
	PageSize    int    `form:"pageSize,default=10"`
	UserID      int64  `form:"userid,optional"`
	ParameterID int64  `form:"paramterid,optional"`
	Desc        string `form:"desc,optional"`
}

type TaskResp struct {
	ID          int64  `json:"id"`
	ParameterID int64  `json:"paramterid"`
	UserID      int64  `json:"userid"`
	Desc        string `json:"desc"`
	CreateAt    string `json:"createAt"`
	UpdateAt    string `json:"updateAt"`
}

type ResultCreateReq struct {
	Result    string `json:"result"`
	UserID    int64  `json:"userid"`
	MachineID int64  `json:"machineid"`
	Desc      string `json:"desc,omitempty"`
}

type ResultUpdateReq struct {
	ID     int64  `json:"id"`
	Result string `json:"result,omitempty"`
	Desc   string `json:"desc,omitempty"`
}

type ResultListReq struct {
	Page      int    `form:"page,default=1"`
	PageSize  int    `form:"pageSize,default=10"`
	UserID    int64  `form:"userid,optional"`
	MachineID int64  `form:"machineid,optional"`
	Desc      string `form:"desc,optional"`
}

type ResultResp struct {
	ID        int64  `json:"id"`
	Result    string `json:"result"`
	UserID    int64  `json:"userid"`
	MachineID int64  `json:"machineid"`
	Desc      string `json:"desc"`
	CreateAt  string `json:"createAt"`
	UpdateAt  string `json:"updateAt"`
}
