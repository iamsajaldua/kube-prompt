// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var scaleOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Select all resources in the namespace of the specified resource types"},
	prompt.Suggest{Text: "--current-replicas", Description: "Precondition for current size. Requires that the current size of the resource match this value in order to scale."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to set a new size"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to set a new size"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "-o", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--output", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--record", Description: "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--replicas", Description: "The new desired number of replicas. Required."},
	prompt.Suggest{Text: "--resource-version", Description: "Precondition for resource version. Requires that the current resource version match this value in order to scale."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before giving up on a scale operation, zero means don't wait. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h)."},
}
