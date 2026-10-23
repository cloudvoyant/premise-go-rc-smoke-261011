package main

import (
	component "github.com/cloudvoyant/premise-go-rc-smoke-261011/v4/libs/components"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private helpers ------------------------------------------------------------
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		g, err := component.NewGreeting(ctx, "consumer", component.Args{Name: pulumi.String("consumer")})
		if err != nil {
			return err
		}
		ctx.Export("message", g.Message)
		return nil
	})
}
