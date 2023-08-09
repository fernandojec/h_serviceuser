package appointment

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type RequestAdd struct {
	HealthcareId    string    `protobuf:"bytes,1,opt,name=healthcare_id,json=healthcareId,proto3" json:"healthcare_id,omitempty"`
	ParamedicId     string    `protobuf:"bytes,3,opt,name=paramedic_id,json=paramedicId,proto3" json:"paramedic_id,omitempty"`
	PatientId       string    `protobuf:"bytes,4,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	AppointmentDate time.Time `protobuf:"bytes,5,opt,name=appointment_date,json=appointmentDate,proto3" json:"appointment_date,omitempty"`
	AppointmentTime time.Time `protobuf:"bytes,6,opt,name=appointment_time,json=appointmentTime,proto3" json:"appointment_time,omitempty"`
	ScheduleSlotId  int64     `protobuf:"varint,10,opt,name=schedule_slot_id,json=scheduleSlotId,proto3" json:"schedule_slot_id,omitempty"`
}

func (m *RequestAdd) ConvertToAppointmentProto() AppointmentProto {
	return AppointmentProto{
		HealthcareId:    m.HealthcareId,
		ParamedicId:     m.ParamedicId,
		PatientId:       m.PatientId,
		AppointmentDate: timestamppb.New(m.AppointmentDate),
		AppointmentTime: timestamppb.New(m.AppointmentTime),
		ScheduleSlotId:  m.ScheduleSlotId,
	}
}

type RequestEdit struct {
	ParamedicId     string    `protobuf:"bytes,3,opt,name=paramedic_id,json=paramedicId,proto3" json:"paramedic_id,omitempty"`
	AppointmentDate time.Time `protobuf:"bytes,5,opt,name=appointment_date,json=appointmentDate,proto3" json:"appointment_date,omitempty"`
	AppointmentTime time.Time `protobuf:"bytes,6,opt,name=appointment_time,json=appointmentTime,proto3" json:"appointment_time,omitempty"`
	ScheduleSlotId  int64     `protobuf:"varint,10,opt,name=schedule_slot_id,json=scheduleSlotId,proto3" json:"schedule_slot_id,omitempty"`
}

func (m *RequestEdit) ConvertToAppointmentProto() AppointmentProto {
	return AppointmentProto{
		ParamedicId:     m.ParamedicId,
		AppointmentDate: timestamppb.New(m.AppointmentDate),
		AppointmentTime: timestamppb.New(m.AppointmentTime),
		ScheduleSlotId:  m.ScheduleSlotId,
	}
}
