package external

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity1OptionalDto
type Entity1OptionalDto struct {
	UniqueId           emigo.Nullable[string]                               `json:"uniqueId" yaml:"uniqueId"`
	Title              emigo.Nullable[string]                               `json:"title" yaml:"title"`
	Items              emigo.ArrayNullable[Entity1OptionalDtoItems]         `json:"items" yaml:"items"`
	Items2             emigo.ArrayNullable[Entity1OptionalDtoItems2]        `json:"items2" yaml:"items2"`
	Items3             emigo.CollectionNullable[Entity2Dto]                 `json:"items3" yaml:"items3"`
	Items4             emigo.CollectionNullable[Entity2Dto]                 `json:"items4" yaml:"items4"`
	Owner              emigo.OneNullable[Entity2Dto]                        `json:"owner" yaml:"owner"`
	Manager            emigo.OneNullable[Entity2Dto]                        `json:"manager" yaml:"manager"`
	Content1           emigo.Nullable[Entity1OptionalDtoContent1]           `json:"content1" yaml:"content1"`
	Content2           emigo.Nullable[Entity1OptionalDtoContent2]           `json:"content2" yaml:"content2"`
	Complex1           Money                                                `json:"complex1" yaml:"complex1"`
	Subtitle           emigo.Nullable[string]                               `json:"subtitle" yaml:"subtitle"`
	IsActive           emigo.Nullable[bool]                                 `json:"isActive" yaml:"isActive"`
	IsFeatured         emigo.Nullable[bool]                                 `json:"isFeatured" yaml:"isFeatured"`
	ViewCount          emigo.Nullable[int]                                  `json:"viewCount" yaml:"viewCount"`
	ViewCountOpt       emigo.Nullable[int]                                  `json:"viewCountOpt" yaml:"viewCountOpt"`
	SmallCount         emigo.Nullable[int32]                                `json:"smallCount" yaml:"smallCount"`
	SmallCountOpt      emigo.Nullable[int32]                                `json:"smallCountOpt" yaml:"smallCountOpt"`
	BigCount           emigo.Nullable[int64]                                `json:"bigCount" yaml:"bigCount"`
	BigCountOpt        emigo.Nullable[int64]                                `json:"bigCountOpt" yaml:"bigCountOpt"`
	Ratio32            emigo.Nullable[float32]                              `json:"ratio32" yaml:"ratio32"`
	Ratio32Opt         emigo.Nullable[float32]                              `json:"ratio32Opt" yaml:"ratio32Opt"`
	Ratio64            emigo.Nullable[float64]                              `json:"ratio64" yaml:"ratio64"`
	Ratio64Opt         emigo.Nullable[float64]                              `json:"ratio64Opt" yaml:"ratio64Opt"`
	Status             emigo.Nullable[string]                               `json:"status" yaml:"status"`
	StatusOpt          emigo.Nullable[string]                               `json:"statusOpt" yaml:"statusOpt"`
	Metadata           emigo.Nullable[map[string]string]                    `json:"metadata" yaml:"metadata"`
	MetadataOpt        emigo.Nullable[map[string]string]                    `json:"metadataOpt" yaml:"metadataOpt"`
	RawSettings        emigo.Nullable[map[string]string]                    `json:"rawSettings" yaml:"rawSettings"`
	Labels             emigo.Nullable[[]string]                             `json:"labels" yaml:"labels"`
	LabelsOpt          emigo.Nullable[[]string]                             `json:"labelsOpt" yaml:"labelsOpt"`
	Misc               interface{}                                          `json:"misc" yaml:"misc"`
	NestedContainer    emigo.Nullable[Entity1OptionalDtoNestedContainer]    `json:"nestedContainer" yaml:"nestedContainer"`
	NestedContainerOpt emigo.Nullable[Entity1OptionalDtoNestedContainerOpt] `json:"nestedContainerOpt" yaml:"nestedContainerOpt"`
}

// The base class definition for items
type Entity1OptionalDtoItems struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Item2    string                 `json:"item2" yaml:"item2"`
}

