package provider

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-zoom/tables"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableZoomUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZoomGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZoomAccountSettingsGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZoomAccountLockSettingsGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZoomMyUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZoomRoleGenerator{}),
	}
}
