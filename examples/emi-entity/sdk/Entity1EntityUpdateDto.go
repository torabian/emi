package external

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity1EntityUpdateDto
type Entity1EntityUpdateDto struct {
	Title              emigo.Nullable[string]                                   `json:"title" yaml:"title"`
	Items              emigo.ArrayNullable[Entity1EntityUpdateDtoItems]         `json:"items" yaml:"items"`
	Items2             emigo.ArrayNullable[Entity1EntityUpdateDtoItems2]        `json:"items2" yaml:"items2"`
	Items3             emigo.CollectionNullable[Entity2Entity]                  `json:"items3" yaml:"items3"`
	Items4             emigo.CollectionNullable[Entity2Entity]                  `json:"items4" yaml:"items4"`
	Owner              emigo.OneNullable[Entity2Entity]                         `json:"owner" yaml:"owner"`
	Manager            emigo.OneNullable[Entity2Entity]                         `json:"manager" yaml:"manager"`
	Content1           emigo.Nullable[Entity1EntityUpdateDtoContent1]           `json:"content1" yaml:"content1"`
	Content2           emigo.Nullable[Entity1EntityUpdateDtoContent2]           `json:"content2" yaml:"content2"`
	Complex1           Money                                                    `json:"complex1" yaml:"complex1"`
	Subtitle           emigo.Nullable[string]                                   `json:"subtitle" yaml:"subtitle"`
	IsActive           emigo.Nullable[bool]                                     `json:"isActive" yaml:"isActive"`
	IsFeatured         emigo.Nullable[bool]                                     `json:"isFeatured" yaml:"isFeatured"`
	ViewCount          emigo.Nullable[int]                                      `json:"viewCount" yaml:"viewCount"`
	ViewCountOpt       emigo.Nullable[int]                                      `json:"viewCountOpt" yaml:"viewCountOpt"`
	SmallCount         emigo.Nullable[int32]                                    `json:"smallCount" yaml:"smallCount"`
	SmallCountOpt      emigo.Nullable[int32]                                    `json:"smallCountOpt" yaml:"smallCountOpt"`
	BigCount           emigo.Nullable[int64]                                    `json:"bigCount" yaml:"bigCount"`
	BigCountOpt        emigo.Nullable[int64]                                    `json:"bigCountOpt" yaml:"bigCountOpt"`
	Ratio32            emigo.Nullable[float32]                                  `json:"ratio32" yaml:"ratio32"`
	Ratio32Opt         emigo.Nullable[float32]                                  `json:"ratio32Opt" yaml:"ratio32Opt"`
	Ratio64            emigo.Nullable[float64]                                  `json:"ratio64" yaml:"ratio64"`
	Ratio64Opt         emigo.Nullable[float64]                                  `json:"ratio64Opt" yaml:"ratio64Opt"`
	Status             emigo.Nullable[string]                                   `json:"status" yaml:"status"`
	StatusOpt          emigo.Nullable[string]                                   `json:"statusOpt" yaml:"statusOpt"`
	Metadata           emigo.Nullable[map[string]string]                        `json:"metadata" yaml:"metadata"`
	MetadataOpt        emigo.Nullable[map[string]string]                        `json:"metadataOpt" yaml:"metadataOpt"`
	RawSettings        emigo.Nullable[map[string]string]                        `json:"rawSettings" yaml:"rawSettings"`
	Labels             emigo.Nullable[[]string]                                 `json:"labels" yaml:"labels"`
	LabelsOpt          emigo.Nullable[[]string]                                 `json:"labelsOpt" yaml:"labelsOpt"`
	Misc               interface{}                                              `json:"misc" yaml:"misc"`
	NestedContainer    emigo.Nullable[Entity1EntityUpdateDtoNestedContainer]    `json:"nestedContainer" yaml:"nestedContainer"`
	NestedContainerOpt emigo.Nullable[Entity1EntityUpdateDtoNestedContainerOpt] `json:"nestedContainerOpt" yaml:"nestedContainerOpt"`
}

// The base class definition for items
type Entity1EntityUpdateDtoItems struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Item2    string                 `json:"item2" yaml:"item2"`
}

