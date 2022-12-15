package provider

import (
	"context"
	"github.com/selefra/selefra-provider-zoom/constants"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"

	"github.com/selefra/selefra-provider-zoom/zoom_client"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Zoom,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var zoomConfig zoom_client.Configs
				err := config.Unmarshal(&zoomConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				if len(zoomConfig.Providers) == 0 {
					zoomConfig.Providers = append(zoomConfig.Providers, zoom_client.Config{})
				}

				clients, err := zoom_client.NewClients(zoomConfig)

				if err != nil {
					clientMeta.ErrorF(constants.Newclientserrs, err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Accountinformationnotfound)
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `##  Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
#accounts:
#  - api_key: # YOUR_API_KEY.
#    api_secret: # YOUR_API_SECRET`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var zoomConfig zoom_client.Configs
				err := config.Unmarshal(&zoomConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_0,
				constants.NA,
				constants.Notsupported,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
