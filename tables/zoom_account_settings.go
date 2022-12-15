package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-zoom/zoom_client"
)

type TableZoomAccountSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZoomAccountSettingsGenerator{}

func (x *TableZoomAccountSettingsGenerator) GetTableName() string {
	return "zoom_account_settings"
}

func (x *TableZoomAccountSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZoomAccountSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZoomAccountSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZoomAccountSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			result, err := zoom_client.GetAccountSettingsOption(ctx, "", taskClient)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- result
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableZoomAccountSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZoomAccountSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("email_notification").ColumnType(schema.ColumnTypeJSON).Description("Email notification settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security").ColumnType(schema.ColumnTypeJSON).Description("Security settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration").ColumnType(schema.ColumnTypeJSON).Description("Integration settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trusted_domains").ColumnType(schema.ColumnTypeJSON).Description("Associated domains allow all users with that email domain to be prompted to join the account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetAccountTrustedDomains(ctx, taskClient)
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
		table_schema_generator.NewColumnBuilder().ColumnName("tsp").ColumnType(schema.ColumnTypeJSON).Description("TSP settings.").
			Extractor(column_value_extractor.StructSelector("TSP")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed_domains").ColumnType(schema.ColumnTypeJSON).Description("Associated domains allow all users with that email domain to be prompted to join the account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetAccountManagedDomains(ctx, task)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Domains")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("in_meeting").ColumnType(schema.ColumnTypeJSON).Description("In meeting settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("meeting_authentication").ColumnType(schema.ColumnTypeJSON).Description("Meeting authentication options applied to the account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := zoom_client.GetAccountSettingsMeetingAuthentication(ctx, taskClient)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("Zoom account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetAccountID(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_meeting").ColumnType(schema.ColumnTypeJSON).Description("Schedule meeting settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("telephony").ColumnType(schema.ColumnTypeJSON).Description("Telephony settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("feature").ColumnType(schema.ColumnTypeJSON).Description("Feature settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recording_authentication").ColumnType(schema.ColumnTypeJSON).Description("Recording authentication options applied to the account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				result, err := zoom_client.GetAccountSettingsRecordingAuthentication(ctx, taskClient)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("meeting_security").ColumnType(schema.ColumnTypeJSON).Description("Meeting security settings applied to the account.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetAccountSettingsMeetingSecurity(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recording").ColumnType(schema.ColumnTypeJSON).Description("Recording settings.").Build(),
	}
}

func (x *TableZoomAccountSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