// The base class definition for items2
type Entity1EntityUpdateDtoItems2 struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Item2    string                 `json:"item2" yaml:"item2"`
}

// The base class definition for content1
type Entity1EntityUpdateDtoContent1 struct {
	Item1 emigo.Nullable[int64] `json:"item1" yaml:"item1"`
}

// The base class definition for content2
type Entity1EntityUpdateDtoContent2 struct {
	Item2 emigo.Nullable[int64] `json:"item2" yaml:"item2"`
}

// The base class definition for nestedContainer
type Entity1EntityUpdateDtoNestedContainer struct {
	NestedInner emigo.Nullable[Entity1EntityUpdateDtoNestedContainerNestedInner] `json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1EntityUpdateDtoNestedContainerNestedInner struct {
	NestedItems emigo.ArrayNullable[Entity1EntityUpdateDtoNestedContainerNestedInnerNestedItems] `json:"nestedItems" yaml:"nestedItems"`
	NestedOwner emigo.OneNullable[Entity2Entity]                                                 `json:"nestedOwner" yaml:"nestedOwner"`
}

// The base class definition for nestedItems
type Entity1EntityUpdateDtoNestedContainerNestedInnerNestedItems struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Label    string                 `json:"label" yaml:"label"`
}

// The base class definition for nestedContainerOpt
type Entity1EntityUpdateDtoNestedContainerOpt struct {
	NestedInner emigo.Nullable[Entity1EntityUpdateDtoNestedContainerOptNestedInner] `json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1EntityUpdateDtoNestedContainerOptNestedInner struct {
	NestedItemsOpt emigo.ArrayNullable[Entity1EntityUpdateDtoNestedContainerOptNestedInnerNestedItemsOpt] `json:"nestedItemsOpt" yaml:"nestedItemsOpt"`
}

// The base class definition for nestedItemsOpt
type Entity1EntityUpdateDtoNestedContainerOptNestedInnerNestedItemsOpt struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Label    string                 `json:"label" yaml:"label"`
}

