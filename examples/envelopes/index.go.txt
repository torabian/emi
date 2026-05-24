// This file will be added to ts-sdk folder, to export the content for writing
package prebuiltsdk

import "embed"

//go:embed *
var Content embed.FS
