package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrintJob An active print job.
// swagger:model PrintJob
type PrintJob struct {

	// datetime cleaned
	DatetimeCleaned strfmt.DateTime `json:"datetime_cleaned,omitempty"`

	// datetime finished
	DatetimeFinished strfmt.DateTime `json:"datetime_finished,omitempty"`

	// datetime started
	DatetimeStarted strfmt.DateTime `json:"datetime_started,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// source application
	SourceApplication string `json:"source_application,omitempty"`

	// source user
	SourceUser string `json:"source_user,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// time elapsed
	TimeElapsed int64 `json:"time_elapsed,omitempty"`

	// time total
	TimeTotal int64 `json:"time_total,omitempty"`

	// uuid
	UUID UUID `json:"uuid,omitempty"`
}

// Validate validates this print job
func (m *PrintJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var printJobTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","printing","pausing","paused","resuming","pre_print","post_print","wait_cleanup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		printJobTypeStatePropEnum = append(printJobTypeStatePropEnum, v)
	}
}

const (
	// PrintJobStateNone captures enum value "none"
	PrintJobStateNone string = "none"
	// PrintJobStatePrinting captures enum value "printing"
	PrintJobStatePrinting string = "printing"
	// PrintJobStatePausing captures enum value "pausing"
	PrintJobStatePausing string = "pausing"
	// PrintJobStatePaused captures enum value "paused"
	PrintJobStatePaused string = "paused"
	// PrintJobStateResuming captures enum value "resuming"
	PrintJobStateResuming string = "resuming"
	// PrintJobStatePrePrint captures enum value "pre_print"
	PrintJobStatePrePrint string = "pre_print"
	// PrintJobStatePostPrint captures enum value "post_print"
	PrintJobStatePostPrint string = "post_print"
	// PrintJobStateWaitCleanup captures enum value "wait_cleanup"
	PrintJobStateWaitCleanup string = "wait_cleanup"
)

// prop value enum
func (m *PrintJob) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, printJobTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PrintJob) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}