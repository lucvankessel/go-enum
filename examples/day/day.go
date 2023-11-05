//go:generate go run --mod=mod github.com/klippa-app/go-enum -case=upper_snake -gql=full -json -bson -xml -ent
package day

type Day string

const (
    Unknown Day = "Uknown" //enum:invalid
    Monday  Day = "MonDay"
    Tuesday Day = "Tuesday"
    Wednesday Day = "wednesDay"
    Thursday Day = "Thursday"
    Friday Day = "Friday"
    Saturday Day = "Saturday"
    Sunday
)