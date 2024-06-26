// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1experimental"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// Sample represents each record value encountered within a profiled program.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSample function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Sample struct {
	orig  *otlpprofiles.Sample
	state *internal.State
}

func newSample(orig *otlpprofiles.Sample, state *internal.State) Sample {
	return Sample{orig: orig, state: state}
}

// NewSample creates a new empty Sample.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewSample() Sample {
	state := internal.StateMutable
	return newSample(&otlpprofiles.Sample{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms Sample) MoveTo(dest Sample) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlpprofiles.Sample{}
}

// LocationIndex returns the LocationIndex associated with this Sample.
func (ms Sample) LocationIndex() pcommon.UInt64Slice {
	return pcommon.UInt64Slice(internal.NewUInt64Slice(&ms.orig.LocationIndex, ms.state))
}

// LocationsStartIndex returns the locationsstartindex associated with this Sample.
func (ms Sample) LocationsStartIndex() uint64 {
	return ms.orig.LocationsStartIndex
}

// SetLocationsStartIndex replaces the locationsstartindex associated with this Sample.
func (ms Sample) SetLocationsStartIndex(v uint64) {
	ms.state.AssertMutable()
	ms.orig.LocationsStartIndex = v
}

// LocationsLength returns the locationslength associated with this Sample.
func (ms Sample) LocationsLength() uint64 {
	return ms.orig.LocationsLength
}

// SetLocationsLength replaces the locationslength associated with this Sample.
func (ms Sample) SetLocationsLength(v uint64) {
	ms.state.AssertMutable()
	ms.orig.LocationsLength = v
}

// StacktraceIdIndex returns the stacktraceidindex associated with this Sample.
func (ms Sample) StacktraceIdIndex() uint32 {
	return ms.orig.StacktraceIdIndex
}

// SetStacktraceIdIndex replaces the stacktraceidindex associated with this Sample.
func (ms Sample) SetStacktraceIdIndex(v uint32) {
	ms.state.AssertMutable()
	ms.orig.StacktraceIdIndex = v
}

// Value returns the Value associated with this Sample.
func (ms Sample) Value() pcommon.Int64Slice {
	return pcommon.Int64Slice(internal.NewInt64Slice(&ms.orig.Value, ms.state))
}

// Label returns the Label associated with this Sample.
func (ms Sample) Label() LabelSlice {
	return newLabelSlice(&ms.orig.Label, ms.state)
}

// Attributes returns the Attributes associated with this Sample.
func (ms Sample) Attributes() pcommon.UInt64Slice {
	return pcommon.UInt64Slice(internal.NewUInt64Slice(&ms.orig.Attributes, ms.state))
}

// Link returns the link associated with this Sample.
func (ms Sample) Link() uint64 {
	return ms.orig.Link
}

// SetLink replaces the link associated with this Sample.
func (ms Sample) SetLink(v uint64) {
	ms.state.AssertMutable()
	ms.orig.Link = v
}

// TimestampsUnixNano returns the TimestampsUnixNano associated with this Sample.
func (ms Sample) TimestampsUnixNano() pcommon.UInt64Slice {
	return pcommon.UInt64Slice(internal.NewUInt64Slice(&ms.orig.TimestampsUnixNano, ms.state))
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms Sample) CopyTo(dest Sample) {
	dest.state.AssertMutable()
	ms.LocationIndex().CopyTo(dest.LocationIndex())
	dest.SetLocationsStartIndex(ms.LocationsStartIndex())
	dest.SetLocationsLength(ms.LocationsLength())
	dest.SetStacktraceIdIndex(ms.StacktraceIdIndex())
	ms.Value().CopyTo(dest.Value())
	ms.Label().CopyTo(dest.Label())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetLink(ms.Link())
	ms.TimestampsUnixNano().CopyTo(dest.TimestampsUnixNano())
}
