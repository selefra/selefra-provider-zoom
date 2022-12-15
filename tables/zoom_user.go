package tables

import (
	"context"

	"github.com/himalayan-institute/zoom-lib-golang"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-zoom/zoom_client"
)

type TableZoomUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZoomUserGenerator{}

func (x *TableZoomUserGenerator) GetTableName() string {
	return "zoom_user"
}

func (x *TableZoomUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZoomUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZoomUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZoomUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zoom_client.Connect(ctx, taskClient.(*zoom_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := zoom.ListUsersOptions{
				PageSize:      300,
				IncludeFields: &[]string{"custom_attributes", "host_key"},
			}
			for {
				result, err := conn.ListUsers(opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range result.Users {
					resultChannel <- i
				}
				if result.NextPageToken == "" {
					break
				}
				opts.NextPageToken = &result.NextPageToken
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableZoomUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZoomUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("cms_user_id").ColumnType(schema.ColumnTypeString).Description("CMS ID of user, only enabled for Kaltura integration.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_ids").ColumnType(schema.ColumnTypeJSON).Description("IDs of groups where the user is a member.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plan_united_type").ColumnType(schema.ColumnTypeString).Description("This field is returned if the user is enrolled in the Zoom United plan.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("personal_meeting_url").ColumnType(schema.ColumnTypeString).Description("User's personal meeting url.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pmi").ColumnType(schema.ColumnTypeBigInt).Description("Personal meeting ID of the user.").
			Extractor(column_value_extractor.StructSelector("PMI")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vanity_url").ColumnType(schema.ColumnTypeString).Description("Personal meeting room URL, if the user has one.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("User's status: pending, active or inactive.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time when user's account was created.").
			Extractor(zoom_client.ExtractorTimestamp("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("jid").ColumnType(schema.ColumnTypeString).Description("User's JID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("job_title").ColumnType(schema.ColumnTypeString).Description("User's job title.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("login_type").ColumnType(schema.ColumnTypeInt).Description("Login type. 0 - Facebook, 1 - Google, 99 - API, 100 - ZOOM, 101 - SSO").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_key").ColumnType(schema.ColumnTypeString).Description("The host key of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_client_version").ColumnType(schema.ColumnTypeString).Description("The last client version that user used to login.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone_numbers").ColumnType(schema.ColumnTypeJSON).Description("User phone number, including verification status.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("language").ColumnType(schema.ColumnTypeString).Description("Default language for the Zoom Web Portal.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pic_url").ColumnType(schema.ColumnTypeString).Description("The URL for user's profile picture.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeInt).Description("User's plan type: 1 - Basic, 2 - Licensed or 3 - On-prem.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("User ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_name").ColumnType(schema.ColumnTypeString).Description("User's last name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company").ColumnType(schema.ColumnTypeString).Description("User's company.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_attributes").ColumnType(schema.ColumnTypeJSON).Description("Custom attributes, if any are assigned.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dept").ColumnType(schema.ColumnTypeString).Description("Department, if provided by the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_login_time").ColumnType(schema.ColumnTypeTimestamp).Description("User's last login time. There is a three-days buffer period for this field. For example, if user first logged in on 2020-01-01 and then logged out and logged in on 2020-01-02, the value of this field will still reflect the login time of 2020-01-01. However, if the user logs in on 2020-01-04, the value of this field will reflect the corresponding login time since it exceeds the three-day buffer period.").
			Extractor(zoom_client.ExtractorTimestamp("LastLoginTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the role assigned to the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verified").ColumnType(schema.ColumnTypeInt).Description("Display whether the user's email address for the Zoom account is verified or not. 1 - Verified user email. 0 - User's email not verified.").
			Extractor(column_value_extractor.StructSelector("Verified")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("Zoom account ID.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetAccountID(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_name").ColumnType(schema.ColumnTypeString).Description("User's first name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("User's email address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Description("User's location.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("use_pmi").ColumnType(schema.ColumnTypeBool).Description("Use Personal Meeting ID for instant meetings.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := zoom_client.GetUser(ctx, taskClient)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timezone").ColumnType(schema.ColumnTypeString).Description("The time zone of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("im_group_ids").ColumnType(schema.ColumnTypeJSON).Description("IDs of IM directory groups where the user is a member.").Build(),
	}
}

func (x *TableZoomUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
