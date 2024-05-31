// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stackmgmt

import (
	"context"
	"reflect"

	"github.com/pulumi-pequod/pequod-mlc-stackmgmt/sdk/v3/go/stackmgmt/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StackSettings struct {
	pulumi.ResourceState
}

// NewStackSettings registers a new resource with the given unique name, arguments, and options.
func NewStackSettings(ctx *pulumi.Context,
	name string, args *StackSettingsArgs, opts ...pulumi.ResourceOption) (*StackSettings, error) {
	if args == nil {
		args = &StackSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StackSettings
	err := ctx.RegisterRemoteComponentResource("stackmgmt:index:StackSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type stackSettingsArgs struct {
	// Stack delete setting for automated purge processing.
	DeleteStack *string `pulumi:"deleteStack"`
	// Drift management setting for refresh or correction.
	DriftManagement *string `pulumi:"driftManagement"`
	// Pulumi access token to set up as a deployment environment variable if provided.
	PulumiAccessToken *string `pulumi:"pulumiAccessToken"`
	// Team to which the stack should be assigned.
	TeamAssignment *string `pulumi:"teamAssignment"`
	// Time to live time setting.
	TtlTime *float64 `pulumi:"ttlTime"`
}

// The set of arguments for constructing a StackSettings resource.
type StackSettingsArgs struct {
	// Stack delete setting for automated purge processing.
	DeleteStack pulumi.StringPtrInput
	// Drift management setting for refresh or correction.
	DriftManagement pulumi.StringPtrInput
	// Pulumi access token to set up as a deployment environment variable if provided.
	PulumiAccessToken pulumi.StringPtrInput
	// Team to which the stack should be assigned.
	TeamAssignment pulumi.StringPtrInput
	// Time to live time setting.
	TtlTime pulumi.Float64PtrInput
}

func (StackSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackSettingsArgs)(nil)).Elem()
}

type StackSettingsInput interface {
	pulumi.Input

	ToStackSettingsOutput() StackSettingsOutput
	ToStackSettingsOutputWithContext(ctx context.Context) StackSettingsOutput
}

func (*StackSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**StackSettings)(nil)).Elem()
}

func (i *StackSettings) ToStackSettingsOutput() StackSettingsOutput {
	return i.ToStackSettingsOutputWithContext(context.Background())
}

func (i *StackSettings) ToStackSettingsOutputWithContext(ctx context.Context) StackSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSettingsOutput)
}

// StackSettingsArrayInput is an input type that accepts StackSettingsArray and StackSettingsArrayOutput values.
// You can construct a concrete instance of `StackSettingsArrayInput` via:
//
//	StackSettingsArray{ StackSettingsArgs{...} }
type StackSettingsArrayInput interface {
	pulumi.Input

	ToStackSettingsArrayOutput() StackSettingsArrayOutput
	ToStackSettingsArrayOutputWithContext(context.Context) StackSettingsArrayOutput
}

type StackSettingsArray []StackSettingsInput

func (StackSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackSettings)(nil)).Elem()
}

func (i StackSettingsArray) ToStackSettingsArrayOutput() StackSettingsArrayOutput {
	return i.ToStackSettingsArrayOutputWithContext(context.Background())
}

func (i StackSettingsArray) ToStackSettingsArrayOutputWithContext(ctx context.Context) StackSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSettingsArrayOutput)
}

// StackSettingsMapInput is an input type that accepts StackSettingsMap and StackSettingsMapOutput values.
// You can construct a concrete instance of `StackSettingsMapInput` via:
//
//	StackSettingsMap{ "key": StackSettingsArgs{...} }
type StackSettingsMapInput interface {
	pulumi.Input

	ToStackSettingsMapOutput() StackSettingsMapOutput
	ToStackSettingsMapOutputWithContext(context.Context) StackSettingsMapOutput
}

type StackSettingsMap map[string]StackSettingsInput

func (StackSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackSettings)(nil)).Elem()
}

func (i StackSettingsMap) ToStackSettingsMapOutput() StackSettingsMapOutput {
	return i.ToStackSettingsMapOutputWithContext(context.Background())
}

func (i StackSettingsMap) ToStackSettingsMapOutputWithContext(ctx context.Context) StackSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackSettingsMapOutput)
}

type StackSettingsOutput struct{ *pulumi.OutputState }

func (StackSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StackSettings)(nil)).Elem()
}

func (o StackSettingsOutput) ToStackSettingsOutput() StackSettingsOutput {
	return o
}

func (o StackSettingsOutput) ToStackSettingsOutputWithContext(ctx context.Context) StackSettingsOutput {
	return o
}

type StackSettingsArrayOutput struct{ *pulumi.OutputState }

func (StackSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackSettings)(nil)).Elem()
}

func (o StackSettingsArrayOutput) ToStackSettingsArrayOutput() StackSettingsArrayOutput {
	return o
}

func (o StackSettingsArrayOutput) ToStackSettingsArrayOutputWithContext(ctx context.Context) StackSettingsArrayOutput {
	return o
}

func (o StackSettingsArrayOutput) Index(i pulumi.IntInput) StackSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StackSettings {
		return vs[0].([]*StackSettings)[vs[1].(int)]
	}).(StackSettingsOutput)
}

type StackSettingsMapOutput struct{ *pulumi.OutputState }

func (StackSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackSettings)(nil)).Elem()
}

func (o StackSettingsMapOutput) ToStackSettingsMapOutput() StackSettingsMapOutput {
	return o
}

func (o StackSettingsMapOutput) ToStackSettingsMapOutputWithContext(ctx context.Context) StackSettingsMapOutput {
	return o
}

func (o StackSettingsMapOutput) MapIndex(k pulumi.StringInput) StackSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StackSettings {
		return vs[0].(map[string]*StackSettings)[vs[1].(string)]
	}).(StackSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackSettingsInput)(nil)).Elem(), &StackSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackSettingsArrayInput)(nil)).Elem(), StackSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackSettingsMapInput)(nil)).Elem(), StackSettingsMap{})
	pulumi.RegisterOutputType(StackSettingsOutput{})
	pulumi.RegisterOutputType(StackSettingsArrayOutput{})
	pulumi.RegisterOutputType(StackSettingsMapOutput{})
}
