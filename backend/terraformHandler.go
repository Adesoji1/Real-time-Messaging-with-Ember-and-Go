// terraformHandler.go

package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-exec/tfexec"
)

type TerraformHandler struct {
	tf *tfexec.Terraform
}

// Define our own printfer interface since tfexec's printfer is not exported.
type printfer interface {
	Printf(format string, v ...interface{})
}

// NewTerraformHandler initializes a new TerraformHandler with the given working directory and Terraform binary path.
func NewTerraformHandler(workingDir string, execPath string) (*TerraformHandler, error) {
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		return nil, fmt.Errorf("error initializing Terraform: %s", err)
	}

	return &TerraformHandler{
		tf: tf,
	}, nil
}

// Initialize runs `terraform init` to initialize the Terraform configuration.
func (th *TerraformHandler) Initialize() error {
	err := th.tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return fmt.Errorf("error running terraform init: %s", err)
	}
	return nil
}

// Apply runs `terraform apply` to apply the Terraform configuration.
func (th *TerraformHandler) Apply() error {
	err := th.tf.Apply(context.Background())
	if err != nil {
		return fmt.Errorf("error applying Terraform configurations: %s", err)
	}
	return nil
}

// SetLogger sets the logger for Terraform executions.
func (th *TerraformHandler) SetLogger(logger printfer) {
	th.tf.SetLogger(logger)
}