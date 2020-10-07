// Copyright (C) 2017 ScyllaDB

package manager

import (
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/scylladb/scylla-mgmt-commons/format"
	"github.com/scylladb/scylla-operator/pkg/api/v1alpha1"
	"github.com/scylladb/scylla-operator/pkg/managerclient"
	"github.com/scylladb/scylla-operator/pkg/util/duration"
	"k8s.io/utils/pointer"
)

type RepairTask v1alpha1.RepairTaskStatus

func (r RepairTask) ToManager() (*managerclient.Task, error) {
	t := &managerclient.Task{
		ID:         r.ID,
		Type:       "repair",
		Enabled:    true,
		Schedule:   new(managerclient.Schedule),
		Properties: make(map[string]interface{}),
	}

	props := t.Properties.(map[string]interface{})

	if r.StartDate != nil {
		startDate, err := format.ParseStartDate(*r.StartDate)
		if err != nil {
			return nil, errors.Wrap(err, "parse start date")
		}
		t.Schedule.StartDate = strfmt.DateTime(startDate)
	}

	if r.Interval != nil {
		if _, err := duration.ParseDuration(*r.Interval); err != nil {
			return nil, errors.Wrap(err, "parse interval")
		}
		t.Schedule.Interval = *r.Interval
	}

	if r.NumRetries != nil {
		t.Schedule.NumRetries = int64(*r.NumRetries)
	}

	if r.Keyspace != nil {
		props["keyspace"] = unescapeFilters(r.Keyspace)
	}
	if r.DC != nil {
		props["dc"] = unescapeFilters(r.DC)
	}
	if r.FailFast != nil {
		if *r.FailFast {
			t.Schedule.NumRetries = 0
			props["fail_fast"] = true
		}
	}
	if r.Intensity != nil {
		props["intensity"] = *r.Intensity
	}
	if r.Parallel != nil {
		props["parallel"] = *r.Parallel
	}
	if r.SmallTableThreshold != nil {
		threshold, err := format.ParseByteCount(*r.SmallTableThreshold)
		if err != nil {
			return nil, errors.Wrap(err, "parse small table threshold")
		}
		props["small_table_threshold"] = threshold
	}

	t.Name = r.Name
	t.Properties = props

	return t, nil
}

func (r *RepairTask) FromManager(t *managerclient.ExtendedTask) error {
	r.ID = t.ID
	r.Name = t.Name
	r.Interval = pointer.StringPtr(t.Schedule.Interval)
	r.StartDate = pointer.StringPtr(t.Schedule.StartDate.String())
	r.NumRetries = pointer.Int64Ptr(t.Schedule.NumRetries)

	props := t.Properties.(map[string]interface{})
	if err := mapstructure.Decode(props, r); err != nil {
		return errors.Wrap(err, "decode properties")
	}

	return nil
}

type BackupTask v1alpha1.BackupTaskStatus

func (b BackupTask) ToManager() (*managerclient.Task, error) {
	t := &managerclient.Task{
		ID:         b.ID,
		Type:       "backup",
		Enabled:    true,
		Schedule:   new(managerclient.Schedule),
		Properties: make(map[string]interface{}),
	}

	props := t.Properties.(map[string]interface{})

	if b.StartDate != nil {
		startDate, err := format.ParseStartDate(*b.StartDate)
		if err != nil {
			return nil, errors.Wrap(err, "parse start date")
		}
		t.Schedule.StartDate = strfmt.DateTime(startDate)
	}

	if b.Interval != nil {
		if _, err := duration.ParseDuration(*b.Interval); err != nil {
			return nil, errors.Wrap(err, "parse interval")
		}
		t.Schedule.Interval = *b.Interval
	}

	if b.NumRetries != nil {
		t.Schedule.NumRetries = int64(*b.NumRetries)
	}

	if b.Keyspace != nil {
		props["keyspace"] = unescapeFilters(b.Keyspace)
	}
	if b.DC != nil {
		props["dc"] = unescapeFilters(b.DC)
	}
	if b.Retention != nil {
		props["retention"] = *b.Retention
	}
	if b.RateLimit != nil {
		props["rate_limit"] = b.RateLimit
	}
	if b.SnapshotParallel != nil {
		props["snapshot_parallel"] = b.SnapshotParallel
	}
	if b.UploadParallel != nil {
		props["upload_parallel"] = b.UploadParallel
	}

	props["location"] = b.Location
	t.Name = b.Name
	t.Properties = props

	return t, nil
}

func (b *BackupTask) FromManager(t *managerclient.ExtendedTask) error {
	b.ID = t.ID
	b.Name = t.Name
	b.Interval = pointer.StringPtr(t.Schedule.Interval)
	b.StartDate = pointer.StringPtr(t.Schedule.StartDate.String())
	b.NumRetries = pointer.Int64Ptr(t.Schedule.NumRetries)

	props := t.Properties.(map[string]interface{})
	if err := mapstructure.Decode(props, b); err != nil {
		return errors.Wrap(err, "decode properties")
	}

	return nil
}

// accommodate for escaping of bash expansions, we can safely remove '\'
// as it's not a valid char in keyspace or table name
func unescapeFilters(strs []string) []string {
	for i := range strs {
		strs[i] = strings.ReplaceAll(strs[i], "\\", "")
	}
	return strs
}
