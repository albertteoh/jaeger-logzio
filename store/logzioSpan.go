package store

import (
	"encoding/json"

	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/plugin/storage/es/spanstore/dbmodel"
)

const (
	spanLogType = "jaegerSpan"
)

type logzioSpan struct {
	TraceID         dbmodel.TraceID        `json:"traceID"`
	OperationName   string                 `json:"operationName,omitempty"`
	SpanID          dbmodel.SpanID         `json:"spanID"`
	References      []dbmodel.Reference    `json:"references"`
	Flags           uint32                 `json:"flags,omitempty"`
	StartTime       uint64                 `json:"startTime"`
	StartTimeMillis uint64                 `json:"startTimeMillis"`
	Timestamp       uint64                 `json:"@timestamp"`
	Duration        uint64                 `json:"duration"`
	Tags            []dbmodel.KeyValue     `json:"JaegerTags,omitempty"`
	Tag             map[string]interface{} `json:"JaegerTag,omitempty"`
	Logs            []dbmodel.Log          `json:"logs"`
	Process         dbmodel.Process        `json:"process,omitempty"`
	Type            string                 `json:"type"`
}

func getTagsValues(tags []model.KeyValue) []string {
	var values []string
	for i := range tags {
		values = append(values, tags[i].VStr)
	}
	return values
}

// TransformToLogzioSpanBytes receives Jaeger span, converts it logzio span and return it as byte array.
// The main differences between Jaeger span and logzio span are arrays which are represented as maps
func TransformToLogzioSpanBytes(span *model.Span) ([]byte, error) {
	spanConverter := dbmodel.NewFromDomain(true, getTagsValues(span.Tags), "@")
	jsonSpan := spanConverter.FromDomainEmbedProcess(span)
	logzioSpan := logzioSpan{
		TraceID:         jsonSpan.TraceID,
		OperationName:   jsonSpan.OperationName,
		SpanID:          jsonSpan.SpanID,
		References:      jsonSpan.References,
		Flags:           jsonSpan.Flags,
		StartTime:       jsonSpan.StartTime,
		StartTimeMillis: jsonSpan.StartTimeMillis,
		Timestamp:       jsonSpan.StartTimeMillis,
		Duration:        jsonSpan.Duration,
		Tags:            jsonSpan.Tags,
		Tag:             jsonSpan.Tag,
		Process:         jsonSpan.Process,
		Type:            spanLogType,
	}
	return json.Marshal(logzioSpan)
}

func (span *logzioSpan) transformToDbModelSpan() *dbmodel.Span {
	return &dbmodel.Span{
		OperationName:	span.OperationName,
		Process:		span.Process,
		Tags:			span.Tags,
		Tag:			span.Tag,
		References:		span.References,
		Logs:			span.Logs,
		Duration:		span.Duration,
		StartTimeMillis:span.StartTimeMillis,
		StartTime:		span.StartTime,
		Flags:			span.Flags,
		SpanID:			span.SpanID,
		TraceID:		span.TraceID,
	}
}
