package tables

import (
	"context"

	"github.com/himalayan-institute/zoom-lib-golang"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-zoom/zoom_client"
)

type TableZoomMyUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZoomMyUserGenerator{}

func (x *TableZoomMyUserGenerator) GetTableName() string {
	return "zoom_my_user"
}

func (x *TableZoomMyUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZoomMyUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZoomMyUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZoomMyUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zoom_client.Connect(ctx, taskClient.(*zoom_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := zoom.GetUserOpts{
				EmailOrID: "me",
			}
			result, err := conn.GetUser(opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- result
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableZoomMyUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZoomMyUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("first_name").ColumnType(schema.ColumnTypeString).Description("User's first name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_name").ColumnType(schema.ColumnTypeString).Description("User's last name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("Zoom account ID that contains this user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("User ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("User's email address.").Build(),
	}
}

func (x *TableZoomMyUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
