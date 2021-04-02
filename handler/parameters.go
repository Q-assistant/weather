package handler

import (
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

type Date struct {
	IsRange bool
	Start   time.Time
	End     time.Time
}

func getDate(v *structpb.ListValue) (*Date, error) {
	date := &Date{}
	strct := v.Values[0].GetStructValue()
	tpe := strct.Fields["type"].GetStringValue()

	if tpe == "datetimerange" {
		date.IsRange = true

		start, err := time.Parse(timeLayout, strct.Fields["values"].GetListValue().Values[0].GetStructValue().Fields["resolution"].GetListValue().Values[0].GetStructValue().Fields["start"].GetStringValue())
		if err != nil {
			return nil, err
		}
		end, err := time.Parse(timeLayout, strct.Fields["values"].GetListValue().Values[0].GetStructValue().Fields["resolution"].GetListValue().Values[0].GetStructValue().Fields["end"].GetStringValue())
		if err != nil {
			return nil, err
		}

		date.Start = start
		date.End = end
	}

	return date, nil
}

func getWeather(v *structpb.ListValue) string {
	return v.Values[0].GetListValue().Values[0].GetStringValue()
}
