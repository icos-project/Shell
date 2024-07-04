/*
ICOS Shell

Testing ControllerAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"log"

	"shellclient/cmd"
	openapiclient "shellclient/pkg/openapi"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ControllerAPIService(t *testing.T) {

	cmd.InitConfigForTesting()
	apiClient := openapiclient.Client

	t.Run("Test ControllerAPIService AddController", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		viper.SetConfigFile("../../../config_client.yml") // Read the config file
		viper.AddConfigPath(".")                          // look for config in the working directory

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		tokenRaw := viper.GetString("auth_token")
		token := tokenRaw[1 : len(tokenRaw)-1]

		controller := *openapiclient.NewController("name_test", "address_test")

		httpRes, err := apiClient.ControllerAPI.AddController(context.Background()).ApiKey(token).Controller(controller).Execute()

		require.Nil(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

	})

	t.Run("Test ControllerAPIService GetControllers", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ControllerAPI.GetControllers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}