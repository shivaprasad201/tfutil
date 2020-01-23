package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

var template = map[string]string{
	"main": `#############################################################################################
	## This file contains the main resource code for creating the resources.
	## Use this file to create and manage resource.

	# Example:
		# resource "type" "name" {
		# <resource configuration>
		# }

		# module "name" {
		# source = "source"
		# <module configuration>	  
		# }
#############################################################################################`,

	"data": `#############################################################################################
	## This file contains the data used by the resources.
	## Use this file for creating data blocks used by resources.
	
	# Example:
		# data "type" "name" {
		# <data configuration>  
		# } 
#############################################################################################`,

	"variables": `#############################################################################################
	## This file contains the variables used by the resources.
	## Use this file for creating variable blocks used by the resource.
	
	# Example:
		# variable "name" {
		#   type = "type"
		#   description = "description"
		#   default = "default"
		# }
#############################################################################################`,

	"outputs": `#############################################################################################
	## This file contains the outputs that are to be displayed as a result of execution.
	## Use this file to declare output blocks to be displayed after a terraform apply.

	# Example:
		# output "name" {
		#   value = "value"
		# }	
#############################################################################################`,

	"locals": `#############################################################################################
	## This file contains the local data used by the resource.
	## Use this file to declare locals that are used by the resources.
	
	# Example:
		# locals {
		# <locals data>
		# }
#############################################################################################`,

	"provider": `#############################################################################################
	## This file contains the provider configuration.
	## Use this file to declare the providers.
	
	# Example:
		# provider "name" {
		# <provider configuration>
		# }
#############################################################################################`,

	"backend": `#############################################################################################
	## This file contains the terraform backend configuration.
	## Use this file to declare the backend to store the state file.
	
	# Example:
		# terraform {
		#  backend "name"{
		# <backend configuration>	
		#  }
		# }	
#############################################################################################`,
}

// InitFiles initalises the set of empty terraform configuration files in a provided directory
func InitFiles(path string) {
	files := []string{"main", "data", "variables", "outputs", "locals", "provider", "backend"}

	for _, fileName := range files {
		filePath := filepath.Join(path, filepath.Base(fmt.Sprintf("%s.tf", fileName)))

		f, err := os.Create(filePath)
		if err != nil {
			HandleError(err, fmt.Sprintf("Error: Cant create file %s", fileName))
		}

		f.WriteString(template[fileName])

	}
}
