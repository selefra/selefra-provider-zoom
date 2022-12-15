package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-zoom/zoom_client"
)

type TableZoomGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZoomGroupGenerator{}

func (x *TableZoomGroupGenerator) GetTableName() string {
	return "zoom_group"
}

func (x *TableZoomGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZoomGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZoomGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZoomGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zoom_client.Connect(ctx, taskClient.(*zoom_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			result, err := conn.ListGroups()
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range result.Groups {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableZoomGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZoomGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("total_members").ColumnType(schema.ColumnTypeInt).Description("Total number of members in the group.").
			Extractor(column_value_extractor.StructSelector("TotalMembers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("Zoom account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetAccountID(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Group ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Group name.").Build(),
	}
}

func (x *TableZoomGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
