package pkg

import (
	azurekeyvaultv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/provider/azure/azurekeyvault/v1"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *azurekeyvaultv1.AzureKeyVaultStackInput) error {
	azureCredential := stackInput.AzureCredential
	//create azure provider using the credentials from the input
	_, err := azure.NewProvider(ctx,
		"azure",
		&azure.ProviderArgs{
			ClientId:       pulumi.String(azureCredential.ClientId),
			ClientSecret:   pulumi.String(azureCredential.ClientSecret),
			SubscriptionId: pulumi.String(azureCredential.SubscriptionId),
			TenantId:       pulumi.String(azureCredential.TenantId),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create azure provider")
	}
	return nil
}
