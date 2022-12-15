package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-zoom/zoom_client"
)

type TableZoomAccountLockSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZoomAccountLockSettingsGenerator{}

func (x *TableZoomAccountLockSettingsGenerator) GetTableName() string {
	return "zoom_account_lock_settings"
}

func (x *TableZoomAccountLockSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZoomAccountLockSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZoomAccountLockSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZoomAccountLockSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			result, err := zoom_client.GetAccountLockSettingsOption(ctx, "", taskClient)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- result
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableZoomAccountLockSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZoomAccountLockSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_meeting").ColumnType(schema.ColumnTypeJSON).Description("Schedule meeting lock settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_notification").ColumnType(schema.ColumnTypeJSON).Description("Email notification lock settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("meeting_security").ColumnType(schema.ColumnTypeJSON).Description("Meeting security lock settings.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetAccountLockSettingsMeetingSecurity(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("Zoom account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetAccountID(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Account ID. Set to 'me' for the master account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := zoom_client.IdString()

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("in_meeting").ColumnType(schema.ColumnTypeJSON).Description("In meeting lock settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recording").ColumnType(schema.ColumnTypeJSON).Description("Recording lock settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("telephony").ColumnType(schema.ColumnTypeJSON).Description("Telephony lock settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tsp").ColumnType(schema.ColumnTypeJSON).Description("TSP lock settings.").Build(),
	}
}

func (x *TableZoomAccountLockSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
