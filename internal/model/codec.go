package model

import "encoding/json"

func Encode[T any](v T) ([]byte, error)  { return json.Marshal(v) }
func Decode[T any](b []byte, v *T) error { return json.Unmarshal(b, v) }
func CloneAlbum(a Album) Album           { b, _ := Encode(a); var c Album; _ = Decode(b, &c); return c }
func CloneRecords(in []Record) []Record  { out := make([]Record, len(in)); copy(out, in); return out }