func (x *Entity1EntityUpdateDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity1EntityUpdateDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
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
func CastEntity1EntityUpdateDtoFromCli(c emigo.CliCastable) Entity1EntityUpdateDto {
	data := Entity1EntityUpdateDto{}
	if c.IsSet("title") {
		emigo.ParseNullable(c.String("title"), &data.Title)
	}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArrayNullable(CastEntity1EntityUpdateDtoItemsFromCli, "items", c)
	}
	if c.IsSet("items2") {
		data.Items2 = emigo.CapturePossibleArrayNullable(CastEntity1EntityUpdateDtoItems2FromCli, "items2", c)
	}
	if c.IsSet("items3") {
		data.Items3 = emigo.CapturePossibleCollectionNullable(CastEntity2EntityFromCli, "items3", c)
	}
	if c.IsSet("items4") {
		data.Items4 = emigo.CapturePossibleCollectionNullable(CastEntity2EntityFromCli, "items4", c)
	}
	if c.IsSet("owner") {
		data.Owner = emigo.CapturePossibleOneNullable(CastEntity2EntityFromCli, "owner", c)
	}
	if c.IsSet("manager") {
		data.Manager = emigo.CapturePossibleOneNullable(CastEntity2EntityFromCli, "manager", c)
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
func GetEntity1EntityUpdateDtoItemsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEntity1EntityUpdateDtoItemsFromCli(c emigo.CliCastable) Entity1EntityUpdateDtoItems {
	data := Entity1EntityUpdateDtoItems{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("item2") {
		data.Item2 = c.String("item2")
	}
	return data
}
func GetEntity1EntityUpdateDtoItems2CliFlags(prefix string) []emigo.CliFlag {
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
func CastEntity1EntityUpdateDtoItems2FromCli(c emigo.CliCastable) Entity1EntityUpdateDtoItems2 {
	data := Entity1EntityUpdateDtoItems2{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("item2") {
		data.Item2 = c.String("item2")
	}
	return data
}
func GetEntity1EntityUpdateDtoContent1CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item1",
			Type: "int64?",
		},
	}
}
func CastEntity1EntityUpdateDtoContent1FromCli(c emigo.CliCastable) Entity1EntityUpdateDtoContent1 {
	data := Entity1EntityUpdateDtoContent1{}
	if c.IsSet("item1") {
		emigo.ParseNullable(c.String("item1"), &data.Item1)
	}
	return data
}
func GetEntity1EntityUpdateDtoContent2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "int64?",
		},
	}
}
func CastEntity1EntityUpdateDtoContent2FromCli(c emigo.CliCastable) Entity1EntityUpdateDtoContent2 {
	data := Entity1EntityUpdateDtoContent2{}
	if c.IsSet("item2") {
		emigo.ParseNullable(c.String("item2"), &data.Item2)
	}
	return data
}
func GetEntity1EntityUpdateDtoNestedContainerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-inner",
			Type: "object?",
		},
	}
}
func CastEntity1EntityUpdateDtoNestedContainerFromCli(c emigo.CliCastable) Entity1EntityUpdateDtoNestedContainer {
	data := Entity1EntityUpdateDtoNestedContainer{}
	if c.IsSet("nested-inner") {
		emigo.ParseNullable(c.String("nested-inner"), &data.NestedInner)
	}
	return data
}
func GetEntity1EntityUpdateDtoNestedContainerNestedInnerCliFlags(prefix string) []emigo.CliFlag {
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
func CastEntity1EntityUpdateDtoNestedContainerNestedInnerFromCli(c emigo.CliCastable) Entity1EntityUpdateDtoNestedContainerNestedInner {
	data := Entity1EntityUpdateDtoNestedContainerNestedInner{}
	if c.IsSet("nested-items") {
		data.NestedItems = emigo.CapturePossibleArrayNullable(CastEntity1EntityUpdateDtoNestedContainerNestedInnerNestedItemsFromCli, "nested-items", c)
	}
	if c.IsSet("nested-owner") {
		data.NestedOwner = emigo.CapturePossibleOneNullable(CastEntity2EntityFromCli, "nested-owner", c)
	}
	return data
}
func GetEntity1EntityUpdateDtoNestedContainerNestedInnerNestedItemsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEntity1EntityUpdateDtoNestedContainerNestedInnerNestedItemsFromCli(c emigo.CliCastable) Entity1EntityUpdateDtoNestedContainerNestedInnerNestedItems {
	data := Entity1EntityUpdateDtoNestedContainerNestedInnerNestedItems{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	return data
}
func GetEntity1EntityUpdateDtoNestedContainerOptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-inner",
			Type: "object?",
		},
	}
}
func CastEntity1EntityUpdateDtoNestedContainerOptFromCli(c emigo.CliCastable) Entity1EntityUpdateDtoNestedContainerOpt {
	data := Entity1EntityUpdateDtoNestedContainerOpt{}
	if c.IsSet("nested-inner") {
		emigo.ParseNullable(c.String("nested-inner"), &data.NestedInner)
	}
	return data
}
func GetEntity1EntityUpdateDtoNestedContainerOptNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items-opt",
			Type: "array?",
		},
	}
}
func CastEntity1EntityUpdateDtoNestedContainerOptNestedInnerFromCli(c emigo.CliCastable) Entity1EntityUpdateDtoNestedContainerOptNestedInner {
	data := Entity1EntityUpdateDtoNestedContainerOptNestedInner{}
	if c.IsSet("nested-items-opt") {
		data.NestedItemsOpt = emigo.CapturePossibleArrayNullable(CastEntity1EntityUpdateDtoNestedContainerOptNestedInnerNestedItemsOptFromCli, "nested-items-opt", c)
	}
	return data
}
func GetEntity1EntityUpdateDtoNestedContainerOptNestedInnerNestedItemsOptCliFlags(prefix string) []emigo.CliFlag {
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
func CastEntity1EntityUpdateDtoNestedContainerOptNestedInnerNestedItemsOptFromCli(c emigo.CliCastable) Entity1EntityUpdateDtoNestedContainerOptNestedInnerNestedItemsOpt {
	data := Entity1EntityUpdateDtoNestedContainerOptNestedInnerNestedItemsOpt{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	return data
}