// The base class definition for items2
type Entity1OptionalDtoItems2 struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Item2    string                 `json:"item2" yaml:"item2"`
}

// The base class definition for content1
type Entity1OptionalDtoContent1 struct {
	Item1 emigo.Nullable[int64] `json:"item1" yaml:"item1"`
}

// The base class definition for content2
type Entity1OptionalDtoContent2 struct {
	Item2 emigo.Nullable[int64] `json:"item2" yaml:"item2"`
}

// The base class definition for nestedContainer
type Entity1OptionalDtoNestedContainer struct {
	NestedInner emigo.Nullable[Entity1OptionalDtoNestedContainerNestedInner] `json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1OptionalDtoNestedContainerNestedInner struct {
	NestedItems emigo.ArrayNullable[Entity1OptionalDtoNestedContainerNestedInnerNestedItems] `json:"nestedItems" yaml:"nestedItems"`
	NestedOwner emigo.OneNullable[Entity2Dto]                                                `json:"nestedOwner" yaml:"nestedOwner"`
}

// The base class definition for nestedItems
type Entity1OptionalDtoNestedContainerNestedInnerNestedItems struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Label    string                 `json:"label" yaml:"label"`
}

// The base class definition for nestedContainerOpt
type Entity1OptionalDtoNestedContainerOpt struct {
	NestedInner emigo.Nullable[Entity1OptionalDtoNestedContainerOptNestedInner] `json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1OptionalDtoNestedContainerOptNestedInner struct {
	NestedItemsOpt emigo.ArrayNullable[Entity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOpt] `json:"nestedItemsOpt" yaml:"nestedItemsOpt"`
}

// The base class definition for nestedItemsOpt
type Entity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOpt struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Label    string                 `json:"label" yaml:"label"`
}

