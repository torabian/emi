package external

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity1Dto
type Entity1Dto struct {
	UniqueId           emigo.Nullable[string]                       `json:"uniqueId" yaml:"uniqueId"`
	Title              string                                       `json:"title" yaml:"title"`
	Items              emigo.Array[Entity1DtoItems]                 `json:"items" yaml:"items"`
	Items2             emigo.ArrayNullable[Entity1DtoItems2]        `json:"items2" yaml:"items2"`
	Items3             emigo.CollectionNullable[Entity2Dto]         `json:"items3" yaml:"items3"`
	Items4             emigo.CollectionNullable[Entity2Dto]         `json:"items4" yaml:"items4"`
	Owner              emigo.OneNullable[Entity2Dto]                `json:"owner" yaml:"owner"`
	Manager            emigo.OneNullable[Entity2Dto]                `json:"manager" yaml:"manager"`
	Content1           Entity1DtoContent1                           `json:"content1" yaml:"content1"`
	Content2           emigo.Nullable[Entity1DtoContent2]           `json:"content2" yaml:"content2"`
	Complex1           Money                                        `json:"complex1" yaml:"complex1"`
	Subtitle           emigo.Nullable[string]                       `json:"subtitle" yaml:"subtitle"`
	IsActive           bool                                         `json:"isActive" yaml:"isActive"`
	IsFeatured         emigo.Nullable[bool]                         `json:"isFeatured" yaml:"isFeatured"`
	ViewCount          int                                          `json:"viewCount" yaml:"viewCount"`
	ViewCountOpt       emigo.Nullable[int]                          `json:"viewCountOpt" yaml:"viewCountOpt"`
	SmallCount         int32                                        `json:"smallCount" yaml:"smallCount"`
	SmallCountOpt      emigo.Nullable[int32]                        `json:"smallCountOpt" yaml:"smallCountOpt"`
	BigCount           int64                                        `json:"bigCount" yaml:"bigCount"`
	BigCountOpt        emigo.Nullable[int64]                        `json:"bigCountOpt" yaml:"bigCountOpt"`
	Ratio32            float32                                      `json:"ratio32" yaml:"ratio32"`
	Ratio32Opt         emigo.Nullable[float32]                      `json:"ratio32Opt" yaml:"ratio32Opt"`
	Ratio64            float64                                      `json:"ratio64" yaml:"ratio64"`
	Ratio64Opt         emigo.Nullable[float64]                      `json:"ratio64Opt" yaml:"ratio64Opt"`
	Status             string                                       `json:"status" yaml:"status"`
	StatusOpt          emigo.Nullable[string]                       `json:"statusOpt" yaml:"statusOpt"`
	Metadata           map[string]string                            `json:"metadata" yaml:"metadata"`
	MetadataOpt        emigo.Nullable[map[string]string]            `json:"metadataOpt" yaml:"metadataOpt"`
	RawSettings        map[string]string                            `json:"rawSettings" yaml:"rawSettings"`
	Labels             []string                                     `json:"labels" yaml:"labels"`
	LabelsOpt          emigo.Nullable[[]string]                     `json:"labelsOpt" yaml:"labelsOpt"`
	Misc               interface{}                                  `json:"misc" yaml:"misc"`
	NestedContainer    Entity1DtoNestedContainer                    `json:"nestedContainer" yaml:"nestedContainer"`
	NestedContainerOpt emigo.Nullable[Entity1DtoNestedContainerOpt] `json:"nestedContainerOpt" yaml:"nestedContainerOpt"`
}

// The base class definition for items
type Entity1DtoItems struct {
	Item2 string `json:"item2" yaml:"item2"`
}

// The base class definition for items2
type Entity1DtoItems2 struct {
	Item2 string `json:"item2" yaml:"item2"`
}

// The base class definition for content1
type Entity1DtoContent1 struct {
	Item1 emigo.Nullable[int64] `json:"item1" yaml:"item1"`
}

// The base class definition for content2
type Entity1DtoContent2 struct {
	Item2 emigo.Nullable[int64] `json:"item2" yaml:"item2"`
}

