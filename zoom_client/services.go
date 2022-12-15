package zoom_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/selefra/selefra-provider-zoom/constants"
	"os"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	zoom "github.com/himalayan-institute/zoom-lib-golang"
)

func Connect(ctx context.Context, zoomConfig *Config) (*zoom.Client, error) {
	apiKey := os.Getenv(constants.ZOOMAPIKEY)
	apiSecret := os.Getenv(constants.ZOOMAPISECRET)
	if zoomConfig.APIKey != constants.Constants_3 {
		apiKey = zoomConfig.APIKey
	}
	if zoomConfig.APISecret != constants.Constants_4 {
		apiSecret = zoomConfig.APISecret
	}

	if apiKey == constants.Constants_5 || apiSecret == constants.Constants_6 {

		return nil, errors.New(constants.Apikeyandapisecretmustbeconfigured)
	}
	return zoom.NewClient(apiKey, apiSecret), nil
}

type zoomAccountID struct {
	AccountID string
}

func GetAccountID(ctx context.Context, taskClient any) (any, error) {

	conn, err := Connect(ctx, taskClient.(*Client).Config)
	if err != nil {
		return nil, err
	}
	opts := zoom.GetUserOpts{
		EmailOrID: constants.Me,
	}
	user, err := conn.GetUser(opts)
	if err != nil {
		return nil, err
	}

	return zoomAccountID{
		AccountID: user.AccountID,
	}, nil
}

func GetAccountLockSettingsOption(ctx context.Context, option string, taskClient any) (interface{}, error) {
	conn, err := Connect(ctx, taskClient.(*Client).Config)
	if err != nil {

		return nil, err
	}
	id := constants.Me

	opts := zoom.GetAccountLockSettingsOpts{
		AccountID: id,
	}
	if option != constants.Constants_7 {
		opts.Option = option
	}

	result, err := conn.GetAccountLockSettings(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {

			return nil, nil
		}

		return nil, err
	}
	return result, nil
}

func IdString() (any, error) {
	return constants.Me, nil
}

func GetAccountManagedDomains(ctx context.Context, taskClient any) (any, error) {
	conn, err := Connect(ctx, taskClient.(*Client).Config)
	if err != nil {

		return nil, err
	}
	id := constants.Constants_8
	opts := zoom.GetAccountManagedDomainsOpts{
		AccountID: id,
	}
	res, err := conn.GetAccountManagedDomains(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {
			return nil, nil
		}
		return nil, err
	}

	if res.Domains == nil {
		res.Domains = []string{}
	}
	return res, nil
}

func GetUser(ctx context.Context, taskClient any) (any, error) {
	conn, err := Connect(ctx, taskClient.(*Client).Config)
	if err != nil {

		return nil, err
	}
	emailOrID := constants.Constants_9
	opts := zoom.GetUserOpts{
		EmailOrID: emailOrID,
	}
	result, err := conn.GetUser(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {

			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func GetRole(ctx context.Context, taskClient any) (any, error) {
	conn, err := Connect(ctx, taskClient.(*Client).Config)
	if err != nil {

		return nil, err
	}
	opts := zoom.GetRoleOpts{
		ID: constants.Constants_10,
	}
	result, err := conn.GetRole(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {

			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func GetAccountLockSettingsMeetingSecurity(ctx context.Context, taskClient any) (interface{}, error) {
	return GetAccountLockSettingsOption(ctx, constants.Meetingsecurity, taskClient)
}

func GetAccountSettingsOption(ctx context.Context, option string, taskClient any) (any, error) {
	conn, err := Connect(ctx, taskClient.(*Client).Config)
	if err != nil {

		return nil, err
	}
	opts := zoom.GetAccountSettingsOpts{
		AccountID: constants.Me,
	}
	if option != constants.Constants_11 {
		opts.Option = option
	}

	result, err := conn.GetAccountSettings(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {

			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

type Authentication struct {
	Enabled               bool        `json:"enabled"`
	AuthenticationOptions interface{} `json:"authentication_options"`
}

func GetAccountSettingsMeetingAuthentication(ctx context.Context, taskClient any) (any, error) {
	result, err := GetAccountSettingsOption(ctx, constants.Meetingauthentication, taskClient)
	if err != nil {
		return nil, err
	}
	settings := result.(zoom.AccountSettings)
	auth := Authentication{
		Enabled:               settings.MeetingAuthentication,
		AuthenticationOptions: settings.AuthenticationOptions,
	}
	return auth, nil
}

func GetAccountSettingsRecordingAuthentication(ctx context.Context, taskClient any) (any, error) {
	result, err := GetAccountSettingsOption(ctx, constants.Recordingauthentication, taskClient)
	if err != nil {
		return nil, err
	}
	settings := result.(zoom.AccountSettings)
	auth := Authentication{
		Enabled:               settings.RecordingAuthentication,
		AuthenticationOptions: settings.AuthenticationOptions,
	}
	return auth, nil
}

func GetAccountSettingsMeetingSecurity(ctx context.Context, taskClient any) (any, error) {
	return GetAccountSettingsOption(ctx, constants.Meetingsecurity, taskClient)
}

func GetAccountTrustedDomains(ctx context.Context, taskClient any) (any, error) {
	conn, err := Connect(ctx, taskClient.(*Client).Config)
	if err != nil {

		return nil, err
	}
	opts := zoom.GetAccountTrustedDomainsOpts{
		AccountID: constants.Me,
	}
	result, err := conn.GetAccountTrustedDomains(opts)
	if err != nil {
		if e, ok := err.(*zoom.APIError); ok && e.Code == 1001 {

			return nil, nil
		}
		return nil, err
	}
	return result, nil
}
func ExtractorTimestamp(path string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		v, err := column_value_extractor.StructSelector(path).Extract(ctx, clientMeta, client, task, row, column, result)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, nil
		}
		ts, ok := v.(zoom.Time)
		if !ok {
			return nil, schema.NewDiagnosticsAddErrorMsg(fmt.Sprintf(constants.UnextectedtypewantedzoomTimehaveT, v))
		}
		return ts.Time, nil
	})
}
