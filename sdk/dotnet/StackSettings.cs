// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pequod.Stackmgmt
{
    [StackmgmtResourceType("stackmgmt:index:StackSettings")]
    public partial class StackSettings : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// Create a StackSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackSettings(string name, StackSettingsArgs? args = null, ComponentResourceOptions? options = null)
            : base("stackmgmt:index:StackSettings", name, args ?? new StackSettingsArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumi-pequod/pequod-mlc-stackmgmt",
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class StackSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Stack delete setting for automated purge processing.
        /// </summary>
        [Input("deleteStack")]
        public Input<string>? DeleteStack { get; set; }

        /// <summary>
        /// Drift management setting for refresh or correction.
        /// </summary>
        [Input("driftManagement")]
        public Input<string>? DriftManagement { get; set; }

        /// <summary>
        /// Pulumi access token to set up as a deployment environment variable if provided.
        /// </summary>
        [Input("pulumiAccessToken")]
        public Input<string>? PulumiAccessToken { get; set; }

        /// <summary>
        /// Team to which the stack should be assigned.
        /// </summary>
        [Input("teamAssignment")]
        public Input<string>? TeamAssignment { get; set; }

        /// <summary>
        /// Number of minutes to let stack live.
        /// </summary>
        [Input("ttlMinutes")]
        public Input<double>? TtlMinutes { get; set; }

        public StackSettingsArgs()
        {
        }
        public static new StackSettingsArgs Empty => new StackSettingsArgs();
    }
}
