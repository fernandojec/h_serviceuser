package schedules

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type RequestAdd struct {
	Healthcareid  string    `protobuf:"bytes,1,opt,name=healthcareid,proto3" json:"healthcareid,omitempty"`
	Paramedicid   string    `protobuf:"bytes,2,opt,name=paramedicid,proto3" json:"paramedicid,omitempty"`
	Schedulestart time.Time `protobuf:"bytes,3,opt,name=schedulestart,proto3" json:"schedulestart,omitempty"`
	Scheduleend   time.Time `protobuf:"bytes,4,opt,name=scheduleend,proto3" json:"scheduleend,omitempty"`
	Duration      int32     `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *RequestAdd) ConvertToScheduleCreateProto() ScheduleCreateProto {
	return ScheduleCreateProto{
		Healthcareid:  m.Healthcareid,
		Paramedicid:   m.Paramedicid,
		Schedulestart: timestamppb.New(m.Schedulestart),
		Scheduleend:   timestamppb.New(m.Scheduleend),
		Duration:      m.Duration,
	}
}

type RequestFindSchedule struct {
	Healthcareid string    `protobuf:"bytes,1,opt,name=healthcareid,proto3" json:"healthcareid,omitempty"`
	Paramedicid  string    `protobuf:"bytes,2,opt,name=paramedicid,proto3" json:"paramedicid,omitempty"`
	Scheduledate time.Time `protobuf:"bytes,3,opt,name=scheduledate,proto3" json:"scheduledate,omitempty"`
}

func (m *RequestFindSchedule) ConvertToScheduleFindProto() FindScheduleProto {
	return FindScheduleProto{
		Healthcareid: m.Healthcareid,
		Paramedicid:  m.Paramedicid,
		Scheduledate: timestamppb.New(m.Scheduledate),
	}
}
