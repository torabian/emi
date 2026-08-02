package external

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity1EntityUpdateInput
type Entity1EntityUpdateInput struct {
	Title              emigo.Nullable[string]                                     `json:"title" yaml:"title"`
	Items              emigo.ArrayNullable[Entity1EntityItems]                    `json:"items" yaml:"items"`
	Items2             emigo.ArrayNullable[Entity1EntityItems2]                   `json:"items2" yaml:"items2"`
	Items3             emigo.CollectionNullable[Entity2Entity]                    `json:"items3" yaml:"items3"`
	Items4             emigo.CollectionNullable[Entity2Entity]                    `json:"items4" yaml:"items4"`
	Owner              emigo.OneNullable[Entity2Entity]                           `json:"owner" yaml:"owner"`
	Manager            emigo.OneNullable[Entity2Entity]                           `json:"manager" yaml:"manager"`
	Content1           emigo.Nullable[Entity1EntityUpdateInputContent1]           `json:"content1" yaml:"content1"`
	Content2           emigo.Nullable[Entity1EntityUpdateInputContent2]           `json:"content2" yaml:"content2"`
	Complex1           *Money                                                     `json:"complex1" yaml:"complex1"`
	Subtitle           emigo.Nullable[string]                                     `json:"subtitle" yaml:"subtitle"`
	IsActive           emigo.Nullable[bool]                                       `json:"isActive" yaml:"isActive"`
	IsFeatured         emigo.Nullable[bool]                                       `json:"isFeatured" yaml:"isFeatured"`
	ViewCount          emigo.Nullable[int]                                        `json:"viewCount" yaml:"viewCount"`
	ViewCountOpt       emigo.Nullable[int]                                        `json:"viewCountOpt" yaml:"viewCountOpt"`
	SmallCount         emigo.Nullable[int32]                                      `json:"smallCount" yaml:"smallCount"`
	SmallCountOpt      emigo.Nullable[int32]                                      `json:"smallCountOpt" yaml:"smallCountOpt"`
	BigCount           emigo.Nullable[int64]                                      `json:"bigCount" yaml:"bigCount"`
	BigCountOpt        emigo.Nullable[int64]                                      `json:"bigCountOpt" yaml:"bigCountOpt"`
	Ratio32            emigo.Nullable[float32]                                    `json:"ratio32" yaml:"ratio32"`
	Ratio32Opt         emigo.Nullable[float32]                                    `json:"ratio32Opt" yaml:"ratio32Opt"`
	Ratio64            emigo.Nullable[float64]                                    `json:"ratio64" yaml:"ratio64"`
	Ratio64Opt         emigo.Nullable[float64]                                    `json:"ratio64Opt" yaml:"ratio64Opt"`
	Status             emigo.Nullable[string]                                     `json:"status" yaml:"status"`
	StatusOpt          emigo.Nullable[string]                                     `json:"statusOpt" yaml:"statusOpt"`
	Metadata           emigo.Nullable[map[string]string]                          `json:"metadata" yaml:"metadata"`
	MetadataOpt        emigo.Nullable[map[string]string]                          `json:"metadataOpt" yaml:"metadataOpt"`
	RawSettings        emigo.Nullable[map[string]string]                          `json:"rawSettings" yaml:"rawSettings"`
	Labels             emigo.Nullable[[]string]                                   `json:"labels" yaml:"labels"`
	LabelsOpt          emigo.Nullable[[]string]                                   `json:"labelsOpt" yaml:"labelsOpt"`
	Misc               interface{}                                                `json:"misc" yaml:"misc"`
	NestedContainer    emigo.Nullable[Entity1EntityUpdateInputNestedContainer]    `json:"nestedContainer" yaml:"nestedContainer"`
	NestedContainerOpt emigo.Nullable[Entity1EntityUpdateInputNestedContainerOpt] `json:"nestedContainerOpt" yaml:"nestedContainerOpt"`
}

// The base class definition for content1
type Entity1EntityUpdateInputContent1 struct {
	Item1 emigo.Nullable[int64] `json:"item1" yaml:"item1"`
}

// The base class definition for content2
type Entity1EntityUpdateInputContent2 struct {
	Item2 emigo.Nullable[int64] `json:"item2" yaml:"item2"`
}

