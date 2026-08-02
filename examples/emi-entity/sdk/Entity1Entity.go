package external

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity1Entity
type Entity1Entity struct {
	Id                    int64                                           `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId              string                                          `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Title                 string                                          `json:"title" yaml:"title"`
	Items                 emigo.Array[Entity1EntityItems]                 `gorm:"-" json:"items" yaml:"items"`
	Items2                emigo.ArrayNullable[Entity1EntityItems2]        `gorm:"-" json:"items2" yaml:"items2"`
	Items3                emigo.Collection[Entity2Entity]                 `gorm:"-" json:"items3" yaml:"items3"`
	Items4                emigo.CollectionNullable[Entity2Entity]         `gorm:"-" json:"items4" yaml:"items4"`
	Owner                 emigo.One[Entity2Entity]                        `gorm:"-" json:"owner" yaml:"owner"`
	Manager               emigo.OneNullable[Entity2Entity]                `gorm:"-" json:"manager" yaml:"manager"`
	Content1              Entity1EntityContent1                           `gorm:"embedded" json:"content1" yaml:"content1"`
	Content2              emigo.Nullable[Entity1EntityContent2]           `gorm:"-" json:"content2" yaml:"content2"`
	Complex1              Money                                           `json:"complex1" yaml:"complex1"`
	Subtitle              emigo.Nullable[string]                          `json:"subtitle" yaml:"subtitle"`
	IsActive              bool                                            `json:"isActive" yaml:"isActive"`
	IsFeatured            emigo.Nullable[bool]                            `json:"isFeatured" yaml:"isFeatured"`
	ViewCount             int                                             `json:"viewCount" yaml:"viewCount"`
	ViewCountOpt          emigo.Nullable[int]                             `json:"viewCountOpt" yaml:"viewCountOpt"`
	SmallCount            int32                                           `json:"smallCount" yaml:"smallCount"`
	SmallCountOpt         emigo.Nullable[int32]                           `json:"smallCountOpt" yaml:"smallCountOpt"`
	BigCount              int64                                           `json:"bigCount" yaml:"bigCount"`
	BigCountOpt           emigo.Nullable[int64]                           `json:"bigCountOpt" yaml:"bigCountOpt"`
	Ratio32               float32                                         `json:"ratio32" yaml:"ratio32"`
	Ratio32Opt            emigo.Nullable[float32]                         `json:"ratio32Opt" yaml:"ratio32Opt"`
	Ratio64               float64                                         `json:"ratio64" yaml:"ratio64"`
	Ratio64Opt            emigo.Nullable[float64]                         `json:"ratio64Opt" yaml:"ratio64Opt"`
	Status                string                                          `json:"status" yaml:"status"`
	StatusOpt             emigo.Nullable[string]                          `json:"statusOpt" yaml:"statusOpt"`
	Metadata              map[string]string                               `gorm:"serializer:json" json:"metadata" yaml:"metadata"`
	MetadataOpt           emigo.Nullable[map[string]string]               `json:"metadataOpt" yaml:"metadataOpt"`
	RawSettings           map[string]string                               `gorm:"serializer:json;type:jsonb" json:"rawSettings" yaml:"rawSettings"`
	Labels                []string                                        `gorm:"serializer:json" json:"labels" yaml:"labels"`
	LabelsOpt             emigo.Nullable[[]string]                        `json:"labelsOpt" yaml:"labelsOpt"`
	Misc                  interface{}                                     `gorm:"serializer:json" json:"misc" yaml:"misc"`
	NestedContainer       Entity1EntityNestedContainer                    `gorm:"embedded" json:"nestedContainer" yaml:"nestedContainer"`
	NestedContainerOpt    emigo.Nullable[Entity1EntityNestedContainerOpt] `gorm:"-" json:"nestedContainerOpt" yaml:"nestedContainerOpt"`
	ItemsRow              []*Entity1EntityItems                           `gorm:"foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE" json:"-" yaml:"-"`
	Items2Row             []*Entity1EntityItems2                          `gorm:"foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE" json:"-" yaml:"-"`
	Items3Row             []*Entity2Entity                                `gorm:"many2many:entity1_items3;foreignKey:Id;references:Id" json:"-" yaml:"-"`
	Items4Row             []*Entity2Entity                                `gorm:"many2many:entity1_items4;foreignKey:Id;references:Id" json:"-" yaml:"-"`
	OwnerId               int64                                           `gorm:"index" json:"-" yaml:"-"`
	OwnerRow              *Entity2Entity                                  `gorm:"foreignKey:OwnerId;references:Id" json:"-" yaml:"-"`
	ManagerId             int64                                           `gorm:"index" json:"-" yaml:"-"`
	ManagerRow            *Entity2Entity                                  `gorm:"foreignKey:ManagerId;references:Id" json:"-" yaml:"-"`
	Content2Row           *Entity1EntityContent2                          `gorm:"embedded" json:"-" yaml:"-"`
	NestedContainerOptRow *Entity1EntityNestedContainerOpt                `gorm:"embedded" json:"-" yaml:"-"`
}

