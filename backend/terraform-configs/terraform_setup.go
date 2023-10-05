package main

import (
    "context"
    "fmt"
    "log"

    "github.com/hashicorp/go-version"
    "github.com/hashicorp/hc-install/product"
    "github.com/hashicorp/hc-install/releases"
    "github.com/hashicorp/terraform-exec/tfexec"
)

func main() {
    installer := &releases.ExactVersion{
        Product: product.Terraform,
        Version: version.Must(version.NewVersion("1.0.6")),
    }

    execPath, err := installer.Install(context.Background())
    if err != nil {
        log.Fatalf("error installing Terraform: %s", err)
    }

    // Point to the terraform-configs directory within your backend directory
    workingDir := "/home/adesoji/real-time-messaging-app/backend/terraform-configs"
    tf, err := tfexec.NewTerraform(workingDir, execPath)
    if err != nil {
        log.Fatalf("error running NewTerraform: %s", err)
    }

    err = tf.Init(context.Background(), tfexec.Upgrade(true))
    if err != nil {
        log.Fatalf("error running Init: %s", err)
    }

    state, err := tf.Show(context.Background())
    if err != nil {
        log.Fatalf("error running Show: %s", err)
    }

    fmt.Println(state.FormatVersion) // "0.1"

    // Apply Terraform configurations
    err = tf.Apply(context.Background())
    if err != nil {
        log.Fatalf("error applying Terraform configurations: %s", err)
    }
}
