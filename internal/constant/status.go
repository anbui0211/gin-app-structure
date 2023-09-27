package constant

const (
	Active   = "active"
	InActive = "inactive"
)

type StatusInfo struct {
	Key   string
	Title string
}

var Status = struct {
	New       StatusInfo
	Rejected  StatusInfo
	Approved  StatusInfo
	Running   StatusInfo
	Completed StatusInfo
	Deleted   StatusInfo
	Fail      StatusInfo
}{
	New: StatusInfo{
		Key:   "new",
		Title: "chờ kiểm duyệt",
	},
	Rejected: StatusInfo{
		Key:   "rejected",
		Title: "đã hủy",
	},
	Approved: StatusInfo{
		Key:   "approved",
		Title: "đã duyệt",
	},
	Running: StatusInfo{
		Key:   "running",
		Title: "đang đối soát",
	},
	Completed: StatusInfo{
		Key:   "completed",
		Title: "đã đối soát",
	},
	Deleted: StatusInfo{
		Key:   "deleted",
		Title: "đã xóa",
	},
	Fail: StatusInfo{
		Key:   "fail",
		Title: "lỗi đối soát",
	},
}