// The base class definition for nestedContainer
type Entity1DtoNestedContainer struct {
	NestedInner Entity1DtoNestedContainerNestedInner `json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1DtoNestedContainerNestedInner struct {
	NestedItems emigo.Array[Entity1DtoNestedContainerNestedInnerNestedItems] `json:"nestedItems" yaml:"nestedItems"`
	NestedOwner emigo.OneNullable[Entity2Dto]                                `json:"nestedOwner" yaml:"nestedOwner"`
}

// The base class definition for nestedItems
type Entity1DtoNestedContainerNestedInnerNestedItems struct {
	Label string `json:"label" yaml:"label"`
}

// The base class definition for nestedContainerOpt
type Entity1DtoNestedContainerOpt struct {
	NestedInner Entity1DtoNestedContainerOptNestedInner `json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1DtoNestedContainerOptNestedInner struct {
	NestedItemsOpt emigo.Array[Entity1DtoNestedContainerOptNestedInnerNestedItemsOpt] `json:"nestedItemsOpt" yaml:"nestedItemsOpt"`
}

// The base class definition for nestedItemsOpt
type Entity1DtoNestedContainerOptNestedInnerNestedItemsOpt struct {
	Label string `json:"label" yaml:"label"`
}

func (x *Entity1Dto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity1DtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "title",
			Type: "string",
		},
		{
			Name: prefix + "items",
			Type: "array",
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
			Name:     prefix + "content1",
			Type:     "object",
			Children: GetEntity1DtoContent1CliFlags("content1-"),
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
			Type: "bool",
		},
		{
			Name: prefix + "is-featured",
			Type: "bool?",
		},
		{
			Name: prefix + "view-count",
			Type: "int",
		},
		{
			Name: prefix + "view-count-opt",
			Type: "int?",
		},
		{
			Name: prefix + "small-count",
			Type: "int32",
		},
		{
			Name: prefix + "small-count-opt",
			Type: "int32?",
		},
		{
			Name: prefix + "big-count",
			Type: "int64",
		},
		{
			Name: prefix + "big-count-opt",
			Type: "int64?",
		},
		{
			Name: prefix + "ratio32",
			Type: "float32",
		},
		{
			Name: prefix + "ratio32-opt",
			Type: "float32?",
		},
		{
			Name: prefix + "ratio64",
			Type: "float64",
		},
		{
			Name: prefix + "ratio64-opt",
			Type: "float64?",
		},
		{
			Name: prefix + "status",
			Type: "enum",
		},
		{
			Name: prefix + "status-opt",
			Type: "enum?",
		},
		{
			Name: prefix + "metadata",
			Type: "map",
		},
		{
			Name: prefix + "metadata-opt",
			Type: "map?",
		},
		{
			Name: prefix + "raw-settings",
			Type: "map",
		},
		{
			Name: prefix + "labels",
			Type: "slice",
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
			Name:     prefix + "nested-container",
			Type:     "object",
			Children: GetEntity1DtoNestedContainerCliFlags("nested-container-"),
		},
		{
			Name: prefix + "nested-container-opt",
			Type: "object?",
		},
	}
}
func CastEntity1DtoFromCli(c emigo.CliCastable) Entity1Dto {
	data := Entity1Dto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastEntity1DtoItemsFromCli, "items", c)
	}
	if c.IsSet("items2") {
		data.Items2 = emigo.CapturePossibleArrayNullable(CastEntity1DtoItems2FromCli, "items2", c)
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
		data.Content1 = CastEntity1DtoContent1FromCli(c)
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
		data.IsActive = bool(c.Bool("is-active"))
	}
	if c.IsSet("is-featured") {
		emigo.ParseNullable(c.String("is-featured"), &data.IsFeatured)
	}
	if c.IsSet("view-count") {
		data.ViewCount = int(c.Int64("view-count"))
	}
	if c.IsSet("view-count-opt") {
		emigo.ParseNullable(c.String("view-count-opt"), &data.ViewCountOpt)
	}
	if c.IsSet("small-count") {
		data.SmallCount = int32(c.Int64("small-count"))
	}
	if c.IsSet("small-count-opt") {
		emigo.ParseNullable(c.String("small-count-opt"), &data.SmallCountOpt)
	}
	if c.IsSet("big-count") {
		data.BigCount = int64(c.Int64("big-count"))
	}
	if c.IsSet("big-count-opt") {
		emigo.ParseNullable(c.String("big-count-opt"), &data.BigCountOpt)
	}
	if c.IsSet("ratio32") {
		data.Ratio32 = float32(c.Float64("ratio32"))
	}
	if c.IsSet("ratio32-opt") {
		emigo.ParseNullable(c.String("ratio32-opt"), &data.Ratio32Opt)
	}
	if c.IsSet("ratio64") {
		data.Ratio64 = float64(c.Float64("ratio64"))
	}
	if c.IsSet("ratio64-opt") {
		emigo.ParseNullable(c.String("ratio64-opt"), &data.Ratio64Opt)
	}
	if c.IsSet("status-opt") {
		emigo.ParseNullable(c.String("status-opt"), &data.StatusOpt)
	}
	if c.IsSet("metadata-opt") {
		emigo.ParseNullable(c.String("metadata-opt"), &data.MetadataOpt)
	}
	if c.IsSet("labels") {
		emigo.InflatePossibleSlice(c.String("labels"), &data.Labels)
	}
	if c.IsSet("labels-opt") {
		emigo.ParseNullable(c.String("labels-opt"), &data.LabelsOpt)
	}
	if c.IsSet("nested-container") {
		data.NestedContainer = CastEntity1DtoNestedContainerFromCli(c)
	}
	if c.IsSet("nested-container-opt") {
		emigo.ParseNullable(c.String("nested-container-opt"), &data.NestedContainerOpt)
	}
	return data
}
func GetEntity1DtoItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "string",
		},
	}
}
func CastEntity1DtoItemsFromCli(c emigo.CliCastable) Entity1DtoItems {
	data := Entity1DtoItems{}
	if c.IsSet("item2") {
		data.Item2 = c.String("item2")
	}
	return data
}
func GetEntity1DtoItems2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "string",
		},
	}
}
func CastEntity1DtoItems2FromCli(c emigo.CliCastable) Entity1DtoItems2 {
	data := Entity1DtoItems2{}
	if c.IsSet("item2") {
		data.Item2 = c.String("item2")
	}
	return data
}
func GetEntity1DtoContent1CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item1",
			Type: "int64?",
		},
	}
}
func CastEntity1DtoContent1FromCli(c emigo.CliCastable) Entity1DtoContent1 {
	data := Entity1DtoContent1{}
	if c.IsSet("item1") {
		emigo.ParseNullable(c.String("item1"), &data.Item1)
	}
	return data
}
func GetEntity1DtoContent2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "int64?",
		},
	}
}
func CastEntity1DtoContent2FromCli(c emigo.CliCastable) Entity1DtoContent2 {
	data := Entity1DtoContent2{}
	if c.IsSet("item2") {
		emigo.ParseNullable(c.String("item2"), &data.Item2)
	}
	return data
}
func GetEntity1DtoNestedContainerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "nested-inner",
			Type:     "object",
			Children: GetEntity1DtoNestedContainerNestedInnerCliFlags("nested-inner-"),
		},
	}
}
func CastEntity1DtoNestedContainerFromCli(c emigo.CliCastable) Entity1DtoNestedContainer {
	data := Entity1DtoNestedContainer{}
	if c.IsSet("nested-inner") {
		data.NestedInner = CastEntity1DtoNestedContainerNestedInnerFromCli(c)
	}
	return data
}
func GetEntity1DtoNestedContainerNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items",
			Type: "array",
		},
		{
			Name: prefix + "nested-owner",
			Type: "one?",
		},
	}
}
func CastEntity1DtoNestedContainerNestedInnerFromCli(c emigo.CliCastable) Entity1DtoNestedContainerNestedInner {
	data := Entity1DtoNestedContainerNestedInner{}
	if c.IsSet("nested-items") {
		data.NestedItems = emigo.CapturePossibleArray(CastEntity1DtoNestedContainerNestedInnerNestedItemsFromCli, "nested-items", c)
	}
	if c.IsSet("nested-owner") {
		data.NestedOwner = emigo.CapturePossibleOneNullable(CastEntity2DtoFromCli, "nested-owner", c)
	}
	return data
}
func GetEntity1DtoNestedContainerNestedInnerNestedItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "label",
			Type: "string",
		},
	}
}
func CastEntity1DtoNestedContainerNestedInnerNestedItemsFromCli(c emigo.CliCastable) Entity1DtoNestedContainerNestedInnerNestedItems {
	data := Entity1DtoNestedContainerNestedInnerNestedItems{}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	return data
}
func GetEntity1DtoNestedContainerOptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "nested-inner",
			Type:     "object",
			Children: GetEntity1DtoNestedContainerOptNestedInnerCliFlags("nested-inner-"),
		},
	}
}
func CastEntity1DtoNestedContainerOptFromCli(c emigo.CliCastable) Entity1DtoNestedContainerOpt {
	data := Entity1DtoNestedContainerOpt{}
	if c.IsSet("nested-inner") {
		data.NestedInner = CastEntity1DtoNestedContainerOptNestedInnerFromCli(c)
	}
	return data
}
func GetEntity1DtoNestedContainerOptNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items-opt",
			Type: "array",
		},
	}
}
func CastEntity1DtoNestedContainerOptNestedInnerFromCli(c emigo.CliCastable) Entity1DtoNestedContainerOptNestedInner {
	data := Entity1DtoNestedContainerOptNestedInner{}
	if c.IsSet("nested-items-opt") {
		data.NestedItemsOpt = emigo.CapturePossibleArray(CastEntity1DtoNestedContainerOptNestedInnerNestedItemsOptFromCli, "nested-items-opt", c)
	}
	return data
}
func GetEntity1DtoNestedContainerOptNestedInnerNestedItemsOptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "label",
			Type: "string",
		},
	}
}
func CastEntity1DtoNestedContainerOptNestedInnerNestedItemsOptFromCli(c emigo.CliCastable) Entity1DtoNestedContainerOptNestedInnerNestedItemsOpt {
	data := Entity1DtoNestedContainerOptNestedInnerNestedItemsOpt{}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	return data
}