func (x *Entity1OptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity1OptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "title",
			Type: "string?",
		},
		{
			Name: prefix + "items",
			Type: "array?",
		},
		{
			Name: prefix + "items2",
			Type: "array?",
		},
		{
			Name: prefix + "items3",
			Type: "collection?",
		},
		{
			Name: prefix + "items4",
			Type: "collection?",
		},
		{
			Name: prefix + "owner",
			Type: "one?",
		},
		{
			Name: prefix + "manager",
			Type: "one?",
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
func CastEntity1OptionalDtoFromCli(c emigo.CliCastable) Entity1OptionalDto {
	data := Entity1OptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("title") {
		emigo.ParseNullable(c.String("title"), &data.Title)
	}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArrayNullable(CastEntity1OptionalDtoItemsFromCli, "items", c)
	}
	if c.IsSet("items2") {
		data.Items2 = emigo.CapturePossibleArrayNullable(CastEntity1OptionalDtoItems2FromCli, "items2", c)
	}
	if c.IsSet("items3") {
		data.Items3 = emigo.CapturePossibleCollectionNullable(CastEntity2DtoFromCli, "items3", c)
	}
	if c.IsSet("items4") {
		data.Items4 = emigo.CapturePossibleCollectionNullable(CastEntity2DtoFromCli, "items4", c)
	}
	if c.IsSet("owner") {
		data.Owner = emigo.CapturePossibleOneNullable(CastEntity2DtoFromCli, "owner", c)
	}
	if c.IsSet("manager") {
		data.Manager = emigo.CapturePossibleOneNullable(CastEntity2DtoFromCli, "manager", c)
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
func GetEntity1OptionalDtoItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "item2",
			Type: "string",
		},
	}
}
func CastEntity1OptionalDtoItemsFromCli(c emigo.CliCastable) Entity1OptionalDtoItems {
	data := Entity1OptionalDtoItems{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("item2") {
		data.Item2 = c.String("item2")
	}
	return data
}
func GetEntity1OptionalDtoItems2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "item2",
			Type: "string",
		},
	}
}
func CastEntity1OptionalDtoItems2FromCli(c emigo.CliCastable) Entity1OptionalDtoItems2 {
	data := Entity1OptionalDtoItems2{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("item2") {
		data.Item2 = c.String("item2")
	}
	return data
}
func GetEntity1OptionalDtoContent1CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item1",
			Type: "int64?",
		},
	}
}
func CastEntity1OptionalDtoContent1FromCli(c emigo.CliCastable) Entity1OptionalDtoContent1 {
	data := Entity1OptionalDtoContent1{}
	if c.IsSet("item1") {
		emigo.ParseNullable(c.String("item1"), &data.Item1)
	}
	return data
}
func GetEntity1OptionalDtoContent2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "int64?",
		},
	}
}
func CastEntity1OptionalDtoContent2FromCli(c emigo.CliCastable) Entity1OptionalDtoContent2 {
	data := Entity1OptionalDtoContent2{}
	if c.IsSet("item2") {
		emigo.ParseNullable(c.String("item2"), &data.Item2)
	}
	return data
}
func GetEntity1OptionalDtoNestedContainerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-inner",
			Type: "object?",
		},
	}
}
func CastEntity1OptionalDtoNestedContainerFromCli(c emigo.CliCastable) Entity1OptionalDtoNestedContainer {
	data := Entity1OptionalDtoNestedContainer{}
	if c.IsSet("nested-inner") {
		emigo.ParseNullable(c.String("nested-inner"), &data.NestedInner)
	}
	return data
}
func GetEntity1OptionalDtoNestedContainerNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items",
			Type: "array?",
		},
		{
			Name: prefix + "nested-owner",
			Type: "one?",
		},
	}
}
func CastEntity1OptionalDtoNestedContainerNestedInnerFromCli(c emigo.CliCastable) Entity1OptionalDtoNestedContainerNestedInner {
	data := Entity1OptionalDtoNestedContainerNestedInner{}
	if c.IsSet("nested-items") {
		data.NestedItems = emigo.CapturePossibleArrayNullable(CastEntity1OptionalDtoNestedContainerNestedInnerNestedItemsFromCli, "nested-items", c)
	}
	if c.IsSet("nested-owner") {
		data.NestedOwner = emigo.CapturePossibleOneNullable(CastEntity2DtoFromCli, "nested-owner", c)
	}
	return data
}
func GetEntity1OptionalDtoNestedContainerNestedInnerNestedItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "label",
			Type: "string",
		},
	}
}
func CastEntity1OptionalDtoNestedContainerNestedInnerNestedItemsFromCli(c emigo.CliCastable) Entity1OptionalDtoNestedContainerNestedInnerNestedItems {
	data := Entity1OptionalDtoNestedContainerNestedInnerNestedItems{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	return data
}
func GetEntity1OptionalDtoNestedContainerOptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-inner",
			Type: "object?",
		},
	}
}
func CastEntity1OptionalDtoNestedContainerOptFromCli(c emigo.CliCastable) Entity1OptionalDtoNestedContainerOpt {
	data := Entity1OptionalDtoNestedContainerOpt{}
	if c.IsSet("nested-inner") {
		emigo.ParseNullable(c.String("nested-inner"), &data.NestedInner)
	}
	return data
}
func GetEntity1OptionalDtoNestedContainerOptNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items-opt",
			Type: "array?",
		},
	}
}
func CastEntity1OptionalDtoNestedContainerOptNestedInnerFromCli(c emigo.CliCastable) Entity1OptionalDtoNestedContainerOptNestedInner {
	data := Entity1OptionalDtoNestedContainerOptNestedInner{}
	if c.IsSet("nested-items-opt") {
		data.NestedItemsOpt = emigo.CapturePossibleArrayNullable(CastEntity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOptFromCli, "nested-items-opt", c)
	}
	return data
}
func GetEntity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "label",
			Type: "string",
		},
	}
}
func CastEntity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOptFromCli(c emigo.CliCastable) Entity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOpt {
	data := Entity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOpt{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	return data
}