// The base class definition for nestedContainer
type Entity1EntityUpdateInputNestedContainer struct {
	NestedInner emigo.Nullable[Entity1EntityUpdateInputNestedContainerNestedInner] `json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1EntityUpdateInputNestedContainerNestedInner struct {
	NestedItems emigo.ArrayNullable[Entity1EntityNestedContainerNestedInnerNestedItems] `json:"nestedItems" yaml:"nestedItems"`
	NestedOwner emigo.OneNullable[Entity2Entity]                                        `json:"nestedOwner" yaml:"nestedOwner"`
}

// The base class definition for nestedContainerOpt
type Entity1EntityUpdateInputNestedContainerOpt struct {
	NestedInner emigo.Nullable[Entity1EntityUpdateInputNestedContainerOptNestedInner] `json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1EntityUpdateInputNestedContainerOptNestedInner struct {
	NestedItemsOpt emigo.ArrayNullable[Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt] `json:"nestedItemsOpt" yaml:"nestedItemsOpt"`
}

func (x *Entity1EntityUpdateInput) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity1EntityUpdateInputCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "title",
			Type: "string?",
		},
		{
			Name: prefix + "items",
			Type: "complex",
		},
		{
			Name: prefix + "items2",
			Type: "complex",
		},
		{
			Name: prefix + "items3",
			Type: "complex",
		},
		{
			Name: prefix + "items4",
			Type: "complex",
		},
		{
			Name: prefix + "owner",
			Type: "complex",
		},
		{
			Name: prefix + "manager",
			Type: "complex",
		},
		{
			Name: prefix + "content1",
			Type: "object?",
		},
		{
			Name: prefix + "content2",
			Type: "object?",
		},
		{
			Name: prefix + "complex1",
			Type: "complex",
		},
		{
			Name: prefix + "subtitle",
			Type: "string?",
		},
		{
			Name: prefix + "is-active",
			Type: "bool?",
		},
		{
			Name: prefix + "is-featured",
			Type: "bool?",
		},
		{
			Name: prefix + "view-count",
			Type: "int?",
		},
		{
			Name: prefix + "view-count-opt",
			Type: "int?",
		},
		{
			Name: prefix + "small-count",
			Type: "int32?",
		},
		{
			Name: prefix + "small-count-opt",
			Type: "int32?",
		},
		{
			Name: prefix + "big-count",
			Type: "int64?",
		},
		{
			Name: prefix + "big-count-opt",
			Type: "int64?",
		},
		{
			Name: prefix + "ratio32",
			Type: "float32?",
		},
		{
			Name: prefix + "ratio32-opt",
			Type: "float32?",
		},
		{
			Name: prefix + "ratio64",
			Type: "float64?",
		},
		{
			Name: prefix + "ratio64-opt",
			Type: "float64?",
		},
		{
			Name: prefix + "status",
			Type: "enum?",
		},
		{
			Name: prefix + "status-opt",
			Type: "enum?",
		},
		{
			Name: prefix + "metadata",
			Type: "map?",
		},
		{
			Name: prefix + "metadata-opt",
			Type: "map?",
		},
		{
			Name: prefix + "raw-settings",
			Type: "map?",
		},
		{
			Name: prefix + "labels",
			Type: "slice?",
		},
		{
			Name: prefix + "labels-opt",
			Type: "slice?",
		},
		{
			Name: prefix + "misc",
			Type: "any",
		},
		{
			Name: prefix + "nested-container",
			Type: "object?",
		},
		{
			Name: prefix + "nested-container-opt",
			Type: "object?",
		},
	}
}
func CastEntity1EntityUpdateInputFromCli(c emigo.CliCastable) Entity1EntityUpdateInput {
	data := Entity1EntityUpdateInput{}
	if c.IsSet("title") {
		emigo.ParseNullable(c.String("title"), &data.Title)
	}
	if c.IsSet("items") {
		if u, ok := any(&data.Items).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("items")))
		}
	}
	if c.IsSet("items2") {
		if u, ok := any(&data.Items2).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("items2")))
		}
	}
	if c.IsSet("items3") {
		if u, ok := any(&data.Items3).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("items3")))
		}
	}
	if c.IsSet("items4") {
		if u, ok := any(&data.Items4).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("items4")))
		}
	}
	if c.IsSet("owner") {
		if u, ok := any(&data.Owner).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("owner")))
		}
	}
	if c.IsSet("manager") {
		if u, ok := any(&data.Manager).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("manager")))
		}
	}
	if c.IsSet("content1") {
		emigo.ParseNullable(c.String("content1"), &data.Content1)
	}
	if c.IsSet("content2") {
		emigo.ParseNullable(c.String("content2"), &data.Content2)
	}
	if c.IsSet("complex1") {
		if u, ok := any(&data.Complex1).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("complex1")))
		}
	}
	if c.IsSet("subtitle") {
		emigo.ParseNullable(c.String("subtitle"), &data.Subtitle)
	}
	if c.IsSet("is-active") {
		emigo.ParseNullable(c.String("is-active"), &data.IsActive)
	}
	if c.IsSet("is-featured") {
		emigo.ParseNullable(c.String("is-featured"), &data.IsFeatured)
	}
	if c.IsSet("view-count") {
		emigo.ParseNullable(c.String("view-count"), &data.ViewCount)
	}
	if c.IsSet("view-count-opt") {
		emigo.ParseNullable(c.String("view-count-opt"), &data.ViewCountOpt)
	}
	if c.IsSet("small-count") {
		emigo.ParseNullable(c.String("small-count"), &data.SmallCount)
	}
	if c.IsSet("small-count-opt") {
		emigo.ParseNullable(c.String("small-count-opt"), &data.SmallCountOpt)
	}
	if c.IsSet("big-count") {
		emigo.ParseNullable(c.String("big-count"), &data.BigCount)
	}
	if c.IsSet("big-count-opt") {
		emigo.ParseNullable(c.String("big-count-opt"), &data.BigCountOpt)
	}
	if c.IsSet("ratio32") {
		emigo.ParseNullable(c.String("ratio32"), &data.Ratio32)
	}
	if c.IsSet("ratio32-opt") {
		emigo.ParseNullable(c.String("ratio32-opt"), &data.Ratio32Opt)
	}
	if c.IsSet("ratio64") {
		emigo.ParseNullable(c.String("ratio64"), &data.Ratio64)
	}
	if c.IsSet("ratio64-opt") {
		emigo.ParseNullable(c.String("ratio64-opt"), &data.Ratio64Opt)
	}
	if c.IsSet("status") {
		emigo.ParseNullable(c.String("status"), &data.Status)
	}
	if c.IsSet("status-opt") {
		emigo.ParseNullable(c.String("status-opt"), &data.StatusOpt)
	}
	if c.IsSet("metadata") {
		emigo.ParseNullable(c.String("metadata"), &data.Metadata)
	}
	if c.IsSet("metadata-opt") {
		emigo.ParseNullable(c.String("metadata-opt"), &data.MetadataOpt)
	}
	if c.IsSet("raw-settings") {
		emigo.ParseNullable(c.String("raw-settings"), &data.RawSettings)
	}
	if c.IsSet("labels") {
		emigo.ParseNullable(c.String("labels"), &data.Labels)
	}
	if c.IsSet("labels-opt") {
		emigo.ParseNullable(c.String("labels-opt"), &data.LabelsOpt)
	}
	if c.IsSet("nested-container") {
		emigo.ParseNullable(c.String("nested-container"), &data.NestedContainer)
	}
	if c.IsSet("nested-container-opt") {
		emigo.ParseNullable(c.String("nested-container-opt"), &data.NestedContainerOpt)
	}
	return data
}
func GetEntity1EntityUpdateInputContent1CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item1",
			Type: "int64?",
		},
	}
}
func CastEntity1EntityUpdateInputContent1FromCli(c emigo.CliCastable) Entity1EntityUpdateInputContent1 {
	data := Entity1EntityUpdateInputContent1{}
	if c.IsSet("item1") {
		emigo.ParseNullable(c.String("item1"), &data.Item1)
	}
	return data
}
func GetEntity1EntityUpdateInputContent2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "int64?",
		},
	}
}
func CastEntity1EntityUpdateInputContent2FromCli(c emigo.CliCastable) Entity1EntityUpdateInputContent2 {
	data := Entity1EntityUpdateInputContent2{}
	if c.IsSet("item2") {
		emigo.ParseNullable(c.String("item2"), &data.Item2)
	}
	return data
}
func GetEntity1EntityUpdateInputNestedContainerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-inner",
			Type: "object?",
		},
	}
}
func CastEntity1EntityUpdateInputNestedContainerFromCli(c emigo.CliCastable) Entity1EntityUpdateInputNestedContainer {
	data := Entity1EntityUpdateInputNestedContainer{}
	if c.IsSet("nested-inner") {
		emigo.ParseNullable(c.String("nested-inner"), &data.NestedInner)
	}
	return data
}
func GetEntity1EntityUpdateInputNestedContainerNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items",
			Type: "complex",
		},
		{
			Name: prefix + "nested-owner",
			Type: "complex",
		},
	}
}
func CastEntity1EntityUpdateInputNestedContainerNestedInnerFromCli(c emigo.CliCastable) Entity1EntityUpdateInputNestedContainerNestedInner {
	data := Entity1EntityUpdateInputNestedContainerNestedInner{}
	if c.IsSet("nested-items") {
		if u, ok := any(&data.NestedItems).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("nested-items")))
		}
	}
	if c.IsSet("nested-owner") {
		if u, ok := any(&data.NestedOwner).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("nested-owner")))
		}
	}
	return data
}
func GetEntity1EntityUpdateInputNestedContainerOptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-inner",
			Type: "object?",
		},
	}
}
func CastEntity1EntityUpdateInputNestedContainerOptFromCli(c emigo.CliCastable) Entity1EntityUpdateInputNestedContainerOpt {
	data := Entity1EntityUpdateInputNestedContainerOpt{}
	if c.IsSet("nested-inner") {
		emigo.ParseNullable(c.String("nested-inner"), &data.NestedInner)
	}
	return data
}
func GetEntity1EntityUpdateInputNestedContainerOptNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items-opt",
			Type: "complex",
		},
	}
}
func CastEntity1EntityUpdateInputNestedContainerOptNestedInnerFromCli(c emigo.CliCastable) Entity1EntityUpdateInputNestedContainerOptNestedInner {
	data := Entity1EntityUpdateInputNestedContainerOptNestedInner{}
	if c.IsSet("nested-items-opt") {
		if u, ok := any(&data.NestedItemsOpt).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("nested-items-opt")))
		}
	}
	return data
}