// The base class definition for items
type Entity1EntityItems struct {
	Item2    string `json:"item2" yaml:"item2"`
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	LinkerId int64  `gorm:"index" json:"linkerId" yaml:"linkerId"`
}

// The base class definition for items2
type Entity1EntityItems2 struct {
	Item2    string `json:"item2" yaml:"item2"`
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	LinkerId int64  `gorm:"index" json:"linkerId" yaml:"linkerId"`
}

// The base class definition for content1
type Entity1EntityContent1 struct {
	Item1 emigo.Nullable[int64] `json:"item1" yaml:"item1"`
}

// The base class definition for content2
type Entity1EntityContent2 struct {
	Item2 emigo.Nullable[int64] `json:"item2" yaml:"item2"`
}

// The base class definition for nestedContainer
type Entity1EntityNestedContainer struct {
	NestedInner Entity1EntityNestedContainerNestedInner `gorm:"embedded" json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1EntityNestedContainerNestedInner struct {
	NestedItems    emigo.Array[Entity1EntityNestedContainerNestedInnerNestedItems] `gorm:"-" json:"nestedItems" yaml:"nestedItems"`
	NestedOwner    emigo.One[Entity2Entity]                                        `gorm:"-" json:"nestedOwner" yaml:"nestedOwner"`
	NestedItemsRow []*Entity1EntityNestedContainerNestedInnerNestedItems           `gorm:"foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE" json:"-" yaml:"-"`
	NestedOwnerId  int64                                                           `gorm:"index" json:"-" yaml:"-"`
	NestedOwnerRow *Entity2Entity                                                  `gorm:"foreignKey:NestedOwnerId;references:Id" json:"-" yaml:"-"`
}

// The base class definition for nestedItems
type Entity1EntityNestedContainerNestedInnerNestedItems struct {
	Label    string `json:"label" yaml:"label"`
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	LinkerId int64  `gorm:"index" json:"linkerId" yaml:"linkerId"`
}

// The base class definition for nestedContainerOpt
type Entity1EntityNestedContainerOpt struct {
	NestedInner Entity1EntityNestedContainerOptNestedInner `gorm:"embedded" json:"nestedInner" yaml:"nestedInner"`
}

// The base class definition for nestedInner
type Entity1EntityNestedContainerOptNestedInner struct {
	NestedItemsOpt    emigo.Array[Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt] `gorm:"-" json:"nestedItemsOpt" yaml:"nestedItemsOpt"`
	NestedItemsOptRow []*Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt           `gorm:"foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE" json:"-" yaml:"-"`
}

// The base class definition for nestedItemsOpt
type Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt struct {
	Label    string `json:"label" yaml:"label"`
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	LinkerId int64  `gorm:"index" json:"linkerId" yaml:"linkerId"`
}

func (x *Entity1Entity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity1EntityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
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
			Type: "collection",
		},
		{
			Name: prefix + "items4",
			Type: "collection?",
		},
		{
			Name: prefix + "owner",
			Type: "one",
		},
		{
			Name: prefix + "manager",
			Type: "one?",
		},
		{
			Name:     prefix + "content1",
			Type:     "object",
			Children: GetEntity1EntityContent1CliFlags("content1-"),
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
			Children: GetEntity1EntityNestedContainerCliFlags("nested-container-"),
		},
		{
			Name: prefix + "nested-container-opt",
			Type: "object?",
		},
		{
			Name: prefix + "items-row",
			Type: "complex",
		},
		{
			Name: prefix + "items2-row",
			Type: "complex",
		},
		{
			Name: prefix + "items3-row",
			Type: "complex",
		},
		{
			Name: prefix + "items4-row",
			Type: "complex",
		},
		{
			Name: prefix + "owner-id",
			Type: "int64",
		},
		{
			Name: prefix + "owner-row",
			Type: "complex",
		},
		{
			Name: prefix + "manager-id",
			Type: "int64",
		},
		{
			Name: prefix + "manager-row",
			Type: "complex",
		},
		{
			Name: prefix + "content2-row",
			Type: "complex",
		},
		{
			Name: prefix + "nested-container-opt-row",
			Type: "complex",
		},
	}
}
func CastEntity1EntityFromCli(c emigo.CliCastable) Entity1Entity {
	data := Entity1Entity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("title") {
		data.Title = c.String("title")
	}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastEntity1EntityItemsFromCli, "items", c)
	}
	if c.IsSet("items2") {
		data.Items2 = emigo.CapturePossibleArrayNullable(CastEntity1EntityItems2FromCli, "items2", c)
	}
	if c.IsSet("items3") {
		data.Items3 = emigo.CapturePossibleCollection(CastEntity2EntityFromCli, "items3", c)
	}
	if c.IsSet("items4") {
		data.Items4 = emigo.CapturePossibleCollectionNullable(CastEntity2EntityFromCli, "items4", c)
	}
	if c.IsSet("owner") {
		data.Owner = emigo.CapturePossibleOne(CastEntity2EntityFromCli, "owner", c)
	}
	if c.IsSet("manager") {
		data.Manager = emigo.CapturePossibleOneNullable(CastEntity2EntityFromCli, "manager", c)
	}
	if c.IsSet("content1") {
		data.Content1 = CastEntity1EntityContent1FromCli(c)
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
		data.NestedContainer = CastEntity1EntityNestedContainerFromCli(c)
	}
	if c.IsSet("nested-container-opt") {
		emigo.ParseNullable(c.String("nested-container-opt"), &data.NestedContainerOpt)
	}
	if c.IsSet("items-row") {
		if u, ok := any(&data.ItemsRow).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("items-row")))
		}
	}
	if c.IsSet("items2-row") {
		if u, ok := any(&data.Items2Row).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("items2-row")))
		}
	}
	if c.IsSet("items3-row") {
		if u, ok := any(&data.Items3Row).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("items3-row")))
		}
	}
	if c.IsSet("items4-row") {
		if u, ok := any(&data.Items4Row).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("items4-row")))
		}
	}
	if c.IsSet("owner-id") {
		data.OwnerId = int64(c.Int64("owner-id"))
	}
	if c.IsSet("owner-row") {
		if u, ok := any(&data.OwnerRow).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("owner-row")))
		}
	}
	if c.IsSet("manager-id") {
		data.ManagerId = int64(c.Int64("manager-id"))
	}
	if c.IsSet("manager-row") {
		if u, ok := any(&data.ManagerRow).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("manager-row")))
		}
	}
	if c.IsSet("content2-row") {
		if u, ok := any(&data.Content2Row).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("content2-row")))
		}
	}
	if c.IsSet("nested-container-opt-row") {
		if u, ok := any(&data.NestedContainerOptRow).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("nested-container-opt-row")))
		}
	}
	return data
}
func GetEntity1EntityItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "string",
		},
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "linker-id",
			Type: "int64",
		},
	}
}
func CastEntity1EntityItemsFromCli(c emigo.CliCastable) Entity1EntityItems {
	data := Entity1EntityItems{}
	if c.IsSet("item2") {
		data.Item2 = c.String("item2")
	}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("linker-id") {
		data.LinkerId = int64(c.Int64("linker-id"))
	}
	return data
}
func GetEntity1EntityItems2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "string",
		},
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "linker-id",
			Type: "int64",
		},
	}
}
func CastEntity1EntityItems2FromCli(c emigo.CliCastable) Entity1EntityItems2 {
	data := Entity1EntityItems2{}
	if c.IsSet("item2") {
		data.Item2 = c.String("item2")
	}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("linker-id") {
		data.LinkerId = int64(c.Int64("linker-id"))
	}
	return data
}
func GetEntity1EntityContent1CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item1",
			Type: "int64?",
		},
	}
}
func CastEntity1EntityContent1FromCli(c emigo.CliCastable) Entity1EntityContent1 {
	data := Entity1EntityContent1{}
	if c.IsSet("item1") {
		emigo.ParseNullable(c.String("item1"), &data.Item1)
	}
	return data
}
func GetEntity1EntityContent2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "item2",
			Type: "int64?",
		},
	}
}
func CastEntity1EntityContent2FromCli(c emigo.CliCastable) Entity1EntityContent2 {
	data := Entity1EntityContent2{}
	if c.IsSet("item2") {
		emigo.ParseNullable(c.String("item2"), &data.Item2)
	}
	return data
}
func GetEntity1EntityNestedContainerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "nested-inner",
			Type:     "object",
			Children: GetEntity1EntityNestedContainerNestedInnerCliFlags("nested-inner-"),
		},
	}
}
func CastEntity1EntityNestedContainerFromCli(c emigo.CliCastable) Entity1EntityNestedContainer {
	data := Entity1EntityNestedContainer{}
	if c.IsSet("nested-inner") {
		data.NestedInner = CastEntity1EntityNestedContainerNestedInnerFromCli(c)
	}
	return data
}
func GetEntity1EntityNestedContainerNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items",
			Type: "array",
		},
		{
			Name: prefix + "nested-owner",
			Type: "one",
		},
		{
			Name: prefix + "nested-items-row",
			Type: "complex",
		},
		{
			Name: prefix + "nested-owner-id",
			Type: "int64",
		},
		{
			Name: prefix + "nested-owner-row",
			Type: "complex",
		},
	}
}
func CastEntity1EntityNestedContainerNestedInnerFromCli(c emigo.CliCastable) Entity1EntityNestedContainerNestedInner {
	data := Entity1EntityNestedContainerNestedInner{}
	if c.IsSet("nested-items") {
		data.NestedItems = emigo.CapturePossibleArray(CastEntity1EntityNestedContainerNestedInnerNestedItemsFromCli, "nested-items", c)
	}
	if c.IsSet("nested-owner") {
		data.NestedOwner = emigo.CapturePossibleOne(CastEntity2EntityFromCli, "nested-owner", c)
	}
	if c.IsSet("nested-items-row") {
		if u, ok := any(&data.NestedItemsRow).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("nested-items-row")))
		}
	}
	if c.IsSet("nested-owner-id") {
		data.NestedOwnerId = int64(c.Int64("nested-owner-id"))
	}
	if c.IsSet("nested-owner-row") {
		if u, ok := any(&data.NestedOwnerRow).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("nested-owner-row")))
		}
	}
	return data
}
func GetEntity1EntityNestedContainerNestedInnerNestedItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "label",
			Type: "string",
		},
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "linker-id",
			Type: "int64",
		},
	}
}
func CastEntity1EntityNestedContainerNestedInnerNestedItemsFromCli(c emigo.CliCastable) Entity1EntityNestedContainerNestedInnerNestedItems {
	data := Entity1EntityNestedContainerNestedInnerNestedItems{}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("linker-id") {
		data.LinkerId = int64(c.Int64("linker-id"))
	}
	return data
}
func GetEntity1EntityNestedContainerOptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "nested-inner",
			Type:     "object",
			Children: GetEntity1EntityNestedContainerOptNestedInnerCliFlags("nested-inner-"),
		},
	}
}
func CastEntity1EntityNestedContainerOptFromCli(c emigo.CliCastable) Entity1EntityNestedContainerOpt {
	data := Entity1EntityNestedContainerOpt{}
	if c.IsSet("nested-inner") {
		data.NestedInner = CastEntity1EntityNestedContainerOptNestedInnerFromCli(c)
	}
	return data
}
func GetEntity1EntityNestedContainerOptNestedInnerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-items-opt",
			Type: "array",
		},
		{
			Name: prefix + "nested-items-opt-row",
			Type: "complex",
		},
	}
}
func CastEntity1EntityNestedContainerOptNestedInnerFromCli(c emigo.CliCastable) Entity1EntityNestedContainerOptNestedInner {
	data := Entity1EntityNestedContainerOptNestedInner{}
	if c.IsSet("nested-items-opt") {
		data.NestedItemsOpt = emigo.CapturePossibleArray(CastEntity1EntityNestedContainerOptNestedInnerNestedItemsOptFromCli, "nested-items-opt", c)
	}
	if c.IsSet("nested-items-opt-row") {
		if u, ok := any(&data.NestedItemsOptRow).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("nested-items-opt-row")))
		}
	}
	return data
}
func GetEntity1EntityNestedContainerOptNestedInnerNestedItemsOptCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "label",
			Type: "string",
		},
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "linker-id",
			Type: "int64",
		},
	}
}
func CastEntity1EntityNestedContainerOptNestedInnerNestedItemsOptFromCli(c emigo.CliCastable) Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt {
	data := Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt{}
	if c.IsSet("label") {
		data.Label = c.String("label")
	}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("linker-id") {
		data.LinkerId = int64(c.Int64("linker-id"))
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// TableName overrides the default gorm table name for Entity1Entity.
func (x *Entity1Entity) TableName() string {
	return "entity1_table"
}
